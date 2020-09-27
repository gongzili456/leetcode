package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	var result [][]int
	length := len(nums)
	sort.Ints(nums)

	for i := 0; i < length-3; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}

		if nums[i]+nums[length-1]+nums[length-2]+nums[length-3] < target {
			continue
		}

		for j := i + 1; j < length-2; j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}

			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}

			if nums[i]+nums[j]+nums[length-1]+nums[length-2] < target {
				continue
			}

			head, end := j+1, length-1

			for head < end {
				temp := nums[i] + nums[j] + nums[head] + nums[end]

				if temp == target {
					result = append(result, append([]int(nil), nums[i], nums[j], nums[head], nums[end]))

					head++
					for head < end && nums[head] == nums[head-1] {
						head++
					}
				} else if temp > target {
					end--
				} else {
					head++
				}
			}
		}
	}

	return result
}

// @lc code=end
