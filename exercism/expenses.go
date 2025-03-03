package expenses

import "errors"

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	out := []Record{}
	for _, rec := range in {
		if predicate(rec) {
			out = append(out, rec)
		}
	}
	return out
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		if p.From <= r.Day && r.Day <= p.To {
			return true
		}
		return false
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		if r.Category == c {
			return true
		}
		return false
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	sum := 0.0
	for _, v := range in {
		if p.From <= v.Day && v.Day <= p.To {
			sum += v.Amount
		}
	}
	return sum
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	hit := false
	// Category exists check
	for _, v := range in {
		if v.Category == c {
			hit = true
		}
	}
	if !hit {
		return 0, errors.New("unknown category " + c)
	}

	// ByDaysPeriod and ByCategory filtering
	recordFiltered := Filter(in, ByDaysPeriodAndCategory(p, c))
	return TotalByPeriod(recordFiltered, p), nil
}

func ByDaysPeriodAndCategory(p DaysPeriod, c string) func(Record) bool {
	return func(r Record) bool {
		daysPeriodCheck := ByDaysPeriod(p)
		categoryCheck := ByCategory(c)
		return daysPeriodCheck(r) && categoryCheck(r)
	}
}
