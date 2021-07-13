package main

import (
	"fmt"
)

type QuestionList struct {
	Parameter
	Answer
}

// Parameter 是参数
// check_text 代表檢查的字串
type Parameter struct {
	check_text string
}

// Answer 是答案
// one 代表第一个答案
type Answer struct {
	one int
}

func main() {

	question_list := []QuestionList{

		{
			Parameter{"abcabcbb"},
			Answer{3},
		},

		{
			Parameter{"bbbbb"},
			Answer{1},
		},

		{
			Parameter{"pwwkew"},
			Answer{3},
		},

		{
			Parameter{""},
			Answer{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3------------------------\n")

	for _, question := range question_list {
		// 問題參數
		Param := question.Parameter
		fmt.Printf("solution 1【input】:%v       【output】:%v\n", Param, lengthOfLongestSubstring1(Param.check_text))
		fmt.Printf("solution 2【input】:%v       【output】:%v\n", Param, lengthOfLongestSubstring2(Param.check_text))
	}
	fmt.Printf("\n\n\n")
}

// 解法1 使用 bit
func lengthOfLongestSubstring1(check_text string) int {
	// 被檢查的文字長度
	check_text_length := len(check_text)

	if check_text_length == 0 {
		// 驗證字串長度為 0 不檢查
		return 0
	}

	// 文字 ASCII 檢查設定，預設都沒有找到此文字 false
	var text_ascii_bit_set [256]bool
	// 右方指標字元 ASCII Code
	var right_character_ascii_code uint8
	// 左方指標字元 ASCII Code
	var left_character_ascii_code uint8

	// 左方字元索引
	left_character_index := 0
	// 右方字元索引
	right_character_index := 0
	// 檢查文字最大不重複字串長度
	check_text_longest_string_length := 0
	// 目前字串長度
	current_text_string_length := 0

	for left_character_index < check_text_length {
		// 左側字元索引小於檢查字串長度，還沒有全部檢查完，繼續檢查

		// 右方指標字元 ASCII Code
		right_character_ascii_code = check_text[right_character_index]
		// 左方指標字元 ASCII Code
		left_character_ascii_code = check_text[left_character_index]

		if text_ascii_bit_set[right_character_ascii_code] {
			// 若右方文字已經有出現過，表示從左方索引到右方索引中間已經有出現過該文字
			// 左方文字設定為未出現過 false
			text_ascii_bit_set[left_character_ascii_code] = false
			// 左方往右推進 1 個字元，繼續檢查搜尋
			left_character_index++
		} else {
			// 右方文字沒有出現過，設定右方文字為已出現過 true
			text_ascii_bit_set[right_character_ascii_code] = true
			// 右方指標往前搜尋
			right_character_index++
		}

		// 設定目前字串長度
		current_text_string_length = right_character_index - left_character_index

		if check_text_longest_string_length < current_text_string_length {
			// 若右方指標在左方指標前面，且長度大於目前檢查文字的最大長度，將目前長度設定為最大長度
			check_text_longest_string_length = current_text_string_length
		}
		if left_character_index+check_text_longest_string_length >= check_text_length || right_character_index >= check_text_length {
			// 1. 左方字元索引 + 目前檢查文字的最大長度 > 文字最大長度：再往右找也找不到更長的文字了
			// 2. 右方字元索引 >= 被檢查的文字長度 : 已經檢查到最後一個字元了
			// 跳出檢查
			break
		}
	}
	return check_text_longest_string_length
}

// 解法2 Sliding Window
func lengthOfLongestSubstring2(check_text string) int {
	// 被檢查的文字長度
	check_text_length := len(check_text)

	if check_text_length == 0 {
		// 驗證字串長度為 0 不檢查
		return 0
	}
	// 建立長度 256 的整數陣列
	var text_ascii_int_flag [256]int

	// 左方字元索引
	left_character_index := 0
	// 右方字元索引
	right_character_index := 0
	// 檢查文字最大不重複字串長度
	check_text_longest_string_length := 0
	// 右方指標字元 ASCII Code
	var right_character_ascii_code uint8
	// 左方指標字元 ASCII Code
	var left_character_ascii_code uint8
	// 目前字串長度
	current_text_string_length := 0

	for left_character_index < check_text_length {
		// 若左側指標小於字串長度，繼續檢查
		if right_character_index >= check_text_length {
			// 若右側指標大於字串長度，表示已經檢查到最後的字串了，不需要再檢查
			break
		}

		// 右方指標字元 ASCII Code
		right_character_ascii_code = check_text[right_character_index]

		if text_ascii_int_flag[right_character_ascii_code] == 0 {
			// 「右側指標 +1 小於字串長度，表示字串還沒全部檢查完」且「此文字未出現過」
			// 標記右方文字出現過
			text_ascii_int_flag[right_character_ascii_code] = 1
			// 繼續往右檢查
			right_character_index++
		} else {
			// 左方指標字元 ASCII Code
			left_character_ascii_code = check_text[left_character_index]
			// 標記左方文字沒出現過
			text_ascii_int_flag[left_character_ascii_code] = 0
			// 將左方指標往前移動
			left_character_index++
		}

		// 設定目前字串長度
		current_text_string_length = right_character_index - left_character_index

		if check_text_longest_string_length < current_text_string_length {
			// 若右方指標在左方指標前面，且長度大於目前檢查文字的最大長度，將目前長度設定為最大長度
			check_text_longest_string_length = current_text_string_length
		}
	}
	return check_text_longest_string_length
}
