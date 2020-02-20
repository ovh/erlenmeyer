package graphite

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/ovh/erlenmeyer/core"
)

const (
	opEq        = "="
	opRegExp    = "=~"
	opNotEq     = "!="
	opNotRegExp = "!=~"
)

var (
	rString                = regexp.MustCompile(`('|").+?('|")`)
	rTransformedString     = regexp.MustCompile(`\$\$\d+`)
	rNumber                = regexp.MustCompile(`\d+(\.\d+)?(e(\.\d+)?)?`)
	rName                  = regexp.MustCompile(fmt.Sprintf(`(%[1]s)|(%[1]s)?(\[(%[2]s)\])+(%[1]s)?`, `[\w\-*]+`, `[\w\-]+`))
	rNameWithCurlyBrackets = regexp.MustCompile(fmt.Sprintf(`(%[1]s)|(%[1]s)?\{(%[1]s)(, ?(%[1]s))*\}(%[1]s)?`, rName))
	rCurlyBrackets         = regexp.MustCompile(fmt.Sprintf(`(\{(%[1]s)(,?(%[1]s))*\})+`, rName))
	rSerie                 = regexp.MustCompile(fmt.Sprintf(`((%[1]s)(\.(%[1]s))*)`, rNameWithCurlyBrackets))
	rArgument              = regexp.MustCompile(fmt.Sprintf(`(%s)|(%s)|(%s)|(%s)`, rSerie, rTransformedString, rString, rNumber))
	rFunction              = regexp.MustCompile(fmt.Sprintf(`\w+\(((%s)(\,\s*?(%s))*)?\)`, rArgument, rArgument))
)

// Function structure using to describe a graphite function call
type Function struct {
	Name       string
	Arguments  []string
	Parameters map[string]string
}

// graphite go parser implementation
// see https://github.com/graphite-project/graphite-web/blob/master/webapp/graphite/render/grammar.py

// Parse a graphite query
func Parse(target, from, until string, node *core.Node) (*core.Node, error) {
	var err error

	stack := make([]Function, 0)
	if ContainsFunction(target) {
		stack, err = ParseQuery(target)
		if err != nil {
			return nil, err
		}
	} else {
		stack = append(stack, Function{
			Name:       "fetch",
			Arguments:  []string{target, from, until},
			Parameters: make(map[string]string),
		})
	}

	for _, element := range stack {
		element.Parameters["target"] = target
		element.Parameters["from"] = from
		element.Parameters["until"] = until
		element.Parameters["func"] = element.Name
	}

	return generateNode(node, stack)
}

// ParseSerie return the node corresponding to the query
func ParseSerie(serie string) (string, map[string]string, error) {
	split := strings.Split(serie, ";")
	serie = toWarpScriptRegExp(split[0])
	labels, err := parseLabels(split[1:])

	return serie, labels, err
}

// ParseQuery and transform it into a call stack
func ParseQuery(query string) ([]Function, error) {
	if !ContainsFunction(query) {
		return nil, fmt.Errorf("No such functions in %s", query)
	}

	i := 0
	arguments := make(map[string]string)
	for ContainsString(query) {
		argument := rString.FindString(query)
		arguments[fmt.Sprintf("$$%d", i)] = strings.Trim(argument, " '\"")
		query = strings.Replace(query, argument, fmt.Sprintf("$$%d", i), 1)
		i++
	}

	functions := make([]Function, 0)
	for ContainsFunction(query) {
		function := rFunction.FindString(query)
		functions = append(functions, ParseFunction(function))
		query = strings.Replace(query, function, swap, 1)
	}

	for i, function := range functions {
		for j, argument := range function.Arguments {
			if strings.HasPrefix(argument, "$$") {
				if val, ok := arguments[argument]; ok {
					functions[i].Arguments[j] = val
				}
			}
		}
	}

	return functions, nil
}

// ContainsFunction detect if the query contains a function
func ContainsFunction(query string) bool {
	return rFunction.MatchString(query)
}

// ContainsString detect if the query contains a string
func ContainsString(query string) bool {
	return rString.MatchString(query)
}

// ContainsNameWithCurlyBrackets detect if the query contains a name with curly brackets
func ContainsNameWithCurlyBrackets(query string) bool {
	return rNameWithCurlyBrackets.MatchString(query)
}

// ContainsCurlyBrackets detect if the query contains curly brackets
func ContainsCurlyBrackets(query string) bool {
	return rCurlyBrackets.MatchString(query)
}

// ContainsSerie detect if the query contains a serie
func ContainsSerie(query string) bool {
	return rSerie.MatchString(query)
}

// ParseFunction transform a graphite function into a struct
func ParseFunction(function string) Function {
	parenthesis := strings.Index(function, "(")

	for ContainsCurlyBrackets(function) {
		name := rCurlyBrackets.FindString(function)
		replace := strings.Replace(name, "{", "(", 1)
		replace = strings.Replace(replace, "}", ")", 1)
		replace = strings.Replace(replace, ",", "|", -1)
		function = strings.Replace(function, name, replace, 1)
	}

	fn := Function{
		Parameters: make(map[string]string),
	}
	fn.Name = function[:parenthesis]
	if len(function[parenthesis+1:len(function)-1]) > 0 {
		fn.Arguments = strings.Split(function[parenthesis+1:len(function)-1], ",")
	} else {
		fn.Arguments = make([]string, 0)
	}

	for i, argument := range fn.Arguments {
		fn.Arguments[i] = strings.Trim(argument, " ")
	}

	return fn
}

func generateNode(node *core.Node, fns []Function) (*core.Node, error) {
	var err error

	last := len(fns) - 1
	fn := fns[last]
	f, err := GetFunction(fn.Name)
	if err != nil {
		return nil, err
	}

	node, err = f(node, fn.Arguments, fn.Parameters)
	if err != nil {
		return nil, err
	}

	if last == 0 {
		return node, nil
	}

	return generateNode(node, fns[:last])
}

func parseLabels(args []string) (map[string]string, error) {
	labels := make(map[string]string)
	for _, label := range args {
		if !strings.Contains(label, opEq) {
			return nil, errors.New("Label filter is not well formatted")
		}

		op := opEq
		if strings.Contains(label, opNotRegExp) {
			op = opNotRegExp
		} else if strings.Contains(label, opRegExp) {
			op = opRegExp
		} else if strings.Contains(label, opNotEq) {
			op = opNotEq
		}

		tmp := strings.Split(label, op)
		name := tmp[0]
		value := tmp[1]

		if op == opNotRegExp || op == opNotEq {
			labels[name] = toWarpScriptRegExp(fmt.Sprintf("~(?!%s)*", value))
		} else if op == opRegExp {
			labels[name] = toWarpScriptRegExp(fmt.Sprintf("~%s", value))
		} else {
			labels[name] = value
		}
	}

	return labels, nil
}

func toWarpScriptRegExp(serie string) string {
	serie = strings.Replace(serie, ".", "\\.", -1)
	serie = strings.Replace(serie, "*", ".*?", -1)

	if !strings.Contains(serie, "~") {
		serie = "~" + serie
	}

	return serie
}
