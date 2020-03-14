package chinese_calendar

import (
	"testing"
)

func TestSexagenaryTerm(t *testing.T) {
	t.Run("test index", func(t *testing.T) {
		for i := -60; i <= 119; i++ {
			s := NewSexagenaryTermFromIndex(i)
			expect := i
			actual := s.Index()
			comparison := actual
			if i < 0 {
				comparison -= 60
				expect += 60
			} else if i >= 60 {
				comparison += 60
				expect -= 60
			}
			if i != comparison {
				t.Fatalf("expect the index of sexagenary term from %d to be %d, got %d", i, expect, actual)
			}
		}
	})
}

//
//func TestSexagenary(t *testing.T) {
//	var sexagenary = []string{
//		"甲子", "乙丑", "丙寅", "丁卯", "戊辰", "己巳", "庚午", "辛未", "壬申", "癸酉",
//		"甲戌", "乙亥", "丙子", "丁丑", "戊寅", "己卯", "庚辰", "辛巳", "壬午", "癸未",
//		"甲申", "乙酉", "丙戌", "丁亥", "戊子", "己丑", "庚寅", "辛卯", "壬辰", "癸巳",
//		"甲午", "乙未", "丙申", "丁酉", "戊戌", "己亥", "庚子", "辛丑", "壬寅", "癸卯",
//		"甲辰", "乙巳", "丙午", "丁未", "戊申", "己酉", "庚戌", "辛亥", "壬子", "癸丑",
//		"甲寅", "乙卯", "丙辰", "丁巳", "戊午", "己未", "庚申", "辛酉", "壬戌", "癸亥",
//	}
//	t.Run("test sexagenaryWords number", func(t *testing.T) {
//		for i := 0; i < 60; i++ {
//			require.Equal(t, sexagenaryWords[i], SexagenaryNumber(i))
//			require.Equal(t, sexagenaryWords[i], SexagenaryNumber(i-3, 3))
//			require.Equal(t, sexagenaryWords[i], SexagenaryNumber(i+60))
//			require.Equal(t, sexagenaryWords[i], SexagenaryNumber(i-60))
//			require.Equal(t, sexagenaryWords[i], SexagenaryNumber(i-120))
//		}
//	})
//
//}
