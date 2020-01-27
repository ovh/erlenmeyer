package core

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

// ToWarpScriptWithTime translate a node to a WarpScript with start and end prefix
func (n *Node) ToWarpScriptWithTime(token string, query string, step string, start Time, end Time) string {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("%v000 ", start) + " 'start' STORE \n")
	b.WriteString(fmt.Sprintf("%v000 ", end) + " 'end' STORE \n")
	return b.String() + n.ToWarpScript(token, query, step)
}

// ToWarpScript translate a node to a WarpScript
func (n *Node) ToWarpScript(token string, query string, step string) string {
	var b bytes.Buffer

	b.WriteString("/* Generating query for " + query + " */ \n")

	if len(step) != 0 {
		b.WriteString(step + " 'step' STORE \n")
		b.WriteString("0 'instant' STORE \n")
	} else {
		b.WriteString("0 'step' STORE \n")
		b.WriteString("1 'instant' STORE \n")
	}

	b.WriteString("'" + token + "' 'token' STORE \n")
	b.WriteString("$token AUTHENTICATE \n")
	// keys of STACKATTRIBUTE can be found here: https://github.com/cityzendata/warp10-platform/blob/master/warp10/src/main/java/io/warp10/script/WarpScriptStack.java
	b.WriteString("'stack.maxops.hard' STACKATTRIBUTE DUP <% ISNULL ! %> <% MAXOPS %> <% DROP %> IFTE\n")
	b.WriteString("'fetch.limit.hard' STACKATTRIBUTE DUP <% ISNULL ! %> <% LIMIT %> <% DROP %> IFTE\n")

	// Generate node WarpScript
	n.toWarpScript(&b)

	// Adding footer
	b.WriteString("\nDUP TYPEOF <% 'GTS' == %> <% [ SWAP ] %> IFT SORT\n")
	b.WriteString("\nUNBUCKETIZE [ SWAP mapper.finite 0 0 0 ] MAP\n")

	return b.String()
}

// InternalToWarpScript translate an internal Node to WarpScript
func (n *Node) InternalToWarpScript(query string) string {
	var b bytes.Buffer

	n.toWarpScript(&b)

	return b.String()
}

// toWarpScript is recursive, be aware ;-)
func (n *Node) toWarpScript(b *bytes.Buffer) {

	if n.Right != nil {

		leftNodeType := fmt.Sprintf("%T", n.Left.Payload)
		rightNodeType := fmt.Sprintf("%T", n.Right.Payload)

		n.Left.toWarpScript(b)
		b.WriteString("'left-" + strconv.Itoa(n.Level) + "' STORE\n")
		n.Right.toWarpScript(b)
		b.WriteString("'right-" + strconv.Itoa(n.Level) + "' STORE\n")

		if strings.Contains(leftNodeType, "NumberLiteralPayload") && strings.Contains(rightNodeType, "NumberLiteralPayload") {
			b.WriteString("$left-" + strconv.Itoa(n.Level) + " 'left' STORE\n")
			b.WriteString("$right-" + strconv.Itoa(n.Level) + " 'right' STORE\n")
			b.WriteString("$left $right \n")
		}

		if !strings.Contains(leftNodeType, "NumberLiteralPayload") && !strings.Contains(rightNodeType, "NumberLiteralPayload") {
			b.WriteString("$left-" + strconv.Itoa(n.Level) + " 'left' STORE\n")
			b.WriteString("$right-" + strconv.Itoa(n.Level) + " 'right' STORE\n")
			b.WriteString("$left $right 2 ->LIST\n")
		}

		if !strings.Contains(leftNodeType, "NumberLiteralPayload") && strings.Contains(rightNodeType, "NumberLiteralPayload") {
			b.WriteString("$left-" + strconv.Itoa(n.Level) + " 'left' STORE\n")
			b.WriteString("$right-" + strconv.Itoa(n.Level) + " 'right' STORE\n")
			b.WriteString("$left\n")
		}

		if strings.Contains(leftNodeType, "NumberLiteralPayload") && !strings.Contains(rightNodeType, "NumberLiteralPayload") {
			b.WriteString("$left-" + strconv.Itoa(n.Level) + " 'left' STORE\n")
			b.WriteString("$right-" + strconv.Itoa(n.Level) + " 'right' STORE\n")
			b.WriteString("$right\n")
		}

		n.Write(b)
	} else {
		if n.Left != nil {
			n.Left.toWarpScript(b)
		}
		n.Write(b)
	}
}

