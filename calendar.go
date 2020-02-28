package chinese_calendar

import (
	"fmt"
	"math"
	"time"
)

// baseTimezone 农历计算的标准时区, 为东经120°标准时
var baseTimezone *time.Location

// sexagenaryDayBase 干支纪日参考时
var sexagenaryDayBase time.Time

// zodiac 生肖
var zodiac = [][2]string{
	{"鼠", "鼠"}, {"牛", "牛"}, {"虎", "虎"}, {"兔", "兔"},
	{"龙", "龍"}, {"蛇", "蛇"}, {"马", "馬"}, {"羊", "羊"},
	{"猴", "猴"}, {"鸡", "雞"}, {"狗", "狗"}, {"猪", "豬"},
}

// celestialStem 天干
var celestialStem = []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}

// terrestrialBranch 地支
var terrestrialBranch = []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}

func init() {
	// 农历计算时区恒为北京时间(东经120°标准时)
	// 这里没有通过loadLocation载入Asia/Shanghai, 而是用FixedZone来固定这个8小时的时差(尽管当前这两者当前是等同的)
	baseTimezone = time.FixedZone("Beijing", 8*60*60)

	// 以1949年10月1日作为基准日, 这一天是甲子日
	sexagenaryDayBase = time.Date(1949, 10, 1, 0, 0, 0, 0, baseTimezone)
}

// Sexagenary 根据给定的干支索引计算指定跨度的干支纪数
// 用于计算干支纪时,日,月,年
// span 指定跨越的数
// startIdx 代表干支纪数序号, 0-59, 0代表甲子, 1代表乙丑, 以此类推
//		不传默认为0
// 例:
//		Sexagenary(2) 返回"丙寅", 即按照"甲子"往后移2: "甲子" -> "乙丑" -> "丙寅"
//		Sexagenary(-3, 1) 返回"壬戌", 即按照"乙丑往前移3: "乙丑" -> "甲子" -> "癸亥" -> "壬戌"
// 因为60干支周的计算实际上是模算数,所以参数中任何小于0或大于59的数都会换算成模60的数再计算干支纪日
func Sexagenary(span int, startIdx ...int) string {
	sIdx := 0
	if len(startIdx) > 0 {
		sIdx = startIdx[0]
	}

	span += sIdx
	if span < 0 {
		span = span%60 + 60
	}
	span = span % 60

	return fmt.Sprintf("%s%s", celestialStem[span%10], terrestrialBranch[span%12])
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
