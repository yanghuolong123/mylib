package help

import (
	"fmt"
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

func ShowTime(t time.Time) (s string) {
	now := time.Now()
	if t.After(now) {
		return fmt.Sprintf("%d秒前", 1)
	}

	du := time.Since(t) / time.Second

	y := time.Duration(365 * 24 * 60 * 60)
	if du > y {
		return fmt.Sprintf("%d年前", du/y)
	}

	m := time.Duration(30 * 24 * 60 * 60)
	if du > m {
		return fmt.Sprintf("%d月前", du/m)
	}

	d := time.Duration(24 * 60 * 60)
	if du > d {
		return fmt.Sprintf("%d天前", du/d)
	}

	h := time.Duration(60 * 60)
	if du > h {
		return fmt.Sprintf("%d小时前", du/h)
	}

	min := time.Duration(60)
	if du > min {
		return fmt.Sprintf("%d分钟前", du/min)
	}

	if du <= 0 {
		return "1秒前"
	}

	return fmt.Sprintf("%d秒前", du)
}
