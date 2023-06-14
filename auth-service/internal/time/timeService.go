package time

import (
	"errors"
	t "time"
)

type TimeService struct {
	loc *t.Location
}

func NewTimeService(zoneId string) (*TimeService, error) {
	loc, err := validateTimeZone(zoneId)
	if err != nil {
		return nil, err
	}
	return &TimeService{loc: loc}, nil
}

func validateTimeZone(zoneId string) (*t.Location, error) {
	if zoneId == "" {
		return nil, errors.New("zone id can't be empty")
	}
	loc, err := t.LoadLocation(zoneId)
	if err != nil {
		return nil, err
	}
	return loc, nil
}

func (time *TimeService) Now() t.Time {
	return t.Now().In(time.loc)
}
