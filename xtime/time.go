package xtime

import "time"

const (
	Nanosecond  time.Duration = 1
	Microsecond               = 1000 * Nanosecond
	Millisecond               = 1000 * Microsecond
	Second                    = 1000 * Millisecond
	Minute                    = 60 * Second
	Hour                      = 60 * Minute
	Day                       = 24 * Hour
	Week                      = 7 * Day
)

func SomeDay(countDay int) time.Duration {
	return time.Duration(countDay) * Day
}

var Month = map[int]func() time.Duration{
	1: func() time.Duration { return 31 * Day },
	2: func() time.Duration {
		if IsLeapYear(time.Now().Year()) {
			return 29 * Day
		}
		return 28 * Day
	},
	3:  func() time.Duration { return 31 * Day },
	4:  func() time.Duration { return 30 * Day },
	5:  func() time.Duration { return 31 * Day },
	6:  func() time.Duration { return 30 * Day },
	7:  func() time.Duration { return 31 * Day },
	8:  func() time.Duration { return 31 * Day },
	9:  func() time.Duration { return 30 * Day },
	10: func() time.Duration { return 31 * Day },
	11: func() time.Duration { return 30 * Day },
	12: func() time.Duration { return 31 * Day },
}

var Year func() time.Duration = func() time.Duration {
	if IsLeapYear(time.Now().Year()) {
		return 366 * Day
	}
	return 365 * Day
}

// IsLeapYear 闰年（Leap Year）是为了弥补因人为历法规定造成的年度天数与地球实际公转周期的时间差而设立的。
// 补上时间差的年份为闰年。闰年共有366天（1月~12月分别为31天、29天、31天、30天、31天、30天、31天、31天、30天、31天、30天、31天）。
// 凡阳历中有闰日（2月29日）的年份，闰余（岁余置闰。阴历每年与回归年相比所差的时日）。
// 注意闰年（公历中的名词）和闰月（农历中的名词）并没有直接的关联，公历只分闰年和平年，
// 平年有365天，闰年有366天（2月中多一天）；平年中也可能有闰月（如2017年是平年，农历有闰月，闰六月）。
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
