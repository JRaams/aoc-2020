package helpers

import (
	"fmt"
	"time"
)

// Measure takes a starting time and message, and logs the amount of time that has passed after a function is done executing
func Measure(start time.Time, message string) {
	elapsed := time.Since(start)
	if len(message) > 0 {
		fmt.Printf("(%s: %s) ", message, elapsed)
	} else {
		fmt.Printf("(%s) ", elapsed)
	}
}
