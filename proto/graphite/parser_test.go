package graphite_test

import (
	"testing"

	"github.com/ovh/erlenmeyer/proto/graphite"
)

type serieTest struct {
	Is       string
	Labels   map[string]string
	ShouldBe string
}

func TestParseSerie(t *testing.T) {
	for _, serie := range seriesTest {
		s, l, e := graphite.ParseSerie(serie.Is)
		if e != nil {
			t.Error(e)

			continue
		}

		if serie.ShouldBe != s {
			t.Errorf("Series are not the same: %s != %s", serie.ShouldBe, s)

			continue
		}

		for k, v := range serie.Labels {
			val, ok := l[k]
			if !ok {
				t.Errorf("Label %s must exists", k)

				continue
			}

			if val != v {
				t.Errorf("Value for label %s is not the good one: %s != %s", k, val, v)
			}
		}
	}
}

var seriesTest = []serieTest{
	{
		Is:       "*",
		ShouldBe: "~.*?",
		Labels:   make(map[string]string),
	},
	{
		Is:       "os.*",
		ShouldBe: "~os\\..*?",
		Labels:   make(map[string]string),
	},
	{
		Is:       "os.*.used",
		ShouldBe: "~os\\..*?\\.used",
		Labels:   make(map[string]string),
	},
	{
		Is:       "os.mem.used;host=dn22",
		ShouldBe: "~os\\.mem\\.used",
		Labels: map[string]string{
			"host": "dn22",
		},
	},
	{
		Is:       "os.mem.used;host=~dn22",
		ShouldBe: "~os\\.mem\\.used",
		Labels: map[string]string{
			"host": "~dn22",
		},
	},
	{
		Is:       "os.mem.used;host!=dn22",
		ShouldBe: "~os\\.mem\\.used",
		Labels: map[string]string{
			"host": "~(?!dn22).*?",
		},
	},
	{
		Is:       "os.mem.used;host!=~dn22",
		ShouldBe: "~os\\.mem\\.used",
		Labels: map[string]string{
			"host": "~(?!dn22).*?",
		},
	},
	{
		Is:       "os.mem.used;host!=~dn22;rack=~*",
		ShouldBe: "~os\\.mem\\.used",
		Labels: map[string]string{
			"host": "~(?!dn22).*?",
			"rack": "~.*?",
		},
	},
}

// type test struct {
// 	Query     string
// 	Functions []graphite.Function
// }

// func TestParsing(t *testing.T) {
// 	for _, test := range tests {
// 		functions, err := graphite.CreateRenderRequest(test.Query, "now", "-1s")
// 		if err != nil {
// 			t.Error(err)

// 			continue
// 		}

// 		if len(functions) != len(test.Functions) {
// 			t.Error("functions have not the same length")

// 			continue
// 		}

// 		for i, function := range test.Functions {
// 			if function.Name != functions[i].Name {
// 				t.Errorf("%s - %d: function has not the same name", function.Name, i)

// 				continue
// 			}

// 			if len(function.Arguments) != len(functions[i].Arguments) {
// 				fmt.Println(function.Arguments, len(function.Arguments), functions[i].Arguments, len(functions[i].Arguments))
// 				t.Errorf("%s - %d: %s: function has not the same argument length", function.Name, i, test.Query)

// 				continue
// 			}

// 			for j, argument := range function.Arguments {
// 				if argument != functions[i].Arguments[j] {
// 					t.Errorf("%s - %d - %d: argument is not the good one", function.Name, i, j)
// 				}
// 			}
// 		}
// 	}
// }

// var tests = []test{
// 	{
// 		Query: "absolute()",
// 		Functions: []graphite.Function{
// 			{
// 				Name:       "absolute",
// 				Arguments:  make([]string, 0),
// 				Parameters: make(map[string]string),
// 			},
// 		},
// 	},
// 	{
// 		Query: "absolute(os.cpu)",
// 		Functions: []graphite.Function{
// 			{
// 				Name:       "absolute",
// 				Arguments:  []string{"os.cpu"},
// 				Parameters: make(map[string]string),
// 			},
// 		},
// 	},
// 	{
// 		Query: "absolute(os.cpu, 'a.time.serie')",
// 		Functions: []graphite.Function{
// 			{
// 				Name:       "absolute",
// 				Arguments:  []string{"os.cpu", "a.time.serie"},
// 				Parameters: make(map[string]string),
// 			},
// 		},
// 	},
// 	{
// 		Query: "absolute(maximumAbove(os.cpu, 5), 'a.time.serie')",
// 		Functions: []graphite.Function{
// 			{
// 				Name:       "maximumAbove",
// 				Arguments:  []string{"os.cpu", "5"},
// 				Parameters: make(map[string]string),
// 			},
// 			{
// 				Name:       "absolute",
// 				Arguments:  []string{swap, "a.time.serie"},
// 				Parameters: make(map[string]string),
// 			},
// 		},
// 	},

