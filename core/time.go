package core

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/common/model"
)

// Most of this code is taken from original time.go in Prometheus repo
// https://github.com/prometheus/common/blob/master/model/time.go
const (
	// MinimumTick is the minimum supported time resolution. This has to be
	// at least time.Second in order for the code below to work.
	minimumTick = time.Millisecond
	// second is the Time duration equivalent to one second.
	second = int64(time.Second / minimumTick)
	// The number of nanoseconds per minimum tick.
	nanosPerTick = int64(minimumTick / time.Nanosecond)

	// The number of microseconds per minimum tick
	microsPerTick = int64(minimumTick / time.Microsecond)

	// Earliest is the earliest Time representable. Handy for
	// initializing a high watermark.
	Earliest = model.Time(math.MinInt64)
	// Latest is the latest Time representable. Handy for initializing
	// a low watermark.
	Latest = model.Time(math.MaxInt64)
)

// Time is the number of milliseconds since the epoch
// (1970-01-01 00:00 UTC) excluding leap seconds.
type Time int64

// Interval describes and interval between two timestamps.
type Interval struct {
	Start, End Time
}

// ParsePromTime returns a core.Time from string
func ParsePromTime(s string) (Time, error) {
	if t, err := strconv.ParseFloat(s, 64); err == nil {
		ts := t * float64(time.Second)
		if ts > float64(math.MaxInt64) || ts < float64(math.MinInt64) {
			return 0, fmt.Errorf("cannot parse %q to a valid timestamp. It overflows int64", s)
		}
		return TimeFromUnixMicro(int64(ts)), nil
	}
	if t, err := time.Parse(time.RFC3339Nano, s); err == nil {
		return TimeFromUnixMicro(t.UnixNano()), nil
	}
	return 0, fmt.Errorf("cannot parse %q to a valid timestamp", s)
}

// ParsePromDuration parses a duration as a string and returns a string.
func ParsePromDuration(s string) (string, error) {
	if d, err := strconv.ParseFloat(s, 64); err == nil {
		ts := d * float64(time.Second)
		if ts > float64(math.MaxInt64) || ts < float64(math.MinInt64) {
			return "", fmt.Errorf("cannot parse %q to a valid duration. It overflows int64", s)
		}
		return s + " s", nil
	}
	if d, err := ParseDuration(s); err == nil {
		return d, nil
	}
	return "", fmt.Errorf("cannot parse %q to a valid duration", s)
}

// Now returns the current time as a Time.
func Now() Time {
	return TimeFromUnixMicro(time.Now().UnixNano())
}

// Before reports whether the Time t is before o.
func (t Time) Before(o Time) bool {
	return t < o
}

// TimeFromUnix returns the Time equivalent to the Unixmodel.Timet
// provided in seconds.
func TimeFromUnix(t int64) Time {
	return Time(t * second)
}

// TimeFromUnixMicro returns the Time equivalent to the Unix Time
// t provided in nanoseconds.
func TimeFromUnixMicro(t int64) Time {
	curTime := t / microsPerTick
	round := 0
	if (curTime % 1000) >= 500 {
		round = 1
	}
	curTime = curTime / 1000
	curTime = curTime + int64(round)
	return Time(curTime * 1000)
}

var durationRE = regexp.MustCompile("^([0-9]+)(y|w|d|h|m|s|ms)$")

// ParseDuration parses a string into a time.Duration, assuming that a year
// always has 365d, a week always has 7d, and a day always has 24h.
func ParseDuration(durationStr string) (string, error) {
	matches := durationRE.FindStringSubmatch(durationStr)
	if len(matches) != 3 {
		return "", fmt.Errorf("not a valid duration string: %q", durationStr)
	}
	return matches[1] + " " + matches[2], nil
}

// TimeFromUnixNano returns the Time equivalent to the Unix Time
// t provided in nanoseconds.
func TimeFromUnixNano(t int64) Time {
	return Time(t / nanosPerTick)
}

// Equal reports whether two Times represent the same instant.
func (t Time) Equal(o Time) bool {
	return t == o
}

// After reports whether the Time t is after o.
func (t Time) After(o Time) bool {
	return t > o
}

// Add returns the Time t + d.
func (t Time) Add(d time.Duration) Time {
	return t + Time(d/minimumTick)
}

// Sub returns the Duration t - o.
func (t Time) Sub(o Time) time.Duration {
	return time.Duration(t-o) * minimumTick
}

