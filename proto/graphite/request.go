package graphite

import (
	"fmt"
	"strconv"

	"github.com/ovh/erlenmeyer/core"
)

const findQueryWarpScript = `
'%s' 'query' STORE
$query '.' SPLIT SIZE 1 - 'level' STORE

[ $token '%s' {} ] FIND

<%% DROP NAME %%> LMAP UNIQUE
<%% DROP '.' SPLIT
	<%% DUP SIZE 1 - $level == %%>
		<%% 1 %%>
		<%% 0 %%>
	IFTE 'leaf' STORE

	<%% $leaf 0 == %%>
		<%% 1 %%>
		<%% 0 %%>
	IFTE 'children' STORE

	[ 0 $level ] SUBLIST 'name' STORE
	$name LIST-> '.' SWAP JOIN 'path' STORE
	$name REVERSE 0 GET 'name' STORE

	{ 'leaf' $leaf 'id' $path 'text' $name 'expandable' $children 'allowChildren' $children }
%%> LMAP UNIQUE
STOP
`

const expandQueryWarpScript = `
'%s' 'query' STORE
$query '.' SPLIT SIZE 1 - 'level' STORE

[ $token '%s' {} ] FIND

<%% DROP NAME %%> LMAP UNIQUE
<%% DROP '.' SPLIT [ 0 $level ] SUBLIST LIST-> '.' SWAP JOIN %%> LMAP UNIQUE
`

// CreateFindRequest return children of a serie given in parameter as mc2 tree
func CreateFindRequest(query string, wildcards bool) (*core.Node, error) {
	if wildcards {
		query += ".*"
	}

	serie, _, err := ParseSerie(query)
	if err != nil {
		return nil, err
	}

	root := core.NewNode(core.WarpScriptPayload{
		WarpScript: fmt.Sprintf(findQueryWarpScript, query, serie),
	})

	return root, nil
}

// CreateExpandRequest return children of a serie given in parameter as mc2 tree
func CreateExpandRequest(query string) (*core.Node, error) {
	serie, _, err := ParseSerie(query)
	if err != nil {
		return nil, err
	}

	root := core.NewNode(core.WarpScriptPayload{
		WarpScript: fmt.Sprintf(expandQueryWarpScript, query, serie),
	})

	return root, nil
}

// CreateRenderRequest return what you ask for as mc2 tree
func CreateRenderRequest(target, from, until string) (*core.Node, error) {
	// default time values
	if len(from) == 0 {
		from = "-24h"
	}

	if len(until) == 0 {
		until = "-0s"
	}

	// parse time value
	start, err := ParseTime([]byte(from))
	if err != nil {
		return nil, err
	}

	end, err := ParseTime([]byte(until))
	if err != nil {
		return nil, err
	}

	from = strconv.FormatInt(start.UnixNano()/1000, 10)
	until = strconv.FormatInt(end.UnixNano()/1000, 10)
	root := core.NewEmptyNode()
	_, err = Parse(target, from, until, root)
	if err != nil {
		return nil, err
	}

	return root, nil
}
