package graphite_test

import (
	"strings"
	"testing"

	"github.com/ovh/erlenmeyer/proto/graphite"
)

type findRequestTest struct {
	Serie          string
	Wildcards      bool
	ShouldContains []string
}

func TestCreateFindRequest(t *testing.T) {
	for _, request := range findRequestsTest {
		tree, err := graphite.CreateFindRequest(request.Serie, request.Wildcards)
		if err != nil {
			t.Error(err)

			continue
		}

		ws := tree.ToWarpScript("", "", "")

		for _, shouldContain := range request.ShouldContains {
			if !strings.Contains(ws, shouldContain) {
				t.Errorf("Query does not contain '%s'", shouldContain)
				t.Errorf("WarpScript %s", ws)
				t.Fail()
			}
		}
	}
}

var findRequestsTest = []findRequestTest{
	{
		Serie:     "os.cpu",
		Wildcards: true,
		ShouldContains: []string{
			"[ $token '~os\\.cpu\\..*?' {} ] FIND",
		},
	},
	{
		Serie:     "os.cpu.*",
		Wildcards: false,
		ShouldContains: []string{
			"[ $token '~os\\.cpu\\..*?' {} ] FIND",
		},
	},

	{
		Serie:     "os.*.used",
		Wildcards: true,
		ShouldContains: []string{
			"[ $token '~os\\..*?\\.used\\..*?' {} ] FIND",
		},
	},

	{
		Serie:     "os.*.used.*",
		Wildcards: false,
		ShouldContains: []string{
			"[ $token '~os\\..*?\\.used\\..*?' {} ] FIND",
		},
	},
	{
		Serie:     "*",
		Wildcards: true,
		ShouldContains: []string{
			"[ $token '~.*?\\..*?' {} ] FIND",
		},
	},
	{
		Serie:     "*.*",
		Wildcards: false,
		ShouldContains: []string{
			"[ $token '~.*?\\..*?' {} ] FIND",
		},
	},
}
