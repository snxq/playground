package longestsubstring

import "strings"

func lengthOfLongestSubstring(s string) int {
	var (
		temp string
		l    int
	)

	for _, c := range strings.Split(s, "") {
		if strings.Contains(temp, c) {
			temp = strings.SplitN(temp, c, 2)[1]
			temp = temp + c
		} else {
			temp = temp + c
			if len(temp) > l {
				l = len(temp)
			}
		}
	}

	return l
}
