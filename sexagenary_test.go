package chinese_calendar

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSexagenary(t *testing.T) {
	t.Run("test sexagenary number", func(t *testing.T) {
		for i := 0; i < 60; i++ {
			require.Equal(t, sexagenary[i], SexagenaryNumber(i))
			require.Equal(t, sexagenary[i], SexagenaryNumber(i-3, 3))
			require.Equal(t, sexagenary[i], SexagenaryNumber(i+60))
			require.Equal(t, sexagenary[i], SexagenaryNumber(i-60))
			require.Equal(t, sexagenary[i], SexagenaryNumber(i-120))
		}
	})

}
