package influxdb

import (
	"fmt"
	"strings"

	"github.com/influxdata/influxql"
	"github.com/ovh/erlenmeyer/core"
)

// parseTransformationsName Parse Influx Transformation methods:
// https://docs.influxdata.com/influxdb/v1.7/query_language/functions/#transformations
func (p *InfluxParser) parseTransformationsName(name string, args []string) (string, error) {
	mc2 := ""

	switch name {
	case "abs":
		mc2 = "[ SWAP mapper.abs 0 0 0 ] MAP\n"
	case "acos":
		mc2 = core.NewSimpleMacroMapper("ACOS")
	case "asin":
		mc2 = core.NewSimpleMacroMapper("ASIN")
	case "atan":
		mc2 = core.NewSimpleMacroMapper("ATAN")
	case "atan2":
		mc2 = `
		MERGE 'x' STORE MERGE 'y' STORE
		$x CLONEEMPTY
		$x TICKLIST ->SET
		$y TICKLIST ->SET
		INTERSECTION SET-> 
		'ticks' STORE
		$ticks
	    [] [] []
		$ticks
		<%
		    DROP
		    'tick' STORE
		    $y $tick ATTICK 4 GET
		    $x $tick ATTICK 4 GET
		    ATAN2
		%>
		LMAP
		MAKEGTS
		2 ->LIST
		MERGE
		1 ->LIST
		NONEMPTY
		`
	case "ceil":
		mc2 = "[ SWAP mapper.ceil 0 0 0 ] MAP\n"
	case "cos":
		mc2 = core.NewSimpleMacroMapper("COS")
	case "cumulative_sum":
		mc2 = "[ SWAP mapper.sum MAXLONG 0 0 ] MAP\n"
	case "derivative":
		if len(args) > 0 {
			mc2 += fmt.Sprintf("[ SWAP mapper.rate 1 1 0 ] MAP\n")
			mc2 += fmt.Sprintf("[ SWAP %s 1 s / TODOUBLE mapper.mul 0 0 0 ] MAP \n", args[0])
		} else {
			mc2 += "[ SWAP mapper.rate 1 1 0 ] MAP\n"
		}
	case "difference":
		mc2 = `
		[ SWAP mapper.delta 1 0 0 ] MAP
		<%
			DROP
			DUP SIZE 1 - 0 MAX -1 * SHRINK
		%>
		LMAP
		`
	case "elapsed":
		mc2 = `
		[ SWAP mapper.tick 0 0 0 ] MAP
		[ SWAP mapper.delta 1 0 0 ] MAP
		<%
			DROP
			DUP SIZE 1 - 0 MAX -1 * SHRINK
		%>
		LMAP
		`
	case "exp":
		mc2 = core.NewSimpleMacroMapper("EXP")
	case "floor":
		mc2 = "[ SWAP mapper.floor 0 0 0 ] MAP\n"
	case "ln":
		mc2 = "[ SWAP e mapper.log 0 0 0 ] MAP\n"
	case "log":
		if len(args) == 0 {
			return "", fmt.Errorf(invalidNumberOfArgs(name, 2, len(args)+1))
		}
		mc2 = fmt.Sprintf("[ SWAP %s mapper.log 0 0 0 ] MAP\n", args[0])
	case "log2":
		mc2 = "[ SWAP 2.0 mapper.log 0 0 0 ] MAP\n"
	case "log10":
		mc2 = "[ SWAP 10.0 mapper.log 0 0 0 ] MAP\n"
	case "moving_average":
		pre := "1"
		if len(args) > 0 {
			pre = args[0]
		}
		mc2 = fmt.Sprintf("[ SWAP mapper.mean %s 0 0 ] MAP\n", pre)
	case "non_negative_derivative":
		mc2 += "[ SWAP mapper.rate 1 1 0 ] MAP\n"
		if len(args) > 0 {
			mc2 += fmt.Sprintf("[ SWAP %s 1 s / TODOUBLE mapper.mul 0 0 0 ] MAP \n", args[0])
		}
		mc2 += "[ SWAP 0.0 mapper.ge 1 1 0 ] MAP\n"
	case "non_negative_difference":
		mc2 = `
		[ SWAP mapper.delta 1 0 0 ] MAP
		<%
			DROP
			DUP SIZE 1 - 0 MAX -1 * SHRINK
		%>
		LMAP
		`
		mc2 += "[ SWAP 0.0 mapper.ge 1 1 0 ] MAP\n"
	case "pow":
		if len(args) == 0 {
			return "", fmt.Errorf(invalidNumberOfArgs(name, 2, len(args)+1))
		}
		mc2 = fmt.Sprintf("[ SWAP %s mapper.exp 0 0 0 ] MAP\n", args[0])
	case "round":
		mc2 = "[ SWAP mapper.round 0 0 0 ] MAP\n"
	case "sin":
		mc2 = core.NewSimpleMacroMapper("SIN")
	case "sqrt":
		mc2 = "[ SWAP mapper.sqrt 0 0 0 ] MAP\n"
	case "tan":
		mc2 = core.NewSimpleMacroMapper("TAN")
	default:
		return "", fmt.Errorf("Unimplemented transformation method %s", name)
	}
	return mc2, nil
}

