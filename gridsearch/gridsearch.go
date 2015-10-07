package main

import (
	"fmt"
	"challengesutils"
)

func main() {
    testCases, input := challengesutils.StdinReaderWithTestCases()
    
    fmt.Println(testCases)
    
    for _, element := range input {
    	fmt.Println(element)
    }
}