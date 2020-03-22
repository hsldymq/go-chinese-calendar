package calendar

import (
	"testing"
	"time"
)

func TestNewTerrestrialBranchFromWord(t *testing.T) {
	inputs := []time.Time{
		time.Date(2020, 1, 1, 23, 0, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 23, 30, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 0, 30, 0, 0, time.UTC),

		time.Date(2020, 1, 1, 1, 0, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 1, 30, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 2, 0, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 2, 30, 0, 0, time.UTC),

		time.Date(2020, 1, 1, 3, 0, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 3, 30, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 4, 0, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 4, 30, 0, 0, time.UTC),

		time.Date(2020, 1, 1, 5, 0, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 5, 30, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 6, 0, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 6, 30, 0, 0, time.UTC),

		time.Date(2020, 1, 1, 7, 0, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 7, 30, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 8, 0, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 8, 30, 0, 0, time.UTC),

		time.Date(2020, 1, 1, 9, 0, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 9, 30, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 10, 0, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 10, 30, 0, 0, time.UTC),

		time.Date(2020, 1, 1, 11, 0, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 11, 30, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 12, 30, 0, 0, time.UTC),

		time.Date(2020, 1, 1, 13, 0, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 13, 30, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 14, 0, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 14, 30, 0, 0, time.UTC),

		time.Date(2020, 1, 1, 15, 0, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 15, 30, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 16, 0, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 16, 30, 0, 0, time.UTC),

		time.Date(2020, 1, 1, 17, 0, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 17, 30, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 18, 0, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 18, 30, 0, 0, time.UTC),

		time.Date(2020, 1, 1, 19, 0, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 19, 30, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 20, 0, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 20, 30, 0, 0, time.UTC),

		time.Date(2020, 1, 1, 21, 0, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 21, 30, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 22, 0, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 22, 30, 0, 0, time.UTC),
	}
	expect := []TerrestrialBranch{
		TerrestrialBranchEnum.Zi, TerrestrialBranchEnum.Zi, TerrestrialBranchEnum.Zi, TerrestrialBranchEnum.Zi,
		TerrestrialBranchEnum.Chou, TerrestrialBranchEnum.Chou, TerrestrialBranchEnum.Chou, TerrestrialBranchEnum.Chou,
		TerrestrialBranchEnum.Yin, TerrestrialBranchEnum.Yin, TerrestrialBranchEnum.Yin, TerrestrialBranchEnum.Yin,
		TerrestrialBranchEnum.Mao, TerrestrialBranchEnum.Mao, TerrestrialBranchEnum.Mao, TerrestrialBranchEnum.Mao,
		TerrestrialBranchEnum.Chen, TerrestrialBranchEnum.Chen, TerrestrialBranchEnum.Chen, TerrestrialBranchEnum.Chen,
		TerrestrialBranchEnum.Si, TerrestrialBranchEnum.Si, TerrestrialBranchEnum.Si, TerrestrialBranchEnum.Si,
		TerrestrialBranchEnum.Wu, TerrestrialBranchEnum.Wu, TerrestrialBranchEnum.Wu, TerrestrialBranchEnum.Wu,
		TerrestrialBranchEnum.Wei, TerrestrialBranchEnum.Wei, TerrestrialBranchEnum.Wei, TerrestrialBranchEnum.Wei,
		TerrestrialBranchEnum.Shen, TerrestrialBranchEnum.Shen, TerrestrialBranchEnum.Shen, TerrestrialBranchEnum.Shen,
		TerrestrialBranchEnum.You, TerrestrialBranchEnum.You, TerrestrialBranchEnum.You, TerrestrialBranchEnum.You,
		TerrestrialBranchEnum.Xu, TerrestrialBranchEnum.Xu, TerrestrialBranchEnum.Xu, TerrestrialBranchEnum.Xu,
		TerrestrialBranchEnum.Hai, TerrestrialBranchEnum.Hai, TerrestrialBranchEnum.Hai, TerrestrialBranchEnum.Hai,
	}

	for idx, each := range inputs {
		tb := NewTerrestrialBranchFromTime(each)
		if tb != expect[idx] {
			t.Fatalf("new celestial stem from time %02d:%02d, should return %d, got %d",
				each.Hour(),
				each.Minute(),
				expect[idx],
				tb,
			)
		}
	}
}

func TestNewTerrestrialBranchFromTime(t *testing.T) {
	inputs := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥", "我"}
	expect := []struct {
		TerrestrialBranch
		Valid bool
	}{
		{TerrestrialBranchEnum.Zi, true},
		{TerrestrialBranchEnum.Chou, true},
		{TerrestrialBranchEnum.Yin, true},
		{TerrestrialBranchEnum.Mao, true},
		{TerrestrialBranchEnum.Chen, true},
		{TerrestrialBranchEnum.Si, true},
		{TerrestrialBranchEnum.Wu, true},
		{TerrestrialBranchEnum.Wei, true},
		{TerrestrialBranchEnum.Shen, true},
		{TerrestrialBranchEnum.You, true},
		{TerrestrialBranchEnum.Xu, true},
		{TerrestrialBranchEnum.Hai, true},
		{TerrestrialBranchEnum.Zi, false},
	}

	for idx, each := range inputs {
		tb, v := NewTerrestrialBranchFromWord(each)
		if tb != expect[idx].TerrestrialBranch || v != expect[idx].Valid {
			t.Fatalf("new celestial stem from word %s, should return %d and %v, got %d and %v",
				each,
				expect[idx].TerrestrialBranch,
				expect[idx].Valid,
				tb,
				v,
			)
		}
	}
}