// fillBucketize Handling of missing values
// https://docs.influxdata.com/influxdb/v1.7/query_language/data_exploration/#group-by-time-intervals-and-fill/
func (p *InfluxParser) fillBucketize() string {
	mc2 := ""
	switch p.Fill {
	case influxql.NullFill:
	case influxql.NoFill:
	case influxql.NumberFill:

		switch p.FillValue.(type) {
		case int:
			mc2 = fmt.Sprintf("[ NaN NaN NaN %d ] FILLVALUE \n", p.FillValue)
		case int64:
			mc2 = fmt.Sprintf("[ NaN NaN NaN %d ] FILLVALUE \n", p.FillValue)
		case float32:
			mc2 = fmt.Sprintf("[ NaN NaN NaN %f ] FILLVALUE \n", p.FillValue)
		case float64:
			mc2 = fmt.Sprintf("[ NaN NaN NaN %f ] FILLVALUE \n", p.FillValue)
		case string:
			mc2 = fmt.Sprintf("[ NaN NaN NaN '%s' ] FILLVALUE \n", p.FillValue)
		}
	case influxql.PreviousFill:
		mc2 = "FILLPREVIOUS \n"
	case influxql.LinearFill:
		mc2 = "INTERPOLATE \n"
	}
	mc2 += "UNBUCKETIZE \n"
	return mc2
}

