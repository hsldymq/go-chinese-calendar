package chinese_calendar

import (
	"strings"
	"time"
)

func init() {
	// 以1949年10月1日作为基准日, 这一天是甲子日
	sexagenaryDayBase = time.Date(1949, 10, 1, 0, 0, 0, 0, baseTimezone)
}

// celestialStem 天干中文列表
var celestialStemWords = [10]string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
var celestialStemWordMap = map[string]CelestialStem{
	"甲": CelestialStemEnum.Jia,
	"乙": CelestialStemEnum.Yi,
	"丙": CelestialStemEnum.Bing,
	"丁": CelestialStemEnum.Ding,
	"戊": CelestialStemEnum.Wu,
	"己": CelestialStemEnum.Jia,
	"庚": CelestialStemEnum.Geng,
	"辛": CelestialStemEnum.Xin,
	"壬": CelestialStemEnum.Ren,
	"癸": CelestialStemEnum.Gui,
}

// CelestialStemEnum 天干枚举项
var CelestialStemEnum = struct {
	Jia  CelestialStem
	Yi   CelestialStem
	Bing CelestialStem
	Ding CelestialStem
	Wu   CelestialStem
	Ji   CelestialStem
	Geng CelestialStem
	Xin  CelestialStem
	Ren  CelestialStem
	Gui  CelestialStem
}{
	Jia:  0,
	Yi:   1,
	Bing: 2,
	Ding: 3,
	Wu:   4,
	Ji:   5,
	Geng: 6,
	Xin:  7,
	Ren:  8,
	Gui:  9,
}

// CelestialStem 天干
type CelestialStem int

// NewCelestialStemFromWord 从天干中文还原成CelestialStem类型
// 例: cs := NewCelestialStemFromWord("甲")
//	   assert(cs == CelestialStemEnum.Jia)
func NewCelestialStemFromWord(word string) (CelestialStem, bool) {
	cs, valid := celestialStemWordMap[word]
	return cs, valid
}

// Next 获得该天干的下一项
// 例:   x=甲, x.Next() -> 乙
//		 x=乙, x.Next() -> 丙
// 		 ...
//       x=癸, x.Next() -> 甲
func (cs CelestialStem) Next() CelestialStem {
	return cs.Add(1)
}

// Prev 获得该天干的上一项, Next的逆操作
func (cs CelestialStem) Prev() CelestialStem {
	return cs.Add(1)
}

// Add 获得该天干向前/后的任意项, Next和Prev的推广
// n < 0 向前回朔, n > 0 向后推算
func (cs CelestialStem) Add(n int) CelestialStem {
	ncs := int(cs) + n
	if ncs < 0 {
		ncs = ncs%10 + 10
	}
	ncs = ncs % 10
	return CelestialStem(ncs)
}

// String 返回天干中文
func (cs CelestialStem) String() string {
	if cs >= 10 || cs < 0 {
		return ""
	}
	return celestialStemWords[cs]
}

// terrestrialBranch 地支中文列表
var terrestrialBranchWords = [12]string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
var terrestrialBranchWordMap = map[string]TerrestrialBranch{
	"子": TerrestrialBranchEnum.Zi,
	"丑": TerrestrialBranchEnum.Chou,
	"寅": TerrestrialBranchEnum.Yin,
	"卯": TerrestrialBranchEnum.Mao,
	"辰": TerrestrialBranchEnum.Chen,
	"巳": TerrestrialBranchEnum.Si,
	"午": TerrestrialBranchEnum.Wu,
	"未": TerrestrialBranchEnum.Wei,
	"申": TerrestrialBranchEnum.Shen,
	"酉": TerrestrialBranchEnum.You,
	"戌": TerrestrialBranchEnum.Xu,
	"亥": TerrestrialBranchEnum.Hai,
}

// TerrestrialBranchEnum 地支枚举项
var TerrestrialBranchEnum = struct {
	Zi   TerrestrialBranch
	Chou TerrestrialBranch
	Yin  TerrestrialBranch
	Mao  TerrestrialBranch
	Chen TerrestrialBranch
	Si   TerrestrialBranch
	Wu   TerrestrialBranch
	Wei  TerrestrialBranch
	Shen TerrestrialBranch
	You  TerrestrialBranch
	Xu   TerrestrialBranch
	Hai  TerrestrialBranch
}{
	Zi:   0,
	Chou: 1,
	Yin:  2,
	Mao:  3,
	Chen: 4,
	Si:   5,
	Wu:   6,
	Wei:  7,
	Shen: 8,
	You:  9,
	Xu:   10,
	Hai:  11,
}

