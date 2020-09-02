package leetcode

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return nil
	}

	// 模拟双端队列，队头保持最大
	deque := []int{}

	// 结果
	res := []int{}

	// 滑动窗口两端边界
	j := 0
	i := j - k + 1

	for j < len(nums) {
		// 队首元素 = 窗口左侧外部元素，删除头部元素 ？
		if i > 0 && deque[0] == nums[i-1] {
			deque = deque[1:]
		}

		// 从队尾删除比当前数字小的元素，保证队列头部为最大值
		for len(deque) > 0 && deque[len(deque)-1] < nums[j] {
			deque = deque[:len(deque)-1]
		}

		// 从队列尾部添加元素，如果队列为空，则是最大值
		deque = append(deque, nums[j])

		// 当滑动窗口生效时，添加结果集
		if i >= 0 {
			res = append(res, deque[0])
		}

		j++
		i++
	}

	return res
}
