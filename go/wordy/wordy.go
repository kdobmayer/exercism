package wordy

import (
	"regexp"
	"strconv"
)

type Op func(int, int) int

func Answer(s string) (int, bool) {
	nums, ops := parseNums(s), parseOps(s)
	if len(nums) == 0 || len(ops) == 0 {
		return 0, false
	}

	answer := nums[0]
	for i, op := range ops {
		answer = op(answer, nums[i+1])
	}
	return answer, true
}

func parseNums(s string) []int {
	var nums []int
	matches := regexp.MustCompile("-?[0-9]+").FindAllStringSubmatch(s, -1)
	for _, m := range matches {
		n, _ := strconv.Atoi(m[0])
		nums = append(nums, n)
	}
	return nums
}

func parseOps(s string) []Op {
	var ops []Op
	matches := regexp.MustCompile("plus|minus|multiplied by|divided by").FindAllStringSubmatch(s, -1)
	for _, m := range matches {
		switch m[0] {
		case "plus":
			ops = append(ops, func(a, b int) int { return a + b })
		case "minus":
			ops = append(ops, func(a, b int) int { return a - b })
		case "multiplied by":
			ops = append(ops, func(a, b int) int { return a * b })
		case "divided by":
			ops = append(ops, func(a, b int) int { return a / b })
		}
	}
	return ops
}