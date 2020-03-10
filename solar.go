package chinese_calendar

import "math"

// solarTerms 24节气
// 通常春分的黄经定义为0°, 所以我们以春分作为开始
var solarTerms = []string{
	"春分", "清明", "谷雨", "立夏", "小满", "芒种",
	"夏至", "小暑", "大暑", "立秋", "处暑", "白露",
	"秋分", "寒露", "霜降", "立冬", "小雪", "大雪",
	"冬至", "小寒", "大寒", "立春", "雨水", "惊蛰",
}

// solarTermsTraditional 24节气繁体
var solarTermsTraditional = []string{
	"春分", "清明", "穀雨", "立夏", "小滿", "芒種",
	"夏至", "小暑", "大暑", "立秋", "處暑", "白露",
	"秋分", "寒露", "霜降", "立冬", "小雪", "大雪",
	"冬至", "小寒", "大寒", "立春", "雨水", "驚蟄",
}

// EclipticLongitude 黄道经度
type EclipticLongitude float64

// SolarTerm 根据黄道经度得到对应节气
func (l EclipticLongitude) SolarTerm(traditional ...bool) string {
	t := false
	if len(traditional) > 0 {
		t = traditional[0]
	}

	longitude := float64(l)
	if longitude >= 360 || longitude < 0 {
		longitude -= math.Floor(longitude/360) * 360
	}

	idx := int(math.Floor(longitude / 15))
	if t {
		return solarTermsTraditional[idx]
	}
	return solarTerms[idx]
}
