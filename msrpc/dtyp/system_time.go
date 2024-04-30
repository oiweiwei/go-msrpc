package dtyp

import "time"

func (s *SystemTime) AsTime() time.Time {
	return time.Date(int(s.Year), time.Month(s.Month), int(s.Day), int(s.Hour), int(s.Minute), int(s.Second), 1000*int(s.Milliseconds), nil)
}
