package helpers

import (
	"time"
)


const (
	INPUT_FORMAT  = "2006-01-02T15:04:05.999999999-07:00"
	OUTPUT_FORMAT = "2006-01-02T15:04:05.000Z"
)

func TimestampToJavaScriptISO(s string) (time.Time, error) {
	t, err := time.Parse(INPUT_FORMAT, s)
	if err != nil {
	  return time.Now(), err
	}
	return t, err
}