// Write write node content
// nolint: gocyclo
func (n *Node) Write(b *bytes.Buffer) {
	switch p := n.Payload.(type) {
	case FetchPayload:
		b.WriteString("[ $token '")
		b.WriteString(p.ClassName)
		b.WriteString("' ")
		b.WriteString(printLabelsAsWarpScriptHash(p.Labels))

		var start, stop uint64
		var errStart, errStop error

		// checking if start < end. If so, we have two timestamp, and we need to calculate the range
		start, errStart = strconv.ParseUint(p.Start, 10, 64)
		stop, errStop = strconv.ParseUint(p.End, 10, 64)

		if errStart == nil && errStop == nil {
			if start < stop {
				b.WriteString(" " + p.End)
				b.WriteString(" " + p.End + " " + p.Start + " - ")
			}
		} else {

			if p.Absent {
				b.WriteString(" " + p.Start + " 15 m - ISO8601")
			} else {
				b.WriteString(" " + p.Start + " ISO8601")
			}
			b.WriteString(" " + p.End + " ISO8601")
		}

		b.WriteString(" ] FETCH \n")
		b.WriteString("DUP <% VALUES SIZE 0 == %> <% NEWGTS '" + p.ClassName + "' RENAME " + printLabelsAsWarpScriptHash(p.Labels) + " RELABEL 1 ->LIST APPEND %> IFT\n")

		if len(p.Offset) > 0 {
			b.WriteString(p.Offset + " TIMESHIFT \n")
		}

	case FindPayload:
		b.WriteString("[ $token '")
		b.WriteString(p.ClassName)
		b.WriteString("' ")
		b.WriteString(printLabelsAsWarpScriptHash(p.Labels))
		b.WriteString("] FIND \n")

	case WarpScriptPayload:
		b.WriteString(p.WarpScript)
		b.WriteString("\n")

	case StorePayload:
		b.WriteString(p.Value)
		b.WriteString(" '")
		b.WriteString(p.Name)
		b.WriteString("'")
		b.WriteString(" STORE\n")

	case BucketizePayload:

		// In case of a rate, FALSE RESETS needs to be applied BEFORE BUCKETIZE.
		// If not, bucketizer.mean will create a mean during the false resets
		if p.ApplyRate {
			b.WriteString("FALSE RESETS\n")
		}

		if p.Absent {
			b.WriteString(p.BucketRange + " [ SWAP " + p.Op + " " + p.LastBucket + " " + p.BucketSpan + " " + p.BucketCount + " 15 m " + p.BucketSpan + " / + ] BUCKETIZE\n")
			b.WriteString("[ SWAP mapper.last 15 m $step / 0 $instant -1 * ] MAP\n")
			b.WriteString(p.BucketRange + " [ SWAP " + p.Op + " " + p.LastBucket + " " + p.BucketSpan + " " + p.BucketCount + " ] BUCKETIZE\n")
		} else {
			b.WriteString(p.BucketRange + " [ SWAP " + p.Op + " " + p.LastBucket + " " + p.BucketSpan + " " + p.BucketCount + " ] BUCKETIZE\n")
			b.WriteString(p.Filler + "\n")
		}
		b.WriteString(p.BucketCount + " 'bucketCount' STORE\n")

	case MapperPayload:
		b.WriteString("[ SWAP ")
		b.WriteString(p.Constant)
		b.WriteString(" mapper.")
		b.WriteString(p.Mapper)
		b.WriteString(" ")
		b.WriteString(p.PreWindow)
		b.WriteString(" ")
		b.WriteString(p.PostWindow)
		b.WriteString(" ")
		b.WriteString(p.Occurrences)
		b.WriteString(" ] MAP\n")

	case AddValuePayload:
		b.WriteString(p.Timestamp)
		b.WriteString(" ")
		b.WriteString(p.Latitude)
		b.WriteString(" ")
		b.WriteString(p.Longitude)
		b.WriteString(" ")
		b.WriteString(p.Elevation)
		b.WriteString(" ")
		b.WriteString(p.Value)
		b.WriteString(" ADDVALUE\n")

	case FillValuePayload:
		b.WriteString("DUP TYPEOF <% 'GTS' == %> <% [ SWAP ] %> IFT \n")
		b.WriteString("[ ")
		b.WriteString(p.Latitude)
		b.WriteString(" ")
		b.WriteString(p.Longitude)
		b.WriteString(" ")
		b.WriteString(p.Elevation)
		b.WriteString(" ")
		b.WriteString(p.Value)
		b.WriteString(" ] FILLVALUE SORT\n")

	case ReducerPayload:
		b.WriteString(" MARK\n")
		for _, l := range p.Labels {
			b.WriteString(l)
			b.WriteString(" ")
		}

		b.WriteString(" COUNTTOMARK ->LIST SWAP DROP 'equivalenceClass' CSTORE\n")
		b.WriteString("[ SWAP $equivalenceClass ")
		if p.Value != "" {
			b.WriteString(p.Value)
			b.WriteString(" ")
		}
		b.WriteString("reducer.")
		b.WriteString(p.Reducer)
		b.WriteString(" ] REDUCE\n")

	case FunctionPayload:
		switch p.Name {
		case "abs":
			b.WriteString("[ SWAP mapper.abs 0 0 0 ] MAP\n")
		case "absent":
			b.WriteString("[ SWAP 0.0 mapper.replace 0 0 0 ] MAP [ NaN NaN NaN 1 ] FILLVALUE ")
		case "ceil":
			b.WriteString("UNBUCKETIZE [ SWAP mapper.ceil 0 0 0 ] MAP\n")
		case "changes":
			// COMPACT will dedup useless values, then we check if lasttick has the same value than penultimate (lasttick is forced by COMPACT), if yes we decrease the size by 1.
			b.WriteString("COMPACT MARK SWAP <% DUP DUP DUP NAME 'name' STORE LABELS 'l' STORE LASTTICK 'lt' STORE\n")
			b.WriteString("VALUES 'list' STORE $list SIZE 's' STORE $list $s 1 - GET $list $s 2 - GET\n")
			b.WriteString("<%  ==  %> <% $s 1 - %> <% $s %> IFTE 'val' STORE NEWGTS $name RENAME $l RELABEL $lt NaN DUP DUP $val SETVALUE \n")
			b.WriteString("%> FOREACH COUNTTOMARK ->LIST SWAP DROP\n")
		case "clamp_max":
			b.WriteString("UNBUCKETIZE [ SWAP " + p.Args[0] + fixScalar() + " mapper.min.x 0 0 0 ] MAP\n")
		case "clamp_min":
			b.WriteString("UNBUCKETIZE [ SWAP " + p.Args[0] + fixScalar() + " mapper.max.x 0 0 0 ] MAP\n")
		case "count_scalar":
			b.WriteString("[ SWAP [ ] reducer.count ] REDUCE [ 0.0 0.0 0 0 ] FILLVALUE\n")
		case "day_of_month":
			b.WriteString("[ SWAP 'UTC' mapper.day 0 0 0 ] MAP\n")
		case "day_of_week":
			b.WriteString("[ SWAP 'UTC' mapper.weekday 0 0 0 ] MAP\n")
			b.WriteString("[ SWAP 7 mapper.mod 0 0 0 ] MAP\n")
		case "days_in_month":
			b.WriteString(warpIsFebruary + "\n" + warpMacroBissex + "\n" + warpMacroDayInMonth + "\n")
			b.WriteString("[ SWAP <%  'mapping_window' STORE  $mapping_window 0 GET  'tick' STORE  $tick TSELEMENTS DUP 0 GET 'year' STORE 1 GET 'month' STORE\n")
			b.WriteString("$month $year @DAYSINMONTH  'days' STORE $tick NaN NaN NaN $days %> MACROMAPPER 0 0 0 ] MAP\n")
		case "delta":
			b.WriteString("[ SWAP mapper.delta $step $range MAX -1 * 0 $bucketCount 1 - -1 * ] MAP\n")
		case "deriv":
			b.WriteString("'deriv method is not supported' MSGFAIL")
			// FIXME
		case "drop_common_labels":
			b.WriteString("DUP [ SWAP [ ] reducer.count ] REDUCE 0 GET LABELS KEYLIST MARK SWAP <% '' %> FOREACH COUNTTOMARK ->MAP SWAP DROP RELABEL\n")
		case "exp":
			b.WriteString(NewSimpleMacroMapper("EXP"))
		case "floor":
			b.WriteString("UNBUCKETIZE [ SWAP mapper.floor 0 0 0 ] MAP\n")
		case "histogram_quantile":
			b.WriteString(p.Args[0] + fixScalar() + " 'QUANTILE' STORE \n" + warpBucketQuantile + "\n" + warpReducerHistogram + "\n")
			b.WriteString("<% 'equivalenceClass' DEFINED ! %> <%  [ ] 'equivalenceClass' STORE %> IFT \n")
			b.WriteString("[ SWAP  [ 'le' ] ->SET $equivalenceClass ->SET SWAP DIFFERENCE SET-> $reducer.histogram MACROREDUCER ] REDUCE\n")
		case "holt_winters":
			b.WriteString(" " + p.Args[0] + fixScalar() + " " + p.Args[1] + fixScalar() + "DOUBLEEXPONENTIALSMOOTHING 0 GET\n")
			b.WriteString("[ SWAP [] op.add ] APPLY\n")
		case "hour":
			b.WriteString("[ SWAP 'UTC' mapper.hour 0 0 0 ] MAP\n")
		case "idelta":
			b.WriteString("[ SWAP mapper.delta 1 0 $bucketCount 1 - -1 * ] MAP\n")
		case "increase":
			b.WriteString("FALSE RESETS\n")
			b.WriteString("[ SWAP mapper.delta $step $range MAX -1 * 0 $bucketCount 1 - -1 * ] MAP\n")
		case "irate":
			b.WriteString("[ SWAP mapper.rate 1 0 $bucketCount 1 - -1 * ] MAP\n")
		case "label_join":
			b.WriteString("<% DROP DUP LABELS 'labels' STORE ")
			b.WriteString("[ ")
			for index, item := range p.Args {
				if index >= 2 {
					b.WriteString(item + fixScalar())
				}
			}
			b.WriteString("] UNIQUE ")
			b.WriteString("<% DROP $labels SWAP GET %> LMAP ")
			b.WriteString("$labels SWAP ")
			b.WriteString(p.Args[1] + fixScalar())
			b.WriteString("JOIN ")
			b.WriteString(p.Args[0] + fixScalar())
			b.WriteString("PUT RELABEL %> LMAP\n")
		case "label_replace":
			b.WriteString(p.Args[0] + fixScalar() + " 'new_label' STORE " + p.Args[1] + fixScalar() + " 'replacement' STORE " + p.Args[2] + fixScalar() + " 'src_label' STORE " + p.Args[3] + fixScalar() + " 'regex' STORE \n")
			b.WriteString("MARK SWAP <%  DUP DUP NAME 'c' STORE LABELS { '__name__' $c  '' '' } APPEND DUP  $src_label GET DUP \n")
			b.WriteString("<% ISNULL %>  <% DROP DROP %> <% DUP $regex MATCH SIZE \n")
			b.WriteString("<% 0  >  %>  <%   $regex $replacement REPLACE  $new_label <% DUP '__name__' == %> <% DROP SWAP DROP RENAME %> <% PUT RELABEL %> IFTE %>  <% DROP DROP %> IFTE\n")
			b.WriteString("%> IFTE %> FOREACH COUNTTOMARK ->LIST SWAP DROP\n")
		case "ln":
			b.WriteString("[ SWAP e mapper.log 0 0 0 ] MAP\n")
		case "log2":
			b.WriteString("[ SWAP 2.0 mapper.log 0 0 0 ] MAP\n")
		case "log10":
			b.WriteString("[ SWAP 10.0 mapper.log 0 0 0 ] MAP\n")

		// Predict_linear works as a mapper in Prom
		// Compute alpha and beta linear regression on a range value ([1m] as example),
		// Apply beta * time + alpha to each points, where time is current time + p seconds (p corresponding to a user param)
		case "predict_linear":

			// prepare Window mapper
			b.WriteString("[ SWAP <% 'mappingWindow' STORE $mappingWindow 0 GET 'tick' STORE $mappingWindow 3 GET [] [] [] $mappingWindow 7 GET MAKEGTS DUP <% VALUES SIZE 2 >= %> \n")

			// Compute predict_linear
			b.WriteString("<% LR 'beta' STORE 'alpha' STORE $tick NaN NaN NaN $alpha $tick " + p.Args[0] + fixScalar() + "+ $beta * + %> <% DROP $mappingWindow 0 GET NaN NaN NaN NULL %> IFTE %> MACROMAPPER $range $step / 0 $instant -1 * ] MAP\n")

		case "minute":
			b.WriteString("[ SWAP 'UTC' mapper.minute 0 0 0 ] MAP\n")
		case "month":
			b.WriteString("[ SWAP 'UTC' mapper.month 0 0 0 ] MAP\n")
		case "rate":
			b.WriteString("FALSE RESETS\n")
			b.WriteString("[ SWAP mapper.rate $step $range MAX -1 * 0 $bucketCount 1 - -1 * ] MAP\n")
		case "resets":
			b.WriteString("FALSE RESETS\n")
		case "round":
			// When setting to_nearest optionnal param: Divide series per it before rounding it to nearest integer and multiplicate it back by it
			if len(p.Args) == 1 {
				b.WriteString("UNBUCKETIZE [ SWAP 1.0 " + p.Args[0] + fixScalar() + " / TODOUBLE mapper.mul 0 0 0 ] MAP\n")
				b.WriteString("UNBUCKETIZE [ SWAP mapper.round 0 0 0 ] MAP\n")
				b.WriteString("UNBUCKETIZE [ SWAP " + p.Args[0] + fixScalar() + " TODOUBLE mapper.mul 0 0 0 ] MAP\n")
			}
			b.WriteString("UNBUCKETIZE [ SWAP mapper.round 0 0 0 ] MAP\n")
		case "scalar":
			b.WriteString("DUP SIZE <% 1 == %> <% VALUES 0 GET 0 GET %> <% DROP NaN %> IFTE\n")
			b.WriteString(" 'value' STORE [ $start $end ] [] [] [] [ $value DUP ] MAKEGTS 'scalar' RENAME\n")
			b.WriteString(" [ SWAP bucketizer.mean $end $step $instant ] BUCKETIZE INTERPOLATE SORT\n")
		case "sort":
			b.WriteString("<% [ SWAP bucketizer.mean 0 0 1 ] BUCKETIZE VALUES 0 GET 0 GET %> SORTBY\n")
		case "sort_desc":
			b.WriteString("<% [ SWAP bucketizer.mean 0 0 1 ] BUCKETIZE VALUES 0 GET 0 GET %> SORTBY REVERSE\n")
		case "sqrt":
			b.WriteString("[ SWAP mapper.sqrt 0 0 0 ] MAP\n")
		case "time":
			b.WriteString(" [ $start $end ] [] [] [] [ 1 DUP ] MAKEGTS 'scalar' RENAME\n")
			b.WriteString(" [ SWAP bucketizer.mean $end $step $instant ] BUCKETIZE INTERPOLATE SORT\n")
			b.WriteString(" [ SWAP mapper.tick 0 0 0 ] MAP [ SWAP 0.000001 mapper.mul 0 0 0 ] MAP \n")
		case "timestamp":
			b.WriteString("[ SWAP mapper.tick 0 0 0 ] MAP\n")
		case "vector":
			b.WriteString("'scalar' STORE  <% $scalar TYPEOF 'LIST' != %> <% [ $start $end ] [] [] [] [ $scalar ] MAKEGTS  'vector' RENAME [ SWAP bucketizer.mean $end $step $instant ] BUCKETIZE INTERPOLATE SORT %> <% $scalar <% DROP 'vector' RENAME %> LMAP %> IFTE\n")
		case "year":
			b.WriteString("[ SWAP 'UTC' mapper.year 0 0 0 ] MAP\n")
		}

	case AggregatePayload:
		convertAggregate(b, p)

	case BinaryExprPayload:
		leftNodeType := fmt.Sprintf("%T", n.Left.Payload)
		rightNodeType := fmt.Sprintf("%T", n.Right.Payload)

		convertBinaryExpr(b, p.Op, leftNodeType, rightNodeType)

	case NumberLiteralPayload:
		b.WriteString(fmt.Sprintf(" %s ", p.Value))

	default:
		panic(fmt.Sprintf("Type %T is not handled", n.Payload))
	}
}