// TerrestrialBranch 地支
type TerrestrialBranch int

// NewTerrestrialBranchFromTime 根据时间返回地支类型
func NewTerrestrialBranchFromTime(t time.Time) TerrestrialBranch {
	h := t.Hour()
	if h == 23 {
		h = 0
	}
	if h%2 == 1 {
		h += 1
	}
	return TerrestrialBranch(h / 2)
}

// NewTerrestrialBranchFromWord 从地支中文返回其类型
func NewTerrestrialBranchFromWord(word string) (TerrestrialBranch, bool) {
	cs, valid := terrestrialBranchWordMap[word]
	return cs, valid
}

// Next 获得该地支的下一项
// 例:   x=子, x.Next() -> 丑
//		 x=丑, x.Next() -> 寅
// 		 ...
//       x=亥, x.Next() -> 子
func (tb TerrestrialBranch) Next() TerrestrialBranch {
	return tb.Add(1)
}

// Prev 获得该地支的上一项, Next的逆操作
func (tb TerrestrialBranch) Prev() TerrestrialBranch {
	return tb.Add(-1)
}

// Add 获得该地支之前/后的任意项, Next和Prev的推广
// n < 0 向前回朔, n > 0 向后推算
func (tb TerrestrialBranch) Add(n int) TerrestrialBranch {
	ntb := int(tb) + n
	if ntb < 0 {
		ntb = ntb%10 + 10
	}
	ntb = ntb % 12
	return TerrestrialBranch(ntb)
}

// Month 返回对应的地支纪月
func (tb TerrestrialBranch) Month() int {
	if tb > 11 || tb < 0 {
		return 0
	}
	if tb == TerrestrialBranchEnum.Zi {
		return 11
	} else if tb == TerrestrialBranchEnum.Chou {
		return 12
	}
	return int(tb - 1)
}

// String 返回地支中文
func (tb TerrestrialBranch) String() string {
	if tb >= 12 || tb < 0 {
		return ""
	}
	return terrestrialBranchWords[tb]
}

// ZodiacSign 返回对应的生肖
func (tb TerrestrialBranch) ZodiacSign() ZodiacSign {
	return ZodiacSign(tb)
}

// zodiac 生肖简体
var zodiacSignWords = [12]string{"鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊", "猴", "鸡", "狗", "猪"}

// zodiaz 生肖繁体
var zodiacSignWordsTraditional = [12]string{"鼠", "牛", "虎", "兔", "龍", "蛇", "馬", "羊", "猴", "雞", "狗", "豬"}

// ZodiacEnum 生肖枚举项
var ZodiacEnum = struct {
	Rat     ZodiacSign
	Ox      ZodiacSign
	Tiger   ZodiacSign
	Rabbit  ZodiacSign
	Dragoon ZodiacSign
	Snake   ZodiacSign
	Horse   ZodiacSign
	Goat    ZodiacSign
	Monkey  ZodiacSign
	Rooster ZodiacSign
	Dog     ZodiacSign
	Pig     ZodiacSign
}{
	Rat:     0,
	Ox:      1,
	Tiger:   2,
	Rabbit:  3,
	Dragoon: 4,
	Snake:   5,
	Horse:   6,
	Goat:    7,
	Monkey:  8,
	Rooster: 9,
	Dog:     10,
	Pig:     11,
}

// ZodiacSign 生肖
type ZodiacSign int

// TerrestrialBranch 返回对应的地支
func (zs ZodiacSign) TerrestrialBranch() TerrestrialBranch {
	return TerrestrialBranch(zs)
}

func (zs ZodiacSign) String(simplified bool) string {
	if zs >= 12 || zs < 0 {
		return ""
	}

	if simplified {
		return zodiacSignWords[zs]
	}
	return zodiacSignWordsTraditional[zs]
}