// parseAggregationName Parse Influx Aggregation methods:
// https://docs.influxdata.com/influxdb/v1.7/query_language/functions/#aggregations
func (p *InfluxParser) parseAggregationName(name string, args []string) (string, error) {
	mc2 := ""
	switch name {
	case "bottom":
		if len(args) == 0 {
			return "", fmt.Errorf(invalidNumberOfArgs(name, 2, len(args)+1))
		}
		mc2 += p.getBottom(p.BucketTime, args[len(args)-1], args[:len(args)-1])
		p.KeepTopLabels = append(p.KeepTopLabels, args[:len(args)-1]...)
	case "count":
		mc2 += fmt.Sprintf("MERGE [ SWAP bucketizer.count %s %s %s ] BUCKETIZE\n", p.End, p.BucketTime, p.BucketCount)
		mc2 += p.fillBucketize()
	case "distinct":
		mc2 += p.getDistinct(p.BucketTime)
	case "first":
		if !p.HasGroupBy {
			mc2 += p.getFirst(p.BucketTime)
		} else {
			mc2 += fmt.Sprintf("MERGE [ SWAP bucketizer.first %s %s %s ] BUCKETIZE\n", p.End, p.BucketTime, p.BucketCount)
			mc2 += p.fillBucketize()
		}
	case "integral":
		if len(args) > 0 {
			mc2 += fmt.Sprintf("%s 'p1' STORE ", args[0])
		} else {
			mc2 += fmt.Sprintf("1 s 'p1' STORE ")
		}
		mc2 += p.integral(p.BucketTime)
	case "last":
		if !p.HasGroupBy {
			mc2 += p.getLast(p.BucketTime)
		} else {
			mc2 += fmt.Sprintf("MERGE [ SWAP bucketizer.last %s %s %s ] BUCKETIZE\n", p.End, p.BucketTime, p.BucketCount)
			mc2 += p.fillBucketize()
		}
	case "max":
		if !p.HasGroupBy {
			mc2 += p.getMax(p.BucketTime)
		} else {
			mc2 += fmt.Sprintf("MERGE [ SWAP bucketizer.max %s %s %s ] BUCKETIZE\n", p.End, p.BucketTime, p.BucketCount)
			mc2 += p.fillBucketize()
		}
	case "mean":
		mc2 += fmt.Sprintf("MERGE [ SWAP bucketizer.mean %s %s %s ] BUCKETIZE\n", p.End, p.BucketTime, p.BucketCount)
		mc2 += p.fillBucketize()
	case "median":
		mc2 += fmt.Sprintf("MERGE [ SWAP bucketizer.median %s %s %s ] BUCKETIZE\n", p.End, p.BucketTime, p.BucketCount)
		mc2 += p.fillBucketize()
	case "min":
		if !p.HasGroupBy {
			mc2 += p.getMin(p.BucketTime)
		} else {
			mc2 += fmt.Sprintf("MERGE [ SWAP bucketizer.min %s %s %s ] BUCKETIZE\n", p.End, p.BucketTime, p.BucketCount)
			mc2 += p.fillBucketize()
		}
	case "mode":
		mc2 += p.mode(p.BucketTime)
		mc2 += fmt.Sprintf("[ SWAP bucketizer.last %s %s %s ] BUCKETIZE\n", p.End, p.BucketTime, p.BucketCount)
		mc2 += p.fillBucketize()
	case "percentile":
		if len(args) == 0 {
			return "", fmt.Errorf(invalidNumberOfArgs(name, 2, len(args)+1))
		}
		mc2 += fmt.Sprintf("MERGE [ SWAP %s TODOUBLE bucketizer.percentile %s %s %s ] BUCKETIZE\n", args[0], p.End, p.BucketTime, p.BucketCount)
		mc2 += p.fillBucketize()
	case "sample":
		if len(args) == 0 {
			return "", fmt.Errorf(invalidNumberOfArgs(name, 2, len(args)+1))
		}
		mc2 += p.getSample(p.BucketTime, args[0])
	case "spread":
		mc2 += "DUP\n"
		mc2 += fmt.Sprintf("[ SWAP bucketizer.max %s %s %s ] BUCKETIZE\n", p.End, p.BucketTime, p.BucketCount)
		mc2 += fmt.Sprintf("[ SWAP [ %s ] reducer.max ] REDUCE\n", strings.Join(p.GroupByTags, " "))
		mc2 += "SWAP\n"
		mc2 += fmt.Sprintf("[ SWAP bucketizer.min %s %s %s ] BUCKETIZE\n", p.End, p.BucketTime, p.BucketCount)
		mc2 += fmt.Sprintf("[ SWAP [ %s ] reducer.min ] REDUCE\n", strings.Join(p.GroupByTags, " "))
		mc2 += "[ SWAP -1 mapper.mul 0 0 0 ] MAP\n"
		mc2 += "APPEND\n"
		mc2 += fmt.Sprintf("[ SWAP [ %s ] reducer.sum ] REDUCE\n", strings.Join(p.GroupByTags, " "))
		mc2 += p.fillBucketize()
	case "stddev":
		mc2 += fmt.Sprintf("MERGE [ SWAP true bucketizer.sd %s %s %s ] BUCKETIZE\n", p.End, p.BucketTime, p.BucketCount)
		mc2 += p.fillBucketize()
	case "sum":
		mc2 += fmt.Sprintf("MERGE [ SWAP bucketizer.sum %s %s %s ] BUCKETIZE\n", p.End, p.BucketTime, p.BucketCount)
		mc2 += p.fillBucketize()
	case "top":
		if len(args) == 0 {
			return "", fmt.Errorf(invalidNumberOfArgs(name, 2, len(args)+1))
		}
		mc2 += p.getTop(p.BucketTime, args[len(args)-1], args[:len(args)-1])
		p.KeepTopLabels = append(p.KeepTopLabels, args[:len(args)-1]...)
	default:
		return "", fmt.Errorf("Unimplemented aggregation method %s", name)
	}

	return mc2, nil
}

// startPartition: Handle Influx Group by tags/time
// https://docs.influxdata.com/influxdb/v1.7/query_language/data_exploration/#the-group-by-clause
func (p *InfluxParser) startPartition() string {
	mc2 := fmt.Sprintf("[ %s ] DUP 'tags' STORE PARTITION", strings.Join(p.GroupByTags, " "))
	mc2 += `
	[] SWAP <% 
		SWAP
		'.InfluxDBName' REMOVE DROP 
		'keys' STORE
		{ '.InfluxDBName' NULL } RELABEL
		{}
		$tags
		<%
			DUP
			$keys
			SWAP 
			GET 
			<% DUP ISNULL %>
			<% DROP DROP CONTINUE %>
			IFT
			SWAP
			PUT
		%>
		FOREACH
		'.InfluxDBName' REMOVE DROP
		'.INFLUXQL_COLUMN_NAME' REMOVE DROP
		SETATTRIBUTES
	`
	return mc2
}

