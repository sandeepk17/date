package date

import (
	"fmt"
	"time"
)

// Date ...
type Date struct {
	datetime interface{}
}

// DateHandler ...
type DateHandler func(d *Date)

// DateFormat ...
const DateFormat = "2006-01-02"

// DateTimeFormat ...
const DateTimeFormat = "2006-01-02 15:04:05"

// NewDate ...
func NewDate(opts ...DateHandler) *Date {
	var d = &Date{}
	for _, o := range opts {
		o(d)
	}
	return d
}

// BindDate ...
func BindDate(dateStr string) DateHandler {
	return func(d *Date) {
		d.datetime, _ = time.Parse(DateFormat, dateStr)
	}
}

// BindDateTime ...
func BindDateTime(datetimeStr string) DateHandler {
	return func(d *Date) {
		d.datetime, _ = time.Parse(DateTimeFormat, datetimeStr)
	}
}

// TodayDate ...
func (d *Date) TodayDate() string {
	return d.now().Format(DateFormat)
}

// TodayDateTime ...
func (d *Date) TodayDateTime() string {
	return d.now().Format(DateTimeFormat)
}

// TodayStartDateTime ...
func (d *Date) TodayStartDateTime() string {
	return fmt.Sprintf("%s 00:00:00", d.TodayDate())
}

// TodayEndDateTime ...
func (d *Date) TodayEndDateTime() string {
	return fmt.Sprintf("%s 23:59:59", d.TodayDate())
}

// YesterdayDate ...
func (d *Date) YesterdayDate() string {
	return d.now().AddDate(0, 0, -1).Format(DateFormat)
}

// YesterdayStartDateTime ...
func (d *Date) YesterdayStartDateTime() string {
	return fmt.Sprintf("%s 00:00:00", d.YesterdayDate())
}

// YesterdayEndDateTime ...
func (d *Date) YesterdayEndDateTime() string {
	return fmt.Sprintf("%s 23:59:59", d.YesterdayDate())
}

// WeekStartDate ...
func (d *Date) WeekStartDate() string {
	return d.weekStart().Format(DateFormat)
}

// LastWeekStartDate ...
func (d *Date) LastWeekStartDate() string {
	return d.weekStart().AddDate(0, 0, -7).Format(DateFormat)
}

// LastWeekEndDate ...
func (d *Date) LastWeekEndDate() string {
	return d.weekStart().AddDate(0, 0, -1).Format(DateFormat)
}

// MonthStartDate ...
func (d *Date) MonthStartDate() string {
	return d.monthStart().Format(DateFormat)
}

// LastMonthStartDate ...
func (d *Date) LastMonthStartDate() string {
	return d.monthStart().AddDate(0, -1, 0).Format(DateFormat)
}

// LastMonthEndDate ...
func (d *Date) LastMonthEndDate() string {
	return d.monthStart().AddDate(0, 0, -1).Format(DateFormat)
}

// YearStartDate ...
func (d *Date) YearStartDate() string {
	y, _, _ := d.date()
	return time.Date(y, 1, 1, 0, 0, 0, 0, time.Local).Format(DateFormat)
}

func (d *Date) weekStart() time.Time {
	return d.now().AddDate(0, 0, -int(d.week())+1)
}

func (d *Date) monthStart() time.Time {
	y, m, _ := d.date()
	return time.Date(y, m, 1, 0, 0, 0, 0, time.Local)
}

func (d *Date) now() time.Time {
	if d.datetime != nil {
		return (d.datetime).(time.Time)
	}
	return time.Now()
}

func (d *Date) week() time.Weekday {
	// 周计算
	return d.now().Weekday()
}

func (d *Date) date() (year int, month time.Month, day int) {
	// 年月日
	year, month, day = d.now().Date()
	return
}
