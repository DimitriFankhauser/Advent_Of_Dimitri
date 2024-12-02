/* okay now that we've done the warm-up, let's do some small things with parameters
This main-function will be used to calculate volatage according to ohm's law (U=R * I)
you receive two parameters, r for resistance AND i for intensity/current.
validate the input (no negative values) and return the voltage
if you receive an invalid input, return -1 instead of the calculation */

package main

import "fmt"

func calculateVoltage(r int64, i int64) int64 {
	return r * i
}

func main() {
	var voltage int64 = calculateVoltage(10, 20)
	fmt.Println(voltage)
}
