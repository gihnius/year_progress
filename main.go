package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// EndOfYear get time of the end of year
func EndOfYear(t time.Time) time.Time {
	format := "2006-01-02 15:04:05"
	strYear := fmt.Sprintf("%d-12-31 23:59:59", t.Year())
	eoy, _ := time.Parse(format, strYear)
	return eoy
}

// ProgressBar print the progress bar
func ProgressBar() {
	passedDays := time.Now().YearDay()
	totalDays := EndOfYear(time.Now()).YearDay()
	progress := math.Round((float64(passedDays) / float64(totalDays)) * 100)
	bar := strings.Repeat("▓", int(progress)) + strings.Repeat("░", (100-int(progress)))
	fmt.Printf("%s %d%%(%d/%d)\n", bar, int(progress), passedDays, totalDays)
}

func main() {
	ProgressBar()
}
