package promremote

import (
	"fmt"

	"github.com/prometheus/prometheus/prompb"
)

type mapper struct {
	Fn    string
	Pre   int64
	Post  int64
	Count int64
}

func (m *mapper) String() string {
	return fmt.Sprintf("[ SWAP mapper.%s %d %d %d ] MAP", m.Fn, m.Pre, m.Post, m.Count)
}

func getMapper(fn string, rh *prompb.ReadHints, nowMs int64) string {
	b := &mapper{
		Fn:  fn,
		Pre: 0,
	}

	if fn == "rate" {
		b.Pre = 1
	}

	return b.String()
}