// getDistinct: Distinct aggregation specific implementation
// https://docs.influxdata.com/influxdb/v1.7/query_language/functions/#distinct
func (p *InfluxParser) getDistinct(bucket string) string {
	return p.genAggregationFunction(bucket, "COMPACT")
}

// getMin: Min aggregation specific implementation
// https://docs.influxdata.com/influxdb/v1.7/query_language/functions/#min
func (p *InfluxParser) getMin(bucket string) string {
	mc2 := `
	DUP [ SWAP bucketizer.min 0 0 1 ] BUCKETIZE
	0 GET VALUES 0 GET 'min' STORE
	[ SWAP $min mapper.eq 0 0 0 ] MAP
	0 GET
	DUP SIZE 1 MIN SHRINK
	`
	return p.genAggregationFunction(bucket, mc2)
}

// getMax: Max aggregation specific implementation
// https://docs.influxdata.com/influxdb/v1.7/query_language/functions/#max
func (p *InfluxParser) getMax(bucket string) string {
	mc2 := `
	DUP [ SWAP bucketizer.max 0 0 1 ] BUCKETIZE
	0 GET VALUES 0 GET 'max' STORE
	[ SWAP $max mapper.eq 0 0 0 ] MAP
	0 GET
	DUP SIZE 1 MIN SHRINK
	`
	return p.genAggregationFunction(bucket, mc2)
}

// getFirst: First aggregation specific implementation
// https://docs.influxdata.com/influxdb/v1.7/query_language/functions/#first
func (p *InfluxParser) getFirst(bucket string) string {
	mc2 := `
	DUP SIZE 1 - 1 MIN SHRINK
	`
	return p.genAggregationFunction(bucket, mc2)
}

// getLast: Last aggregation specific implementation
// https://docs.influxdata.com/influxdb/v1.7/query_language/functions/#last
func (p *InfluxParser) getLast(bucket string) string {
	mc2 := `
	DUP SIZE 1 - -1 * -1 MAX SHRINK
	`
	return p.genAggregationFunction(bucket, mc2)
}

// getSample: Sample aggregation specific implementation
// https://docs.influxdata.com/influxdb/v1.7/query_language/functions/#sample
func (p *InfluxParser) getSample(bucket string, limit string) string {
	mc2 := `
	DUP CLONEEMPTY SWAP
	DUP LABELS ".chunkid" GET TOLONG 'endTick' STORE 
	DUP
	TICKLIST
	SWAP
	VALUES
	2 ->LIST
	ZIP
	DUP SIZE 'size' STORE
	SHUFFLE
	`

	mc2 += fmt.Sprintf("[ 0 %s 1 - $size 1 - MIN ]", limit)

	mc2 += `
	SUBLIST
	<%
		'sample' STORE
		$sample 0 GET NaN NaN NaN $sample 1 GET ADDVALUE
	%>
	FOREACH
	`
	return p.genAggregationFunction(bucket, mc2)
}

// getBottom: Bottom aggregation specific implementation
// https://docs.influxdata.com/influxdb/v1.7/query_language/functions/#bottom
func (p *InfluxParser) getBottom(bucket string, limit string, args []string) string {
	prefix := ""
	suffix := ""

	if len(args) > 0 {
		scriptLimit := limit
		limit = "1"

		topPartition := "'" + strings.Join(args, "' '") + "'"

		prefix += `
		[]
		SWAP
		`
		prefix += fmt.Sprintf("[ %s ] PARTITION\n", topPartition)
		prefix += `
		<% 
			SWAP	
			DROP
		`

		suffix += `
			+
		%>
		FOREACH
		FLATTEN
		LASTSORT
		DUP SIZE 'size' STORE
		`
		suffix += fmt.Sprintf("[ 0 %s 1 - $size MIN ]\n", scriptLimit)
		suffix += "SUBLIST\n"
	}

	mc2 := ""
	if p.HasGroupBy {
		mc2 += `
		DUP CLONEEMPTY SWAP
		DUP LABELS ".chunkid" GET TOLONG 'endTick' STORE
		VALUESORT 
		DUP
		TICKLIST DUP SIZE 'size' STORE  `
	} else {
		mc2 += `
		DUP CLONEEMPTY SWAP
        VALUESORT
		DUP TICKLIST DUP SIZE 'size' STORE
		`
	}

	mc2 += fmt.Sprintf("[ 0 %s 1 - $size 1 - MIN ]", limit)

	mc2 += `
	SUBLIST
	SWAP
	VALUES
	`

	mc2 += fmt.Sprintf("[ 0 %s 1 - $size 1 - MIN ]", limit)

	mc2 += `
	SUBLIST [] SWAP [] SWAP [] SWAP
	MAKEGTS
	2 ->LIST
	MERGE
	`
	return prefix + p.genAggregationFunction(bucket, mc2) + suffix
}

