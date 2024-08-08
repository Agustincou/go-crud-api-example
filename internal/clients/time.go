package clients

import "time"

type Time interface {
	Now() time.Time
}

type systemTime struct{}

func NewSystemTime() Time {
	return &systemTime{}
}

func (s *systemTime) Now() time.Time {
	return time.Now()
}