// Time returns the time.Time representation of t.
func (t Time) Time() time.Time {
	return time.Unix(int64(t)/second, (int64(t)%second)*nanosPerTick)
}

// Unix returns t as a Unix time, the number of seconds elapsed
// since January 1, 1970 UTC.
func (t Time) Unix() int64 {
	return int64(t) / second
}

// UnixNano returns t as a Unix time, the number of nanoseconds elapsed
// since January 1, 1970 UTC.
func (t Time) UnixNano() int64 {
	return int64(t) * nanosPerTick
}

// The number of digits after the dot.
var dotPrecision = int(math.Log10(float64(second)))

// String returns a string representation of the Time.
func (t Time) String() string {
	return strconv.FormatFloat(float64(t)/float64(second), 'f', -1, 64)
}

// MarshalJSON implements the json.Marshaler interface.
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(t.String()), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *Time) UnmarshalJSON(b []byte) error {
	p := strings.Split(string(b), ".")
	switch len(p) {
	case 1:
		v, err := strconv.ParseInt(string(p[0]), 10, 64)
		if err != nil {
			return err
		}
		*t = Time(v * second)

	case 2:
		v, err := strconv.ParseInt(string(p[0]), 10, 64)
		if err != nil {
			return err
		}
		v *= second

		prec := dotPrecision - len(p[1])
		if prec < 0 {
			p[1] = p[1][:dotPrecision]
		} else if prec > 0 {
			p[1] = p[1] + strings.Repeat("0", prec)
		}

		va, err := strconv.ParseInt(p[1], 10, 32)
		if err != nil {
			return err
		}

		*t = Time(v + va)

	default:
		return fmt.Errorf("invalid time %q", string(b))
	}
	return nil
}

// Duration wraps time.Duration. It is used to parse the custom duration format
// from YAML.
// This type should not propagate beyond the scope of input/output processing.
type Duration time.Duration

// Set implements pflag/flag.Value
func (d *Duration) Set(s string) error {
	var err error
	*d, err = ParseStringDuration(s)
	return err
}

// Type implements pflag.Value
func (d *Duration) Type() string {
	return "duration"
}

// ParseStringDuration parses a string into a time.Duration, assuming that a year
// always has 365d, a week always has 7d, and a day always has 24h.
func ParseStringDuration(durationStr string) (Duration, error) {
	matches := durationRE.FindStringSubmatch(durationStr)
	if len(matches) != 3 {
		return 0, fmt.Errorf("not a valid duration string: %q", durationStr)
	}
	var (
		n, _ = strconv.Atoi(matches[1])
		dur  = time.Duration(n) * time.Millisecond
	)
	switch unit := matches[2]; unit {
	case "y":
		dur *= 1000 * 60 * 60 * 24 * 365
	case "w":
		dur *= 1000 * 60 * 60 * 24 * 7
	case "d":
		dur *= 1000 * 60 * 60 * 24
	case "h":
		dur *= 1000 * 60 * 60
	case "m":
		dur *= 1000 * 60
	case "s":
		dur *= 1000
	case "ms":
		// Value already correct
	default:
		return 0, fmt.Errorf("invalid time unit in duration string: %q", unit)
	}
	return Duration(dur), nil
}

func (d Duration) String() string {
	var (
		ms   = int64(time.Duration(d) / time.Millisecond)
		unit = "ms"
	)
	factors := map[string]int64{
		"y":  1000 * 60 * 60 * 24 * 365,
		"w":  1000 * 60 * 60 * 24 * 7,
		"d":  1000 * 60 * 60 * 24,
		"h":  1000 * 60 * 60,
		"m":  1000 * 60,
		"s":  1000,
		"ms": 1,
	}

	switch int64(0) {
	case ms % factors["y"]:
		unit = "y"
	case ms % factors["w"]:
		unit = "w"
	case ms % factors["d"]:
		unit = "d"
	case ms % factors["h"]:
		unit = "h"
	case ms % factors["m"]:
		unit = "m"
	case ms % factors["s"]:
		unit = "s"
	}
	return fmt.Sprintf("%v%v", ms/factors[unit], unit)
}

// MarshalYAML implements the yaml.Marshaler interface.
func (d Duration) MarshalYAML() (interface{}, error) {
	return d.String(), nil
}

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (d *Duration) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}
	dur, err := ParseStringDuration(s)
	if err != nil {
		return err
	}
	*d = dur
	return nil
}
