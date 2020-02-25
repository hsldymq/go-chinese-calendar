package chinese_calendar

import (
	"fmt"
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

// Sexagenary 计算从甲子到指定跨度的干支纪元值
// 用于计算干支纪时,日,月,年
// 例:
//		当span == 2, 即甲子日往后推算2, 即"丙寅"
//		当span == -3, 甲子日往前推算3, 即"辛酉"
// 因为60干支周的计算实际上是模算数,所以参数中任何小于0或大于59的数都会换算成模60的数再计算干支纪日
func Sexagenary(span int) string {
	if span < 0 {
		span = span%60 + 60
	}
	span = span % 60

	return fmt.Sprintf("%s%s", celestialStem[span%10], terrestrialBranch[span%12])
}
