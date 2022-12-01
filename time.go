package robi

import (
	"strconv"
	"time"
)

type Timer struct {
}

func (t *Timer) ParseDate(dateStr string) interface{} {

	layout := "2006-01-02"
	tt, err := time.Parse(layout, dateStr)
	if err != nil {
		return err
	}

	return tt
}

func (t *Timer) DateEndTime(tt time.Time) interface{} {
	return time.Date(tt.Year(), tt.Month(), tt.Day()+1, 0, 0, 0, 0, tt.Location())
}

func (t *Timer) FormatDate(tt time.Time) string {
	return tt.Format("2006 年 01 月 02 日")
}

func (t *Timer) DateStartTime(tt time.Time) interface{} {
	return time.Date(tt.Year(), tt.Month(), tt.Day(), 0, 0, 0, 0, tt.Location())
}

func (t *Timer) FromUnix(sec int64) interface{} {
	return time.Unix(sec, 0)
}

func (t *Timer) FromUnixStr(sec string) interface{} {
	i, err := strconv.ParseInt(sec, 10, 64)
	if err != nil {
		return err
	}
	return time.Unix(i, 0)
}

func (t *Timer) GetDateStr(tt time.Time, format string) string {
	return tt.Format(format)
}

func (t *Timer) ParseDateStr(dateStr, format string) interface{} {
	tt, err := time.Parse(format, dateStr)
	if err != nil {
		return err
	}

	return tt
}

func (t *Timer) Yesterday(tt time.Time) interface{} {
	return tt.Add(-24 * time.Hour)
}
