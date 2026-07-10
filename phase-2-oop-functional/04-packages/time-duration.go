package phase2packages

import (
	"fmt"
	"time"
)

// time-duration.go demonstrates the time package
// Essential for any Go program dealing with time and delays

// ===== TIME BASICS =====

// timeBasics demonstrates basic time operations
func TimeBasics() {
	fmt.Println("\n=== Time Basics ===")

	// Current time
	now := time.Now()
	fmt.Println("Current time:", now)
	fmt.Printf("Type: %T\n", now)

	// Components
	fmt.Printf("Year: %d\n", now.Year())
	fmt.Printf("Month: %v\n", now.Month())
	fmt.Printf("Day: %d\n", now.Day())
	fmt.Printf("Hour: %d\n", now.Hour())
	fmt.Printf("Minute: %d\n", now.Minute())
	fmt.Printf("Second: %d\n", now.Second())
}

// ===== DURATION =====

// durationBasics demonstrates duration operations
func DurationBasics() {
	fmt.Println("\n=== Duration ===")

	// Create durations
	d1 := 5 * time.Second
	d2 := 100 * time.Millisecond
	d3 := 2 * time.Hour

	fmt.Printf("5 seconds: %v\n", d1)
	fmt.Printf("100 milliseconds: %v\n", d2)
	fmt.Printf("2 hours: %v\n", d3)

	// Convert to different units
	fmt.Printf("5 seconds in nanoseconds: %d\n", d1.Nanoseconds())
	fmt.Printf("5 seconds in milliseconds: %d\n", d1.Milliseconds())
	fmt.Printf("5 seconds in seconds: %.2f\n", d1.Seconds())
}

// ===== FORMATTING TIME =====

// formatTime demonstrates time formatting
func FormatTime() {
	fmt.Println("\n=== Formatting Time ===")

	now := time.Now()

	// Format using predefined layouts
	fmt.Println("RFC3339:", now.Format(time.RFC3339))
	fmt.Println("RFC1123:", now.Format(time.RFC1123))
	fmt.Println("Kitchen:", now.Format(time.Kitchen))
	fmt.Println("Stamp:", now.Format(time.Stamp))

	// Custom format (using reference time: Mon Jan 2 15:04:05 MST 2006)
	custom := now.Format("2006-01-02 15:04:05")
	fmt.Println("Custom (YYYY-MM-DD HH:MM:SS):", custom)

	custom2 := now.Format("Monday, January 2, 2006")
	fmt.Println("Custom (Day, Month D, YYYY):", custom2)
}

// ===== PARSING TIME =====

// parseTime demonstrates parsing time from strings
func ParseTime() {
	fmt.Println("\n=== Parsing Time ===")

	// Parse using a layout
	layout := "2006-01-02"
	timeStr := "2024-07-10"


t, err := time.Parse(layout, timeStr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Parsed '%s': %v\n", timeStr, t)
	}

	// Another example
	layout2 := "Mon Jan 2 15:04:05 MST 2006"
	timeStr2 := "Wed Jul 10 14:30:45 UTC 2024"

	t2, err := time.Parse(layout2, timeStr2)
	if err == nil {
		fmt.Printf("Parsed '%s': %v\n", timeStr2, t2)
	}
}

// ===== TIME ARITHMETIC =====

// timeArithmetic demonstrates adding/subtracting times
func TimeArithmetic() {
	fmt.Println("\n=== Time Arithmetic ===")

	now := time.Now()
	fmt.Println("Now:", now.Format("2006-01-02 15:04:05"))

	// Add duration
	next := now.Add(24 * time.Hour)
	fmt.Println("Tomorrow:", next.Format("2006-01-02 15:04:05"))

	// Subtract duration
	last := now.Add(-7 * 24 * time.Hour)
	fmt.Println("Last week:", last.Format("2006-01-02 15:04:05"))

	// Difference between times
	diff := next.Sub(now)
	fmt.Printf("Difference: %v\n", diff)
}

// ===== SLEEP AND TICKERS =====

// sleepExample demonstrates time.Sleep
func SleepExample() {
	fmt.Println("\n=== Sleep ===")

	fmt.Println("Start:", time.Now().Format("15:04:05"))
	fmt.Println("Sleeping for 2 seconds...")
	time.Sleep(2 * time.Second)
	fmt.Println("End:", time.Now().Format("15:04:05"))
}

// ===== UNIX TIME =====

// unixTime demonstrates Unix timestamps
func UnixTime() {
	fmt.Println("\n=== Unix Time ===")

	now := time.Now()

	// Get Unix timestamp
	unixtimestamp := now.Unix()
	fmt.Printf("Unix timestamp: %d\n", unixtimestamp)

	unixtimestampNano := now.UnixNano()
	fmt.Printf("Unix nano timestamp: %d\n", unixtimestampNano)

	// Convert Unix timestamp back to time
	t := time.Unix(unixtimestamp, 0)
	fmt.Printf("From Unix timestamp: %v\n", t)
}

// ===== COMPARISONS =====

// timeComparison demonstrates comparing times
func TimeComparison() {
	fmt.Println("\n=== Time Comparison ===")


t1 := time.Now()
	time.Sleep(100 * time.Millisecond)
	t2 := time.Now()

	fmt.Printf("t1 < t2: %v\n", t1.Before(t2))
	fmt.Printf("t2 > t1: %v\n", t2.After(t1))
	fmt.Printf("t1 == t1: %v\n", t1.Equal(t1))
}

// PackageTimePlayground runs all time examples
func PackageTimePlayground() {
	fmt.Println("\n========== TIME PACKAGE ==========")
	TimeBasics()
	DurationBasics()
	FormatTime()
	ParseTime()
	TimeArithmetic()
	SleepExample()
	UnixTime()
	TimeComparison()
}
