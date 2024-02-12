package log

import (
	"fmt"
	"time"
)

// Println prints multiple values along with the current time and date.
func Println(values ...interface{}) {
	currentTime := time.Now()
	timeFormatted := currentTime.Format("2006-01-02 15:04:05") // Customize the format as needed
    fmt.Println("{")
	fmt.Println(timeFormatted)
	// Print each value
	for _, val := range values {
		fmt.Print(val, " ")
	}
	// Add a newline at the end
	fmt.Println()
	fmt.Println("}")
}
