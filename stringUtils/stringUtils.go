package stringUtils

func IndexOfString(arr []string, search string) int {
	for i,v := range arr {
		if v==search {
			return i
		}
	}
	return -1;
}
