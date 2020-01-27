package promremote

import (
	"fmt"

	"github.com/prometheus/prometheus/prompb"
)

func formatLabel(matchers []*prompb.LabelMatcher) (string, string) {
	s := "~.*"
	l := "{ "

	if matchers == nil {
		return s, l + "}"
	}

	for _, matcher := range matchers {
		if matcher == nil {
			continue
		}

		// Classname
		if matcher.GetName() == prometheusClassNameLabel {

			// RegExp
			if matcher.GetType() == prompb.LabelMatcher_RE {
				s = "~" + matcher.GetValue()
			} else if matcher.GetType() == prompb.LabelMatcher_NRE {
				s = fmt.Sprintf("~(?!%s)", matcher.GetValue())
			} else if matcher.GetType() == prompb.LabelMatcher_NEQ {
				s = fmt.Sprintf("~(?!%s)", matcher.GetValue())
			} else if matcher.GetType() == prompb.LabelMatcher_EQ {
				s = matcher.GetValue()
			}
			continue
		}

		// RegExp
		if matcher.GetType() == prompb.LabelMatcher_RE || matcher.GetType() == prompb.LabelMatcher_NRE {
			l += fmt.Sprintf("'%s' '~%s' ", matcher.GetName(), matcher.GetValue())
		} else {
			l += fmt.Sprintf("'%s' '%s' ", matcher.GetName(), matcher.GetValue())
		}

	}

	l += "}"

	log.Debugf("formatLabel: %+v => %s%s", matchers, s, l)
	return s, l
}
