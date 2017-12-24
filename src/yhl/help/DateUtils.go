package help

import (
	"time"
)

const (
	DateFormat        string = "2006-01-02"
	DatetimeFormat    string = "2006-01-02 15:04:05"
	DatetimeNumFormat string = "20060102150405"
)

func Date() (year, month, day int) {
	t := time.Now()
	year, m, day := t.Date()
	month = int(m)

	return

}
