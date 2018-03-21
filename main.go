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

// Calc get the values
func Calc() (percent, passedDays, totalDays int) {
	passedDays = time.Now().YearDay()
	totalDays = EndOfYear(time.Now()).YearDay()
	percent = int(math.Round((float64(passedDays) / float64(totalDays)) * 100))
	return
}

// ProgressBar print the progress bar
func ProgressBar() {
	percent, passedDays, totalDays := Calc()
	bar := strings.Repeat("▓", percent) + strings.Repeat("░", (100-percent))
	fmt.Printf("%s %d%%(%d/%d)\n", bar, percent, passedDays, totalDays)
}

// ProgressGrid print as grid
func ProgressGrid() {
	percent, passedDays, totalDays := Calc()
	for i := 1; i <= totalDays; i++ {
		if i <= passedDays {
			fmt.Print("■")
		} else {
			fmt.Print("□")
		}
		if i%21 == 0 {
			fmt.Print("\n")
		}
	}
	fmt.Printf(" %d%%(%d/%d)\n", percent, passedDays, totalDays)
}

func main() {
	ProgressBar()
	ProgressGrid()
}
