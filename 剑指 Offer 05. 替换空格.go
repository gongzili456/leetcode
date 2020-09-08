package leetcode

import "strings"

func replaceSpace(s string) string {
	ans := make([]string, len(s))

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			ans = append(ans, string(s[i]))
			continue
		}

		ans = append(ans, "%20")
	}

	return strings.Join(ans, "")
}