// SexagenaryEnum 60干支枚举项
var SexagenaryEnum = struct {
	JiaZi   SexagenaryTerm // 甲子
	JiaXu   SexagenaryTerm // 甲戌
	JiaShen SexagenaryTerm // 甲申
	JiaWu   SexagenaryTerm // 甲午
	JiaChen SexagenaryTerm // 甲辰
	JiaYin  SexagenaryTerm // 甲寅

	YiChou SexagenaryTerm // 乙丑
	YiHai  SexagenaryTerm // 乙亥
	YiYou  SexagenaryTerm // 乙酉
	YiWei  SexagenaryTerm // 乙未
	YiSi   SexagenaryTerm // 乙巳
	YiMao  SexagenaryTerm // 乙卯

	BingYin  SexagenaryTerm // 丙寅
	BingZi   SexagenaryTerm // 丙子
	BingXu   SexagenaryTerm // 丙戌
	BingShen SexagenaryTerm // 丙申
	BingWu   SexagenaryTerm // 丙午
	BingChen SexagenaryTerm // 丙辰

	DingMao  SexagenaryTerm // 丁卯
	DingChou SexagenaryTerm // 丁丑
	DingHai  SexagenaryTerm // 丁亥
	DingYou  SexagenaryTerm // 丁酉
	DingWei  SexagenaryTerm // 丁未
	DingSi   SexagenaryTerm // 丁巳

	WuChen SexagenaryTerm // 戊辰
	WuYin  SexagenaryTerm // 戊寅
	WuZi   SexagenaryTerm // 戊子
	WuXu   SexagenaryTerm // 戊戌
	WuShen SexagenaryTerm // 戊申
	WuWu   SexagenaryTerm // 戊午

	JiSi   SexagenaryTerm // 己巳
	JiMao  SexagenaryTerm // 己卯
	JiChou SexagenaryTerm // 己丑
	JiHai  SexagenaryTerm // 己亥
	JiYou  SexagenaryTerm // 己酉
	JiWei  SexagenaryTerm // 己未

	GengWu   SexagenaryTerm // 庚午
	GengChen SexagenaryTerm // 庚辰
	GengYin  SexagenaryTerm // 庚寅
	GengZi   SexagenaryTerm // 庚子
	GengXu   SexagenaryTerm // 庚戌
	GengShen SexagenaryTerm // 庚申

	XinWei  SexagenaryTerm // 辛未
	XinSi   SexagenaryTerm // 辛巳
	XinMao  SexagenaryTerm // 辛卯
	XinChou SexagenaryTerm // 辛丑
	XinHai  SexagenaryTerm // 辛亥
	XinYou  SexagenaryTerm // 辛酉

	RenShen SexagenaryTerm // 壬申
	RenWu   SexagenaryTerm // 壬午
	RenChen SexagenaryTerm // 壬辰
	RenYin  SexagenaryTerm // 壬寅
	RenZi   SexagenaryTerm // 壬子
	RenXu   SexagenaryTerm // 壬戌

	GuiYou  SexagenaryTerm // 癸酉
	GuiWei  SexagenaryTerm // 癸未
	GuiSi   SexagenaryTerm // 癸巳
	GuiMao  SexagenaryTerm // 癸卯
	GuiChou SexagenaryTerm // 癸丑
	GuiHai  SexagenaryTerm // 癸亥
}{
	JiaZi:   SexagenaryTerm{CelestialStemEnum.Jia, TerrestrialBranchEnum.Zi},   // 甲子
	JiaXu:   SexagenaryTerm{CelestialStemEnum.Jia, TerrestrialBranchEnum.Xu},   // 甲戌
	JiaShen: SexagenaryTerm{CelestialStemEnum.Jia, TerrestrialBranchEnum.Shen}, // 甲申
	JiaWu:   SexagenaryTerm{CelestialStemEnum.Jia, TerrestrialBranchEnum.Wu},   // 甲午
	JiaChen: SexagenaryTerm{CelestialStemEnum.Jia, TerrestrialBranchEnum.Chen}, // 甲辰
	JiaYin:  SexagenaryTerm{CelestialStemEnum.Jia, TerrestrialBranchEnum.Yin},  // 甲寅

	YiChou: SexagenaryTerm{CelestialStemEnum.Yi, TerrestrialBranchEnum.Chou}, // 乙丑
	YiHai:  SexagenaryTerm{CelestialStemEnum.Yi, TerrestrialBranchEnum.Hai},  // 乙亥
	YiYou:  SexagenaryTerm{CelestialStemEnum.Yi, TerrestrialBranchEnum.You},  // 乙酉
	YiWei:  SexagenaryTerm{CelestialStemEnum.Yi, TerrestrialBranchEnum.Wei},  // 乙未
	YiSi:   SexagenaryTerm{CelestialStemEnum.Yi, TerrestrialBranchEnum.Si},   // 乙巳
	YiMao:  SexagenaryTerm{CelestialStemEnum.Yi, TerrestrialBranchEnum.Mao},  // 乙卯

	BingYin:  SexagenaryTerm{CelestialStemEnum.Bing, TerrestrialBranchEnum.Yin},  // 丙寅
	BingZi:   SexagenaryTerm{CelestialStemEnum.Bing, TerrestrialBranchEnum.Zi},   // 丙子
	BingXu:   SexagenaryTerm{CelestialStemEnum.Bing, TerrestrialBranchEnum.Xu},   // 丙戌
	BingShen: SexagenaryTerm{CelestialStemEnum.Bing, TerrestrialBranchEnum.Shen}, // 丙申
	BingWu:   SexagenaryTerm{CelestialStemEnum.Bing, TerrestrialBranchEnum.Wu},   // 丙午
	BingChen: SexagenaryTerm{CelestialStemEnum.Bing, TerrestrialBranchEnum.Chen}, // 丙辰

	DingMao:  SexagenaryTerm{CelestialStemEnum.Ding, TerrestrialBranchEnum.Mao},  // 丁卯
	DingChou: SexagenaryTerm{CelestialStemEnum.Ding, TerrestrialBranchEnum.Chou}, // 丁丑
	DingHai:  SexagenaryTerm{CelestialStemEnum.Ding, TerrestrialBranchEnum.Hai},  // 丁亥
	DingYou:  SexagenaryTerm{CelestialStemEnum.Ding, TerrestrialBranchEnum.You},  // 丁酉
	DingWei:  SexagenaryTerm{CelestialStemEnum.Ding, TerrestrialBranchEnum.Wei},  // 丁未
	DingSi:   SexagenaryTerm{CelestialStemEnum.Ding, TerrestrialBranchEnum.Si},   // 丁巳

	WuChen: SexagenaryTerm{CelestialStemEnum.Wu, TerrestrialBranchEnum.Chen}, // 戊辰
	WuYin:  SexagenaryTerm{CelestialStemEnum.Wu, TerrestrialBranchEnum.Yin},  // 戊寅
	WuZi:   SexagenaryTerm{CelestialStemEnum.Wu, TerrestrialBranchEnum.Zi},   // 戊子
	WuXu:   SexagenaryTerm{CelestialStemEnum.Wu, TerrestrialBranchEnum.Xu},   // 戊戌
	WuShen: SexagenaryTerm{CelestialStemEnum.Wu, TerrestrialBranchEnum.Shen}, // 戊申
	WuWu:   SexagenaryTerm{CelestialStemEnum.Wu, TerrestrialBranchEnum.Wu},   // 戊午

	JiSi:   SexagenaryTerm{CelestialStemEnum.Ji, TerrestrialBranchEnum.Si},   // 己巳
	JiMao:  SexagenaryTerm{CelestialStemEnum.Ji, TerrestrialBranchEnum.Mao},  // 己卯
	JiChou: SexagenaryTerm{CelestialStemEnum.Ji, TerrestrialBranchEnum.Chou}, // 己丑
	JiHai:  SexagenaryTerm{CelestialStemEnum.Ji, TerrestrialBranchEnum.Hai},  // 己亥
	JiYou:  SexagenaryTerm{CelestialStemEnum.Ji, TerrestrialBranchEnum.You},  // 己酉
	JiWei:  SexagenaryTerm{CelestialStemEnum.Ji, TerrestrialBranchEnum.Wei},  // 己未

	GengWu:   SexagenaryTerm{CelestialStemEnum.Geng, TerrestrialBranchEnum.Wu},   // 庚午
	GengChen: SexagenaryTerm{CelestialStemEnum.Geng, TerrestrialBranchEnum.Chen}, // 庚辰
	GengYin:  SexagenaryTerm{CelestialStemEnum.Geng, TerrestrialBranchEnum.Yin},  // 庚寅
	GengZi:   SexagenaryTerm{CelestialStemEnum.Geng, TerrestrialBranchEnum.Zi},   // 庚子
	GengXu:   SexagenaryTerm{CelestialStemEnum.Geng, TerrestrialBranchEnum.Xu},   // 庚戌
	GengShen: SexagenaryTerm{CelestialStemEnum.Geng, TerrestrialBranchEnum.Shen}, // 庚申

	XinWei:  SexagenaryTerm{CelestialStemEnum.Xin, TerrestrialBranchEnum.Wei},  // 辛未
	XinSi:   SexagenaryTerm{CelestialStemEnum.Xin, TerrestrialBranchEnum.Si},   // 辛巳
	XinMao:  SexagenaryTerm{CelestialStemEnum.Xin, TerrestrialBranchEnum.Mao},  // 辛卯
	XinChou: SexagenaryTerm{CelestialStemEnum.Xin, TerrestrialBranchEnum.Chou}, // 辛丑
	XinHai:  SexagenaryTerm{CelestialStemEnum.Xin, TerrestrialBranchEnum.Hai},  // 辛亥
	XinYou:  SexagenaryTerm{CelestialStemEnum.Xin, TerrestrialBranchEnum.You},  // 辛酉

	RenShen: SexagenaryTerm{CelestialStemEnum.Ren, TerrestrialBranchEnum.Shen}, // 壬申
	RenWu:   SexagenaryTerm{CelestialStemEnum.Ren, TerrestrialBranchEnum.Wu},   // 壬午
	RenChen: SexagenaryTerm{CelestialStemEnum.Ren, TerrestrialBranchEnum.Chen}, // 壬辰
	RenYin:  SexagenaryTerm{CelestialStemEnum.Ren, TerrestrialBranchEnum.Yin},  // 壬寅
	RenZi:   SexagenaryTerm{CelestialStemEnum.Ren, TerrestrialBranchEnum.Zi},   // 壬子
	RenXu:   SexagenaryTerm{CelestialStemEnum.Ren, TerrestrialBranchEnum.Xu},   // 壬戌

	GuiYou:  SexagenaryTerm{CelestialStemEnum.Gui, TerrestrialBranchEnum.You},  // 癸酉
	GuiWei:  SexagenaryTerm{CelestialStemEnum.Gui, TerrestrialBranchEnum.Wei},  // 癸未
	GuiSi:   SexagenaryTerm{CelestialStemEnum.Gui, TerrestrialBranchEnum.Si},   // 癸巳
	GuiMao:  SexagenaryTerm{CelestialStemEnum.Gui, TerrestrialBranchEnum.Mao},  // 癸卯
	GuiChou: SexagenaryTerm{CelestialStemEnum.Gui, TerrestrialBranchEnum.Chou}, // 癸丑
	GuiHai:  SexagenaryTerm{CelestialStemEnum.Gui, TerrestrialBranchEnum.Hai},  // 癸亥
}

