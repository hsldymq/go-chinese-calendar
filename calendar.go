package chinese_calendar

import (
	"time"
)

// baseTimezone 农历计算的标准时区, 为东经120°标准时
var baseTimezone *time.Location

// sexagenaryDayBase 干支纪日参考时
var sexagenaryDayBase time.Time

func init() {
	// 农历计算时区恒为北京时间(东经120°标准时)
	// 这里没有通过loadLocation载入Asia/Shanghai, 而是用FixedZone来固定这个8小时的时差(尽管当前这两者当前是等同的)
	baseTimezone = time.FixedZone("Beijing", 8*60*60)
}
