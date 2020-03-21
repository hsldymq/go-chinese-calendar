package calendar

import "testing"

func TestNewCelestialStemFromWord(t *testing.T) {
	inputs := []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸", "我"}
	expect := []struct {
		CelestialStem
		Valid bool
	}{
		{CelestialStemEnum.Jia, true},
		{CelestialStemEnum.Yi, true},
		{CelestialStemEnum.Bing, true},
		{CelestialStemEnum.Ding, true},
		{CelestialStemEnum.Wu, true},
		{CelestialStemEnum.Ji, true},
		{CelestialStemEnum.Geng, true},
		{CelestialStemEnum.Xin, true},
		{CelestialStemEnum.Ren, true},
		{CelestialStemEnum.Gui, true},
		{CelestialStemEnum.Jia, false},
	}

	for idx, each := range inputs {
		c, v := NewCelestialStemFromWord(each)
		if c != expect[idx].CelestialStem || v != expect[idx].Valid {
			t.Fatalf("new celestial stem from word %s, should return %d and %v, got %d and %v",
				each,
				expect[idx].CelestialStem,
				expect[idx].Valid,
				c,
				v,
			)
		}
	}
}

func TestCelestialStem(t *testing.T) {
	t.Run("test Add method", func(t *testing.T) {
		actual := CelestialStemEnum.Jia
		ns := []int{
			1, 1, 1, 1, 1,
			1, 1, 1, 1, 1,
			-1, -1, -1, -1, -1,
			-1, -1, -1, -1, -1,
			10, 20, -10, -20,
			15, 5, -15, -5,
		}
		expect := []CelestialStem{
			CelestialStemEnum.Yi, CelestialStemEnum.Bing, CelestialStemEnum.Ding, CelestialStemEnum.Wu, CelestialStemEnum.Ji,
			CelestialStemEnum.Geng, CelestialStemEnum.Xin, CelestialStemEnum.Ren, CelestialStemEnum.Gui, CelestialStemEnum.Jia,
			CelestialStemEnum.Gui, CelestialStemEnum.Ren, CelestialStemEnum.Xin, CelestialStemEnum.Geng, CelestialStemEnum.Ji,
			CelestialStemEnum.Wu, CelestialStemEnum.Ding, CelestialStemEnum.Bing, CelestialStemEnum.Yi, CelestialStemEnum.Jia,
			CelestialStemEnum.Jia, CelestialStemEnum.Jia, CelestialStemEnum.Jia, CelestialStemEnum.Jia,
			CelestialStemEnum.Ji, CelestialStemEnum.Jia, CelestialStemEnum.Ji, CelestialStemEnum.Jia,
		}

		for idx, n := range ns {
			original := actual
			actual = actual.Add(n)
			if actual != expect[idx] {
				t.Fatalf("add %d to %d should return %d, got %d",
					n,
					original,
					expect[idx],
					actual,
				)
			}
		}
	})

	t.Run("test String method", func(t *testing.T) {
		tbs := [10]CelestialStem{
			CelestialStemEnum.Jia, CelestialStemEnum.Yi,
			CelestialStemEnum.Bing, CelestialStemEnum.Ding,
			CelestialStemEnum.Wu, CelestialStemEnum.Ji,
			CelestialStemEnum.Geng, CelestialStemEnum.Xin,
			CelestialStemEnum.Ren, CelestialStemEnum.Gui,
		}
		expect := [10]string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}

		for idx, each := range tbs {
			actual := each.String()
			if actual != expect[idx] {
				t.Fatalf("string of %d shoud be %s, got %s", each, expect[idx], actual)
			}
		}
	})
}
