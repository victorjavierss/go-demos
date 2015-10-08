package main

import (
	"challengesutils"
	"strings"
	"strconv"
	"regexp"
	"fmt"
)

func main() {

    _, input := challengesutils.StdinReaderWithTestCases()

    for v := 0; v < len(input); v++ {
        
        //fmt.Println("========== Case ===============")
        haystackSize := strings.Split(input[v], " ")

        haystackRows,_ := strconv.Atoi(haystackSize[0])
      	haystackOffset := v + haystackRows + 1;

       	haystack := append([]string{}, input[v+1:haystackOffset]...)
        
        //fmt.Printf("Haystack: %v \n", haystack)
        
		needleSize := strings.Split(input[v+haystackRows+1], " ")
		needleRows,_ := strconv.Atoi(needleSize[0])
		
		needleOffset := v+haystackRows+2;
		needle := append([]string{}, input[needleOffset:needleRows+needleOffset]...)
	
		//fmt.Printf("Needle: %v \n", needle)
	
		if searchIn(haystack, needle) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

		v += needleRows + haystackRows + 1
		
    }
}

func searchIn( haystack []string,  needle []string) bool {
	
	var keyMatches [][]int

	result := true

	rowsToIterate := len(haystack) - len(needle)

	for haystackIndex := 0; haystackIndex < rowsToIterate; haystackIndex++ {

		haystackRow := haystack[haystackIndex]
		
		if strings.Index(haystackRow, needle[0]) > -1 {
			
			for needleIndex, needleRow := range needle {
				needleKey := regexp.MustCompile(needleRow)

				matches := needleKey.FindAllStringSubmatchIndex(haystack[haystackIndex+needleIndex], -1)
				
				//fmt.Printf("Search [%v] => [%v] : [%v]\n", needleRow, haystack[haystackIndex+needleIndex], matches != nil)
				
				if matches != nil {
					if needleIndex == 0 {
						keyMatches = matches
					} else {
						keyMatches = getKeyMatches(keyMatches, matches)
						result = len(keyMatches) > 0
					}
				} else {
					result = false
				}
			}
			
			//fmt.Printf("has the full pattern: %v \n", result)
			
			if result {
				break
			}
		}
	}
	
	return result
}

func getKeyMatches( originalKeys [][]int, newKeys [][]int) [][]int {
	var matchesIntersect [][]int

	aMap := map[int]int{}

	//fmt.Printf("original Keys: %v \n", originalKeys)
	//fmt.Printf("Keys to Find: %v \n", newKeys)

	for i, keys := range originalKeys {
		if _, inMap := aMap[ keys[0] ]; ! inMap {
		    aMap[ keys[0] ] = i
		}
	}

	for _, keys := range newKeys {
		
		if i, inMap := aMap[ keys[0] ]; inMap {
		    matchesIntersect =  append(matchesIntersect, originalKeys[i])
		}
	}

	//fmt.Printf("Intersect: %v \n", matchesIntersect)
	
	return matchesIntersect
}