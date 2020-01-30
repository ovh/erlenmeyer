package graphite

import (
	"errors"
	"fmt"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"

	"github.com/ovh/erlenmeyer/core"
)

// ----------------------------------------------------------------------------
// Graphite function implementation/transpiler
//
// Documentation :
//
// https://github.com/graphite-project/graphite-web/blob/master/webapp/graphite/render/functions.py
//

const (
	swap = "SWAP"
)

var (
	functions = map[string]func(*core.Node, []string, map[string]string) (*core.Node, error){
		"fetch":                       fetch,
		"find":                        find,
		"bucketize":                   bucketize,
		"name":                        name,
		"noOp":                        noOp,
		"avg":                         averageSeries,               // aggregate.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.averageSeries
		"absolute":                    absolute,                    // math.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.absolute
		"aggregate":                   aggregate,                   // aggregate.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.aggregate
		"aggregateLine":               aggregateLine,               // aggregate.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.aggregateLine
		"aggregateWithWildcards":      aggregateWithWildcards,      // aggregate.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.aggregateWithWildcards
		"alias":                       alias,                       // name.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.alias
		"aliasByMetric":               aliasByMetric,               // name.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.aliasByMetric
		"aliasByNode":                 aliasByTags,                 // name.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.aliasByNode
		"aliasByTags":                 aliasByTags,                 // name.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.aliasByTags
		"aliasQuery":                  noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.aliasQuery
		"aliasSub":                    aliasSub,                    // name.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.aliasSub
		"alpha":                       noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.alpha
		"applyByNode":                 noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.applyByNode
		"areaBetween":                 noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.areaBetween
		"asPercent":                   noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.asPercent
		"averageAbove":                averageAbove,                // filter.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.averageAbove
		"averageBelow":                averageBelow,                // filter.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.averageBelow
		"averageOutsidePercentile":    noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.averageOutsidePercentile
		"averageSeries":               averageSeries,               // aggregate.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.averageSeries
		"averageSeriesWithWildcards":  averageSeriesWithWildcards,  // aggregate.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.averageSeriesWithWildcards
		"cactiStyle":                  noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.cactiStyle
		"changed":                     noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.changed
		"color":                       noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.color
		"consolidateBy":               consolidateBy,               // bucketize.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.consolidateBy
		"constantLine":                constantLine,                // yield.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.constantLine
		"countSeries":                 countSeries,                 // reduce.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.countSeries
		"cumulative":                  cumulative,                  // bucketize.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.cumulative
		"currentAbove":                currentAbove,                // filter.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.currentAbove
		"currentBelow":                currentBelow,                // filter.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.currentBelow
		"dashed":                      noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.dashed
		"delay":                       delay,                       // time.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.delay
		"derivative":                  derivative,                  // mapper.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.derivative
		"diffSeries":                  diffSeries,                  // aggregate.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.diffSeries
		"divideSeries":                divideSeries,                // mapper.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.divideSeries
		"divideSeriesLists":           divideSeriesLists,           // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.divideSeriesLists
		"drawAsInfinite":              drawAsInfinite,              // mapper.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.drawAsInfinite
		"events":                      noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.events
		"exclude":                     exclude,                     // filter.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.exclude
		"exponentialMovingAverage":    noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.exponentialMovingAverage
		"fallbackSeries":              noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.fallbackSeries
		"grep":                        grep,                        // filter.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.grep
		"group":                       noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.group
		"groupByNode":                 groupByNode,                 // aggregate.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.groupByNode
		"groupByNodes":                aggregateWithWildcards,      // aggregate.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.groupByNodes
		"groupByTags":                 noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.groupByTags
		"highestAverage":              highestAverage,              // filter.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.highestAverage
		"highestCurrent":              highestCurrent,              // filter.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.highestCurrent
		"highestMax":                  highestMax,                  // filter.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.highestMax
		"hitcount":                    hitcount,                    // bucketize.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.hitcount
		"holtWintersAberration":       noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.holtWintersAberration
		"holtWintersConfidenceArea":   noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.holtWintersConfidenceArea
		"holtWintersConfidenceBands":  noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.holtWintersConfidenceBands
		"holtWintersForecast":         noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.holtWintersForecast
		"identity":                    timeFunction,                // yield.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.identity
		"integral":                    integral,                    // math.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.integral
		"integralByInterval":          noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.integralByInterval
		"interpolate":                 interpolate,                 // math.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.interpolate
		"invert":                      invert,                      // mapper.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.invert
		"isNonNull":                   noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.isNonNull
		"keepLastValue":               keepLastValue,               // operate.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.keepLastValue
		"legendValue":                 noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.legendValue
		"limit":                       limit,                       // filter.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.limit
		"lineWidth":                   noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.lineWidth
		"linearRegression":            noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.linearRegression
		"linearRegressionAnalysis":    noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.linearRegressionAnalysis
		"logarithm":                   logarithm,                   // math.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.logarithm
		"log":                         logarithm,                   // math.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.logarithm
		"lowestAverage":               lowestAverage,               // filter.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.lowestAverage
		"lowestCurrent":               lowestCurrent,               // filter.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.lowestCurrent
		"mapSeries":                   noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.mapSeries
		"maxSeries":                   maxSeries,                   // aggregate.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.maxSeries
		"maximumAbove":                maximumAbove,                // filter.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.maximumAbove
		"maximumBelow":                maximumBelow,                // filter.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.maximumBelow
		"minMax":                      minMax,                      // max.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.minMax
		"minSeries":                   minSeries,                   // aggregate.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.minSeries
		"minimumAbove":                minimumAbove,                // filter.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.minimumAbove
		"minimumBelow":                minimumBelow,                // filter.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.minimumBelow
		"mostDeviant":                 noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.mostDeviant
		"movingAverage":               noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.movingAverage
		"movingMax":                   noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.movingMax
		"movingMedian":                noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.movingMedian
		"movingMin":                   noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.movingMin
		"movingSum":                   noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.movingSum
		"movingWindow":                noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.movingWindow
		"multiplySeries":              multiplySeries,              // aggregate.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.multiplySeries
		"multiplySeriesWithWildcards": multiplySeriesWithWildcards, // aggregate.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.multiplySeriesWithWildcards
		"nPercentile":                 noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.nPercentile
		"nonNegativeDerivative":       noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.nonNegativeDerivative
		"offset":                      offset,                      // mapper.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.offset
		"offsetToZero":                noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.offsetToZero
		"perSecond":                   perSecond,                   // mapper.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.perSecond
		"percentileOfSeries":          noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.percentileOfSeries
		"pow":                         pow,                         // map.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.pow
		"powSeries":                   noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.powSeries
		"randomWalkFunction":          randomWalkFunction,          // yield.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.randomWalkFunction
		"randomWalk":                  randomWalkFunction,          // yield.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.randomWalkFunction
		"rangeOfSeries":               rangeOfSeries,               // aggregate.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.rangeOfSeries
		"reduceSeries":                noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.reduceSeries
		"removeAbovePercentile":       noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.removeAbovePercentile
		"removeAboveValue":            removeAboveValue,            // mapper.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.removeAboveValue
		"removeBelowPercentile":       noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.removeBelowPercentile
		"removeBelowValue":            removeBelowValue,            // mapper.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.removeBelowValue
		"removeBetweenPercentile":     noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.removeBetweenPercentile
		"removeEmptySeries":           removeEmptySeries,           // filter.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.removeEmptySeries
		"roundFunction":               noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.roundFunction
		"scale":                       scale,                       // mapper.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.scale
		"scaleToSeconds":              scaleToSeconds,              // mapper.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.scaleToSeconds
		"secondYAxis":                 noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.secondYAxis
		"seriesByTag":                 seriesByTag,                 // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.seriesByTag
		"setXFilesFactor":             noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.setXFilesFactor
		"sinFunction":                 sinFunction,                 // yield.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.sinFunction
		"sin":                         sinFunction,                 // yield.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.sinFunction
		"smartSummarize":              noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.smartSummarize
		"sortByMaxima":                sortByMaxima,                // sort.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.sortByMaxima
		"sortByMinima":                sortByMinima,                // sort.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.sortByMinima
		"sortByName":                  sortByName,                  // sort.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.sortByName
		"sortByTotal":                 sortByTotal,                 // sort.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.sortByTotal
		"squareRoot":                  squareRoot,                  // math.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.squareRoot
		"stacked":                     noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.stacked
		"stddevSeries":                stddevSeries,                // aggregate.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.stddevSeries
		"stdev":                       stdev,                       // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.stdev
		"substr":                      substr,                      // name.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.substr
		"sumSeries":                   sumSeries,                   // aggregate.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.sumSeries
		"sumSeriesWithWildcards":      sumSeriesWithWildcards,      // aggregate.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.sumSeriesWithWildcards
		"summarize":                   summarize,                   // bucketize.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.summarize
		"threshold":                   threshold,                   // yield.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.threshold
		"timeFunction":                timeFunction,                // yield.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.timeFunction
		"time":                        timeFunction,                // yield.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.timeFunction
		"timeShift":                   timeShift,                   // time.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.timeShift
		"timeSlice":                   timeSlice,                   // time.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.timeSlice
		"timeStack":                   noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.timeStack
		"transformNull":               transformNull,               // operate.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.transformNull
		"unique":                      unique,                      // sort.go - http://graphite.readthedocs.io/en/latest/functions.html#render.functions.unique
		"useSeriesAbove":              noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.useSeriesAbove
		"verticalLine":                noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.verticalLine
		"weightedAverage":             noOp,                        // http://graphite.readthedocs.io/en/latest/functions.html#render.functions.weightedAverage
	}
)