var macroFilterAggregateFooter = `
RELABEL

[] 'equivalenceClass' STORE
<%
   DROP
   DUP
   LABELS KEYLIST
   $equivalenceClass
   APPEND UNIQUE 'equivalenceClass' STORE
%>
LMAP
` + "\n"

func getMacroFilterAggregate(labels []string) string {
	return printLabelsAsWarpScriptMaps(labels) + macroFilterAggregateFooter
}

var simpleSupportedAggregator = map[string]string{
	"sum":      "reducer.sum",
	"min":      "reducer.min",
	"max":      "reducer.max",
	"avg":      "reducer.mean.exclude-nulls",
	"stddev":   "reducer.sd",
	"stdvar":   "reducer.var",
	"count":    "reducer.count.exclude-nulls",
	"quantile": "reducer.percentile",
}

// convertAggregate is transforming a prom aggregation into a MC2
// nolint: gocyclo
func convertAggregate(b *bytes.Buffer, p AggregatePayload) {
	// Filtering using without
	if p.Without && len(p.Grouping) > 0 {
		b.WriteString(getMacroFilterAggregate(p.Grouping))
		b.WriteString("[ SWAP $equivalenceClass reducer.sum ] REDUCE\n")
	}

	if reducer, ok := simpleSupportedAggregator[p.Op]; ok {
		// Starting writing reduce
		if p.Op == "quantile" {
			b.WriteString("[ SWAP DROP $left [ ")
		} else {
			b.WriteString("[ SWAP [ ")
		}

		if !p.Without && len(p.Grouping) > 0 {
			b.WriteString(printLabelsAsWarpScriptList(p.Grouping))
		}

		if p.Without && len(p.Grouping) > 0 {
			b.WriteString("] DROP $equivalenceClass ")
		} else {
			b.WriteString("] DUP 'equivalenceClass' STORE ")
		}

		if p.Op == "quantile" {
			b.WriteString(" <% " + p.Param + " 0.0 < " + p.Param + " 1.0 > || %> <% 'quantile expects a number included between [0,1]' MSGFAIL %> IFT " + p.Param + " 100.0 * ")
		}

		b.WriteString(reducer)
		b.WriteString(" ] REDUCE\n")
		// Keep only labels in the equivalence class like does promQL
		b.WriteString("MARK SWAP  <%  DUP LABELS { } SWAP  <% 'v' STORE 'k' STORE <% $equivalenceClass $k CONTAINS %> <%  DROP { $k $v } APPEND  %> <% DROP %> IFTE %> FOREACH SWAP { NULL NULL } RELABEL SWAP RELABEL %> FOREACH COUNTTOMARK ->LIST SWAP DROP\n")
	} else {
		// Advanced reduction
		switch p.Op {
		case "count_values":
			b.WriteString(p.Param + " VALUESPLIT FLATTEN\n")
			b.WriteString("[ SWAP bucketizer.count 0 0 1 ] BUCKETIZE\n")
			b.WriteString("[ SWAP [ " + p.Param + " ] reducer.sum ] REDUCE\n")
			b.WriteString("[ SWAP bucketizer.sum 0 0 1 ] BUCKETIZE\n")
		case "topk":
			b.WriteString("<% [ SWAP bucketizer.mean 0 0 1 ] BUCKETIZE VALUES 0 GET 0 GET %> SORTBY REVERSE [ 0 " + p.Param + " 1 - ] SUBLIST\n")
		case "bottomk":
			b.WriteString("[ SWAP bucketizer.min 0 0 1 ] BUCKETIZE\n")
			b.WriteString("<% [ SWAP bucketizer.mean 0 0 1 ] BUCKETIZE VALUES 0 GET 0 GET %> SORTBY [ 0 " + p.Param + " 1 - ] SUBLIST\n")
		default:
			log.Errorf("Aggregator not supported: %s", p.Op)
		}
	}
}

