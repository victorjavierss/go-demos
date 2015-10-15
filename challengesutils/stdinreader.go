package challengesutils

import (
	"io/ioutil"
	"strconv"
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
    noTestCases, _ := strconv.Atoi( strings.Trim(allInput[0], "￼") )
    testCases := append([]string{}, allInput[1:len(allInput)]...)
    return noTestCases, testCases
}

func IsPalindrome(word string) bool {
	j := len(word) - 1
	for i := 0 ; i <= j ; i++ {
        if word[i] != word[j] {
        	return false
        }
        j--
    }
	return true;
}