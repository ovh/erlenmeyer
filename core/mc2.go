package core

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

// ShouldRemoveNameLabel WarpScript attribute key to remove name label
const ShouldRemoveNameLabel = "SHOULD_REMOVE_NAME_LABEL"

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
	b.WriteString("'gts.limit.hard' STACKATTRIBUTE DUP <% ISNULL ! %> <% MAXGTS %> <% DROP %> IFTE\n")

	// Generate node WarpScript
	n.toWarpScript(&b)

	// Adding footer
	b.WriteString("\nDUP TYPEOF <% 'GTS' == %> <% [ SWAP ] %> IFT \n")

	b.WriteString("\nDUP TYPEOF <% 'LIST' != %> <% \n")
	b.WriteString("\t 'gts_values' STORE NEWGTS $start NaN NaN NaN $gts_values ADDVALUE $end NaN NaN NaN $gts_values ADDVALUE \n")
	b.WriteString("\t [ SWAP bucketizer.last $end $step 0 ] BUCKETIZE FILLPREVIOUS FILLNEXT \n")
	b.WriteString("\t { '" + ShouldRemoveNameLabel + "' 'true' } SETATTRIBUTES \n")
	b.WriteString("%> IFT \n")
	b.WriteString("\nNONEMPTY SORT UNBUCKETIZE [ SWAP mapper.tostring 0 0 0 ] MAP\n")

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

		if len(n.ChildLabels) == 0 {
			b.WriteString(" [] 'child_labels' STORE \n")
		} else {
			b.WriteString(" [ '" + strings.Join(n.ChildLabels, "' '") + "' ] 'child_labels' STORE \n")
		}

		n.Write(b)
	} else {
		if n.Left != nil {
			n.Left.toWarpScript(b)
		}
		if len(n.ChildLabels) == 0 {
			b.WriteString(" [] 'child_labels' STORE \n")
		} else {
			b.WriteString(" [ '" + strings.Join(n.ChildLabels, "' '") + "' ] 'child_labels' STORE \n")
		}
		n.Write(b)
	}
}

