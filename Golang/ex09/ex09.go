package main

import (
	"fmt"
	"math"
	"time"
)

// The elves worry they can't build all presents until christmas.
// Have a look at daytime-library and create a function that calculates the number of days until the next christmas

func main() {
	currentTime := time.Now().Local()
	christmasEveCurrentYear := time.Date(currentTime.Year(), 12, 24, 18, 0, 0, 0, time.Local)

	if currentTime.Before(christmasEveCurrentYear) {
		fmt.Println("Days until christmas THIS YEAR", math.Round(christmasEveCurrentYear.Sub(currentTime).Hours()/24))
	} else {
		updatedYear := time.Date(christmasEveCurrentYear.Year()+1, christmasEveCurrentYear.Month(), christmasEveCurrentYear.Day(), christmasEveCurrentYear.Hour(),
			christmasEveCurrentYear.Minute(), christmasEveCurrentYear.Second(), christmasEveCurrentYear.Nanosecond(), time.Local)
		fmt.Println("Days until christmas NEXT YEAR", math.Round(updatedYear.Sub(currentTime).Hours()/24))
	}

}
