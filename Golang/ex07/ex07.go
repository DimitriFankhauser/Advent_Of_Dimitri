package main

import (
	"bufio"
	"log"
	"os"
)

/*
The elves are working very hard to build and package all the presents for the nice kids.
Some presents take up more space than others, which is why they need to be packaged in smaller and bigger crates (this would be way funnier, if this challenge were written in Rust)
In Example a Candybar (7 letters) takes up more space than bunny (5 letters)
Now that we can read input from files, we will also write to files. You will receive a list of Presents by reading from the presents.txt.
If the present takes up a lot of space (>5 letters), write them into the file big_crates.txt. If they take up less space write them into the list small_crates.
If you run into trouble or want to reset the challenge, run the unit-test "reset-challenge" in sol07
The elves thank you for your service*/

func openFile(path string) *os.File {
	var file, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0660)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return file

}
func main() {
	file0, _ := os.Open("./ex07/presents.txt")
	file1 := openFile("./ex07/small_crates.txt")
	file2 := openFile("./ex07/big_crates.txt")

	//defer means close the file as soon as the surrounding function ends
	defer file0.Close()
	defer file1.Close()
	defer file2.Close()

	scanner := bufio.NewScanner(file0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 5 {
			// It would be recommended to log errors here, because we are writing to a resource.
			file1.WriteString(line + "\n")
		} else {
			file2.WriteString(line + "\n")
		}
	}

}
