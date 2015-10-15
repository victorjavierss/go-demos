package main

import (
	"challengesutils"
	"strings"
	"fmt"
)

func main() {
	testCases, input := challengesutils.StdinReaderWithTestCases()
	testCases = testCases*2
	for i := 0; i < testCases; i+=2 {
		needle := input[i]
		haystack := input[i+1]
		if hasString(haystack, needle) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}	
	}
}

func hasString(needle string, haystack string) bool {
	
	// strings.ContainsAny(haystack, needle); 
	
	for _, char := range needle {
		if strings.IndexRune(haystack, char) > -1 {
			return true	
		}
	}
	
	return false;
}