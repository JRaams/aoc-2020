package helpers

import (
	"log"
	"time"
)

// Measure takes a starting time and message, and logs the amount of time that has passed after a function is done executing
func Measure(start time.Time, message string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", message, elapsed)
}