// ----------------------------------------------------------------------------
// export all functions across an unique interface

// GetFunction return the function if exists
func GetFunction(name string) (func(*core.Node, []string, map[string]string) (*core.Node, error), error) {
	fn, ok := functions[name]
	if ok {
		requestFunctionCounter.With(prometheus.Labels{
			"function": name,
		}).Inc()
		return fn, nil
	}

	return nil, fmt.Errorf("The function %s does not exist", name)
}

// ----------------------------------------------------------------------------
// helper functions

func fetch(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 3 {
		return nil, errors.New("Missing arguments")
	}

	serie, labels, err := ParseSerie(args[0])
	if err != nil {
		return nil, err
	}

	kwargs["span"] = "60 s" // nolint: goconst
	kwargs["count"] = "0"
	kwargs["end"] = args[2]

	op := "mean"
	consolidate, ok := kwargs["consolidate"]

	if ok {
		op = consolidate
	}

	node, err = bucketize(node, []string{swap, op}, kwargs)
	if err != nil {
		return nil, err
	}

	node.Left = core.NewNode(core.FetchPayload{
		ClassName: serie,
		Start:     args[1],
		End:       args[2],
		Labels:    labels,
	})

	item, ok := kwargs["node"]

	if ok && item == "true" {
		return node, nil
	}

	return node.Left, nil
}

