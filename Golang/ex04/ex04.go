package main

import (
	"errors"
	"fmt"
	"strconv"
)

/*
"//|\||"
"//\/\/"
"|||"
*/
//
//* is 42
// | is 124
// / is 47
// \ is 92

// One of the reasons I love Go: errors are NOT exceptions, but are treated as normal return values
func addOrnaments(tree string) (string, error) {
	var result = ""
	var straightBranchesFound = false
	//in golang strings are slices of bytes NOT slices of Chars, as we would expect them to be coming from other languages.
	//This means the a string like "ABC" is actually a slice with numeric values
	// the Pipesymbol we are looking for has the number 124 as it has in ASCII
	// runes are an alias for the type int32. Therefor we are looking for the ascii-value instead of "|"
	for _, runeValue := range tree {
		if runeValue == 124 {
			result += "*" + strconv.QuoteRune(runeValue) + "*"
			straightBranchesFound = true
		} else {
			result += strconv.QuoteRune(runeValue)
		}
	}
	if !straightBranchesFound {
		return "", errors.New("no straight branches found")
	}
	return result, nil
}

func main() {
	tree0 := `//|\||`
	tree1 := `//\/\/`
	tree2 := `|||`

	fmt.Println(addOrnaments(tree0))
	res, err := addOrnaments(tree1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	fmt.Println(addOrnaments(tree2))

}
