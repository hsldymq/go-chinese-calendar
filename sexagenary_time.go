package calendar

import (
	"time"

	"github.com/hsldymq/go-chinese-calendar/sexagenary"
)

type SexagenaryTime struct {
	Year  sexagenary.SexagenaryTerm
	Month sexagenary.SexagenaryTerm
	Day   sexagenary.SexagenaryTerm
	Hour  sexagenary.SexagenaryTerm
}

func NewSexagenaryTime(t time.Time, Timezone ...*time.Location) SexagenaryTime {
	tz := baseTimezone
	if len(Timezone) > 0 {
		tz = Timezone[0]
	}
	t.In(tz)
	st := SexagenaryTime{}

	return st
}
