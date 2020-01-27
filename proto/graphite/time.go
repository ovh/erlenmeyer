package graphite

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Jamesxql/at"
	"github.com/ovh/erlenmeyer/core"
)

// ----------------------------------------------------------------------------
// helpers

// see definition of time here
// http://readthedocs.io/en/latest/render_api.html#from-until

// manage relative time
var (
	relativeTime       = regexp.MustCompile(`^-(\d+)(s|min|h|d|w|mon|y)$`)
	relativeTimeMapper = map[string]time.Duration{
		"s":   time.Second,
		"min": time.Minute,
		"h":   time.Hour,
		"d":   time.Hour * 24,
		"w":   time.Hour * 24 * 7,
		"mon": time.Hour * 24 * 30,
		"y":   time.Hour * 24 * 365,
	}

	convert    = regexp.MustCompile(`^(\d+\.?\d*)(y|w|mon|d|h|min|s)$`)
	unitMapper = map[string]string{
		"y":   "365 d *",
		"w":   "w",
		"mon": "30 d *",
		"d":   "d",
		"h":   "h",
		"min": "m",
		"s":   "s",
	}
)

// manage absolute time
var (
	absoluteTime        = regexp.MustCompile(`^(\d{2}):(\d{2})_(\d{4})(\d{2})(\d{2})|(\d{4})(\d{2})(\d{2})|(\d{2})/(\d{2})/(\d{2})$`)
	absoluteTimeFormat1 = regexp.MustCompile(`^(\d{2}):(\d{2})_(\d{4})(\d{2})(\d{2})$`)
	absoluteTimeFormat2 = regexp.MustCompile(`^(\d{4})(\d{2})(\d{2})$`)
	absoluteTimeFormat3 = regexp.MustCompile(`^(\d{2})/(\d{2})/(\d{2})$`)
	timeStamp           = regexp.MustCompile(`^\d+$`)
)

const (
	now = "now"
)

// ParseTime provided in the query
func ParseTime(t []byte) (*time.Time, error) {
	if string(t) == now {
		t = []byte("-0s")
	}

	if relativeTime.Match(t) {
		return parseRelativeTime(t)
	}

	if absoluteTime.Match(t) {
		return parseAbsoluteTime(t)
	}

	return parseAtTime(t)
}

// https://github.com/Jamesxql/at
func parseAtTime(t []byte) (*time.Time, error) {
	schedule, err := at.Parse(string(t))
	if err != nil {
		return nil, err
	}

	rt := schedule.At(time.Now())

	return &rt, nil
}

func parseRelativeTime(t []byte) (*time.Time, error) {
	matches := relativeTime.FindSubmatch(t)[1:]
	number, err := strconv.Atoi(string(matches[0]))
	if err != nil {
		return nil, err
	}

	duration, ok := relativeTimeMapper[string(matches[1])]
	if !ok {
		return nil, errors.New("Time unit does not exist")
	}

	rt := time.Now().Add(-1 * time.Duration(number) * duration)

	return &rt, nil
}

func parseAbsoluteTime(t []byte) (*time.Time, error) {
	if absoluteTimeFormat1.Match(t) {
		return parseAbsoluteTimeFormat1(t)
	}

	if absoluteTimeFormat2.Match(t) {
		return parseAbsoluteTimeFormat2(t)
	}

	if absoluteTimeFormat3.Match(t) {
		return parseAbsoluteTimeFormat3(t)
	}

	if timeStamp.Match(t) {
		return parseTimeStamp(t)
	}

	return nil, errors.New("Time format is not supported")
}

// parse format HH:MM_YYYYMMDD
func parseAbsoluteTimeFormat1(t []byte) (*time.Time, error) {
	location, _ := time.LoadLocation("Local") // nolint: gas
	rt, err := time.ParseInLocation("15:04_20060102", string(t), location)
	if err != nil {
		return nil, err
	}

	return &rt, nil
}

// parse format YYYYMMDD
func parseAbsoluteTimeFormat2(t []byte) (*time.Time, error) {
	location, _ := time.LoadLocation("Local") // nolint: gas
	rt, err := time.ParseInLocation("20060102", string(t), location)
	if err != nil {
		return nil, err
	}

	return &rt, nil
}

// parse format MM/DD/YY
func parseAbsoluteTimeFormat3(t []byte) (*time.Time, error) {
	location, _ := time.LoadLocation("Local") // nolint: gas
	rt, err := time.ParseInLocation("01/02/06", string(t), location)
	if err != nil {
		return nil, err
	}

	return &rt, nil
}

func parseTimeStamp(t []byte) (*time.Time, error) {
	ts, err := strconv.ParseInt(string(t), 10, 64)
	if err != nil {
		return nil, err
	}

	rt := time.Unix(ts, 0)

	return &rt, nil
}

func convertTimeToWarpScript(t string) (*string, error) {
	if !convert.MatchString(t) {
		return nil, fmt.Errorf("Date %s has not the correct format", t)
	}

	matches := convert.FindStringSubmatch(t)[1:]
	wsTime := fmt.Sprintf("%s %s", matches[0], unitMapper[matches[1]])

	return &wsTime, nil
}

// ----------------------------------------------------------------------------
// graphite functions implementations

func delay(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 2 {
		return nil, errors.New("The delay take two parameter which are a series or a list of series and a number")
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: fmt.Sprintf("%s TIMESHIFT", args[1]),
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func timeShift(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 2 {
		return nil, errors.New("The timeShift function take two parameters which are a list of series and a time")
	}

	sign := "-"
	if strings.HasPrefix(args[1], "+") {
		sign = "+"
		args[1] = args[1][1:]
	} else if strings.HasPrefix(args[1], "-") {
		args[1] = args[1][1:]
	}

	t, err := convertTimeToWarpScript(args[1])
	if err != nil {
		return nil, err
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: fmt.Sprintf("%s %s1 * TIMESHIFT", *t, sign),
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func timeSlice(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 2 {
		return nil, errors.New("The timeSlice function take at least two parameters which are a list of series and a date")
	}

	start := args[1]
	from, err := ParseTime([]byte(start))
	if err != nil {
		return nil, err
	}

	end := now
	if len(args) >= 3 {
		end = args[2]
	}

	until, err := ParseTime([]byte(end))
	if err != nil {
		return nil, err
	}

	start = strconv.FormatInt(from.UnixNano()/1000, 10)
	end = strconv.FormatInt(until.UnixNano()/1000, 10)
	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: fmt.Sprintf(`%s ISO8601 %s ISO8601 TIMECLIP`, end, start),
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}