// 	{
// 		Query: "absolute(divideSeries(os.cpu, os.net), 'a.time.serie')",
// 		Functions: []graphite.Function{
// 			{
// 				Name:       "divideSeries",
// 				Arguments:  []string{"os.cpu", "os.net"},
// 				Parameters: make(map[string]string),
// 			},
// 			{
// 				Name:       "absolute",
// 				Arguments:  []string{swap, "a.time.serie"},
// 				Parameters: make(map[string]string),
// 			},
// 		},
// 	},
// 	{
// 		Query: "reduceSeries(divideSeries(os.cpu, os.net), maximumAbove(os.cpu.used, 5), 'sum', 'a.time.serie')",
// 		Functions: []graphite.Function{
// 			{
// 				Name:       "divideSeries",
// 				Arguments:  []string{"os.cpu", "os.net"},
// 				Parameters: make(map[string]string),
// 			},
// 			{
// 				Name:       "maximumAbove",
// 				Arguments:  []string{"os.cpu.used", "5"},
// 				Parameters: make(map[string]string),
// 			},
// 			{
// 				Name:       "reduceSeries",
// 				Arguments:  []string{swap, swap, "sum", "a.time.serie"},
// 				Parameters: make(map[string]string),
// 			},
// 		},
// 	},
// 	{
// 		Query: "reduceSeries(divideSeries(os.cpu.*, os.net), maximumAbove(os.*.used, 5), 'sum', 'a.time.serie')",
// 		Functions: []graphite.Function{
// 			{
// 				Name:       "divideSeries",
// 				Arguments:  []string{"os.cpu.*", "os.net"},
// 				Parameters: make(map[string]string),
// 			},
// 			{
// 				Name:       "maximumAbove",
// 				Arguments:  []string{"os.*.used", "5"},
// 				Parameters: make(map[string]string),
// 			},
// 			{
// 				Name:       "reduceSeries",
// 				Arguments:  []string{swap, swap, "sum", "a.time.serie"},
// 				Parameters: make(map[string]string),
// 			},
// 		},
// 	},
// 	{
// 		Query: "reduceSeries(divideSeries(os.cpu.*, os.net.[0-9]), maximumAbove(os.*.used, 5), os.cpu[0-9].used, 'a.time.serie')",
// 		Functions: []graphite.Function{
// 			{
// 				Name:       "divideSeries",
// 				Arguments:  []string{"os.cpu.*", "os.net.[0-9]"},
// 				Parameters: make(map[string]string),
// 			},
// 			{
// 				Name:       "maximumAbove",
// 				Arguments:  []string{"os.*.used", "5"},
// 				Parameters: make(map[string]string),
// 			},
// 			{
// 				Name:       "reduceSeries",
// 				Arguments:  []string{swap, swap, "os.cpu[0-9].used", "a.time.serie"},
// 				Parameters: make(map[string]string),
// 			},
// 		},
// 	},
// 	{
// 		Query: "divideSeries(os.cpu.*, os.net.{dsd,top})",
// 		Functions: []graphite.Function{
// 			{
// 				Name:       "divideSeries",
// 				Arguments:  []string{"os.cpu.*", "os.net.(dsd|top)"},
// 				Parameters: make(map[string]string),
// 			},
// 		},
// 	},
// 	{
// 		Query: "divideSeries(os.cpu.*, os.net[s-a].a{dsd,top}ez[dsd])",
// 		Functions: []graphite.Function{
// 			{
// 				Name:       "divideSeries",
// 				Arguments:  []string{"os.cpu.*", "os.net[s-a].a(dsd|top)ez[dsd]"},
// 				Parameters: make(map[string]string),
// 			},
// 		},
// 	},
// 	{
// 		Query: "reduceSeries(divideSeries(os.cpu.*, os.disk.{s,m}d[0-9][0-9]), maximumAbove(os.*.used, 5), os.cpu[0-9].used, 'a.time.serie')",
// 		Functions: []graphite.Function{
// 			{
// 				Name:       "divideSeries",
// 				Arguments:  []string{"os.cpu.*", "os.disk.(s|m)d[0-9][0-9]"},
// 				Parameters: make(map[string]string),
// 			},
// 			{
// 				Name:       "maximumAbove",
// 				Arguments:  []string{"os.*.used", "5"},
// 				Parameters: make(map[string]string),
// 			},
// 			{
// 				Name:       "reduceSeries",
// 				Arguments:  []string{swap, swap, "os.cpu[0-9].used", "a.time.serie"},
// 				Parameters: make(map[string]string),
// 			},
// 		},
// 	},
// }