// IsValid check if a label name string is Valid
func IsValid(labelName string) bool {
	if len(labelName) == 0 {
		return false
	}
	for i, b := range labelName {
		if !((b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || b == '_' || b == '.' || (b >= '0' && b <= '9' && i > 0)) {
			return false
		}
	}
	return true
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

			if p.Instant {
				b.WriteString(" " + p.Start + " " + p.End + " ")
			} else {
				if p.Absent {
					b.WriteString(" " + p.Start + " 15 m - ISO8601")
				} else {
					b.WriteString(" " + p.Start + " ISO8601")
				}
				b.WriteString(" " + p.End + " ISO8601")
			}
		}

		b.WriteString(" ] FETCH \n")
		b.WriteString("false 'empty' STORE \n")
		b.WriteString("DUP <% VALUES SIZE 0 == %> <% NEWGTS '" + p.ClassName + "' RENAME " + printLabelsAsWarpScriptHash(p.Labels) + " RELABEL 1 ->LIST APPEND true 'empty' STORE %> IFT\n")

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

		if p.BucketCount == "1" && p.BucketSpan == "0" {
			b.WriteString("<% $empty ! %>\n")
			b.WriteString("<%\n")
		}

		if p.Absent {
			b.WriteString(p.BucketRange + " [ SWAP " + p.Op + " " + p.LastBucket + " " + p.BucketSpan + " " + p.BucketCount + " 2 - 15 m " + p.BucketSpan + " / + ] BUCKETIZE\n")
			b.WriteString("[ SWAP mapper.last 15 m $step / 0 $instant -1 * ] MAP\n")
			b.WriteString(p.BucketRange + " [ SWAP " + p.Op + " " + p.LastBucket + " " + p.BucketSpan + " " + p.BucketCount + " 2 - ] BUCKETIZE\n")
		} else {
			b.WriteString(p.PreBucketize + "\n")
			b.WriteString(p.BucketRange + " [ SWAP " + p.Op + " " + p.LastBucket + " " + p.BucketSpan + " " + p.BucketCount + " 2 - ] BUCKETIZE\n")
			b.WriteString(p.Filler + "\n")
		}

		if p.BucketCount == "1" && p.BucketSpan == "0" {
			b.WriteString("%>\n")
			b.WriteString("IFT\n")
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
		b.WriteString(p.Suffix)

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
			b.WriteString("[ SWAP mapper.abs 0 0 0 ] MAP { '" + ShouldRemoveNameLabel + "' 'true' } SETATTRIBUTES\n")
		case "absent":
			b.WriteString("[ SWAP 0.0 mapper.replace 0 0 0 ] MAP [ NaN NaN NaN 1 ] FILLVALUE ")
		case "ceil":
			b.WriteString("UNBUCKETIZE [ SWAP mapper.ceil 0 0 0 ] MAP { '" + ShouldRemoveNameLabel + "' 'true' } SETATTRIBUTES\n")
		case "changes":
			// COMPACT will dedup useless values, then we check if lasttick has the same value than penultimate (lasttick is forced by COMPACT), if yes we decrease the size by 1.
			b.WriteString("COMPACT MARK SWAP <% DUP DUP DUP NAME 'name' STORE LABELS 'l' STORE LASTTICK 'lt' STORE\n")
			b.WriteString("VALUES 'list' STORE $list SIZE 's' STORE $list $s 1 - GET $list $s 2 - GET\n")
			b.WriteString("<%  ==  %> <% $s 1 - %> <% $s %> IFTE 'val' STORE NEWGTS $name RENAME $l RELABEL $lt NaN DUP DUP $val SETVALUE \n")
			b.WriteString("%> FOREACH COUNTTOMARK ->LIST SWAP DROP\n")
		case "clamp_max":
			b.WriteString("UNBUCKETIZE [ SWAP " + p.Args[0] + fixScalar() + " mapper.min.x 0 0 0 ] MAP\n")
			b.WriteString("{ '" + ShouldRemoveNameLabel + "' 'true' } SETATTRIBUTES \n")
		case "clamp_min":
			b.WriteString("UNBUCKETIZE [ SWAP " + p.Args[0] + fixScalar() + " mapper.max.x 0 0 0 ] MAP\n")
			b.WriteString("{ '" + ShouldRemoveNameLabel + "' 'true' } SETATTRIBUTES \n")
		case "count_scalar":
			b.WriteString("[ SWAP [ ] reducer.count ] REDUCE [ 0.0 0.0 0 0 ] FILLVALUE\n")
		case "day_of_month":
			b.WriteString("DEPTH <% 0 == %> <% \n")
			b.WriteString("\t NEWGTS $start NaN NaN NaN $start ADDVALUE $end NaN NaN NaN $end ADDVALUE \n")
			b.WriteString("\t [ SWAP bucketizer.last $end $step 0 ] BUCKETIZE FILLPREVIOUS FILLNEXT \n")
			b.WriteString("[ SWAP 'UTC' mapper.day 0 0 0 ] MAP \n")
			b.WriteString("%> <% \n")
			b.WriteString(`
			[ 
				SWAP 
				<% 
					DUP 0 GET SWAP 
					7 GET 0 GET 1 s * TOLONG
					TSELEMENTS 
					2 GET 
					NaN SWAP NaN SWAP NaN SWAP
				%>
				MACROMAPPER
				0
				0
				0
			] MAP
			`)
			b.WriteString("%> IFTE \n")
			b.WriteString("{ '" + ShouldRemoveNameLabel + "' 'true' } SETATTRIBUTES \n")

		case "day_of_week":
			b.WriteString("DEPTH <% 0 == %> <% \n")
			b.WriteString("\t NEWGTS $start NaN NaN NaN $start ADDVALUE $end NaN NaN NaN $end ADDVALUE \n")
			b.WriteString("\t [ SWAP bucketizer.last $end $step 0 ] BUCKETIZE FILLPREVIOUS FILLNEXT \n")
			b.WriteString("[ SWAP 'UTC' mapper.weekday 0 0 0 ] MAP\n")
			b.WriteString("[ SWAP 7 mapper.mod 0 0 0 ] MAP\n")
			b.WriteString("%> <% \n")
			b.WriteString(`
			[ 
				SWAP 
				<% 
					DUP 0 GET SWAP 
					7 GET 0 GET 1 s * TOLONG
					TSELEMENTS 
					8 GET 
					NaN SWAP NaN SWAP NaN SWAP
				%>
				MACROMAPPER
				0
				0
				0
			] MAP
			`)
			b.WriteString("%> IFTE \n")
			b.WriteString("{ '" + ShouldRemoveNameLabel + "' 'true' } SETATTRIBUTES\n")
		case "days_in_month":
			b.WriteString("DEPTH <% 0 == %> <% \n")
			b.WriteString("\t NEWGTS $start NaN NaN NaN $start ADDVALUE $end NaN NaN NaN $end ADDVALUE \n")
			b.WriteString("\t [ SWAP bucketizer.last $end $step 0 ] BUCKETIZE FILLPREVIOUS FILLNEXT \n")
			b.WriteString("%> IFT \n")
			b.WriteString(warpIsFebruary + "\n" + warpMacroBissex + "\n" + warpMacroDayInMonth + "\n")
			b.WriteString("[ SWAP <%  'mapping_window' STORE  $mapping_window 0 GET  'tick' STORE  $tick TSELEMENTS DUP 0 GET 'year' STORE 1 GET 'month' STORE\n")
			b.WriteString("$month $year @DAYSINMONTH  'days' STORE $tick NaN NaN NaN $days %> MACROMAPPER 0 0 0 ] MAP { '" + ShouldRemoveNameLabel + "' 'true' } SETATTRIBUTES\n")
		case "delta":
			b.WriteString("[ SWAP mapper.delta $step $range MAX -1 * 0 $bucketCount 1 - -1 * ] MAP\n")
		case "deriv":
			b.WriteString("'deriv method is not supported' MSGFAIL")
			// FIXME
		case "drop_common_labels":
			b.WriteString("DUP [ SWAP [ ] reducer.count ] REDUCE 0 GET LABELS KEYLIST MARK SWAP <% '' %> FOREACH COUNTTOMARK ->MAP SWAP DROP RELABEL\n")
		case "exp":
			b.WriteString(NewSimpleMacroMapper("EXP"))
			b.WriteString(" { '" + ShouldRemoveNameLabel + "' 'true' } SETATTRIBUTES \n")
		case "floor":
			b.WriteString("UNBUCKETIZE [ SWAP mapper.floor 0 0 0 ] MAP { '" + ShouldRemoveNameLabel + "' 'true' } SETATTRIBUTES\n")
		case "histogram_quantile":
			b.WriteString(p.Args[0] + fixScalar() + " 'QUANTILE' STORE \n" + warpBucketQuantile + "\n" + warpReducerHistogram + "\n")
			b.WriteString("<% 'equivalenceClass' DEFINED ! %> <%  [ ] 'equivalenceClass' STORE %> IFT \n")
			b.WriteString("[ SWAP  [ 'le' ] ->SET $equivalenceClass ->SET SWAP DIFFERENCE SET-> $reducer.histogram MACROREDUCER ] REDUCE\n")
			b.WriteString(" { '" + ShouldRemoveNameLabel + "' 'true' } SETATTRIBUTES \n")
		case "holt_winters":
			b.WriteString(" " + p.Args[0] + fixScalar() + " " + p.Args[1] + fixScalar() + "DOUBLEEXPONENTIALSMOOTHING 0 GET\n")
			b.WriteString("[ SWAP [] op.add ] APPLY\n")
		case "hour":
			b.WriteString("DEPTH <% 0 == %> <% \n")
			b.WriteString("\t NEWGTS $start NaN NaN NaN $start ADDVALUE $end NaN NaN NaN $end ADDVALUE \n")
			b.WriteString("\t [ SWAP bucketizer.last $end $step 0 ] BUCKETIZE FILLPREVIOUS FILLNEXT \n")
			b.WriteString("[ SWAP 'UTC' mapper.hour 0 0 0 ] MAP\n")
			b.WriteString("%> <% \n")
			b.WriteString(`
			[ 
				SWAP 
				<% 
					DUP 0 GET SWAP 
					7 GET 0 GET 1 s * TOLONG
					TSELEMENTS 
					3 GET 
					NaN SWAP NaN SWAP NaN SWAP
				%>
				MACROMAPPER
				0
				0
				0
			] MAP
			`)
			b.WriteString("%> IFTE \n")
			b.WriteString("{ '" + ShouldRemoveNameLabel + "' 'true' } SETATTRIBUTES\n")
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
			log.Warn(p.Args[0])
			promLabel := strings.Trim(p.Args[0], " [] 'child_labels' STORE \n")
			if !IsValid(strings.Trim(strings.TrimSpace(promLabel), "\"")) {
				b.WriteString("'invalid destination label name in label_replace(): ' " + p.Args[0] + " + MSGFAIL\n")
			}
			b.WriteString(p.Args[0] + " 'new_label' STORE " + p.Args[1] + " 'replacement' STORE " + p.Args[2] + " 'src_label' STORE " + p.Args[3] + " 'regex' STORE \n")
			b.WriteString("MARK SWAP <%  DUP DUP NAME 'c' STORE LABELS { '__name__' $c  '' '' } APPEND DUP  $src_label GET DUP \n")
			b.WriteString("<% ISNULL %>  \n")
			b.WriteString("<% DROP DUP $new_label GET DUP <% ISNULL %> <% DROP DROP %> <% <% 'source-value-' $regex 0 13  SUBSTRING != %> <% DROP '' %> IFT $regex $replacement REPLACE $new_label <% DUP '__name__' == %> <% DROP SWAP DROP RENAME %> <% PUT RELABEL %> IFTE %> IFTE %>  \n")
			b.WriteString("<% DUP $regex MATCH SIZE \n")
			b.WriteString("<% 0  >  %>  <%   $regex $replacement REPLACE  $new_label <% DUP '__name__' == %> <% DROP SWAP DROP RENAME %> <% PUT RELABEL %> IFTE %>  <% DROP DROP %> IFTE\n")
			b.WriteString("%> IFTE %> FOREACH COUNTTOMARK ->LIST SWAP DROP\n")
		case "ln":
			b.WriteString("[ SWAP e mapper.log 0 0 0 ] MAP { '" + ShouldRemoveNameLabel + "' 'true' } SETATTRIBUTES\n")
		case "log2":
			b.WriteString("[ SWAP 2.0 mapper.log 0 0 0 ] MAP { '" + ShouldRemoveNameLabel + "' 'true' } SETATTRIBUTES\n")
		case "log10":
			b.WriteString("[ SWAP 10.0 mapper.log 0 0 0 ] MAP { '" + ShouldRemoveNameLabel + "' 'true' } SETATTRIBUTES\n")

		// Predict_linear works as a mapper in Prom
		// Compute alpha and beta linear regression on a range value ([1m] as example),
		// Apply beta * time + alpha to each points, where time is current time + p seconds (p corresponding to a user param)
		case "predict_linear":

			// prepare Window mapper
			b.WriteString("[ SWAP <% 'mappingWindow' STORE $mappingWindow 0 GET 'tick' STORE $mappingWindow 3 GET [] [] [] $mappingWindow 7 GET MAKEGTS DUP <% VALUES SIZE 2 >= %> \n")

			// Compute predict_linear
			b.WriteString("<% LR 'beta' STORE 'alpha' STORE $tick NaN NaN NaN $alpha $tick " + p.Args[0] + fixScalar() + "+ $beta * + %> <% DROP $mappingWindow 0 GET NaN NaN NaN NULL %> IFTE %> MACROMAPPER $range $step / 0 $instant -1 * ] MAP\n")

		case "minute":
			b.WriteString("DEPTH <% 0 == %> <% \n")
			b.WriteString("\t NEWGTS $start NaN NaN NaN $start ADDVALUE $end NaN NaN NaN $end ADDVALUE \n")
			b.WriteString("\t [ SWAP bucketizer.last $end $step 0 ] BUCKETIZE FILLPREVIOUS FILLNEXT \n")
			b.WriteString("\t [ SWAP 'UTC' mapper.minute 0 0 0 ] MAP \n")
			b.WriteString("%> <% \n")
			b.WriteString(`
			[ 
				SWAP 
				<% 
					DUP 0 GET SWAP 
					7 GET 0 GET 1 s * TOLONG
					TSELEMENTS 
					4 GET 
					NaN SWAP NaN SWAP NaN SWAP
				%>
				MACROMAPPER
				0
				0
				0
			] MAP
			`)

			b.WriteString("%> IFTE \n")
			b.WriteString("{ '" + ShouldRemoveNameLabel + "' 'true' } SETATTRIBUTES \n")

		case "month":
			b.WriteString("DEPTH <% 0 == %> <% \n")
			b.WriteString("\t NEWGTS $start NaN NaN NaN $start ADDVALUE $end NaN NaN NaN $end ADDVALUE \n")
			b.WriteString("\t [ SWAP bucketizer.last $end $step 0 ] BUCKETIZE FILLPREVIOUS FILLNEXT \n")
			b.WriteString("[ SWAP 'UTC' mapper.month 0 0 0 ] MAP \n")
			b.WriteString("%> <% \n")
			b.WriteString(`
			[ 
				SWAP 
				<% 
					DUP 0 GET SWAP 
					7 GET 0 GET 1 s * TOLONG
					TSELEMENTS 
					1 GET 
					NaN SWAP NaN SWAP NaN SWAP
				%>
				MACROMAPPER
				0
				0
				0
			] MAP
			`)
			b.WriteString("%> IFTE \n")
			b.WriteString("{ '" + ShouldRemoveNameLabel + "' 'true' } SETATTRIBUTES\n")
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
			b.WriteString("UNBUCKETIZE [ SWAP mapper.round 0 0 0 ] MAP { '" + ShouldRemoveNameLabel + "' 'true' } SETATTRIBUTES\n")
		case "scalar":
			b.WriteString("DUP SIZE <% 1 == %> <% VALUES 0 GET 0 GET %> <% DROP NaN %> IFTE\n")
			b.WriteString(" 'value' STORE [ $start $end ] [] [] [] [ $value DUP ] MAKEGTS 'scalar' RENAME { '" + ShouldRemoveNameLabel + "' 'true' } SETATTRIBUTES\n")
			b.WriteString(" [ SWAP bucketizer.mean $end $step $instant ] BUCKETIZE INTERPOLATE SORT\n")
		case "sort":
			b.WriteString("<% [ SWAP bucketizer.mean 0 0 1 ] BUCKETIZE VALUES 0 GET 0 GET %> SORTBY\n")
		case "sort_desc":
			b.WriteString("<% [ SWAP bucketizer.mean 0 0 1 ] BUCKETIZE VALUES 0 GET 0 GET %> SORTBY REVERSE\n")
		case "sqrt":
			b.WriteString("[ SWAP mapper.sqrt 0 0 0 ] MAP { '" + ShouldRemoveNameLabel + "' 'true' } SETATTRIBUTES\n")
		case "time":
			b.WriteString(" [ $start $end ] [] [] [] [ 1 DUP ]  MAKEGTS 'scalar' RENAME { '" + ShouldRemoveNameLabel + "' 'true' } SETATTRIBUTES\n")
			b.WriteString(" [ SWAP bucketizer.mean $end $step $instant ] BUCKETIZE INTERPOLATE SORT\n")
			b.WriteString(" [ SWAP mapper.tick 0 0 0 ] MAP [ SWAP 0.000001 mapper.mul 0 0 0 ] MAP \n")
		case "timestamp":
			b.WriteString(" UNBUCKETIZE [ SWAP mapper.tick 0 0 0 ] MAP [ SWAP 0.000001 mapper.mul 0 0 0 ] MAP { '" + ShouldRemoveNameLabel + "' 'true' } SETATTRIBUTES\n")
		case "vector":
			b.WriteString("'scalar' STORE  <% $scalar TYPEOF 'LIST' != %> <% [ $start $end ] [] [] [] [ $scalar ] MAKEGTS  'vector' RENAME [ SWAP bucketizer.mean $end $step $instant ] BUCKETIZE INTERPOLATE SORT %> <% $scalar <% DROP 'vector' RENAME %> LMAP %> IFTE { '" + ShouldRemoveNameLabel + "' 'true' } SETATTRIBUTES\n")
		case "year":
			b.WriteString("DEPTH <% 0 == %> <% \n")
			b.WriteString("\t NEWGTS $start NaN NaN NaN $start ADDVALUE $end NaN NaN NaN $end ADDVALUE \n")
			b.WriteString("\t [ SWAP bucketizer.last $end $step 0 ] BUCKETIZE FILLPREVIOUS FILLNEXT \n")
			b.WriteString("[ SWAP 'UTC' mapper.year 0 0 0 ] MAP \n")
			b.WriteString("%> <% \n")
			b.WriteString(`
			[ 
				SWAP 
				<% 
					DUP 0 GET SWAP 
					7 GET 0 GET 1 s * TOLONG
					TSELEMENTS 
					0 GET 
					NaN SWAP NaN SWAP NaN SWAP
				%>
				MACROMAPPER
				0
				0
				0
			] MAP
			`)
			b.WriteString("%> IFTE \n")
			b.WriteString("{ '" + ShouldRemoveNameLabel + "' 'true' } SETATTRIBUTES\n")
		}

	case AggregatePayload:
		convertAggregate(b, p)

	case BinaryExprPayload:
		leftNodeType := fmt.Sprintf("%T", n.Left.Payload)
		rightNodeType := fmt.Sprintf("%T", n.Right.Payload)

		if len(p.FilteredLabels) == 0 {
			b.WriteString(" [] 'ignoringLabels' STORE \n")
			b.WriteString("[ 'hash_945fa9bc3027d7025e3' ] 'hashlabel' STORE \n")

		} else if p.IsOn {
			b.WriteString(" [] 'ignoringLabels' STORE \n")
			b.WriteString(" [ '" + strings.Join(p.FilteredLabels, "' '") + "' ] 'hashlabel' STORE \n")
		} else if p.IsIgnoring {
			b.WriteString(" [ '" + strings.Join(p.FilteredLabels, "' '") + "' ] 'ignoringLabels' STORE \n")
			b.WriteString("[ 'hash_945fa9bc3027d7025e3' ] 'hashlabel' STORE \n")
		}

		if len(p.IncludeLabels) == 0 {
			b.WriteString(" [] 'include_labels' STORE \n")
		} else {
			b.WriteString(" [ '" + strings.Join(p.IncludeLabels, "' '") + "' ] 'include_labels' STORE \n")
		}
		if p.Card == "many-to-one" {
			// " [ 'many-to-one' ] 'joinedLabels' STORE \n"
		} else if p.Card == "one-to-many" {
			b.WriteString(" [ 'one-to-many' ] 'joinedLabels' STORE \n")
		} else {
			b.WriteString(" [] 'joinedLabels' STORE \n")
		}
		if p.ReturnBool {
			b.WriteString("true 'return_bool' STORE \n")
		} else {
			b.WriteString("false 'return_bool' STORE \n")
		}
		convertBinaryExpr(b, p.Op, leftNodeType, rightNodeType, p.Card)
		if p.ReturnBool {
			b.WriteString(" [ NaN NaN NaN 0 ] FILLVALUE UNBUCKETIZE [ SWAP mapper.toboolean 0 0 0 ] MAP [ SWAP mapper.todouble 0 0 0 ] MAP  \n")
			b.WriteString("{ '" + ShouldRemoveNameLabel + "' 'true' } SETATTRIBUTES \n")
		}
		b.WriteString("NONEMPTY \n")

	case NumberLiteralPayload:
		switch p.Value {
		case fmt.Sprintf("%v", math.Inf(1)):
			b.WriteString(" 1.0 0.0 / ")
		case fmt.Sprintf("%v", math.Inf(-1)):
			b.WriteString(" -1.0 0.0 / ")
		case "NaN":
			b.WriteString(" NaN ")
		default:
			if _, err := strconv.ParseFloat(p.Value, 64); err == nil {
				b.WriteString(fmt.Sprintf(" %s TODOUBLE ", p.Value))
			} else {
				b.WriteString(fmt.Sprintf(" %s ", p.Value))
			}
		}

	case UnaryExprPayload:
		switch p.Op {
		case "-":
			b.WriteString("\nDUP TYPEOF <% 'LIST' != %> <% \n")
			b.WriteString("\t 'gts_values' STORE NEWGTS $start NaN NaN NaN $gts_values ADDVALUE $end NaN NaN NaN $gts_values ADDVALUE \n")
			b.WriteString("\t [ SWAP bucketizer.last $end $step 0 ] BUCKETIZE FILLPREVIOUS FILLNEXT \n")
			b.WriteString("%> IFT \n")
			b.WriteString(" [ SWAP -1 mapper.mul 0 0 0 ] MAP\n")
			b.WriteString("\t { '" + ShouldRemoveNameLabel + "' 'true' } SETATTRIBUTES \n")
		}
	default:
		panic(fmt.Sprintf("Type %T is not handled", n.Payload))
	}
}

