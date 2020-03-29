package calendar

import (
	"time"

	"github.com/hsldymq/go-chinese-calendar/sexagenary"
)

type SexagenaryTime time.Time

func NewSexagenaryTime(t time.Time, Timezone ...*time.Location) SexagenaryTime {
	tz := baseTimezone
	if len(Timezone) > 0 {
		tz = Timezone[0]
	}
	return SexagenaryTime(t.In(tz))
}

func (t SexagenaryTime) Year() sexagenary.SexagenaryTerm {
	// TODO
	return sexagenary.SexagenaryTermEnum.JiaZi
}

func (t SexagenaryTime) Month() sexagenary.SexagenaryTerm {
	// TODO
	return sexagenary.SexagenaryTermEnum.JiaZi
}

func (t SexagenaryTime) Day() sexagenary.SexagenaryTerm {
	// TODO
	return sexagenary.SexagenaryTermEnum.JiaZi
}

func (t SexagenaryTime) Hour() sexagenary.SexagenaryTerm {
	// TODO
	return sexagenary.SexagenaryTermEnum.JiaZi
}
