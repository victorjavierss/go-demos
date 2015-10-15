package main

import (
	"challengesutils"
	"fmt"
//	"strings"
)

func main() {
	//_, input := challengesutils.StdinReaderWithTestCases()
	
	const minLetter = 97
	
	word := "cba"
			// aacd
			// aabd
			// aaad
			// aaac
			// aaab
			// aaaa
			
	minTicks := 0
	
	if ! challengesutils.IsPalindrome(word) {
		
		ticksForward := 1
		
		copyWord := word
		
		for i := 0; i < len(word); i++ {
			for letter := copyWord[i]-1; letter >= minLetter; letter-- {
				copyWord = copyWord[0:i] + string(letter) + copyWord[i+1:len(word)]
				if challengesutils.IsPalindrome(copyWord) {
					break
				} else {
					ticksForward++
				}
			} 
		}
		
		ticksBackwards := 1
		copyWord = word
		
		for i := len(word)-1; i > 0; i-- {
			for letter := copyWord[i]-1; letter >= minLetter; letter-- {
				copyWord = copyWord[0:i-1] + string(letter) + copyWord[i:len(word)]
				if challengesutils.IsPalindrome(copyWord) {
					break
				} else {
					ticksBackwards++
				}
			}
		}
		
		
		if ticksBackwards < ticksForward {
			minTicks = ticksBackwards
		} else {
			minTicks = ticksForward	
		}

	} 
	
	fmt.Printf("%v", minTicks);
	
}