func fixScalar() string {
	return " 'scalar' STORE <% $scalar TYPEOF 'LIST' == $scalar TYPEOF 'GTS' == || %> <% [ $scalar bucketizer.last 0 0 1 ] BUCKETIZE FLATTEN 0 GET VALUES 0 GET %> <% $scalar %> IFTE\n"
}

// binaryExprEquivalence is holding the right warpscript for the right type of operation
// for a binaryExpr
type binaryExprEquivalence struct {
	ScalarToScalar string // 12 /13
	VectorToScalar string // http_request /12
	ScalarToVector string // http_request /12
	VectorToVector string // GTS1 / GTS2
}

var binaryExprEquivalences = map[string]binaryExprEquivalence{
	"+": {
		ScalarToScalar: " + ",
		VectorToScalar: "[ SWAP $right TODOUBLE mapper.add 0 0 0 ] MAP\n",
		ScalarToVector: "[ SWAP $left TODOUBLE mapper.add 0 0 0 ] MAP\n",
		VectorToVector: "[ SWAP  DUP 0 GET @HASHLABELS SWAP 1 GET @HASHLABELS [ 'hash_945fa9bc3027d7025e3' ] op.add ]  APPLY NONEMPTY { 'hash_945fa9bc3027d7025e3' '' } RELABEL \n",
	},
	"-": {
		ScalarToScalar: " - ",
		VectorToScalar: "[ SWAP 0 $right TODOUBLE - mapper.add 0 0 0 ] MAP\n",
		ScalarToVector: "[ SWAP [ SWAP -1 mapper.mul 0 0 0 ] MAP $left TODOUBLE mapper.add 0 0 0 ] MAP\n",
		VectorToVector: "[ SWAP  DUP 0 GET @HASHLABELS '%2B.tosub' RENAME SWAP 1 GET @HASHLABELS [ 'hash_945fa9bc3027d7025e3' ] op.sub ]  APPLY NONEMPTY { 'hash_945fa9bc3027d7025e3' '' } RELABEL \n",
	},
	"*": {
		ScalarToScalar: " * ",
		VectorToScalar: "[ SWAP $right TODOUBLE mapper.mul 0 0 0 ] MAP\n",
		ScalarToVector: "[ SWAP $left TODOUBLE mapper.mul 0 0 0 ] MAP\n",
		VectorToVector: "[ SWAP  DUP 0 GET @HASHLABELS SWAP 1 GET @HASHLABELS [ 'hash_945fa9bc3027d7025e3' ] op.mul ]  APPLY NONEMPTY { 'hash_945fa9bc3027d7025e3' '' } RELABEL \n",
	},
	"/": {
		ScalarToScalar: " / ",
		VectorToScalar: "[ SWAP 1 $right TODOUBLE / mapper.mul 0 0 0 ] MAP\n",
		ScalarToVector: warpHashLabels + " [ SWAP @HASHLABELS DUP [ SWAP $left mapper.replace 0 0 0 ] MAP SWAP @HASHLABELS [ 'hash_945fa9bc3027d7025e3' ] op.div ]  APPLY NONEMPTY { 'hash_945fa9bc3027d7025e3' '' } RELABEL\n",
		VectorToVector: "[ SWAP  DUP 0 GET @HASHLABELS '%2B.todiv' RENAME SWAP 1 GET @HASHLABELS [ 'hash_945fa9bc3027d7025e3' ] op.div ]  APPLY NONEMPTY { 'hash_945fa9bc3027d7025e3' '' } RELABEL \n",
	},
	"%": {
		ScalarToScalar: " % ",
		VectorToScalar: NewSimpleMacroMapper("$right %"),
		ScalarToVector: NewSimpleMacroMapper("$left %"),
		VectorToVector: "'modulo across GTS not supported' MSGFAIL", // FIXME:
	},
	"^": {
		ScalarToScalar: " ** ",
		VectorToScalar: "[ SWAP $right TODOUBLE mapper.pow 0 0 0 ] MAP\n",
		ScalarToVector: "'pow with scalar across GTS not supported' MSGFAIL\n", // FIXME:
		VectorToVector: "'pow across GTS not supported' MSGFAIL\n",             // FIXME:
	},
	">": {
		ScalarToScalar: " > ",
		VectorToScalar: "[ SWAP $right DUP TYPEOF <% 'LONG' == %> <% TODOUBLE %> IFT mapper.gt 0 0 0 ] MAP\n NONEMPTY\n",
		ScalarToVector: "[ SWAP $left DUP TYPEOF <% 'LONG' == %> <% TODOUBLE %> IFT mapper.le 0 0 0 ] MAP\n NONEMPTY\n",
		VectorToVector: "DUP 0 GET SWAP <% DROP @HASHLABELS DUP APPEND %> LMAP \n DUP 0 GET 'init' STORE \n [ SWAP DUP 0 GET SWAP 1 GET [ 'hash_945fa9bc3027d7025e3' ] op.gt ]  APPLY NONEMPTY \n <% DROP [ [ ROT DUP LABELS 'hash_945fa9bc3027d7025e3' GET 'intHash' STORE  ]  [ $init [] { 'hash_945fa9bc3027d7025e3' $intHash } filter.bylabels ] FILTER [] op.mask ] APPLY %> LMAP \n FLATTEN NONEMPTY { 'hash_945fa9bc3027d7025e3' '' } RELABEL \n",
	},
	"<": {
		ScalarToScalar: " < ",
		VectorToScalar: "[ SWAP $right DUP TYPEOF <% 'LONG' == %> <% TODOUBLE %> IFT mapper.lt 0 0 0 ] MAP\n NONEMPTY\n",
		ScalarToVector: "[ SWAP $left DUP TYPEOF <% 'LONG' == %> <% TODOUBLE %> IFT mapper.ge 0 0 0 ] MAP\n NONEMPTY\n",
		VectorToVector: "DUP 0 GET SWAP <% DROP @HASHLABELS DUP APPEND %> LMAP \n DUP 0 GET 'init' STORE \n [ SWAP DUP 0 GET SWAP 1 GET [ 'hash_945fa9bc3027d7025e3' ] op.lt ]  APPLY NONEMPTY \n <% DROP [ [ ROT DUP LABELS 'hash_945fa9bc3027d7025e3' GET 'intHash' STORE  ]  [ $init [] { 'hash_945fa9bc3027d7025e3' $intHash } filter.bylabels ] FILTER [] op.mask ] APPLY %> LMAP \n FLATTEN NONEMPTY { 'hash_945fa9bc3027d7025e3' '' } RELABEL \n",
	},
	"==": {
		ScalarToScalar: " == ",
		VectorToScalar: "[ SWAP $right DUP TYPEOF <% 'LONG' == %> <% TODOUBLE %> IFT mapper.eq 0 0 0 ] MAP\n NONEMPTY\n",
		ScalarToVector: "[ SWAP $left DUP TYPEOF <% 'LONG' == %> <% TODOUBLE %> IFT mapper.eq 0 0 0 ] MAP\n NONEMPTY\n",
		VectorToVector: "DUP 0 GET SWAP <% DROP @HASHLABELS DUP APPEND %> LMAP \n DUP 0 GET 'init' STORE \n [ SWAP DUP 0 GET SWAP 1 GET [ 'hash_945fa9bc3027d7025e3' ] op.eq ]  APPLY NONEMPTY \n <% DROP [ [ ROT DUP LABELS 'hash_945fa9bc3027d7025e3' GET 'intHash' STORE  ]  [ $init [] { 'hash_945fa9bc3027d7025e3' $intHash } filter.bylabels ] FILTER [] op.mask ] APPLY %> LMAP \n FLATTEN NONEMPTY { 'hash_945fa9bc3027d7025e3' '' } RELABEL \n",
	},
	"!=": {
		ScalarToScalar: " != ",
		VectorToScalar: "[ SWAP $right DUP TYPEOF <% 'LONG' == %> <% TODOUBLE %> IFT mapper.ne 0 0 0 ] MAP\n NONEMPTY\n",
		ScalarToVector: "[ SWAP $left DUP TYPEOF <% 'LONG' == %> <% TODOUBLE %> IFT mapper.ne 0 0 0 ] MAP\n NONEMPTY\n",
		VectorToVector: "DUP 0 GET SWAP <% DROP @HASHLABELS DUP APPEND %> LMAP \n DUP 0 GET 'init' STORE \n [ SWAP DUP 0 GET SWAP 1 GET [ 'hash_945fa9bc3027d7025e3' ] op.ne ]  APPLY NONEMPTY \n <% DROP [ [ ROT DUP LABELS 'hash_945fa9bc3027d7025e3' GET 'intHash' STORE  ]  [ $init [] { 'hash_945fa9bc3027d7025e3' $intHash } filter.bylabels ] FILTER [] op.mask ] APPLY %> LMAP \n FLATTEN NONEMPTY { 'hash_945fa9bc3027d7025e3' '' } RELABEL \n",
	},
	">=": {
		ScalarToScalar: " >= ",
		VectorToScalar: "[ SWAP $right DUP TYPEOF <% 'LONG' == %> <% TODOUBLE %> IFT mapper.ge 0 0 0 ] MAP\n NONEMPTY\n",
		ScalarToVector: "[ SWAP $left DUP TYPEOF <% 'LONG' == %> <% TODOUBLE %> IFT mapper.lt 0 0 0 ] MAP\n NONEMPTY\n",
		VectorToVector: "DUP 0 GET SWAP <% DROP @HASHLABELS DUP APPEND %> LMAP \n DUP 0 GET 'init' STORE \n [ SWAP DUP 0 GET SWAP 1 GET [ 'hash_945fa9bc3027d7025e3' ] op.ge ]  APPLY NONEMPTY \n <% DROP [ [ ROT DUP LABELS 'hash_945fa9bc3027d7025e3' GET 'intHash' STORE  ]  [ $init [] { 'hash_945fa9bc3027d7025e3' $intHash } filter.bylabels ] FILTER [] op.mask ] APPLY %> LMAP \n FLATTEN NONEMPTY { 'hash_945fa9bc3027d7025e3' '' } RELABEL \n",
	},
	"<=": {
		ScalarToScalar: " <= ",
		VectorToScalar: "[ SWAP $right DUP TYPEOF <% 'LONG' == %> <% TODOUBLE %> IFT mapper.le 0 0 0 ] MAP\n NONEMPTY\n",
		ScalarToVector: "[ SWAP $left DUP TYPEOF <% 'LONG' == %> <% TODOUBLE %> IFT mapper.gt 0 0 0 ] MAP\n NONEMPTY\n",
		VectorToVector: "DUP 0 GET SWAP <% DROP @HASHLABELS DUP APPEND %> LMAP \n DUP 0 GET 'init' STORE \n [ SWAP DUP 0 GET SWAP 1 GET [ 'hash_945fa9bc3027d7025e3' ] op.le ]  APPLY NONEMPTY \n <% DROP [ [ ROT DUP LABELS 'hash_945fa9bc3027d7025e3' GET 'intHash' STORE  ]  [ $init [] { 'hash_945fa9bc3027d7025e3' $intHash } filter.bylabels ] FILTER [] op.mask ] APPLY %> LMAP \n FLATTEN NONEMPTY { 'hash_945fa9bc3027d7025e3' '' } RELABEL \n",
	},
	"and": {
		VectorToVector: `CLEAR [ $left DUP TYPEOF <% 'GTS' == %> <% [ SWAP ] %> IFT
[]
$right DUP TYPEOF <% 'LIST' == %> <% [ 0 GET ] %> IFT LABELS
filter.bylabels ] FILTER\n
`,
	},
	"or": {
		VectorToVector: `CLEAR
$left DUP TYPEOF 'LIST' <% == %> <% 0 GET %> IFT LABELS 'left_labels'  STORE
$right DUP TYPEOF 'LIST' <% == %> <% 0 GET %> IFT LABELS 'right_labels' STORE

$left_labels <% 'value' STORE 'key' STORE  $value $key $right_labels SWAP GET <% == %> <% $right_labels $key REMOVE DROP 'right_labels' STORE  %> IFT %> FOREACH
[ $right DUP TYPEOF <% 'GTS' == %> <% [ SWAP ] %> IFT [] $right_labels filter.bylabels ] FILTER
$left 2 ->LIST FLATTEN
`,
	},
	"unless": {
		VectorToVector: `CLEAR
$left DUP TYPEOF 'LIST' <% == %> <% 0 GET %> IFT LABELS 'left_labels'  STORE
$right DUP TYPEOF 'LIST' <% == %> <% 0 GET %> IFT LABELS 'right_labels' STORE

$left_labels <% 'value' STORE 'key' STORE  $value $key $right_labels SWAP GET <% == %> <% $right_labels $key REMOVE DROP 'right_labels' STORE  %> IFT %> FOREACH
[ $left $right 2 ->LIST FLATTEN [] $right_labels filter.bylabels ] FILTER
$left 2 ->LIST FLATTEN
`,
	},
}

