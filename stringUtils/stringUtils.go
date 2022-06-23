package stringUtils

import (
	"bufio"
	"fmt"
	"strings"
)

// RemoveFirst
/*
Remove the prefix characters and first separator

s -> ABCDEF
separator -> CD
-----
output:
EF
*/
func RemoveFirst(s string, separator string) (string, bool) {
	index := strings.Index(s, separator)
	if index == -1 {
		return s, false
	} else {
		return s[index+len(separator):], true
	}
}

func IndexOfString(arr []string, search string) int {
	for i, v := range arr {
		if v == search {
			return i
		}
	}
	return -1
}

// SplitCpExcel
/*
 split the {msg} by row with "\n" and split each row by "\n"

 input:
 "a\tb\tc\nd\te\tf"

 output:
 [
 [a,b,c]
 [d,e,f]
 ]
*/
func SplitCpExcel(msg string, trimSpace bool) (line [][]string) {
	var lines [][]string
	tempSplit := strings.Split(msg, "\n")
	for index, _ := range tempSplit {
		if strings.HasSuffix(tempSplit[index], "\r") {
			tempSplit[index] = strings.TrimRight(tempSplit[index], "\r")
		}
		line := tempSplit[index]

		if index == len(tempSplit)-1 && tempSplit[index] == "" {
		} else {
			if trimSpace {
				var tempLine []string
				for _, word := range strings.Split(line, "\t") {
					tempLine = append(tempLine, strings.Trim(word, " "))
				}
				lines = append(lines, tempLine)
			} else {
				lines = append(lines, strings.Split(line, "\t"))
			}
		}
	}
	return
}

// StringToBufIoReader
/*
Convert a string into *bufio.Reader
*/
func StringToBufIoReader(str string) *bufio.Reader {
	return bufio.NewReader(strings.NewReader(str))
}

func FormatNum(index int, format string) string {
	return fmt.Sprintf(format, index)
}
func StringConcat(first string, separator string, second string, withNewLine bool) string {
	if withNewLine {
		return fmt.Sprintf("%s%s%s%s", first, separator, second, "\n")
	} else {
		return fmt.Sprintf("%s%s%s", first, separator, second)
	}
}
