package prom

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/ovh/erlenmeyer/core"

	"github.com/ovh/erlenmeyer/proto/prom/promql"
)

type testStruct struct {
	Query          string
	ShouldContains []string
}

func TestWarpScriptGeneration(t *testing.T) {

	context := Context{}
	var err error

	context.Start, err = core.ParsePromTime("2015-07-01T20:10:30.781Z")
	if err != nil {
		panic(err)
	}
	context.End, err = core.ParsePromTime("2015-07-01T21:20:30.123Z")
	if err != nil {
		panic(err)
	}
	context.Step, err = core.ParsePromDuration("5s")
	if err != nil {
		panic(err)
	}

	for _, test := range tests {

		context.Query = test.Query
		context.Expr, err = promql.ParseExpr(context.Query)
		if err != nil {
			panic(err)
		}

		log.Printf("XXXXXXXXXXX %+v\n", context.Expr)

		evaluator := evaluator{}

		tree := evaluator.GenerateQueryTree(context)

		fmt.Println(fmt.Sprintf("XXXXXXXX     %+v", tree))

		mc2 := tree.ToWarpScript("abcd", context.Query, context.Step)
		mc2 += "\n[ SWAP mapper.tostring 0 0 0 ] MAP\n"
		fmt.Println("XXXXXXXX     generated: begin")
		log.Println(mc2)
		fmt.Println("XXXXXXXX     generated: end")

		for _, shouldContain := range test.ShouldContains {
			if !strings.Contains(mc2, shouldContain) {
				t.Errorf("Error testing query '%s'", test.Query)
				t.Errorf("final mc2: \n'%s'", mc2)
				t.Errorf("looking for '%s'", shouldContain)
				return
			}
		}
		log.Println("-----------------------------")
	}
}

