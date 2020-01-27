package opentsdb

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var timeFormats = []string{
	"2006/01/02-15:04:05",
	"2006/01/02 15:04:05",
	"2006/01/02-15:04",
	"2006/01/02 15:04",
	"2006/01/02",
}

var durationUnits = map[string]int64{
	"ms":  int64(time.Millisecond),
	"s":   int64(time.Second),
	"m":   int64(time.Minute),
	"h":   int64(time.Hour),
	"d":   int64(24 * time.Hour),
	"w":   int64(7 * 24 * time.Hour),
	"n":   int64(30 * 24 * time.Hour),
	"y":   int64(365 * 24 * time.Hour),
	"all": int64(0),
}

var durationPattern = "([0-9]+)(ms|s|m|h|d|w|n|y|all)"

// Duration to parse is with commas
var durationRe, _ = regexp.Compile("^" + durationPattern + "-ago$")

// seconds optionally followed by ms on 3 digits
var msecsRe, _ = regexp.Compile(`^([0-9]+)(\.([0-9]{3}))?$`)

// DecodeDuration convert a ts + unit into Duration
func DecodeDuration(value string, unit string) time.Duration {
	// Expect the params to be valid, since they've passed the duration pattern above
	n, _ := strconv.ParseInt(value, 10, 64) // nolint: gas
	return time.Duration(n * durationUnits[unit])
}

// ParseOpenTsdbTimeStamp parse OpenTSDB times
func ParseOpenTsdbTimeStamp(str string) int64 {
	pos := strings.IndexRune(str, '.')
	if pos == -1 {
		tstamp, _ := strconv.ParseInt(str, 10, 64) // nolint: gas
		return int64toMicroSec(tstamp) / 1000
	}

	secs, _ := strconv.ParseInt(str[:pos], 10, 64)    // nolint: gas
	msecs, _ := strconv.ParseInt(str[pos+1:], 10, 64) // nolint: gas
	return secs*1000 + msecs
}

//-------------------------------------------------------------------------------------------------
// OpenTSDB format can be either absolute (timestamp) or relative (duration)

// TSDBTime is overloading json/schema time parsing to support opentsdb format
type TSDBTime struct {
	time.Time
}

// UnmarshalJSON is a Custom json parser
func (t *TSDBTime) UnmarshalJSON(b []byte) error {
	return t.Parse(string(b))
}

// Convert an int expressed either in seconds or milliseconds into microseconds
func int64toMicroSec(timestamp int64) int64 {
	const microsPerSec = 1000000
	const microsPerMilli = 1000

	timeMicros := timestamp
	if timeMicros < 0xFFFFFFFF {
		// If less than 2^32, assume it's in seconds
		// (in millis that would be Thu Feb 19 18:02:47 CET 1970)
		timeMicros *= microsPerSec
	} else {
		timeMicros *= microsPerMilli
	}

	return timeMicros
}

// Parse a string representing a date
// See http://opentsdb.net/docs/build/html/user_guide/query/dates.html
func (t *TSDBTime) Parse(data string) error {
	// for relative time, we need to remove quotes
	if strings.ContainsAny(data, "\"") {
		// max is 2 replacements
		data = strings.Replace(data, "\"", "", 2)
	}
	splits := durationRe.FindStringSubmatch(data)
	if splits != nil {
		// Relative time
		t.Time = time.Now().Add(-DecodeDuration(splits[1], splits[2]))
		return nil
	}
	// Try integer
	if msecsRe.MatchString(data) {
		timeStamp := ParseOpenTsdbTimeStamp(data)
		t.Time = time.Unix(timeStamp/1000, timeStamp%1000)
		return nil
	}
	// Try the various time formats
	for _, fmt := range timeFormats {
		date, err := time.Parse(fmt, data)
		if err == nil {
			// Conversion succeeded
			t.Time = date
			return nil
		}
	}
	// None matched
	return errors.New("Invalid date: '" + data + "'")
}
