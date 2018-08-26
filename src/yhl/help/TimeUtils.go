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

func TimeToStr(t time.Time, f string) string {
	return t.Format(f)
}

func Date() (year, month, day int) {
	t := time.Now()
	year, m, day := t.Date()
	month = int(m)

	return

}

func GetDateBegin(t time.Time) time.Time {
	y, m, d := t.Date()

	return time.Date(y, m, d, 0, 0, 0, 0, time.Local)
}

func GetDateEnd(t time.Time) time.Time {
	y, m, d := t.Date()

	return time.Date(y, m, d, 24, 0, 0, 0, time.Local)
}

func GetMonthBegin(t time.Time) time.Time {
	y, m, _ := t.Date()

	//return time.Date(y, m, 1, 0, 0, 0, 0, time.UTC)
	return time.Date(y, m, 1, 0, 0, 0, 0, time.Local)
}

func GetMonthEnd(t time.Time) time.Time {
	return GetMonthBegin(t).AddDate(0, 1, 0)
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

func LocalTime(t time.Time) time.Time {
	//return t.Local()
	local, _ := time.LoadLocation("Local")

	return t.In(local)
}
