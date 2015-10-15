package main

import (
	"challengesutils"
	"fmt"
)

func main() {
	_, input := challengesutils.StdinReaderWithTestCases()
	for _, word := range input {
		if challengesutils.IsPalindrome(word) {
			fmt.Println("-1")
		} else {
			for i,_ := range word {
				wordb := word[0:i] + word[i+1:len(word)]
				if challengesutils.IsPalindrome(wordb) {
					fmt.Printf("%v\n", i)
					break
				}
			}
		}	
	}	
}