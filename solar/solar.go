package solar

import "math"

const (
	// 初候
	FirstPentad = 0
	// 次候
	SecondPentad = 1
	// 末候
	ThirdPentad = 2
)

// Pentad 候
// 一个节气分三候: 初候,伺候,末候
// 三候将一个节气均匀分为各5天, 个别候为6天
type Pentad int

func (p Pentad) String() string {
	if p >= 3 || p < 0 {
		return ""
	}
	return [3]string{"初候", "次候", "末候"}[p]
}

// solarTerms 24节气中文简体
var solarTerms = [24]string{
	"春分", "清明", "谷雨", "立夏", "小满", "芒种",
	"夏至", "小暑", "大暑", "立秋", "处暑", "白露",
	"秋分", "寒露", "霜降", "立冬", "小雪", "大雪",
	"冬至", "小寒", "大寒", "立春", "雨水", "惊蛰",
}

// solarTermsTraditional 24节气中文繁体
var solarTermsTraditional = [24]string{
	"春分", "清明", "穀雨", "立夏", "小滿", "芒種",
	"夏至", "小暑", "大暑", "立秋", "處暑", "白露",
	"秋分", "寒露", "霜降", "立冬", "小雪", "大雪",
	"冬至", "小寒", "大寒", "立春", "雨水", "驚蟄",
}

// SolarTerm 24节气
type SolarTerm int

// 返回往后(或往面)第n个节气
// nth > 0往后推, nth < 0 往前推
func (st SolarTerm) Move(byNth int) SolarTerm {
	t := (st + SolarTerm(byNth)) % 24
	if t < 0 {
		t = SolarTerm(24)
	}
	return t
}

// IsMidTerm 该节气是否是中气
func (st SolarTerm) IsMidTerm() bool {
	l := map[SolarTerm]bool{
		SolarTermEnum.TheWinterSolstice:     true,
		SolarTermEnum.GreaterCold:           true,
		SolarTermEnum.RainWater:             true,
		SolarTermEnum.TheSpringEquinox:      true,
		SolarTermEnum.GrainRain:             true,
		SolarTermEnum.LesserFullnessOfGrain: true,
		SolarTermEnum.TheSummerSolstice:     true,
		SolarTermEnum.GreaterHeat:           true,
		SolarTermEnum.TheEndOfHeat:          true,
		SolarTermEnum.TheAutumnEquinox:      true,
		SolarTermEnum.FrostsDescent:         true,
		SolarTermEnum.LesserSnow:            true,
	}

	return l[st]
}

func (st SolarTerm) String(simplified bool) string {
	if int(st) >= 24 || int(st) < 0 {
		return ""
	}

	if simplified {
		return solarTerms[st]
	}
	return solarTermsTraditional[st]
}

// EclipticLongitude 黄道经度, 通常春分的黄经定义为0°
type EclipticLongitude float64

// SolarTerm 根据黄道经度得到对应节气
func (l EclipticLongitude) SolarTerm() SolarTerm {
	longitude := float64(l)
	if longitude >= 360 || longitude < 0 {
		longitude -= math.Floor(longitude/360) * 360
	}
	idx := int(math.Floor(longitude / 15))
	return SolarTerm(idx)
}

// SolarTermEnum 24节气枚举
// 英文命名参照中国气象局: http://www.cma.gov.cn/2011xzt/essjqzt/jqhz/jqhz02/201312/t20131213_233952.html
var SolarTermEnum = struct {
	TheSpringEquinox      SolarTerm // 春分
	PureBrightness        SolarTerm // 清明
	GrainRain             SolarTerm // 谷雨
	TheBeginningOfSummer  SolarTerm // 立夏
	LesserFullnessOfGrain SolarTerm // 小满
	GrainInBeard          SolarTerm // 芒种
	TheSummerSolstice     SolarTerm // 夏至
	LesserHeat            SolarTerm // 小暑
	GreaterHeat           SolarTerm // 大暑
	TheBeginningOfAutumn  SolarTerm // 立秋
	TheEndOfHeat          SolarTerm // 处暑
	WhiteDew              SolarTerm // 白露
	TheAutumnEquinox      SolarTerm // 秋分
	ColdDew               SolarTerm // 寒露
	FrostsDescent         SolarTerm // 霜降
	TheBeginningOfWinter  SolarTerm // 立冬
	LesserSnow            SolarTerm // 小雪
	GreaterSnow           SolarTerm // 大雪
	TheWinterSolstice     SolarTerm // 冬至
	LesserCold            SolarTerm // 小寒
	GreaterCold           SolarTerm // 大寒
	TheBeginningOfSpring  SolarTerm // 立春
	RainWater             SolarTerm // 雨水
	TheWakingOfInsects    SolarTerm // 惊蛰
}{
	TheSpringEquinox:      0,
	PureBrightness:        1,
	GrainRain:             2,
	TheBeginningOfSummer:  3,
	LesserFullnessOfGrain: 4,
	GrainInBeard:          5,
	TheSummerSolstice:     6,
	LesserHeat:            7,
	GreaterHeat:           8,
	TheBeginningOfAutumn:  9,
	TheEndOfHeat:          10,
	WhiteDew:              11,
	TheAutumnEquinox:      12,
	ColdDew:               13,
	FrostsDescent:         14,
	TheBeginningOfWinter:  15,
	LesserSnow:            16,
	GreaterSnow:           17,
	TheWinterSolstice:     18,
	LesserCold:            19,
	GreaterCold:           20,
	TheBeginningOfSpring:  21,
	RainWater:             22,
	TheWakingOfInsects:    23,
}