// getTop: Top aggregation specific implementation
// https://docs.influxdata.com/influxdb/v1.7/query_language/functions/#top
func (p *InfluxParser) getTop(bucket string, limit string, args []string) string {

	prefix := ""
	suffix := ""

	if len(args) > 0 {
		scriptLimit := limit
		limit = "1"

		topPartition := "'" + strings.Join(args, "' '") + "'"

		prefix += `
		[]
		SWAP
		`
		prefix += fmt.Sprintf("[ %s ] PARTITION\n", topPartition)
		prefix += `
		<% 
			SWAP	
			DROP
		`

		suffix += `
			+
		%>
		FOREACH
		FLATTEN
		LASTSORT
		REVERSE
		DUP SIZE 'size' STORE
		`
		suffix += fmt.Sprintf("[ 0 %s 1 - $size MIN ]\n", scriptLimit)
		suffix += "SUBLIST\n"
	}

	mc2 := ""
	if p.HasGroupBy {
		mc2 += `
		DUP CLONEEMPTY SWAP
		DUP LABELS ".chunkid" GET TOLONG 'endTick' STORE
		RVALUESORT 
		DUP
		TICKLIST DUP SIZE 'size' STORE  `
	} else {
		mc2 += `
		DUP CLONEEMPTY SWAP
        RVALUESORT
		DUP TICKLIST DUP SIZE 'size' STORE
		`
	}

	mc2 += fmt.Sprintf("[ 0 %s 1 - $size MIN ]", limit)

	mc2 += `
	SUBLIST
	SWAP
	VALUES
	`

	mc2 += fmt.Sprintf("[ 0 %s 1 - $size MIN ]", limit)

	mc2 += `
	SUBLIST [] SWAP [] SWAP [] SWAP
	MAKEGTS
	2 ->LIST
	MERGE
	`

	return prefix + p.genAggregationFunction(bucket, mc2) + suffix
}

// mode: Mode aggregation specific implementation
// https://docs.influxdata.com/influxdb/v1.7/query_language/functions/#mode
func (p *InfluxParser) mode(bucket string) string {
	return p.genAggregationFunction(bucket, `
		DUP
		MODE 
		DUP
		<% SIZE 0 > %>
		<% 
			DROP
			DUP FIRSTTICK SWAP DUP ROT ATTICK 4 GET 'value' STORE
			CLONEEMPTY $endTick NaN NaN NaN $value ADDVALUE
		%> 
		<%
			DROP
		%>
		IFTE
	`)
}

// integral: Integral aggregation specific implementation
// https://docs.influxdata.com/influxdb/v1.7/query_language/functions/#integral
func (p *InfluxParser) integral(bucket string) string {
	return p.genAggregationFunction(bucket, `
		<% DUP VALUES SIZE 0 > %>
		<% 
			[ SWAP bucketizer.first 0 $p1 0 ] BUCKETIZE FILLPREVIOUS
			[ SWAP mapper.sum MAXLONG 0 0 ] MAP
			[ SWAP bucketizer.last 0 0 1 ] BUCKETIZE
		%> 
		IFT
	`)
}

// genAggregationFunction: Influx aggregation generic case
func (p *InfluxParser) genAggregationFunction(bucket string, appliedFunctions string) string {
	res := fmt.Sprintf("%s 'bucket' STORE \n", bucket)

	if bucket == "0" || bucket == "" {
		res += `
			MERGE
			1 ->LIST
			NONEMPTY
			<% 
				DROP
				$end 'endTick' STORE
		`
	} else {

		res += `
			MERGE
			$end $bucket / $bucket * $bucket 0 0 ".chunkid" false CHUNK
			<%
                DROP
                DUP LABELS ".chunkid" GET TOLONG 'endTick' STORE
		`
	}
	res += appliedFunctions
	res += `
	%>
	LMAP
	FLATTEN
	MERGE
	1 ->LIST
	NONEMPTY
	`
	return res
}

