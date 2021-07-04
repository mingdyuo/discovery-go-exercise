package chapter3

import (
	"fmt"
	"sort"
)

func findStr(strs []string, target string) bool {
	sort.Slice(strs, func(i, j int) bool {
		return strs[i] < strs[j]
	}) // 문제에서는 정리되어 있다고 가정
	start, end := 0, len(strs)-1
	for start <= end {
		mid := (start + end) / 2
		if target == strs[mid] {
			return true
		} else if target < strs[mid] {
			end = mid - 1
		} else if strs[mid] < target {
			start = mid + 1
		}
	}
	if strs[end] == target {
		return true
	}
	return false
}

func FindInit() {
	strs := []string{
		"solution", "always", "build", "perfect",
		"details", "problem", "finding", "step", "ahead", "rule"}
	fmt.Println(findStr(strs, "step"))
	fmt.Println(findStr(strs, "perfect"))
	fmt.Println(findStr(strs, "wow"))
}
