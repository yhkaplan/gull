package date

import (
	"time"
)

const dateFormat = "2006-01-02"

var timeZone = time.Local

func Parse(dateStr string) (time.Time, error) {
	t, err := time.ParseInLocation(dateFormat, dateStr, timeZone)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}
