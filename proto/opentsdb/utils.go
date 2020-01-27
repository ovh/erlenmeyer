package opentsdb

import (
	"reflect"

	"github.com/gorilla/schema"
)

// OpenTsdbDecoder openTSDB schema decoder
var OpenTsdbDecoder = schema.NewDecoder()

// Custom query string parser
func init() {
	OpenTsdbDecoder.RegisterConverter(&TSDBTime{}, func(str string) reflect.Value {
		r := TSDBTime{}

		err := r.Parse(str)
		if err == nil {
			return reflect.ValueOf(r)
		}
		return reflect.Value{} // Invalid value
	})
}

// WarpResultByteCount returns the bytes count of a response
func WarpResultByteCount(response *QueryResponse) int {
	return len(response.Metric) + sizeOfTags(response.Tags) + sizeOfAggregateTags(response.AggregateTags) + len(response.DPs)*16
}

func sizeOfTags(tags map[string]string) (size int) {
	for key, value := range tags {
		size += len(key) + len(value)
	}
	return
}

func sizeOfAggregateTags(tags []string) (size int) {
	for _, tag := range tags {
		size += len(tag)
	}
	return
}

func isValidFillPolicy(fillPolicy string) bool {
	switch fillPolicy {
	case "", none, "nan", "zero":
		return true
	default:
		return false
	}
}

func noBraces(l int, r int) bool {
	return l == -1 && r == -1
}

func checkBraceIndexes(l int, r int, m int) bool {
	/*
	   Indexes must be either:
	   - Both negative which means there is no braces at all
	   - The left one strictly greater than 0, and lower than the right minus 3 (because
	     there must be room for at least a tag definition) and the right must be in the
	     last position.
	*/
	return noBraces(l, r) || (l >= 0) && l < r-3 && r == m
}