// convertBinaryExpr is using binaryExprEquivalences to write the right warpscript according to the situation:
// Binary arithmetic operators are defined between scalar/scalar, vector/scalar, scalar/vector, and vector/vector value pairs.
func convertBinaryExpr(b *bytes.Buffer, op string, leftNodeType string, rightNodeType string) {

	switch {
	case strings.Contains(leftNodeType, "NumberLiteralPayload") && strings.Contains(rightNodeType, "NumberLiteralPayload"):
		b.WriteString(binaryExprEquivalences[op].ScalarToScalar)
		b.WriteString(" 'value' STORE [ $start $end ] [] [] [] [ $value DUP ] MAKEGTS 'scalar' RENAME\n")
		b.WriteString(" [ SWAP bucketizer.mean $end $step $instant ] BUCKETIZE INTERPOLATE SORT \n")
	case !strings.Contains(leftNodeType, "NumberLiteralPayload") && strings.Contains(rightNodeType, "NumberLiteralPayload"):
		b.WriteString(binaryExprEquivalences[op].VectorToScalar)
	case !strings.Contains(rightNodeType, "NumberLiteralPayload") && strings.Contains(leftNodeType, "NumberLiteralPayload"):
		b.WriteString(binaryExprEquivalences[op].ScalarToVector)
	case !strings.Contains(leftNodeType, "NumberLiteralPayload") && !strings.Contains(rightNodeType, "NumberLiteralPayload"):
		b.WriteString(warpHashLabels + "\n" + binaryExprEquivalences[op].VectorToVector)
	}
}