func TestTerrestrialBranch(t *testing.T) {
	t.Run("test Move method", func(t *testing.T) {
		actual := TerrestrialBranchEnum.Zi
		ns := []int{
			1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1,
			-1, -1, -1, -1, -1, -1,
			-1, -1, -1, -1, -1, -1,
			12, 24, -12, -24,
			18, 6, -18, -6,
		}
		expect := []TerrestrialBranch{
			TerrestrialBranchEnum.Chou, TerrestrialBranchEnum.Yin, TerrestrialBranchEnum.Mao, TerrestrialBranchEnum.Chen, TerrestrialBranchEnum.Si, TerrestrialBranchEnum.Wu,
			TerrestrialBranchEnum.Wei, TerrestrialBranchEnum.Shen, TerrestrialBranchEnum.You, TerrestrialBranchEnum.Xu, TerrestrialBranchEnum.Hai, TerrestrialBranchEnum.Zi,
			TerrestrialBranchEnum.Hai, TerrestrialBranchEnum.Xu, TerrestrialBranchEnum.You, TerrestrialBranchEnum.Shen, TerrestrialBranchEnum.Wei, TerrestrialBranchEnum.Wu,
			TerrestrialBranchEnum.Si, TerrestrialBranchEnum.Chen, TerrestrialBranchEnum.Mao, TerrestrialBranchEnum.Yin, TerrestrialBranchEnum.Chou, TerrestrialBranchEnum.Zi,
			TerrestrialBranchEnum.Zi, TerrestrialBranchEnum.Zi, TerrestrialBranchEnum.Zi, TerrestrialBranchEnum.Zi,
			TerrestrialBranchEnum.Wu, TerrestrialBranchEnum.Zi, TerrestrialBranchEnum.Wu, TerrestrialBranchEnum.Zi,
		}

		for idx, n := range ns {
			original := actual
			actual = actual.Move(n)
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

	t.Run("test Month method", func(t *testing.T) {
		tbs := [12]TerrestrialBranch{
			TerrestrialBranchEnum.Zi, TerrestrialBranchEnum.Chou, TerrestrialBranchEnum.Yin, TerrestrialBranchEnum.Mao,
			TerrestrialBranchEnum.Chen, TerrestrialBranchEnum.Si, TerrestrialBranchEnum.Wu, TerrestrialBranchEnum.Wei,
			TerrestrialBranchEnum.Shen, TerrestrialBranchEnum.You, TerrestrialBranchEnum.Xu, TerrestrialBranchEnum.Hai,
		}
		expect := [12]int{
			11, 12, 1, 2,
			3, 4, 5, 6,
			7, 8, 9, 10,
		}

		for idx, each := range tbs {
			actual := each.Month()
			if actual != expect[idx] {
				t.Fatalf("month of %d shoud be %d, got %d", each, expect[idx], actual)
			}
		}
	})

	t.Run("test ZodiacSign method", func(t *testing.T) {
		tbs := [12]TerrestrialBranch{
			TerrestrialBranchEnum.Zi, TerrestrialBranchEnum.Chou, TerrestrialBranchEnum.Yin, TerrestrialBranchEnum.Mao,
			TerrestrialBranchEnum.Chen, TerrestrialBranchEnum.Si, TerrestrialBranchEnum.Wu, TerrestrialBranchEnum.Wei,
			TerrestrialBranchEnum.Shen, TerrestrialBranchEnum.You, TerrestrialBranchEnum.Xu, TerrestrialBranchEnum.Hai,
		}
		expect := [12]ZodiacSign{
			ZodiacSignEnum.Rat, ZodiacSignEnum.Ox, ZodiacSignEnum.Tiger, ZodiacSignEnum.Rabbit,
			ZodiacSignEnum.Dragoon, ZodiacSignEnum.Snake, ZodiacSignEnum.Horse, ZodiacSignEnum.Goat,
			ZodiacSignEnum.Monkey, ZodiacSignEnum.Rooster, ZodiacSignEnum.Dog, ZodiacSignEnum.Pig,
		}

		for idx, each := range tbs {
			actual := each.ZodiacSign()
			if actual != expect[idx] {
				t.Fatalf("zodiac sign of %d shoud be %d, got %d", each, expect[idx], actual)
			}
		}
	})

	t.Run("test String method", func(t *testing.T) {
		tbs := [12]TerrestrialBranch{
			TerrestrialBranchEnum.Zi, TerrestrialBranchEnum.Chou, TerrestrialBranchEnum.Yin, TerrestrialBranchEnum.Mao,
			TerrestrialBranchEnum.Chen, TerrestrialBranchEnum.Si, TerrestrialBranchEnum.Wu, TerrestrialBranchEnum.Wei,
			TerrestrialBranchEnum.Shen, TerrestrialBranchEnum.You, TerrestrialBranchEnum.Xu, TerrestrialBranchEnum.Hai,
		}
		expect := [12]string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}

		for idx, each := range tbs {
			actual := each.String()
			if actual != expect[idx] {
				t.Fatalf("string of %d shoud be %s, got %s", each, expect[idx], actual)
			}
		}
	})

}
