package challengesutils

import (
	"io/ioutil"
	"strings"
	"os"
)


func StdinReaderMatrix() [][]string {
	
    initialMatrix := StdinReaderArray()
    finalMatrix := make([][]string, len(initialMatrix))
    for index, element := range initialMatrix {
  		finalMatrix[index] = strings.Split(element, " ")
	}
    return finalMatrix;
}


func StdinReaderArray() []string {
	bytes, _ := ioutil.ReadAll(os.Stdin)
    stdinInput := string(bytes)
    return strings.Split(stdinInput, "\n")
}


func StdinReaderWithTestCases() (int, []string) {
    
    allInput := StdinReaderArray()
    
    return 0, allInput
}