var macroFilterAggregateFooter = `
'without_map' STORE
$child_labels
<%
    'child_key' STORE
    <%
        $without_map $child_key CONTAINSKEY
    %>
    <%
		$child_key REMOVE DROP $without_keys $child_key + 'without_keys' STORE 'without_map' STORE
    %>
    <%
        DROP
    %>
    IFTE
%>
FOREACH
$without_map
RELABEL

[] 'equivalenceClass' STORE
<%
   DROP
   DUP
   LABELS 
   $without_keys
   <%
      REMOVE DROP
   %>
   FOREACH
   KEYLIST
   $equivalenceClass
   APPEND UNIQUE 'equivalenceClass' STORE
%>
LMAP
` + "\n"

func getMacroFilterAggregate(labels []string, suffix string) string {
	return printLabelsAsWarpScriptMaps(labels) + "\n" + suffix + "\n" + macroFilterAggregateFooter
}

var simpleSupportedAggregator = map[string]string{
	"sum":      "reducer.sum",
	"min":      "reducer.min",
	"max":      "reducer.max",
	"avg":      "reducer.mean.exclude-nulls",
	"stddev":   "false reducer.sd",
	"stdvar":   "false reducer.var",
	"count":    "reducer.count.exclude-nulls",
	"quantile": "reducer.percentile",
}

