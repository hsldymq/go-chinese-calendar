package chinese_calendar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSexagenary(t *testing.T) {
	t.Run("test sexagenary day", func(t *testing.T) {
		sexagennary := []string{
			"甲子", "乙丑", "丙寅", "丁卯", "戊辰", "己巳", "庚午", "辛未", "壬申", "癸酉",
			"甲戌", "乙亥", "丙子", "丁丑", "戊寅", "己卯", "庚辰", "辛巳", "壬午", "癸未",
			"甲申", "乙酉", "丙戌", "丁亥", "戊子", "己丑", "庚寅", "辛卯", "壬辰", "癸巳",
			"甲午", "乙未", "丙申", "丁酉", "戊戌", "己亥", "庚子", "辛丑", "壬寅", "癸卯",
			"甲辰", "乙巳", "丙午", "丁未", "戊申", "己酉", "庚戌", "辛亥", "壬子", "癸丑",
			"甲寅", "乙卯", "丙辰", "丁巳", "戊午", "己未", "庚申", "辛酉", "壬戌", "癸亥",
		}

		for i := 0; i < 60; i++ {
			require.Equal(t, sexagennary[i], Sexagenary(i))
			require.Equal(t, sexagennary[i], Sexagenary(i-3, 3))
			require.Equal(t, sexagennary[i], Sexagenary(i+60))
			require.Equal(t, sexagennary[i], Sexagenary(i-60))
			require.Equal(t, sexagennary[i], Sexagenary(i-120))
		}
	})

}

func TestJulianDayNumber(t *testing.T) {
	t.Run("haha", func(t *testing.T) {
		// 测试用例来自于<<Astronomical Algorithms>>第二版
		cases := map[float64][6]int{
			2451545.0:  {2000, 1, 1, 12, 0, 0},
			2451179.5:  {1999, 1, 1, 0, 0, 0},
			2446822.5:  {1987, 1, 27, 0, 0, 0},
			2446966.0:  {1987, 6, 19, 12, 0, 0},
			2447187.5:  {1988, 1, 27, 0, 0, 0},
			2447332.0:  {1988, 6, 19, 12, 0, 0},
			2436116.31: {1957, 10, 4, 19, 26, 24},
			2415020.5:  {1900, 1, 1, 0, 0, 0},
			2305447.5:  {1600, 1, 1, 0, 0, 0},
			2305812.5:  {1600, 12, 31, 0, 0, 0},
			//2026871.8:  {837, 4, 10, 7, 12, 0},
			//1842713.0: {333, 1, 27, 12, 0, 0},
			//1676496.5: {-123, 12, 31, 0, 0, 0},
			//1676497.5: {-122, 1, 1, 0, 0, 0},
			//1356001.0: {-1000, 7, 12, 12, 0, 0},
			//1355866.5: {-1000, 2, 29, 0, 0, 0},
			//1355671.4: {-1001, 8, 17, 21, 36, 0},
			//0.0:       {-4712, 1, 1, 12, 0, 0},
		}

		for expect, date := range cases {
			require.Equal(t, expect, julianDay(time.Date(date[0], time.Month(date[1]), date[2], date[3], date[4], date[5], 0, time.UTC)))
		}

	})
}