var tests = []testStruct{
	{
		Query: `container.rate`,
		ShouldContains: []string{
			"[ $token 'container.rate' {}  1435781430781000 0 -  ISO8601 1435785630123000 0 -  ISO8601 ] FETCH",
		},
	},
	{
		Query: `sort_desc(sum(rate(container_cpu_user_seconds_total{image!=""}[1m])) by (name))`,
		ShouldContains: []string{
			"[ SWAP mapper.rate $step $range MAX -1 * 0 $bucketCount 1 - -1 * ] MAP",
			"[ SWAP [ 'name' ] DUP 'equivalenceClass' STORE reducer.sum ] REDUCE",
			"<% [ SWAP bucketizer.mean 0 0 1 ] BUCKETIZE VALUES 0 GET 0 GET %> SORTBY REVERSE",
		},
	},
	{
		Query: `(sum(node_memory_MemTotal) - sum(node_memory_MemFree+node_memory_Buffers+node_memory_Cached) ) / sum(node_memory_MemTotal) * 100`,
		ShouldContains: []string{
			"<% DUP LABELS ->JSON 'UTF-8' ->BYTES SHA1 ->HEX  'hash_945fa9bc3027d7025e3' SWAP 2 ->MAP RELABEL 1 ->LIST $l SWAP APPEND 'l' STORE %> FOREACH",
			"[ SWAP $right TODOUBLE mapper.mul 0 0 0 ] MAP",
		},
	},
	{
		Query: `http_requests_total{job="apiserver", handler="/comments"}`,
		ShouldContains: []string{
			"'handler'  '/comments'",
			"'job'  'apiserver'",
			"1435781430781000 0 -  ISO8601 1435785630123000 0 -  ISO8601 ] FETCH"},
	},
	{
		Query: `http_requests_total[5m] offset 1w`,
		ShouldContains: []string{
			"[ $token 'http_requests_total' {}  1435781430781000 604800000000 -  ISO8601 1435785630123000 604800000000 -  ISO8601 ] FETCH",
			"300000000 'range' STORE [ SWAP bucketizer.last 1435785630123000 604800000000 -  5 s  1435785630123000 1435781430781000 - 5 s / TOLONG ] BUCKETIZE",
		},
	},
	{
		Query: `abs(http_requests_total)`,
		ShouldContains: []string{
			"[ SWAP mapper.abs 0 0 0 ] MAP"},
	},
	{
		Query: `log10(http_requests_total offset 1w)`,
		ShouldContains: []string{
			"[ SWAP 10.0 mapper.log 0 0 0 ] MAP"},
	},
	{
		Query: `changes(http_requests_total[5m])`,
		ShouldContains: []string{
			"COMPACT MARK SWAP <% DUP DUP DUP NAME 'name' STORE LABELS 'l' STORE LASTTICK 'lt' STORE",
			"VALUES 'list' STORE $list SIZE 's' STORE $list $s 1 - GET $list $s 2 - GET",
			"<%  ==  %> <% $s 1 - %> <% $s %> IFTE 'val' STORE NEWGTS $name RENAME $l RELABEL $lt NaN DUP DUP $val SETVALUE ",
			"%> FOREACH COUNTTOMARK ->LIST SWAP DROP"},
	},
	{
		Query: `12 + 13`,
		ShouldContains: []string{
			"12 'left-0' STORE",
			"13 'right-0' STORE",
			"$left-0 'left' STORE",
			"$right-0 'right' STORE",
			"$left $right",
			"+"},
	},
	{
		Query: `http_request + 12`,
		ShouldContains: []string{
			"'left-0' STORE",
			"12 'right-0' STORE",
			"$left",
			"[ SWAP $right TODOUBLE mapper.add 0 0 0 ] MAP"},
	},
	{
		Query: `http_request + http_bytes`,
		ShouldContains: []string{
			"[ $token 'http_request' {}  1435781430781000 0 -  ISO8601 1435785630123000 0 -  ISO8601 ] FETCH",
			"$left $right 2 ->LIST",
			"<% DUP LABELS ->JSON 'UTF-8' ->BYTES SHA1 ->HEX  'hash_945fa9bc3027d7025e3' SWAP 2 ->MAP RELABEL 1 ->LIST $l SWAP APPEND 'l' STORE %> FOREACH",
			"[ SWAP  DUP 0 GET @HASHLABELS SWAP 1 GET @HASHLABELS [ 'hash_945fa9bc3027d7025e3' ] op.add ]  APPLY NONEMPTY { 'hash_945fa9bc3027d7025e3' '' } RELABEL"},
	},
	{
		Query: `12 - 13`,
		ShouldContains: []string{
			"12 'left-0' STORE",
			"13 'right-0' STORE",
			"$left $right",
			"-"},
	},
	{
		Query: `http_request - 12`,
		ShouldContains: []string{
			"'left-0' STORE",
			"12 'right-0' STORE",
			"$left",
			"[ SWAP 0 $right TODOUBLE - mapper.add 0 0 0 ] MAP"},
	},
	{
		Query: `http_request - http_bytes`,
		ShouldContains: []string{
			"[ $token 'http_request' {}  1435781430781000 0 -  ISO8601 1435785630123000 0 -  ISO8601 ] FETCH",
			"$left $right 2 ->LIST",
			"[ SWAP  DUP 0 GET @HASHLABELS '%2B.tosub' RENAME SWAP 1 GET @HASHLABELS [ 'hash_945fa9bc3027d7025e3' ] op.sub ]  APPLY NONEMPTY { 'hash_945fa9bc3027d7025e3' '' } RELABEL"},
	},
	{
		Query: `12 * 13`,
		ShouldContains: []string{
			"12 'left-0' STORE",
			"13 'right-0' STORE",
			"$left $right",
			"*"},
	},
	{
		Query: `http_request * 12`,
		ShouldContains: []string{
			"'left-0' STORE",
			"12 'right-0' STORE",
			"$left",
			"[ SWAP $right TODOUBLE mapper.mul 0 0 0 ] MAP"},
	},
	{
		Query: `http_request * http_bytes`,
		ShouldContains: []string{
			"'left-0' STORE",
			"'right-0' STORE",
			"$left $right 2 ->LIST",
			"[ SWAP  DUP 0 GET @HASHLABELS SWAP 1 GET @HASHLABELS [ 'hash_945fa9bc3027d7025e3' ] op.mul ]  APPLY NONEMPTY { 'hash_945fa9bc3027d7025e3' '' } RELABEL"},
	},
	{
		Query: `12 / 13`,
		ShouldContains: []string{
			"12 'left-0' STORE",
			"13 'right-0' STORE",
			"$left $right",
			"/"},
	},
	{
		Query: `http_request / 12`,
		ShouldContains: []string{
			"'left-0' STORE",
			"12 'right-0' STORE",
			"$left",
			"[ SWAP 1 $right TODOUBLE / mapper.mul 0 0 0 ] MAP"},
	},
	{
		Query: `http_request / http_bytes`,
		ShouldContains: []string{
			"'left' STORE",
			"'right' STORE",
			"$left $right 2 ->LIST",
			"[ SWAP  DUP 0 GET @HASHLABELS '%2B.todiv' RENAME SWAP 1 GET @HASHLABELS [ 'hash_945fa9bc3027d7025e3' ] op.div ]  APPLY NONEMPTY { 'hash_945fa9bc3027d7025e3' '' } RELABEL"},
	},
	{
		Query: `12 % 13`,
		ShouldContains: []string{
			"12 'left-0' STORE",
			"13 'right-0' STORE",
			"$left $right",
			"%"},
	},
	{
		Query: `http_request % 12`,
		ShouldContains: []string{
			"'left-0' STORE",
			"12 'right-0' STORE",
			"$left",
			"$mapping_window 7 GET 0 GET $right %"},
	},
	{
		Query: `12 ^ 13`,
		ShouldContains: []string{
			"12 'left-0' STORE",
			"13 'right-0' STORE",
			"$left $right",
			"**"},
	},
	{
		Query: `http_request^2`,
		ShouldContains: []string{
			"'left-0' STORE",
			"2 'right-0' STORE",
			"$left",
			"[ SWAP $right TODOUBLE mapper.pow 0 0 0 ] MAP"},
	},
	{
		Query: `floor(http_request)`,
		ShouldContains: []string{
			"[ SWAP mapper.floor 0 0 0 ] MAP",
		},
	},
	{
		Query: `sqrt(http_request)`,
		ShouldContains: []string{
			"[ SWAP mapper.sqrt 0 0 0 ] MAP",
		},
	},
	{
		Query: `clamp_max(http_requests_total,4)`,
		ShouldContains: []string{
			"[ SWAP  4  'scalar' STORE <% $scalar TYPEOF 'LIST' == $scalar TYPEOF 'GTS' == || %> <% [ $scalar bucketizer.last 0 0 1 ] BUCKETIZE FLATTEN 0 GET VALUES 0 GET %> <% $scalar %> IFTE",
			"mapper.min.x 0 0 0 ] MAP"},
	},
	{
		Query: `clamp_min(http_requests_total,8)`,
		ShouldContains: []string{
			"[ SWAP  8  'scalar' STORE <% $scalar TYPEOF 'LIST' == $scalar TYPEOF 'GTS' == || %> <% [ $scalar bucketizer.last 0 0 1 ] BUCKETIZE FLATTEN 0 GET VALUES 0 GET %> <% $scalar %> IFTE",
			"mapper.max.x 0 0 0 ] MAP"},
	},
	{
		Query: `count_scalar(http_requests_total)`,
		ShouldContains: []string{
			"[ SWAP [ ] reducer.count ] REDUCE"},
	},
	{
		Query: `drop_common_labels(http_requests_total)`,
		ShouldContains: []string{
			"DUP [ SWAP [ ] reducer.count ] REDUCE 0 GET LABELS KEYLIST MARK SWAP <% '' %> FOREACH COUNTTOMARK ->MAP SWAP DROP RELABEL"},
	},
	{
		Query: `exp(http_request)`,
		ShouldContains: []string{
			"$mapping_window 7 GET 0 GET EXP"},
	},
	{
		Query: `day_of_month(http_requests_total)`,
		ShouldContains: []string{
			"[ SWAP 'UTC' mapper.day 0 0 0 ] MAP"},
	},
	{
		Query: `day_of_week(http_requests_total)`,
		ShouldContains: []string{
			"[ SWAP 'UTC' mapper.weekday 0 0 0 ] MAP",
			"[ SWAP 7 mapper.mod 0 0 0 ] MAP"},
	},
	{
		Query: `days_in_month(http_requests_total)`,
		ShouldContains: []string{
			"[ SWAP <%  'mapping_window' STORE  $mapping_window 0 GET  'tick' STORE  $tick TSELEMENTS DUP 0 GET 'year' STORE 1 GET 'month' STORE",
			"$month $year @DAYSINMONTH  'days' STORE $tick NaN NaN NaN $days %> MACROMAPPER 0 0 0 ] MAP"},
	},
	{
		Query: `holt_winters(http_requests_total[5m],0.8,0.6)`,
		ShouldContains: []string{
			"0.8  'scalar' STORE <% $scalar TYPEOF 'LIST' == $scalar TYPEOF 'GTS' == || %> <% [ $scalar bucketizer.last 0 0 1 ] BUCKETIZE FLATTEN 0 GET VALUES 0 GET %> <% $scalar %> IFTE",
			"0.6  'scalar' STORE <% $scalar TYPEOF 'LIST' == $scalar TYPEOF 'GTS' == || %> <% [ $scalar bucketizer.last 0 0 1 ] BUCKETIZE FLATTEN 0 GET VALUES 0 GET %> <% $scalar %> IFTE",
			"DOUBLEEXPONENTIALSMOOTHING 0 GET"},
	},
	{
		Query: `hour(http_requests_total)`,
		ShouldContains: []string{
			"[ SWAP 'UTC' mapper.hour 0 0 0 ] MAP"},
	},
	{
		Query: `sum(http_requests_total) by (instance, host)`,
		ShouldContains: []string{
			"[ SWAP [ 'instance' 'host' ] DUP 'equivalenceClass' STORE reducer.sum ] REDUCE"},
	},
	{
		Query: `sum(http_requests_total)`,
		ShouldContains: []string{
			"[ SWAP [ ] DUP 'equivalenceClass' STORE reducer.sum ] REDUCE"},
	},
	{
		Query: `sum(http_requests_total) without (instance)`,
		ShouldContains: []string{
			"{ 'instance' '' }",
			"[ SWAP [ ] DROP $equivalenceClass reducer.sum ] REDUCE",
			"UNBUCKETIZE [ SWAP mapper.finite 0 0 0 ] MAP"},
	},
	{
		Query: `sum(http_requests_total) without (instance, host)`,
		ShouldContains: []string{
			"{ 'instance' '' 'host' '' }",
			"[ SWAP [ ] DROP $equivalenceClass reducer.sum ] REDUCE",
			"UNBUCKETIZE [ SWAP mapper.finite 0 0 0 ] MAP"},
	},
	{
		Query: `sum(http_requests_total) by (instance, host)`,
		ShouldContains: []string{
			"[ SWAP [ 'instance' 'host' ] DUP 'equivalenceClass' STORE reducer.sum ] REDUCE"},
	},
	{
		Query: `min(http_requests_total)`,
		ShouldContains: []string{
			"[ SWAP [ ] DUP 'equivalenceClass' STORE reducer.min ] REDUCE"},
	},
	{
		Query: `max(http_requests_total)`,
		ShouldContains: []string{
			"[ SWAP [ ] DUP 'equivalenceClass' STORE reducer.max ] REDUCE"},
	},
	{
		Query: `avg(http_requests_total)`,
		ShouldContains: []string{
			"[ SWAP [ ] DUP 'equivalenceClass' STORE reducer.mean.exclude-nulls ] REDUCE"},
	},
	{
		Query: `stddev(http_requests_total)`,
		ShouldContains: []string{
			"[ SWAP [ ] DUP 'equivalenceClass' STORE true reducer.sd ] REDUCE"},
	},
	{
		Query: `stdvar(http_requests_total)`,
		ShouldContains: []string{
			"[ SWAP [ ] DUP 'equivalenceClass' STORE reducer.var ] REDUCE"},
	},
	{
		Query: `count(http_requests_total)`,
		ShouldContains: []string{
			"[ SWAP [ ] DUP 'equivalenceClass' STORE reducer.count.exclude-nulls ] REDUCE"},
	},
	{
		Query: `round(http_requests_total)`,
		ShouldContains: []string{
			"[ SWAP mapper.round 0 0 0 ] MAP"},
	},
	{
		Query: `topk(5, http_requests_total)`,
		ShouldContains: []string{
			" [ SWAP bucketizer.last 1435785630123000 0 -  5 s  1435785630123000 1435781430781000 - 5 s / TOLONG ] BUCKETIZE",
			"<% [ SWAP bucketizer.mean 0 0 1 ] BUCKETIZE VALUES 0 GET 0 GET %> SORTBY REVERSE [ 0 5 1 - ] SUBLIST"},
	},
	{
		Query: `bottomk(5, http_requests_total)`,
		ShouldContains: []string{
			" [ SWAP bucketizer.last 1435785630123000 0 -  5 s  1435785630123000 1435781430781000 - 5 s / TOLONG ] BUCKETIZE",
			"<% [ SWAP bucketizer.mean 0 0 1 ] BUCKETIZE VALUES 0 GET 0 GET %> SORTBY [ 0 5 1 - ] SUBLIST"},
	},
	{
		Query: `scalar(http_requests_total)`,
		ShouldContains: []string{
			"DUP SIZE <% 1 == %> <% VALUES 0 GET 0 GET %> <% DROP NaN %> IFTE"},
	},
	{
		Query: `sort(http_requests_total)`,
		ShouldContains: []string{
			"<% [ SWAP bucketizer.mean 0 0 1 ] BUCKETIZE VALUES 0 GET 0 GET %> SORTBY"},
	},
	{
		Query: `sort_desc(http_requests_total)`,
		ShouldContains: []string{
			"<% [ SWAP bucketizer.mean 0 0 1 ] BUCKETIZE VALUES 0 GET 0 GET %> SORTBY REVERSE"},
	},
	{
		Query: `time()`,
		ShouldContains: []string{
			"[ $start $end ] [] [] [] [ 1 DUP ] { '" + core.ShouldRemoveNameLabel + "' 'true' } SETATTRIBUTES MAKEGTS 'scalar' RENAME",
			"[ SWAP bucketizer.mean $end $step $instant ] BUCKETIZE INTERPOLATE SORT",
			"[ SWAP mapper.tick 0 0 0 ] MAP [ SWAP 0.000001 mapper.mul 0 0 0 ] MAP"},
	},
	{
		Query: `vector(42)`,
		ShouldContains: []string{
			"42 'scalar' STORE  <% $scalar TYPEOF 'LIST' != %> <% [ $start $end ] [] [] [] [ $scalar ] MAKEGTS  'vector' RENAME [ SWAP bucketizer.mean $end $step $instant ] BUCKETIZE INTERPOLATE SORT %> <% $scalar <% DROP 'vector' RENAME %> LMAP %> IFTE"},
	},
	{
		Query: `year(http_requests_total)`,
		ShouldContains: []string{
			"[ SWAP 'UTC' mapper.year 0 0 0 ] MAP"},
	},
	{
		Query: `month(http_requests_total)`,
		ShouldContains: []string{
			"[ SWAP 'UTC' mapper.month 0 0 0 ] MAP"},
	},
	{
		Query: `minute(http_requests_total)`,
		ShouldContains: []string{
			"[ SWAP 'UTC' mapper.minute 0 0 0 ] MAP"},
	},
	{
		Query: `histogram_quantile(0.9, sum(rate(http_requests_total[5m])) by (le))`,
		ShouldContains: []string{
			"[ SWAP [ 'le' ] DUP 'equivalenceClass' STORE reducer.sum ] REDUCE",
			"0.9  'scalar' STORE <% $scalar TYPEOF 'LIST' == $scalar TYPEOF 'GTS' == || %> <% [ $scalar bucketizer.last 0 0 1 ] BUCKETIZE FLATTEN 0 GET VALUES 0 GET %> <% $scalar %> IFTE",
			"'QUANTILE' STORE",
			"$m $lastKey GET $QUANTILE * 'rank' STORE",
			"$metricsBucketMap @bucketQuantile 'result' STORE",
		},
	},
	{
		Query: `histogram_quantile(0.95, rate(http_requests_total[5m]))`,
		ShouldContains: []string{
			"0.95  'scalar' STORE <% $scalar TYPEOF 'LIST' == $scalar TYPEOF 'GTS' == || %> <% [ $scalar bucketizer.last 0 0 1 ] BUCKETIZE FLATTEN 0 GET VALUES 0 GET %> <% $scalar %> IFTE",
			"'QUANTILE' STORE",
		},
	},
	{
		Query: `increase(http_requests_total{job="api-server"}[5m])`,
		ShouldContains: []string{
			"FALSE RESETS",
			"[ SWAP mapper.delta $step $range MAX -1 * 0 $bucketCount 1 - -1 * ] MAP"},
	},
	{
		Query: `avg_over_time(http_requests_total[12m])`,
		ShouldContains: []string{
			" [ SWAP bucketizer.last 1435785630123000 0 -  5 s  1435785630123000 1435781430781000 - 5 s / TOLONG ] BUCKETIZE",
		},
	},
	{
		Query: `min_over_time(http_requests_total[12m])`,
		ShouldContains: []string{
			"720000000 'range' STORE [ SWAP bucketizer.last 1435785630123000 0 -  5 s  1435785630123000 1435781430781000 - 5 s / TOLONG ] BUCKETIZE",
		},
	},
	{
		Query: `max_over_time(http_requests_total[12m])`,
		ShouldContains: []string{
			"720000000 'range' STORE [ SWAP bucketizer.last 1435785630123000 0 -  5 s  1435785630123000 1435781430781000 - 5 s / TOLONG ] BUCKETIZE",
		},
	},
	{
		Query: `sum_over_time(http_requests_total[12m])`,
		ShouldContains: []string{
			"720000000 'range' STORE [ SWAP bucketizer.last 1435785630123000 0 -  5 s  1435785630123000 1435781430781000 - 5 s / TOLONG ] BUCKETIZE",
		},
	},
	{
		Query: `stddev_over_time(http_requests_total[12m])`,
		ShouldContains: []string{
			"DUP <% VALUES SIZE 0 == %> <% NEWGTS 'http_requests_total' RENAME {}  RELABEL 1 ->LIST APPEND %> IFT",
			"720000000 'range' STORE [ SWAP bucketizer.last 1435785630123000 0 -  5 s  1435785630123000 1435781430781000 - 5 s / TOLONG ] BUCKETIZE",
			"[ SWAP true mapper.sd $range $step MAX -1 * 0 0 ] MAP"},
	},
	{
		Query: `stdvar_over_time(http_requests_total[12m])`,
		ShouldContains: []string{
			"720000000 'range' STORE [ SWAP bucketizer.last 1435785630123000 0 -  5 s  1435785630123000 1435781430781000 - 5 s / TOLONG ] BUCKETIZE",
			"[ SWAP true mapper.var $range $step MAX -1 * 0 0 ] MAP"},
	},
	{
		Query: `label_replace(http_requests_total{instance="localhost:9090"},"host","$1-$2","instance","(.*):(.*)")`,
		ShouldContains: []string{
			"\"host\"  'scalar' STORE <% $scalar TYPEOF 'LIST' == $scalar TYPEOF 'GTS' == || %> <% [ $scalar bucketizer.last 0 0 1 ] BUCKETIZE FLATTEN 0 GET VALUES 0 GET %> <% $scalar %> IFTE",
			"'new_label' STORE  \"$1-$2\"  'scalar' STORE <% $scalar TYPEOF 'LIST' == $scalar TYPEOF 'GTS' == || %> <% [ $scalar bucketizer.last 0 0 1 ] BUCKETIZE FLATTEN 0 GET VALUES 0 GET %> <% $scalar %> IFTE",
			"'replacement' STORE  \"instance\"  'scalar' STORE <% $scalar TYPEOF 'LIST' == $scalar TYPEOF 'GTS' == || %> <% [ $scalar bucketizer.last 0 0 1 ] BUCKETIZE FLATTEN 0 GET VALUES 0 GET %> <% $scalar %> IFTE",
			"'src_label' STORE  \"(.*):(.*)\"  'scalar' STORE <% $scalar TYPEOF 'LIST' == $scalar TYPEOF 'GTS' == || %> <% [ $scalar bucketizer.last 0 0 1 ] BUCKETIZE FLATTEN 0 GET VALUES 0 GET %> <% $scalar %> IFTE",
			"'regex' STORE",
			"MARK SWAP <%  DUP DUP NAME 'c' STORE LABELS { '__name__' $c  '' '' } APPEND DUP  $src_label GET DUP",
			"<% ISNULL %>  <% DROP DROP %> <% DUP $regex MATCH SIZE",
			"<% 0  >  %>  <%   $regex $replacement REPLACE  $new_label <% DUP '__name__' == %> <% DROP SWAP DROP RENAME %> <% PUT RELABEL %> IFTE %>  <% DROP DROP %> IFTE",
			"%> IFTE %> FOREACH COUNTTOMARK ->LIST SWAP DROP"},
	},
	{
		Query: `http_request > 10`,
		ShouldContains: []string{
			"[ SWAP $right DUP TYPEOF <% 'LONG' == %> <% TODOUBLE %> IFT mapper.gt 0 0 0 ] MAP",
			"NONEMPTY"},
	},
	{
		Query: `http_request < 10`,
		ShouldContains: []string{
			"[ SWAP $right DUP TYPEOF <% 'LONG' == %> <% TODOUBLE %> IFT mapper.lt 0 0 0 ] MAP",
			"NONEMPTY"},
	},
	{
		Query: `http_request >= 10`,
		ShouldContains: []string{
			"[ SWAP $right DUP TYPEOF <% 'LONG' == %> <% TODOUBLE %> IFT mapper.ge 0 0 0 ] MAP",
			"NONEMPTY"},
	},
	{
		Query: `http_request <= 10`,
		ShouldContains: []string{
			"[ SWAP $right DUP TYPEOF <% 'LONG' == %> <% TODOUBLE %> IFT mapper.le 0 0 0 ] MAP",
			"NONEMPTY"},
	},
	{
		Query: `http_request == 10`,
		ShouldContains: []string{
			"[ SWAP $right DUP TYPEOF <% 'LONG' == %> <% TODOUBLE %> IFT mapper.eq 0 0 0 ] MAP",
			"NONEMPTY"},
	},
	{
		Query: `http_request != 10`,
		ShouldContains: []string{
			"[ SWAP $right DUP TYPEOF <% 'LONG' == %> <% TODOUBLE %> IFT mapper.ne 0 0 0 ] MAP",
			"NONEMPTY"},
	},
	{
		Query: `http_request_1 and http_request_2`,
		ShouldContains: []string{
			"CLEAR [ $left DUP TYPEOF <% 'GTS' == %> <% [ SWAP ] %> IFT",
			"$right DUP TYPEOF <% 'LIST' == %> <% [ 0 GET ] %> IFT LABELS",
			"filter.bylabels ] FILTER",
		},
	},
	{
		Query: `http_request_1 or http_request_2`,
		ShouldContains: []string{
			"$left DUP TYPEOF 'LIST' <% == %> <% 0 GET %> IFT LABELS 'left_labels'  STORE",
			"$right DUP TYPEOF 'LIST' <% == %> <% 0 GET %> IFT LABELS 'right_labels' STORE",
			"[ $right DUP TYPEOF <% 'GTS' == %> <% [ SWAP ] %> IFT [] $right_labels filter.bylabels ] FILTER",
			"$left_labels <% 'value' STORE 'key' STORE  $value $key $right_labels SWAP GET <% == %> <% $right_labels $key REMOVE DROP 'right_labels' STORE  %> IFT %> FOREACH",
			"[ $right DUP TYPEOF <% 'GTS' == %> <% [ SWAP ] %> IFT [] $right_labels filter.bylabels ] FILTER",
			"$left 2 ->LIST FLATTEN"},
	},
	{
		Query: `http_request_1 unless http_request_2`,
		ShouldContains: []string{
			"$left_labels <% 'value' STORE 'key' STORE  $value $key $right_labels SWAP GET <% == %> <% $right_labels $key REMOVE DROP 'right_labels' STORE  %> IFT %> FOREACH",
			"[ $left $right 2 ->LIST FLATTEN [] $right_labels filter.bylabels ] FILTER"},
	},
	{
		Query:          `100.0 - 100 * (node_filesystem_free{instance=~'undefined',device !~'tmpfs',device!~'by-uuid'} / node_filesystem_size{instance=~'undefined',device !~'tmpfs',device!~'by-uuid'})`,
		ShouldContains: []string{""},
	},
	{
		Query: `sum(rate(container_cpu_usage_seconds_total{}[1m] )) / count(node_cpu{mode="system"}) * 100`,
		ShouldContains: []string{
			"",
		},
	},
	{
		Query: `histogram_quantile(0.95, sum(rate(forwarder_execution_duration_seconds_bucket{cluster_id="s1.sbg.functions",status="200",hostname="api1.s1.sbg.functions"}[1m])) by (le, function_id, hostname))`,
		ShouldContains: []string{
			"0.95",
			"'QUANTILE' STORE",
			"$reverseMap SWAP DUP TODOUBLE TOSTRING DUP $reg MATCH",
			"$metricsBucketMap @bucketQuantile 'result' STORE",
			"[ SWAP  [ 'le' ] ->SET $equivalenceClass ->SET SWAP DIFFERENCE SET-> $reducer.histogram MACROREDUCER ] REDUCE",
		},
	},
}
