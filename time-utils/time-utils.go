package timeutils

import "time"

type MonthRange struct {
	Start time.Time
	Stop  time.Time
}

func MonthChunks(start, end time.Time) []MonthRange {
	// inclusive start, exclusive stop
	if !start.Before(end) {
		return nil
	}
	var out []MonthRange
	cur := time.Date(start.Year(), start.Month(), 1, 0, 0, 0, 0, time.UTC)
	// include partial first month
	if cur.Before(start) {
		cur = time.Date(start.Year(), start.Month(), 1, 0, 0, 0, 0, time.UTC)
	}
	for cur.Before(end) {
		next := cur.AddDate(0, 1, 0)
		s := cur
		if s.Before(start) {
			s = start
		}
		e := next
		if e.After(end) {
			e = end
		}
		out = append(out, MonthRange{Start: s, Stop: e})
		cur = next
	}
	return out
}
