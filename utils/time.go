package utils

import (
	"time"
)

const (
	// DateTimeLayout is usual for tests date format.
	DateTimeLayout = "2006-01-02 15:04"
	// DateTimeSecLayout is same as DateTimeLayout but with seconds.
	DateTimeSecLayout = "2006-01-02 15:04:05"
)

// MustParse panic if template not valid.
func MustParse(layout, t string) time.Time {
	result, err := time.Parse(layout, t)
	if err != nil {
		panic(err)
	}

	return result
}

// MustParseDatetime parse time with DateTimeLayout or DateTimeSecLayout template.
// Panic when time.Parse returns error.
// MustParseDatetime("2020-01-01 14:00")
func MustParseDatetime(t string) time.Time {
	timeLen := len(t)
	if timeLen <= len(DateTimeLayout) {
		return MustParse(DateTimeLayout, t)
	} else if timeLen <= len(DateTimeSecLayout) {
		return MustParse(DateTimeSecLayout, t)
	} else {
		panic("invalid template string")
	}
}
