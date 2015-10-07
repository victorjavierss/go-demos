package challengesutils

import (
	"io/ioutil"
	"strings"
	"os"
	"strconv"
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
    
    noTestCases, _ := strconv.Atoi(allInput[0])

    testCases := make([]string, len(allInput)-1)

    for v := 1; v < len(allInput); v++ {
        testCases[v-1] = allInput[v]
    }
    
    return noTestCases, testCases
}