package main

import (
	"fmt"
	"challengesutils"
)

func main() {
   
    //var a [][]string = challengesutils.StdinReaderMatrix();
    
    c := challengesutils.StdinReaderArray()
    
    fmt.Println(c)
}