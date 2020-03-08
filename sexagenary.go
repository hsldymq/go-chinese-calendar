package chinese_calendar

import (
	"time"
)

// zodiac 生肖简体
var zodiac = []string{"鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊", "猴", "鸡", "狗", "猪"}

// zodiaz 生肖繁体
var zodiacTraditional = []string{"鼠", "牛", "虎", "兔", "龍", "蛇", "馬", "羊", "猴", "雞", "狗", "豬"}

// celestialStem 天干
var celestialStem = []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}

// terrestrialBranch 地支
var terrestrialBranch = []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}

// sexagenary 干支周列表
var sexagenary = []string{
	"甲子", "乙丑", "丙寅", "丁卯", "戊辰", "己巳", "庚午", "辛未", "壬申", "癸酉",
	"甲戌", "乙亥", "丙子", "丁丑", "戊寅", "己卯", "庚辰", "辛巳", "壬午", "癸未",
	"甲申", "乙酉", "丙戌", "丁亥", "戊子", "己丑", "庚寅", "辛卯", "壬辰", "癸巳",
	"甲午", "乙未", "丙申", "丁酉", "戊戌", "己亥", "庚子", "辛丑", "壬寅", "癸卯",
	"甲辰", "乙巳", "丙午", "丁未", "戊申", "己酉", "庚戌", "辛亥", "壬子", "癸丑",
	"甲寅", "乙卯", "丙辰", "丁巳", "戊午", "己未", "庚申", "辛酉", "壬戌", "癸亥",
}

func init() {
	// 以1949年10月1日作为基准日, 这一天是甲子日
	sexagenaryDayBase = time.Date(1949, 10, 1, 0, 0, 0, 0, baseTimezone)
}

// Sexagenary 干支周列表
func Sexagenary() []string {
	return sexagenary
}

// SexagenaryNumber 根据给定的干支索引计算指定跨度的干支纪数
// 用于计算干支纪时,日,月,年
// span 指定跨越的数
// startIdx 代表干支纪数序号, 0-59, 0代表甲子, 1代表乙丑, 以此类推
//		不传默认为0
// 例:
//		SexagenaryNumber(2) 返回"丙寅", 即按照"甲子"往后移2: "甲子" -> "乙丑" -> "丙寅"
//		SexagenaryNumber(-3, 1) 返回"壬戌", 即按照"乙丑往前移3: "乙丑" -> "甲子" -> "癸亥" -> "壬戌"
// 因为60干支周的计算实际上是模算数,所以参数中任何小于0或大于59的数都会换算成模60的数再计算干支纪日
func SexagenaryNumber(span int, startIdx ...int) string {
	sIdx := 0
	if len(startIdx) > 0 {
		sIdx = startIdx[0]
	}

	span += sIdx
	if span < 0 {
		span = span%60 + 60
	}
	span = span % 60

	return sexagenary[span]
}
