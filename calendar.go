package calendar

import (
	"time"
)

// baseTimezone 农历计算的标准时区, 为东经120°标准时
var baseTimezone *time.Location

// sexagenaryYearBase 干支年参考时
var sexagenaryYearBase time.Time

// sexagenaryDayBase 干支纪日参考时
var sexagenaryDayBase time.Time

func init() {
	// 农历计算时区恒为北京时间(东经120°标准时)
	// 这里没有通过loadLocation载入Asia/Shanghai, 而是用FixedZone来固定这个8小时的时差(尽管当前这两者当前是等同的)
	baseTimezone = time.FixedZone("UTF+8", 8*60*60)

	// 1984年2月2日开始的农历年为甲子年
	sexagenaryYearBase = time.Date(1984, 2, 2, 0, 0, 0, 0, baseTimezone)

	// 以1949年10月1日作为基准日, 这一天是甲子日
	sexagenaryDayBase = time.Date(1949, 10, 1, 0, 0, 0, 0, baseTimezone)
}
