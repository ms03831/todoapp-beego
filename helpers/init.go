package helpers

import (
	"time"
)


const (
	INPUT_FORMAT  = "2006-01-02T15:04:05.999999999-07:00"
	OUTPUT_FORMAT = "2006-01-02T15:04:05.000Z"
)

func TimestampToJavaScriptISO(s string) (string, error) {
	t, err := time.Parse(INPUT_FORMAT, s)
	if err != nil {
	  return "", err
	}
	return t.UTC().Format(OUTPUT_FORMAT), nil
}