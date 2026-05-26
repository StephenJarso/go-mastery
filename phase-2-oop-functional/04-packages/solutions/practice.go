package solutions

import (
	"strings"
	"strconv"
	"time"
	"errors"
)


func FormatURLPath(title string) string {
	lower := strings.ToLower(title)
	fields := strings.Fields(lower)
	joined := strings.Join(fields, "-")
	
	// Remove non-alphanumeric except hyphen
	var sb strings.Builder
	for _, r := range joined {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

func ParseAndSumCSV(line string) (int, error) {
	parts := strings.Split(line, ",")
	sum := 0
	for _, p := range parts {
		trimmed := strings.TrimSpace(p)
		if trimmed == "" {
			continue
		}
		val, err := strconv.Atoi(trimmed)
		if err != nil {
			return 0, err
		}
		sum += val
	}
	return sum, nil
}

func DaysSinceBirth(birthdate string, now time.Time) (int, error) {
	t, err := time.Parse("2006-01-02", birthdate)
	if err != nil {
		return 0, err
	}
	if t.After(now) {
		return 0, errors.New("birthdate is in the future")
	}
	duration := now.Sub(t)
	return int(duration.Hours() / 24), nil
}