// convertAggregate is transforming a prom aggregation into a MC2
// nolint: gocyclo
func convertAggregate(b *bytes.Buffer, p AggregatePayload) {
	// Filtering using without
	b.WriteString("[] 'without_keys' STORE\n")
	if p.Without && len(p.Grouping) > 0 {
		suffix := ""
		if p.Op == "topk" || p.Op == "bottomk" {
			suffix = "DROP {} \n"
		}
		b.WriteString(getMacroFilterAggregate(p.Grouping, suffix))
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
		} else if p.Without && len(p.Grouping) == 0 {
			b.WriteString("] DROP NULL [] 'equivalenceClass' STORE ")
		} else {
			b.WriteString("] DUP $child_labels APPEND 'equivalenceClass' STORE ")
		}

		if p.Op == "quantile" {
			b.WriteString(" <% " + p.Param + " 0.0 < " + p.Param + " 1.0 > || %> <% 'quantile expects a number included between [0,1]' MSGFAIL %> IFT " + p.Param + " 100.0 * ")
		}

		b.WriteString(reducer)
		b.WriteString(" ] REDUCE\n")

		// When Without is set and grouping equals 0 skip labels refactoring
		if !(p.Without && len(p.Grouping) == 0) {
			// Keep also child nested needed labels
			b.WriteString("$equivalenceClass $without_keys APPEND 'equivalenceClass' STORE\n")
			// Keep only labels in the equivalence class like does promQL
			b.WriteString("MARK SWAP  <%  DUP LABELS { } SWAP  <% 'v' STORE 'k' STORE <% $equivalenceClass $k CONTAINS %> <%  DROP { $k $v } APPEND  %> <% DROP %> IFTE %> FOREACH SWAP { NULL NULL } RELABEL SWAP RELABEL %> FOREACH COUNTTOMARK ->LIST SWAP DROP\n")
		}
		// Set a Warp10 attribute to indicate that the series name should be removed
		b.WriteString("\t { '" + ShouldRemoveNameLabel + "' 'true' } SETATTRIBUTES \n")
	} else {
		// Advanced reduction
		switch p.Op {
		case "count_values":
			b.WriteString(p.Param + " VALUESPLIT FLATTEN\n")
			b.WriteString("[ SWAP bucketizer.count 0 0 1 ] BUCKETIZE\n")
			b.WriteString("[ SWAP [ " + p.Param + " ] reducer.sum ] REDUCE\n")
			b.WriteString("[ SWAP bucketizer.sum 0 0 1 ] BUCKETIZE\n")
		case "topk", "bottomk":
			sortSuffix := ""
			if p.Op == "topk" {
				sortSuffix = "REVERSE"
			}
			if len(p.Grouping) > 0 || (p.Without && len(p.Grouping) == 0) {
				if p.Without && len(p.Grouping) == 0 {
					b.WriteString("NULL PARTITION [] SWAP \n")
				} else if p.Without && len(p.Grouping) > 0 {
					b.WriteString("DUP <% DROP LABELS \n")
					for _, labelKey := range p.Grouping {
						b.WriteString("'" + labelKey + "' REMOVE DROP \n")
					}
					b.WriteString("KEYLIST %> LMAP FLATTEN UNIQUE \n")
					b.WriteString("PARTITION [] SWAP \n")
				} else {
					b.WriteString("[ " + printLabelsAsWarpScriptList(p.Grouping) + " ] PARTITION [] SWAP \n")
				}
				b.WriteString("<% \n")
				b.WriteString("\t SWAP DROP \n")
				b.WriteString("\t ")
			}
			b.WriteString("<% [ SWAP bucketizer.mean 0 0 1 ] BUCKETIZE VALUES 0 GET 0 GET %> SORTBY " + sortSuffix + " [ 0 " + p.Param + " 1 - ] SUBLIST \n")
			if len(p.Grouping) > 0 || (p.Without && len(p.Grouping) == 0) {

				b.WriteString("\t + \n")
				b.WriteString("%> FOREACH FLATTEN \n")
			}
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
	GroupLeft      string // GTS1 group_left / GTS2
	GroupRight     string // GTS1 group_left / GTS2 group_right
}

func groupOnPer(side, operator string, addLeftSuffix bool) string {
	reverse := "$right"
	opSeries := "$full_series $reverse_tmp"
	if addLeftSuffix {
		opSeries = "$full_series '%2B.suffix' RENAME $reverse_tmp"
	}

	applySuffix := ""
	switch operator {
	case "op.gt", "op.ge", "op.le", "op.lt", "op.eq", "op.ne":
		if side == "$left" {
			applySuffix = `
				'maskSeries' STORE
				[ 
					$maskSeries  
					$full_series 
					[] op.mask 
				] APPLY 
			`
		} else {
			applySuffix = `
				'maskSeries' STORE
				[ 
					$maskSeries  
					$reverse_tmp 
					NULL op.mask 
				] APPLY 
			`
		}
	}

	if side == "$right" {
		reverse = "$left"
		if addLeftSuffix {
			opSeries = "$reverse_tmp '%2B.suffix' RENAME $full_series"
		} else {
			opSeries = "$reverse_tmp $full_series"
		}
	}

	mc2 := "DROP " + side + ` @HASHLABELS
   <%
	   DROP
	   
	   DUP CLONEEMPTY 'empty_series' STORE
	   DUP 1 ->LIST 'full_series' STORE
	   LABELS 
	   $hashlabel SUBMAP 'submap_labels' STORE 
	   [ ` + reverse + ` @HASHLABELS [] $submap_labels filter.bylabels ] FILTER 'reverse_tmp' STORE
	   [ ` + opSeries + ` [] ` + operator + ` ]  APPLY ` + applySuffix + `
	   $include_labels PARTITION
	   [] 
       SWAP
	   <%
		    SWAP $include_labels SUBMAP 'sub_labels' STORE
            '' 'series_name' STORE 
            true 'same_name' STORE
            DUP 
            <% SIZE 0 > %>
			<% 
				DUP 
				0 GET NAME 'series_name' STORE  
				DUP
				<% 
				<%
					NAME $series_name !=
				%>
				<%
					false 'same_name' STORE
				%>
				IFT
				%>
				FOREACH
			%> IFT  
			$empty_series + REVERSE  MERGE $sub_labels RELABEL DEDUP
			<% $same_name $empty_series NAME $series_name == && ! %> 
			<%
				{ '` + ShouldRemoveNameLabel + `' 'true' } SETATTRIBUTES
			%> 
			IFT
		    +
	   %>
	   FOREACH
   %>
   LMAP
   FLATTEN { 'hash_945fa9bc3027d7025e3' '' } RELABEL
`
	return mc2
}

func getSingleOperatorScript(operator string) string {
	promString := ""
	switch operator {
	case "%":
		promString = "modulo"
	case "**":
		promString = "pow"
	}
	mc2 := `
	'inputs' STORE
	true 'shouldRemoveName' STORE 
	$inputs 0 GET SIZE 1 == 'scalarZeroInput' STORE
	$inputs 0 GET <% NAME 'scalar' == $scalarZeroInput && 'scalarZeroInput' STORE %> FOREACH
	$inputs 1 GET SIZE 1 == 'scalarOneInput' STORE
	$inputs 1 GET <% NAME 'scalar' == $scalarOneInput && 'scalarOneInput' STORE %> FOREACH
	<%
		$scalarZeroInput $scalarOneInput || 
	%>
	<% 
		<%
			$scalarZeroInput
		%>
		<%
			$inputs 1 GET 
			$inputs 0 GET 'raw' STORE
		%> 
		<%
			$inputs 0 GET 
			$inputs 1 GET 'raw' STORE
		%> 
		IFTE
		[] SWAP
		NULL PARTITION
		<%
			SWAP DROP DUP CLONEEMPTY 'metaSeries' STORE 'filteredSeries' STORE
			<%
				$scalarZeroInput
			%>
			<%
				$raw CLONE SWAP DROP $filteredSeries APPEND 
			%>
			<%
				$filteredSeries $raw CLONE SWAP DROP APPEND
			%> 
			IFTE
			
			[
              SWAP
              []
              <%
                DUP 0 GET 
                SWAP 
                7 GET 
                DUP 0 GET SWAP 1 GET ` + operator + `
                NaN SWAP
                NaN SWAP
                NaN SWAP
              %> 
              MACROREDUCER
            ] REDUCE
			$metaSeries SWAP + MERGE
			+
		%>
		FOREACH
		FLATTEN
	%>
	<%
        '` + promString + ` across GTS not supported' MSGFAIL
    %>
	IFTE
	`
	return mc2
}

func getComparatorScript(operator string) string {
	mc2 := `

	'inputs' STORE

	false 'shouldRemoveName' STORE 
	$inputs 0 GET SIZE 1 == 'scalarZeroInput' STORE
	$inputs 0 GET <% NAME 'scalar' == $scalarZeroInput && 'scalarZeroInput' STORE %> FOREACH
	$inputs 1 GET SIZE 1 == 'scalarOneInput' STORE
	$inputs 1 GET <% NAME 'scalar' == $scalarOneInput && 'scalarOneInput' STORE %> FOREACH
	<%
		$scalarZeroInput $scalarOneInput || 
	%>
	<% 
		<%
			$scalarZeroInput
		%>
		<%
			$inputs 1 GET 
			$inputs 0 GET 'raw' STORE
		%> 
		<%
			$inputs 0 GET 
			$inputs 1 GET 'raw' STORE
		%> 
		IFTE
		[] SWAP
		NULL PARTITION
		<%
			SWAP DROP DUP CLONEEMPTY 'metaSeries' STORE 'filteredSeries' STORE
			<%
				$scalarZeroInput
			%>
			<%
				[ $raw $filteredSeries [] ` + operator + ` ]  APPLY
				%>
			<%
				[ $filteredSeries $raw [] ` + operator + ` ]  APPLY
			%> 
			IFTE
			$metaSeries SWAP + MERGE
			'maskSeries' STORE
			[ 
					[ $maskSeries ]
					$filteredSeries 
					[] op.mask 
			] APPLY 
			+
		%>
		FOREACH
		FLATTEN
	%>
	<% 
		true 'shouldRemoveName' STORE 
		$inputs 
		<% 
			DROP 
			<%
				// Case were hashlabels contains on labels
				$hashlabel SIZE 1 == $hashlabel 'hash_945fa9bc3027d7025e3' CONTAINS SWAP DROP AND !
			%>
			<%
				$hashlabel
				PARTITION
				[]
				SWAP
				<%
					SWAP DROP 
					MERGE DEDUP
					+
				%>
				FOREACH
			%>
			IFT 
			@HASHLABELS 
			DUP 
			APPEND 
		%> LMAP 'inputs' STORE
		$inputs 0 GET 
		true 'skipInputZero' STORE
		DUP <% NAME 'scalar' == $skipInputZero && 'skipInputZero' STORE %> FOREACH
		<% $skipInputZero %> <% DROP $inputs 1 GET %> IFT
		'init' STORE 
		[ $inputs 0 GET $inputs 1 GET [ 'hash_945fa9bc3027d7025e3' ] ` + operator + ` ]  APPLY NONEMPTY { 'hash_945fa9bc3027d7025e3' '' } RELABEL @HASHLABELS
		<% 
			$return_bool !
		%>
		<%
			<% DROP 
				[ SWAP DUP LABELS 'hash_945fa9bc3027d7025e3' GET 'intHash' STORE  ] 'maskSeries' STORE 
				[ $init [] { 'hash_945fa9bc3027d7025e3' $intHash } filter.bylabels ] FILTER 'filterSeries' STORE
				<% $filterSeries SIZE 0 == %>
				<% $filterSeries %>
				<%
					[ 
						$maskSeries
						$filterSeries 
						[] op.mask 
					] APPLY 
				%> IFTE
			%> LMAP 
		%>
		IFT
		FLATTEN { 'hash_945fa9bc3027d7025e3' '' } RELABEL
	%>
	IFTE
	`
	return mc2
}

var binaryExprEquivalences = map[string]binaryExprEquivalence{
	"+": {
		ScalarToScalar: " + ",
		VectorToScalar: "[ SWAP $right TODOUBLE mapper.add 0 0 0 ] MAP\n",
		ScalarToVector: "[ SWAP $left TODOUBLE mapper.add 0 0 0 ] MAP\n",
		VectorToVector: "[ SWAP  DUP 0 GET @HASHLABELS SWAP 1 GET @HASHLABELS $hashlabel op.add ]  APPLY { 'hash_945fa9bc3027d7025e3' '' } RELABEL \n",
		GroupLeft:      groupOnPer("$left", "op.add", false),
		GroupRight:     groupOnPer("$right", "op.add", false),
	},
	"-": {
		ScalarToScalar: " - ",
		VectorToScalar: "[ SWAP 0 $right TODOUBLE - mapper.add 0 0 0 ] MAP\n",
		ScalarToVector: "[ SWAP [ SWAP -1 mapper.mul 0 0 0 ] MAP $left TODOUBLE mapper.add 0 0 0 ] MAP\n",
		VectorToVector: "[ SWAP  DUP 0 GET @HASHLABELS '%2B.tosub' RENAME SWAP 1 GET @HASHLABELS $hashlabel op.sub ]  APPLY { 'hash_945fa9bc3027d7025e3' '' } RELABEL \n",
		GroupLeft:      groupOnPer("$left", "op.sub", true),
		GroupRight:     groupOnPer("$right", "op.sub", true),
	},
	"*": {
		ScalarToScalar: " * ",
		VectorToScalar: "[ SWAP $right TODOUBLE mapper.mul 0 0 0 ] MAP\n",
		ScalarToVector: "[ SWAP $left TODOUBLE mapper.mul 0 0 0 ] MAP\n",
		VectorToVector: "[ SWAP  DUP 0 GET @HASHLABELS SWAP 1 GET @HASHLABELS $hashlabel op.mul ]  APPLY { 'hash_945fa9bc3027d7025e3' '' } RELABEL \n",
		GroupLeft:      groupOnPer("$left", "op.mul", false),
		GroupRight:     groupOnPer("$right", "op.mul", false),
	},
	"/": {
		ScalarToScalar: " / ",
		VectorToScalar: "[ SWAP 1 $right TODOUBLE / mapper.mul 0 0 0 ] MAP\n",
		ScalarToVector: warpHashLabels + " [ SWAP DUP [ SWAP $left mapper.replace 0 0 0 ] MAP @HASHLABELS  SWAP @HASHLABELS $hashlabel op.div ]  APPLY { 'hash_945fa9bc3027d7025e3' '' } RELABEL\n",
		VectorToVector: "[ SWAP  DUP 0 GET @HASHLABELS '%2B.todiv' RENAME SWAP 1 GET @HASHLABELS $hashlabel op.div ]  APPLY { 'hash_945fa9bc3027d7025e3' '' } RELABEL \n",
		GroupLeft:      groupOnPer("$left", "op.div", true),
		GroupRight:     groupOnPer("$right", "op.div", true),
	},
	"%": {
		ScalarToScalar: " % ",
		VectorToScalar: NewSimpleMacroMapper("TODOUBLE $right TODOUBLE %"),
		ScalarToVector: NewSimpleMacroMapper("TODOUBLE $left TODOUBLE SWAP %"),
		VectorToVector: getSingleOperatorScript("%"),                  // FIXME: still not supported between GTS
		GroupLeft:      "'modulo across GTS not supported' MSGFAIL\n", // FIXME:
		GroupRight:     "'modulo across GTS not supported' MSGFAIL\n", // FIXME:
	},
	"^": {
		ScalarToScalar: " ** ",
		VectorToScalar: "[ SWAP $right TODOUBLE mapper.pow 0 0 0 ] MAP\n",
		ScalarToVector: NewSimpleMacroMapper("TODOUBLE $left TODOUBLE SWAP **"),
		VectorToVector: getSingleOperatorScript("**"),              // FIXME: still not supported between GTS
		GroupLeft:      "'pow across GTS not supported' MSGFAIL\n", // FIXME:
		GroupRight:     "'pow across GTS not supported' MSGFAIL\n", // FIXME:
	},
	">": {
		ScalarToScalar: " > ",
		VectorToScalar: "[ SWAP $right DUP TYPEOF <% 'LONG' == %> <% TODOUBLE %> IFT mapper.gt 0 0 0 ] MAP\n",
		ScalarToVector: "[ SWAP $left DUP TYPEOF <% 'LONG' == %> <% TODOUBLE %> IFT mapper.le 0 0 0 ] MAP\n",
		VectorToVector: getComparatorScript("op.gt"),
		GroupLeft:      groupOnPer("$left", "op.gt", false),
		GroupRight:     groupOnPer("$right", "op.gt", false),
	},
	"<": {
		ScalarToScalar: " < ",
		VectorToScalar: "[ SWAP $right DUP TYPEOF <% 'LONG' == %> <% TODOUBLE %> IFT mapper.lt 0 0 0 ] MAP\n",
		ScalarToVector: "[ SWAP $left DUP TYPEOF <% 'LONG' == %> <% TODOUBLE %> IFT mapper.ge 0 0 0 ] MAP\n",
		VectorToVector: getComparatorScript("op.lt"),
		GroupLeft:      groupOnPer("$left", "op.lt", false),
		GroupRight:     groupOnPer("$right", "op.lt", false),
	},
	"==": {
		ScalarToScalar: " == ",
		VectorToScalar: "[ SWAP $right DUP TYPEOF <% 'LONG' == %> <% TODOUBLE %> IFT mapper.eq 0 0 0 ] MAP\n",
		ScalarToVector: "[ SWAP $left DUP TYPEOF <% 'LONG' == %> <% TODOUBLE %> IFT mapper.eq 0 0 0 ] MAP\n",
		VectorToVector: getComparatorScript("op.eq"),
		GroupLeft:      groupOnPer("$left", "op.eq", false),
		GroupRight:     groupOnPer("$right", "op.eq", false),
	},
	"!=": {
		ScalarToScalar: " != ",
		VectorToScalar: "[ SWAP $right DUP TYPEOF <% 'LONG' == %> <% TODOUBLE %> IFT mapper.ne 0 0 0 ] MAP\n",
		ScalarToVector: "[ SWAP $left DUP TYPEOF <% 'LONG' == %> <% TODOUBLE %> IFT mapper.ne 0 0 0 ] MAP\n",
		VectorToVector: getComparatorScript("op.ne"),
		GroupLeft:      groupOnPer("$left", "op.ne", false),
		GroupRight:     groupOnPer("$right", "op.ne", false),
	},
	">=": {
		ScalarToScalar: " >= ",
		VectorToScalar: "[ SWAP $right DUP TYPEOF <% 'LONG' == %> <% TODOUBLE %> IFT mapper.ge 0 0 0 ] MAP\n",
		ScalarToVector: "[ SWAP $left DUP TYPEOF <% 'LONG' == %> <% TODOUBLE %> IFT mapper.lt 0 0 0 ] MAP\n",
		VectorToVector: getComparatorScript("op.ge"),
		GroupLeft:      groupOnPer("$left", "op.ge", false),
		GroupRight:     groupOnPer("$right", "op.ge", false),
	},
	"<=": {
		ScalarToScalar: " <= ",
		VectorToScalar: "[ SWAP $right DUP TYPEOF <% 'LONG' == %> <% TODOUBLE %> IFT mapper.le 0 0 0 ] MAP\n",
		ScalarToVector: "[ SWAP $left DUP TYPEOF <% 'LONG' == %> <% TODOUBLE %> IFT mapper.gt 0 0 0 ] MAP\n",
		VectorToVector: getComparatorScript("op.le"),
		GroupLeft:      groupOnPer("$left", "op.le", false),
		GroupRight:     groupOnPer("$right", "op.le", false),
	},
	"and": {
		VectorToVector: `CLEAR [ $left DUP TYPEOF <% 'GTS' == %> <% [ SWAP ] %> IFT
[]
$right DUP TYPEOF <% 'LIST' == %> <% [ 0 GET ] %> IFT LABELS
filter.bylabels ] FILTER\n
`,
		GroupLeft: `CLEAR [ $left DUP TYPEOF <% 'GTS' == %> <% [ SWAP ] %> IFT
	[]
	$right DUP TYPEOF <% 'LIST' == %> <% [ 0 GET ] %> IFT LABELS
	filter.bylabels ] FILTER\n
	`, // FIXME: VectorToVector is used to not break script
		GroupRight: `CLEAR [ $left DUP TYPEOF <% 'GTS' == %> <% [ SWAP ] %> IFT
	[]
	$right DUP TYPEOF <% 'LIST' == %> <% [ 0 GET ] %> IFT LABELS
	filter.bylabels ] FILTER\n
	`, // FIXME: VectorToVector is used to not break script
	},
	"or": {
		VectorToVector: `CLEAR
$left DUP TYPEOF 'LIST' <% == %> <% 0 GET %> IFT LABELS 'left_labels'  STORE
$right DUP TYPEOF 'LIST' <% == %> <% 0 GET %> IFT LABELS 'right_labels' STORE

$left_labels <% 'value' STORE 'key' STORE  $value $key $right_labels SWAP GET <% == %> <% $right_labels $key REMOVE DROP 'right_labels' STORE  %> IFT %> FOREACH
[ $right DUP TYPEOF <% 'GTS' == %> <% [ SWAP ] %> IFT [] $right_labels filter.bylabels ] FILTER
$left 2 ->LIST FLATTEN
`,
		GroupLeft: `CLEAR
$left DUP TYPEOF 'LIST' <% == %> <% 0 GET %> IFT LABELS 'left_labels'  STORE
$right DUP TYPEOF 'LIST' <% == %> <% 0 GET %> IFT LABELS 'right_labels' STORE

$left_labels <% 'value' STORE 'key' STORE  $value $key $right_labels SWAP GET <% == %> <% $right_labels $key REMOVE DROP 'right_labels' STORE  %> IFT %> FOREACH
[ $right DUP TYPEOF <% 'GTS' == %> <% [ SWAP ] %> IFT [] $right_labels filter.bylabels ] FILTER
$left 2 ->LIST FLATTEN
`, // FIXME: VectorToVector is used to not break script
		GroupRight: `CLEAR
$left DUP TYPEOF 'LIST' <% == %> <% 0 GET %> IFT LABELS 'left_labels'  STORE
$right DUP TYPEOF 'LIST' <% == %> <% 0 GET %> IFT LABELS 'right_labels' STORE

$left_labels <% 'value' STORE 'key' STORE  $value $key $right_labels SWAP GET <% == %> <% $right_labels $key REMOVE DROP 'right_labels' STORE  %> IFT %> FOREACH
[ $right DUP TYPEOF <% 'GTS' == %> <% [ SWAP ] %> IFT [] $right_labels filter.bylabels ] FILTER
$left 2 ->LIST FLATTEN
`, // FIXME: VectorToVector is used to not break script
	},
	"unless": {
		VectorToVector: `CLEAR
$left DUP TYPEOF 'LIST' <% == %> <% 0 GET %> IFT LABELS 'left_labels'  STORE
$right DUP TYPEOF 'LIST' <% == %> <% 0 GET %> IFT LABELS 'right_labels' STORE

$left_labels <% 'value' STORE 'key' STORE  $value $key $right_labels SWAP GET <% == %> <% $right_labels $key REMOVE DROP 'right_labels' STORE  %> IFT %> FOREACH
[ $left $right 2 ->LIST FLATTEN [] $right_labels filter.bylabels ] FILTER
$left 2 ->LIST FLATTEN
`,
		GroupLeft: `CLEAR
$left DUP TYPEOF 'LIST' <% == %> <% 0 GET %> IFT LABELS 'left_labels'  STORE
$right DUP TYPEOF 'LIST' <% == %> <% 0 GET %> IFT LABELS 'right_labels' STORE

$left_labels <% 'value' STORE 'key' STORE  $value $key $right_labels SWAP GET <% == %> <% $right_labels $key REMOVE DROP 'right_labels' STORE  %> IFT %> FOREACH
[ $left $right 2 ->LIST FLATTEN [] $right_labels filter.bylabels ] FILTER
$left 2 ->LIST FLATTEN
`, // FIXME: VectorToVector is used to not break script
		GroupRight: `CLEAR
$left DUP TYPEOF 'LIST' <% == %> <% 0 GET %> IFT LABELS 'left_labels'  STORE
$right DUP TYPEOF 'LIST' <% == %> <% 0 GET %> IFT LABELS 'right_labels' STORE

$left_labels <% 'value' STORE 'key' STORE  $value $key $right_labels SWAP GET <% == %> <% $right_labels $key REMOVE DROP 'right_labels' STORE  %> IFT %> FOREACH
[ $left $right 2 ->LIST FLATTEN [] $right_labels filter.bylabels ] FILTER
$left 2 ->LIST FLATTEN
`, // FIXME: VectorToVector is used to not break script
	},
}

// convertBinaryExpr is using binaryExprEquivalences to write the right warpscript according to the situation:
// Binary arithmetic operators are defined between scalar/scalar, vector/scalar, scalar/vector, and vector/vector value pairs.
func convertBinaryExpr(b *bytes.Buffer, op string, leftNodeType string, rightNodeType string, card string) {

	switch {
	case strings.Contains(leftNodeType, "NumberLiteralPayload") && strings.Contains(rightNodeType, "NumberLiteralPayload"):
		b.WriteString(binaryExprEquivalences[op].ScalarToScalar)
		b.WriteString(" TODOUBLE 'value' STORE [ $start $end ] [] [] [] [ $value DUP ] MAKEGTS { '" + ShouldRemoveNameLabel + "' 'true' } SETATTRIBUTES 'scalar' RENAME\n")
		b.WriteString(" [ SWAP bucketizer.mean $end $step $instant ] BUCKETIZE INTERPOLATE SORT \n")
	case !strings.Contains(leftNodeType, "NumberLiteralPayload") && strings.Contains(rightNodeType, "NumberLiteralPayload"):
		b.WriteString(binaryExprEquivalences[op].VectorToScalar)
		switch op {
		case "!=", ">", ">=", "<=", "<", "==":
		default:
			b.WriteString("{ '" + ShouldRemoveNameLabel + "' 'true' } SETATTRIBUTES \n")
		}
	case !strings.Contains(rightNodeType, "NumberLiteralPayload") && strings.Contains(leftNodeType, "NumberLiteralPayload"):
		b.WriteString(binaryExprEquivalences[op].ScalarToVector)
		switch op {
		case "!=", ">", ">=", "<=", "<", "==":
		default:
			b.WriteString("{ '" + ShouldRemoveNameLabel + "' 'true' } SETATTRIBUTES \n")
		}
	case !strings.Contains(leftNodeType, "NumberLiteralPayload") && !strings.Contains(rightNodeType, "NumberLiteralPayload"):
		if card == "many-to-one" {
			b.WriteString(warpHashLabels + "\n" + binaryExprEquivalences[op].GroupLeft)
		} else if card == "one-to-many" {
			b.WriteString(warpHashLabels + "\n" + binaryExprEquivalences[op].GroupRight)
		} else {
			b.WriteString("true 'shouldRemoveName' STORE \n")
			b.WriteString("DUP <% <% <% ATTRIBUTES '" + ShouldRemoveNameLabel + "' CONTAINSKEY %> <% '" + ShouldRemoveNameLabel + "' GET TOBOOLEAN $shouldRemoveName && 'shouldRemoveName' STORE %> <% DROP %> IFTE %> FOREACH %> FOREACH \n")
			b.WriteString(warpHashLabels + "\n" + binaryExprEquivalences[op].VectorToVector)
			b.WriteString("<% $shouldRemoveName %> <% { '" + ShouldRemoveNameLabel + "' 'true' } SETATTRIBUTES  %> IFT \n")
		}
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
 []  'ignoringLabels' CSTORE
 [] 'l' STORE
 <% DUP LABELS $ignoringLabels <% REMOVE DROP %> FOREACH ->JSON 'UTF-8' ->BYTES SHA1 ->HEX  'hash_945fa9bc3027d7025e3' SWAP 2 ->MAP RELABEL 1 ->LIST $l SWAP APPEND 'l' STORE %> FOREACH
 $l
%> 'HASHLABELS' STORE`

// NewSimpleMacroMapper is creating a macromapper that is only modifying the value.
func NewSimpleMacroMapper(op string) string {
	return simpleMacroMapperHeader + op + simpleMacroMapperFooter
}
