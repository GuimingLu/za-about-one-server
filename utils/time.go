


package utils

import (
	"time"
)

func MillisFromTime(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func TimeFromMillis(millis int64) time.Time {
	return time.Unix(0, millis*int64(time.Millisecond))
}

func StartOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func EndOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 999999999, t.Location())
}

func Yesterday() time.Time {
	return time.Now().AddDate(0, 0, -1)
}

func TimeBeforeHours(hours string)  int64{
	now := time.Now()
	d, _ := time.ParseDuration(hours)
	d1 := now.Add(d)
	return MillisFromTime(d1)
}
