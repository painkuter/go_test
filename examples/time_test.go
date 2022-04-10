package examples

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	DateTimeLayout    = "2006-01-02 15:04"
	DateTimeSecLayout = "2006-01-02 15:04:05"
)

func TestFromExt(t *testing.T) {
	expected := toTime(63780807540)
	fmt.Println(expected.String())
	actual := MustParseDatetime("2022-02-18 18:59:00")
	fmt.Println(actual.String())
	assert.Equal(t, expected, actual)
}

// time.Time.ext to time.Time
func toTime(n int64) time.Time {
	n = n - (1969*365+1969/4-1969/100+1969/400)*24*60*60
	return time.Date(1970, 1, 1, 0, 0, int(n), 0, time.UTC)
}

func TestTimeZone(t *testing.T) {
	testTime := MustParseDatetime("2022-02-20 19:00:00")
	fmt.Println(testTime.String())
	fmt.Println(testTime.Year(), testTime.Month(), testTime.Day())
	loc := time.FixedZone("test zone", 18000)
	testTimeLoc := testTime.In(loc)
	fmt.Println(testTimeLoc.String())
	fmt.Println(testTimeLoc.Year(), testTimeLoc.Month(), testTimeLoc.Day())
}

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

func MustParse(layout, t string) time.Time {
	result, err := time.Parse(layout, t)
	if err != nil {
		panic(err)
	}

	return result
}
