package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	t := ParseAndError("Schedule", layout, date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	targetTime := ParseAndError("HasPassed", layout, date)
	nowTime := time.Now()
	result := targetTime.Compare(nowTime)
	switch result {
	case -1:
		// past
		return true
	case 0:
		// equal, so target date is not past
		return false
	case 1:
		// future, so target date is not past
		return false
	default:
		panic("Not existed.")
	}
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
// afternoon (>= 12:00 and < 18:00)
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t := ParseAndError("IsAfternoonAppointment", layout, date)
	hour := t.Hour()
	if 12 <= hour && hour < 18 {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	t := ParseAndError("Description", layout, date)
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.",
		t.Weekday(), t.Month(), t.Day(), t.Year(), t.Hour(), t.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	layout := "1/2/2006 15:04:05"
	now := time.Now()
	date := fmt.Sprintf("9/15/%d 00:00:00", now.Year())
	t := ParseAndError("Description", layout, date)
	return t
}

// Parse and error output
// If error occurs, panic output funcName.
func ParseAndError(funcName, layout, date string) time.Time {
	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("err ", err.Error())
		panic("Error in " + funcName)
	}
	return t
}
