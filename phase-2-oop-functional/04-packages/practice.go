package phase2packages

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// User represents a user profile
type User struct {
	Name string
	Age  int
}

// FormatUser returns a formatted string representation of User
// Format: "User: <NAME> (Age: <AGE>)" where NAME is uppercase
func FormatUser(u User) string {
	return fmt.Sprintf("User: %s (Age: %d)", strings.ToUpper(u.Name), u.Age)
}

// ParseAndSum converts two string numbers to integers and returns their sum.
// If either fails to parse, return an error.
func ParseAndSum(s1, s2 string) (int, error) {
	v1, err := strconv.Atoi(s1)
	if err != nil {
		return 0, err
	}
	v2, err := strconv.Atoi(s2)
	if err != nil {
		return 0, err
	}
	return v1 + v2, nil
}

// DaysUntilBirthday calculates the number of days from referenceTime until the next birthday.
// birthdayStr is in "YYYY-MM-DD" format.
func DaysUntilBirthday(referenceTime time.Time, birthdayStr string) (int, error) {
	bday, err := time.Parse("2006-01-02", birthdayStr)
	if err != nil {
		return 0, err
	}

	// Set next birthday to be in the same year as referenceTime
	nextBday := time.Date(referenceTime.Year(), bday.Month(), bday.Day(), 0, 0, 0, 0, referenceTime.Location())
	
	// If birthday already passed this year, set to next year
	refDate := time.Date(referenceTime.Year(), referenceTime.Month(), referenceTime.Day(), 0, 0, 0, 0, referenceTime.Location())
	if nextBday.Before(refDate) {
		nextBday = nextBday.AddDate(1, 0, 0)
	}

	diff := nextBday.Sub(refDate)
	days := int(diff.Hours() / 24)
	return days, nil
}
