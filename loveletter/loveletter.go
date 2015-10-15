package main

import (
	"challengesutils"
	"fmt"
)

func main() {
	_, input := challengesutils.StdinReaderWithTestCases()
	
	const minLetter = 97
	
	
	for _, word := range input {
			
	minTicks := 0
	
	if ! challengesutils.IsPalindrome(word) {

			ticksForward := 1
			
			copyWord := word
			
			loopForward:
			for i := 0; i < len(word); i++ {
				for letter := copyWord[i]-1; letter >= minLetter; letter-- {
					copyWord = copyWord[0:i] + string(letter) + copyWord[i+1:len(word)]
					if challengesutils.IsPalindrome(copyWord) {
						break loopForward
					} else {
						ticksForward++
					}
				} 
			}
			
			ticksBackwards := 1
			copyWord = word
		//	fmt.Printf("FWD: %v\n", ticksForward)
			
			loopBackwards: 
			for i := len(word)-1; i > 0; i-- {
				for letter := copyWord[i]-1; letter >= minLetter; letter-- {
					copyWord = copyWord[0:i] + string(letter) + copyWord[i+1:len(word)]
					if challengesutils.IsPalindrome(copyWord) {
						break loopBackwards
					} else {
						ticksBackwards++
					}
				}
			}

		//	fmt.Printf("BAC: %v\n", ticksBackwards)

			if ticksBackwards < ticksForward {
				minTicks = ticksBackwards
			} else {
				minTicks = ticksForward	
			}
		} 
		
		fmt.Printf("%v\n", minTicks);
	}
}