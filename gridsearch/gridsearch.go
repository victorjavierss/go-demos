package main

import (
	"fmt"
	"challengesutils"
)

func main() {
   
    //var a [][]string = challengesutils.StdinReaderMatrix();
    
    testCases, input := challengesutils.StdinReaderWithTestCases()
    
    fmt.Println(testCases)
    
    for _, element := range input {
    	fmt.Println(element)
    }
    
}