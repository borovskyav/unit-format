package unit

import (
	"fmt"
	"math"
)

type Time float64

const (
	Nanosecond  Time = 1
	Microsecond      = Nanosecond * 1000
	Millisecond      = Microsecond * 1000
	Second           = Millisecond * 1000
	Minute           = Second * 60
	Hour             = Minute * 60
	Day              = Hour * 24
	Week             = Day * 7
	Year             = Day * 365
)

var timeUnitTable = map[Time]func(value Time) string{
	Nanosecond:  func(value Time) string { return "ns" },
	Microsecond: func(value Time) string { return "µs" },
	Millisecond: func(value Time) string { return "ms" },
	Second:      func(value Time) string { return "sec" },
	Minute:      func(value Time) string { return "min" },
	Hour:        func(value Time) string { return declensions(float64(value), "hour", "hours", "hours") },
	Day:         func(value Time) string { return declensions(float64(value), "day", "days", "days") },
	Week:        func(value Time) string { return declensions(float64(value), "week", "weeks", "weeks") },
	Year:        func(value Time) string { return declensions(float64(value), "year", "years", "years") },
}

func Nanoseconds(value float64) string {
	return formatTimeValue(value, Nanosecond)
}

func Microseconds(value float64) string {
	return formatTimeValue(value, Microsecond)
}

func Milliseconds(value float64) string {
	return formatTimeValue(value, Millisecond)
}

func Seconds(value float64) string {
	return formatTimeValue(value, Second)
}

func formatTimeValue(value float64, time Time) string {
	ns := Time(math.Abs(value * float64(time)))
	switch {
	case ns < Microsecond:
		return formatValueWithUnit(ns, Nanosecond, 3)
	case ns < Millisecond:
		return formatValueWithUnit(ns, Microsecond, 3)
	case ns < Second:
		return formatValueWithUnit(ns, Millisecond, 3)
	case ns < Minute:
		return formatValueWithUnit(ns, Second, 3)
	case ns < Hour:
		return formatValueWithUnit(ns, Minute, 2)
	case ns < Day:
		return formatValueWithUnit(ns, Hour, 2)
	case ns < Week:
		return formatValueWithUnit(ns, Day, 2)
	case ns < Year:
		return formatValueWithUnit(ns, Week, 2)
	default:
		return formatValueWithUnit(ns, Year, 2)
	}

}

func formatValueWithUnit(value Time, unit Time, decimals int) string {
	return fmt.Sprintf("%v %s", Round(float64(value/unit), decimals), timeUnitTable[unit](value))
}

func Round(value float64, decimals int) float64 {
	multiplier := math.Pow10(decimals)
	return math.Round(value*multiplier) / multiplier
}
