package promremote

import (
	"fmt"

	"github.com/prometheus/prometheus/prompb"
)

type bucketize struct {
	Fn    string
	End   int64
	Span  int64
	Count int64
}

func (b *bucketize) String() string {
	return fmt.Sprintf("[ SWAP bucketizer.%s %d %d %d ] BUCKETIZE", b.Fn, b.End, b.Span, b.Count)
}

func getBucketizer(rh *prompb.ReadHints, nowMs int64) string {
	b := &bucketize{
		Fn:    rh.GetFunc(),
		Span:  rh.GetStepMs() * 1000,
		End:   nowMs * 1000,
		Count: 0,
	}

	return b.String()
}