func find(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 1 {
		return nil, errors.New("Missing arguments")
	}

	serie, labels, err := ParseSerie(args[0])
	if err != nil {
		return nil, err
	}

	node.Left = core.NewNode(core.FindPayload{
		ClassName: serie,
		Labels:    labels,
	})

	return node.Left, nil
}

func noOp(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	log.WithFields(log.Fields{
		"func": kwargs["func"],
	}).Warn("Call to a function which is not implemented")

	return node, nil
}

// ----------------------------------------------------------------------------
// graphite functions implementations

func seriesByTag(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	name := "*"
	labels := args
	match := false
	for idx, label := range labels {
		if strings.HasPrefix(label, "name"+opNotRegExp) {
			name = fmt.Sprintf("~(?!%s)", label[7:]) // len('name!=~') == 7
			match = true
		} else if strings.HasPrefix(label, "name"+opNotEq) {
			name = fmt.Sprintf("~(?!%s)", label[6:]) // len('name!=') == 6
			match = true
		} else if strings.HasPrefix(label, "name"+opRegExp) {
			name = fmt.Sprintf("~%s", label[6:]) // len('name=~') == 6
			match = true
		} else if strings.HasPrefix(label, "name"+opEq) {
			name = fmt.Sprintf("%s", label[5:]) // len('name=') == 5
			match = true
		}

		if match {
			if idx == 0 {
				labels = labels[1:]
			} else if idx == len(labels)-1 {
				labels = labels[:len(labels)-2]
			} else {
				labels = append(labels[:idx-1], labels[idx:]...)
			}
			break
		}
	}

	serie := strings.Trim(fmt.Sprintf("%s;%s", name, strings.Join(labels, ";")), ";")

	return fetch(node, []string{serie, kwargs["from"], kwargs["until"]}, kwargs)
}
