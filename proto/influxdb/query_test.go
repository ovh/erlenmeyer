package influxdb

import (
	"strings"
	"testing"

	"github.com/influxdata/influxql"
)

func TestCreateDatabaseQuery(t *testing.T) {
	var tests = []struct {
		q          string
		expected   string
		shouldFail bool
	}{
		{"CREATE DATABASE UNICORN", "CREATE DATABASE UNICORN", false},
		{"CREATE DATABASES UNICORN", "CREATE DATABASE UNICORN", true},
	}

	for _, test := range tests {

		qr := strings.NewReader(test.q)
		p := influxql.NewParser(qr)
		q, err := p.ParseQuery()
		if err != nil && !test.shouldFail {
			t.Errorf("Expected nil, got error")
			continue
		}
		if err != nil {
			continue
		}
		if q.String() != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, q.String())
		}
	}
}
