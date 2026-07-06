package intermediate

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("current time:", time.Now())

	specificTime := time.Date(2026, time.July, 5, 12, 0, 0, 0, time.UTC)
	fmt.Println("specific time:", specificTime)

	// time.Parse parses a time string using a layout based on Go's reference time:
	// "Mon Jan 2 15:04:05 MST 2006". The layout tells Go how to interpret the input
	// string — each component in the layout (year, month, day, hour, minute, etc.)
	// must match the corresponding part of the input string.
	// Below are several examples showing how different layouts parse the same date
	// (and time) in different formats.

	parsedTime, _ := time.Parse("2006-01-02", "2020-05-01") // 4-digit year, zero-padded month & day
	fmt.Println("parsed (2006-01-02):", parsedTime)

	parsedTime1, _ := time.Parse("06-01-02", "20-05-01") // 2-digit year (06), zero-padded month & day
	fmt.Println("parsed (06-01-02):", parsedTime1)

	parsedTime2, _ := time.Parse("06-1-2", "20-5-1") // 2-digit year, non-zero-padded month & day
	fmt.Println("parsed (06-1-2):", parsedTime2)

	parsedTime3, _ := time.Parse("06-1-2 15-04", "20-5-1 18-03") // adds hour (15) and minute (04)
	fmt.Println("parsed (06-1-2 15-04):", parsedTime3)

	// time.Format is the inverse of time.Parse: it takes a time.Time value and
	// formats it into a string using the same reference-time-based layout.
	// Here "Monday 06-01-02 04-15" produces output like "Weekday YY-MM-DD MM-HH"
	// (full weekday name, 2-digit year, zero-padded month/day, minute then hour).
	t := time.Now()
	fmt.Println("formatted time:", t.Format("Monday 06-01-02 04-15"))

	oneDayLater := t.Add(time.Hour * 24)
	fmt.Println("one day later:", oneDayLater)
	fmt.Println("weekday of one day later:", oneDayLater.Weekday())

	fmt.Println("rounded time:", t.Round(time.Hour))

	loc, _ := time.LoadLocation("Asia/Kolkata")
	t = time.Date(2026, time.July, 6, 14, 16, 40, 00, time.UTC)

	tLocal := t.In(loc)

	roundedTime := t.Round(time.Hour)
	roundedTimeLocal := roundedTime.In(loc)

	fmt.Println("original time (UTC):", t)
	fmt.Println("original time (Local):", tLocal)
	fmt.Println("rounded time (UTC):", roundedTime)
	fmt.Println("rounded time (Local):", roundedTimeLocal)

	fmt.Println("truncated time (to hour):", time.Now().Truncate(time.Hour))

	loc, _ = time.LoadLocation("America/New_York")
	timeInNY := time.Now().In(loc)
	fmt.Println("new york time:", timeInNY)

	t1 := time.Date(2026, time.July, 4, 12, 0, 0, 0, time.UTC)
	t2 := time.Date(2026, time.July, 4, 18, 0, 0, 0, time.UTC)

	duration := t2.Sub(t1)
	fmt.Println("duration:", duration)

	fmt.Println("is t2 after t1?:", t2.After(t1))
}