// formatResultSeries: Influx Output expected format
func formatResultSeries(influxResponse *InfluxResponse) string {
	mc2 := `
	NONEMPTY
	DUP 'set' STORE
	<%
		DROP
		NAME
	%>
	LMAP
	UNIQUE 
	`

	if influxResponse.SLimit != 0 {
		mc2 += fmt.Sprintf("[ %d %d ] SUBLIST\n", influxResponse.SOffset, influxResponse.SLimit)
	} else if influxResponse.SOffset != 0 {
		mc2 += fmt.Sprintf("DUP SIZE [ %d ROT ] SUBLIST\n", influxResponse.SOffset)
	}

	mc2 += `
	<%
		DROP
		'name' STORE
		[ $set [] $name filter.byclass ] FILTER
	%>
	LMAP
	DUP 
	<%
		SIZE 0 >
	%>
	<%
		LIST-> DROP
	%>
	<%
	%>
	IFTE
	`
	if influxResponse.Dedup {
		mc2 += "DEDUP\n"
	}

	if influxResponse.Offset != 0 {
		mc2 += "<% " + fmt.Sprintf("DROP DUP SIZE %d - -1 * 0 MIN SHRINK ", influxResponse.Offset) + "%> LMAP \n"
	}

	if influxResponse.Limit != 0 {
		mc2 += "<% " + fmt.Sprintf("DROP DUP %d SWAP SIZE MIN SHRINK ", influxResponse.Limit) + "%> LMAP \n"
	}
	return mc2
}

// shiftSeries: SHIT series by a time duration
func (p *InfluxParser) shiftSeries() string {
	if p.Shift == "" {
		return ""
	}
	return fmt.Sprintf("%s TIMESHIFT \n", p.Shift)
}

// renameAndSetValueLabels: Manipulate series Meta to clean Tags
func removeAllLabels(keepLabels []string) string {
	res := "<% DROP {} 'saveLabels' STORE DUP LABELS\n"

	for _, key := range keepLabels {
		res += fmt.Sprintf("'%s' 'key' STORE\n", key)
		res += "DUP <% $key CONTAINSKEY %> <% $saveLabels SWAP $key GET $key PUT 'saveLabels' STORE %> <% DROP %> IFTE\n"
	}
	res += "DROP { NULL NULL } RELABEL \n"
	res += "$saveLabels RELABEL \n"
	res += "%> LMAP\n"
	return res
}

// renameAndSetValueLabels: Manipulate series Meta to set expected Tags
func (p *InfluxParser) renameAndSetInfluxLabels() string {
	mc2 := "<% "
	mc2 += fmt.Sprintf("DROP 'series' STORE $series DUP NAME '%s' \n", p.Separator)
	mc2 += " <% DUP '' == %> <% DROP 1 ->LIST %> <% SPLIT %> IFTE\n"
	mc2 += "<% DUP SIZE 1 > %> \n"
	mc2 += "<% "
	mc2 += fmt.Sprintf("DUP SIZE 1 - REMOVE DROP '%s' JOIN RENAME", p.Separator)
	mc2 += " %> \n"
	mc2 += `
		<% DROP %> IFTE 
	%> LMAP
	`
	return mc2
}

// renameAndSetInfluxStarLabels: Manipulate series Meta to set expected Tags in case of a SELECT *
func (p *InfluxParser) renameAndSetInfluxStarLabels() string {
	mc2 := "<% "
	mc2 += fmt.Sprintf("DROP 'series' STORE $series DUP NAME '%s' \n", p.Separator)
	mc2 += " <% DUP '' == %> <% DROP 1 ->LIST %> <% SPLIT %> IFTE\n"
	mc2 += "<% DUP SIZE 1 > %> "
	mc2 += "<%"
	mc2 += fmt.Sprintf(" DUP SIZE 1 - REMOVE 'influxLabel' STORE '%s' JOIN RENAME { '.INFLUXQL_COLUMN_NAME' $influxLabel } RELABEL\n", p.Separator)
	mc2 += " %>"
	mc2 += `
		<% 'influxLabel' STORE { '.INFLUXQL_COLUMN_NAME' $influxLabel } RELABEL %> IFTE 
	%> LMAP
	`
	return mc2
}