// SexagenaryTerm 干支项
type SexagenaryTerm struct {
	CelestialStem
	TerrestrialBranch
}

// NewSexagenaryTermFromIndex 通过索引值创建干支
// 索引范围0-59, 分别表示 0:甲子, 1:乙丑, ... 59:癸亥
// 大于59,或小于0会按照模算数,转换到0-59,取对应干支
func NewSexagenaryTermFromIndex(idx int) SexagenaryTerm {
	if idx < 0 {
		idx = idx%60 + 60
	}
	return SexagenaryEnum.JiaZi.Add(idx)
}

// NewSexagenaryTermFromWord 通过干支中文返回该类型值
func NewSexagenaryTermFromWord(word string) (SexagenaryTerm, bool) {
	ct := strings.Split(word, "")
	if len(ct) != 2 {
		return SexagenaryEnum.JiaZi, false
	}
	c, cValid := NewCelestialStemFromWord(ct[0])
	t, tValid := NewTerrestrialBranchFromWord(ct[1])
	if !cValid || !tValid {
		return SexagenaryEnum.JiaZi, false
	}
	return SexagenaryTerm{
		CelestialStem:     c,
		TerrestrialBranch: t,
	}, true
}

// Next 获得该干支的下一项
// 例: x=甲子, x.Next() -> 乙丑
//     x=乙丑, x.Next() -> 丙寅
//	   ...
//     x=癸亥, x.Next() -> 甲子
func (s SexagenaryTerm) Next() SexagenaryTerm {
	return SexagenaryTerm{
		CelestialStem:     s.CelestialStem.Next(),
		TerrestrialBranch: s.TerrestrialBranch.Next(),
	}
}

