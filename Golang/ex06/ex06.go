package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	//the path is dependant upon the working directory from which the script is ran from.
	//In this case that is /Golang
	file, err := os.Open("./ex06/present.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	//defer means close the file as soon as the surrounding function ends
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		log.Println(line)
	}

}
