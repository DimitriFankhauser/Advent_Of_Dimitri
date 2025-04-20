package main

import (
	"fmt"
	"time"
)

/*Santa likes to drink coffee with his ice-bear friends on wednesdays.
To be in sync for the upcoming 100 years Santa needs a list of every year up to 2124 when Christmas Eve is on a Wednesday. */

func main() {
	for i := 2024; i <= 2124; i++ {
		c := time.Date(i, 12, 24, 10, 10, 0, 0, time.Local)
		if c.Weekday() == time.Wednesday {
			fmt.Println(c.Year())
		}
	}

}