func printLabelsAsWarpScriptHash(labels map[string]string) string {
	s := "{"
	for key, value := range labels {
		s += " '" + key + "' "
		s += " '" + value + "' "
	}
	s += "} "
	return s
}

func printLabelsAsWarpScriptList(labels []string) string {
	if len(labels) == 0 {
		return ""
	}

	s := ""
	for _, label := range labels {
		s += "'"
		s += label
		s += "'"
		s += " "
	}
	return s
}

func printLabelsAsWarpScriptMaps(labels []string) string {
	if len(labels) == 0 {
		return ""
	}

	s := "{ "
	for _, label := range labels {
		s += fmt.Sprintf("'%s' '' ", label)
	}
	s += "}"
	return s
}

var simpleMacroMapperHeader = `[ SWAP
<%
  'mapping_window' STORE    //  Storing macro input information
  $mapping_window 0 GET         // Tick
  $mapping_window 4 GET 0 GET   // Latitude
  $mapping_window 5 GET 0 GET   // Longitude
  $mapping_window 6 GET 0 GET   // Elevation
  $mapping_window 7 GET 0 GET `

var simpleMacroMapperFooter = "\n%> MACROMAPPER 0 0 0 ] MAP\n"

var warpMacroBissex = `<% 'y' STORE
  <% $y 4 %  0 == $y 100 %  0 != && %>
  <% true %>
  <%  
      <% $y 400 %  0 == %>
      <% true %>
      <% false %> IFTE
  %> IFTE
%> 'BISSEX' STORE`

