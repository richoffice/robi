package robi

import "time"

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
