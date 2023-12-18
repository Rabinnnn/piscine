package main

import (
	"fmt"
	"strings"
)

func Join(strs []string, sep string) string {
	// Use strings.Join to concatenate the strings with the separator
	result := strings.Join(strs, sep)

	return result
}

func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(Join(toConcat, ":"))
}


/***ALTERNATIVE THAT DOESN'T IMPORT STRINGS***/

/***
func Join(strs []string, sep string) string {
	// Check if the slice is empty
	if len(strs) == 0 {
		return ""
	}

	// Initialize the result with the first string
	result := strs[0]

	// Iterate through the remaining strings and concatenate with the separator
	for _, s := range strs[1:] {
		result += sep + s
	}

	return result
}
***/
