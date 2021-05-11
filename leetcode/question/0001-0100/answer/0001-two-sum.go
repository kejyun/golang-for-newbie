package main

import (
	"fmt"
)

type question struct {
	// 參數
	parameter
	// 答案
	answer
}

// 參數
type parameter struct {
	nums_list       []int
	final_sum_value int
}

// 答案
type answer struct {
	one []int
}

func main() {
	question_list := []question{
		{
			parameter{[]int{3, 2, 4}, 6},
			answer{[]int{1, 2}},
		},

		{
			parameter{[]int{3, 2, 4}, 5},
			answer{[]int{0, 1}},
		},

		{
			parameter{[]int{0, 8, 7, 3, 3, 4, 2}, 11},
			answer{[]int{1, 3}},
		},

		{
			parameter{[]int{0, 1}, 1},
			answer{[]int{0, 1}},
		},

		{
			parameter{[]int{0, 3}, 5},
			answer{[]int{}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1------------------------\n")
	for _, question := range question_list {
		// _, p := question.answer, question.parameter
		param := question.parameter
		fmt.Printf("【input】:%v       【output】:%v\n", param, twoSum(param.nums_list, param.final_sum_value))
	}
}

func twoSum(nums_list []int, final_sum_value int) []int {
	// 數字反向對應表
	num_reverse_mapping := make(map[int]int)
	for key, value := range nums_list {
		// 判斷是否其他數值存在
		if other_value_key, is_other_value_exist := num_reverse_mapping[final_sum_value-value]; is_other_value_exist {
			// 回傳其他數值鍵值 + 目前數值鍵值
			return []int{other_value_key, key}
		}
		// 設定數值反向對應表：數值 => 鍵值
		num_reverse_mapping[value] = key
	}
	return nil
}