var warpMacroDayInMonth = `<% 'y' STORE 'm' STORE
  { 
  '1' '31'
  '2' '28'
  '3' '31'
  '4' '30'
  '5' '31'
  '6' '30'
  '7' '31'
  '8' '31'
  '9' '30'
  '10' '31'
  '11' '30'
  '12' '31'
  }
  $m TOSTRING GET 
  $y @BISSEX  <% $m @ISFEB  &&   %> <% 1  + %> IFT 
  TOLONG
%> 'DAYSINMONTH' STORE`

var warpIsFebruary = `<%
  <% 2 == %> <% true %> <% false %> IFTE
%> 'ISFEB' STORE`

var warpBucketQuantile = `
<%
    
    '([0-9]*\.[0]*)' MATCHER 'reg' STORE      // Is it an integer?
    {} 'newM' STORE
    <%
        'val' STORE 'key' STORE
        <% 
            $val ISNULL ! 
            $key ISNULL ! 
            &&
        %>
        <% $newM { $key $val } APPEND 'newM' STORE %>
        IFT
    %>
    FOREACH
    $newM DUP 'm' STORE
    {} 'reverseMap' STORE   // Will used to avoid E notation
    KEYLIST DUP
	<% 
		$reverseMap SWAP DUP TODOUBLE TOSTRING DUP $reg MATCH
		<% SIZE 0 != %> <% TODOUBLE TOLONG TOSTRING  %> IFT  // Deal with integer (to avoid having "2.0" values instead of "2")
		SWAP 2 ->MAP APPEND DROP
	%> FOREACH 
	<% DROP TODOUBLE  %> LMAP LSORT     // Sort numerical values in ascending order

	[ ] SWAP

	// Now we convert values to sort numbers in their string forms, then converting them back to string
	<%
		TOSTRING DUP $reg MATCH                             // Check if it's an integer
		<% SIZE 0 != %> <% TODOUBLE TOLONG TOSTRING  %> IFT  // If it's, cast back to Double, then to Long, then back to String
		$reverseMap SWAP GET +!                              // Append into list structure
	%> FOREACH  'sortedList' STORE

	// Ensure Monotonicity as described in https://github.com/prometheus/prometheus/blob/73dc96e7f543d55c1600098617390eaba31dd938/promql/quantile.go
    $sortedList 
    0 'max' STORE
    <% 
         DUP  $m SWAP GET  'val' STORE 'key' STORE
        <% $val $max > %>
        <% $val 'max' STORE %>
        <% $val $max < %>
        <% $m $val $key PUT DROP %>
        <%  %>
        2 SWITCH
    %> FOREACH 
    
    $sortedList DUP SIZE 'size' STORE 
    $size 1 -  GET 'lastKey' STORE
    $m $lastKey GET $QUANTILE * 'rank' STORE
    
    
    <% 
        $sortedList SWAP DUP 'pos' STORE GET $m SWAP GET                                // Retrieve bucket[i].count from m map
        <% $rank >= %> <% $pos BREAK %>  IFT                                            // Compare to rank, if >=, print position and exit
    %> 'searchFirst' STORE 
    
    0 $size 1 - $searchFirst FOR 'b' STORE                                              // Iterate over 0 to len(bucket)-1 applyting @searchFirst macro
    <% $size 1 - $b ==  %> <% $sortedList $size 1 - GET  RETURN %> IFT                  // Return according to PromQL definition
    <% 0 $b == $sortedList 0 GET TODOUBLE 0 <= && %> <% $sortedList 0 GET RETURN %> IFT
    
	0 'bucketStart' STORE
    $sortedList $b GET 'bucketEnd' STORE
    $m $bucketEnd GET 'count'  STORE
    
    <% $b 0 > %>
    <%
        $sortedList $b 1 - GET 'bucketStart' STORE
        $m $bucketStart GET TODOUBLE 'bucketStartCount' STORE
        $count $bucketStartCount - 'count' STORE
        $rank $bucketStartCount - 'rank' STORE
    %> IFT 
    
    $bucketEnd TODOUBLE $bucketStart TODOUBLE -  $rank $count / * $bucketStart TODOUBLE +
    
%> 'bucketQuantile' STORE `

