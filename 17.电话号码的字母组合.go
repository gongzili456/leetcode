package leetcode

/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
func letterCombinations(digits string) []string {
	m := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	ans := []string{}

	var backtrack func(comb, nextDigit string)
	backtrack = func(comb, nextDigit string) {
		if len(nextDigit) == 0 {
			ans = append(ans, comb)
			return
		}

		digit := nextDigit[0:1]
		letters := m[digit]

		for i := 0; i < len(letters); i++ {
			backtrack(comb+letters[i:i+1], nextDigit[1:])
		}
	}

	if len(digits) > 0 {
		backtrack("", digits)
	}

	return ans
}

// @lc code=end