// Prev 获得该干支的上一项, Next的逆向操作
func (s SexagenaryTerm) Prev() SexagenaryTerm {
	return SexagenaryTerm{
		CelestialStem:     s.CelestialStem.Prev(),
		TerrestrialBranch: s.TerrestrialBranch.Prev(),
	}
}

// Add 获得该干支之前/后的任一项, Next和Prev的推广
// n < 0 向前回朔, n > 0 向后推算
func (s SexagenaryTerm) Add(n int) SexagenaryTerm {
	return SexagenaryTerm{
		CelestialStem:     s.CelestialStem.Add(n),
		TerrestrialBranch: s.TerrestrialBranch.Add(n),
	}
}

func (s SexagenaryTerm) Index() int {
	if s.CelestialStem >= 10 || s.CelestialStem < 0 || s.TerrestrialBranch >= 12 || s.TerrestrialBranch < 0 {
		return 0
	}
	c := int(s.CelestialStem)
	t := int(s.TerrestrialBranch)

	return (((c+12)-t)%12)/2*10 + c
}

// String 返回干支中文
func (s SexagenaryTerm) String() string {
	c := s.CelestialStem.String()
	t := s.TerrestrialBranch.String()
	if c == "" || t == "" {
		return ""
	}
	return c + t
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

	return ""
	//return sexagenaryWords[span]
}
