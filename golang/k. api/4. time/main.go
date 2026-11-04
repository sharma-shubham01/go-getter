package main

import (
	"fmt"
	"time"
)

// Wall Clock Time
// 1. Wall clock time represents real world time
// 2. Wall clock time can jump forwards and backwards
// 3. It gives you specific time relative to standard like UTC

// Monotonic Time
// 1. A monotonic time is an ever-increasing value
// 2. It doesn't jump forwards and backwards
// 3. Used to measure elapsed time since
// 4. Example - epoch time

// time package
// 1. Go provides support for time with standard lib time package
// 2. The representation of time in Go uses both - wall clock and monotonic time
// 3. For time telling operations, Go use wall clock time like current time
// 4. While performing time measuring operations, a monotonic clock reading is used
// 5. Two main types provided by Go -> Time and Duration

// time.Duration
// 1. Duration is represented as a span of time with int64 nanosecond count
// 2. It's the smallest unit of time in Go

// 3. Duration constants defined in the time package -
// a. Nanosecond Duration = 1
// b. Microsecond = 1000 * Nanosecond
// c. Millisecond = 1000 * Microsecond
// d. Second = 1000 * Millisecond
// e. Minute = 60 * Second
// f. Hour = 60 * Minute

// 4. Example -> d := 2 * time.Hour + 15 * time.Minute // duration - 2 hrs and 15 minutes
// 5. Duration can also be parsed using ParseDuration function

// 6. Duration also have methods to get individual time components like hours, minutes, seconds etc
// 7. Example -> d.Second(), d.Minutes() and d.Hours() - to get duration in their respective units

// 8. Functions Since() and Until() take Time and return Duration

// time.Time
// 1. Time struct is used to represent a specific point in time
// 2. It stores year, month, day, hour, minute, seconds, nanoseconds along with the location information
// 3. We should use Time as value and not a pointer

// 4. The Sub() method subtracts two instants of time producing a duration
// 5. The Add() method adds Time and Duration to produce Time

// 6. The zero value of type Time is January 1, year 1, 00:00:00.000000000 UTC
// 7. The IsZero() method should be used to check for the zero value of time

// 8. The equality(==) operator compares the time instant but also associated location & monotonic clock reading
// 9. Since time has location information, Time should be compared using Before(), After() and Equal() methods

func main() {
	// Get the current time
	now := time.Now()
	fmt.Println("Current time:", now)

	// Calculate time 2 hours from now
	twoHours := 2 * time.Hour
	futureTime := now.Add(twoHours)
	fmt.Println("Time 2 hours from now:", futureTime)

	// Measure the execution time of a code block
	start := time.Now()
	a := 1
	for i := 0; i <= 10000; i++ {
		a++
	}
	elapsed := time.Since(start)
	fmt.Println("Loop execution time:", elapsed)
}
