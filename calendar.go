package chinese_calendar

import (
	"math"
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

// julianDay 计算给定时间所对应的儒略日
// 算法参考Jean Meeus的<<Astronomical Algorithms>>第二版中的公式
// 只能用于1582年10月15日及以后的时间, 因为在这天之前需要用儒略历来计算儒略日, 而golang的时间计算方式始终按照格里历计算.
func julianDay(t time.Time) float64 {
	t.In(time.FixedZone("", 0))
	year := float64(t.Year())
	month := float64(t.Month())
	secondsOfDay := float64(t.Hour()*3600) + float64(t.Minute()*60) + float64(t.Second())
	if t.Month() < 3 {
		year -= 1
		month += 12
	}
	a := math.Floor(year / 100)
	b := math.Floor(2 - a + math.Floor(a/4))
	d := float64(t.Day()) + secondsOfDay/86400
	//if year < 1000 {
	//	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	//	b = 0
	//}

	return math.Floor(365.25*(year+4716)) +
		math.Floor(30.6001*(month+1)) +
		d + b - 1524.5
}
