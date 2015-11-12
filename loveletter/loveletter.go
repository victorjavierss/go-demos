package main

import (
	"challengesutils"
	"fmt"
	"math"
)

func main() {
	_, input := challengesutils.StdinReaderWithTestCases()
	
	const minLetter = 97
	
	
	for _, word := range input {
	
		j := len(word) - 1
		ticks := 0

		for i := 0 ; i <= j ; i++ {
        	if word[i] != word[j] {
        		charA := int(word[i]) 
        		charB := int(word[j])
        		ticks += int( math.Abs( float64(charA) - float64(charB)  ) )	
        	}
        	j--
    	}
		
		fmt.Printf("%v\n", ticks)
	}
}