var warpReducerHistogram = `<% 
{ } 'metricsBucketMap' STORE // initialize map for buckets
'input' STORE
$input 1 GET  'list_class' STORE 
$input 2 GET DUP SIZE 1 - REMOVE DROP 'list_labels' STORE 
$input 7 GET  'list_values' STORE 
[ $list_class $list_labels $list_values ] ZIP // associate corresponding class, labels and value per GTS
<%
	'l' STORE
	//$l 0 GET 'this_class' STORE
	$l 1 GET 'this_labels' STORE
	$this_labels 'le' GET DUP
	<% '%2bInf' ==  %> <% DROP 'Infinity' %> IFT 'this_le' STORE // Replace golang '+Inf' with 'Infinity'
	$l 2 GET 'this_value' STORE
	$metricsBucketMap { $this_le $this_value } APPEND DROP
%> FOREACH  

$metricsBucketMap @bucketQuantile 'result' STORE
CLEAR

$input 0 GET NaN NaN NaN $result

%> 'reducer.histogram' STORE`

var warpHashLabels = `<% 
 [] 'l' STORE
 <% DUP LABELS ->JSON 'UTF-8' ->BYTES SHA1 ->HEX  'hash_945fa9bc3027d7025e3' SWAP 2 ->MAP RELABEL 1 ->LIST $l SWAP APPEND 'l' STORE %> FOREACH
 $l
%> 'HASHLABELS' STORE`

// NewSimpleMacroMapper is creating a macromapper that is only modifying the value.
func NewSimpleMacroMapper(op string) string {
	return simpleMacroMapperHeader + op + simpleMacroMapperFooter
}
