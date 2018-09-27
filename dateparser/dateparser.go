package dateparser

import (
	"time"
)

const dateFormat = "2006-01-02"

func Parse(dateStr string) (time.Time, error) {
	t, err := time.Parse(dateFormat, dateStr)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}
