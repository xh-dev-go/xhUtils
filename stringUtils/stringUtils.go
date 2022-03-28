package stringUtils

import "strings"

func IndexOfString(arr []string, search string) int {
	for i, v := range arr {
		if v == search {
			return i
		}
	}
	return -1
}

func SplitCpExcel(msg string, trimSpace bool) [][]string {
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
	return lines
}
