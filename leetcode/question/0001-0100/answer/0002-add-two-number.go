package main

import (
	"fmt"
)

// ListNode define
type ListNode struct {
	Val  int
	Next *ListNode
}

// 問題
type question struct {
	parameter
	answer
}

// para 是参数
// one 代表第一个参数
type parameter struct {
	one     []int
	another []int
}

// ans 是答案
// one 代表第一个答案
type answer struct {
	one []int
}

func main() {

	// 問題清單
	question_list := []question{

		{
			parameter{[]int{}, []int{}},
			answer{[]int{}},
		},

		{
			parameter{[]int{1}, []int{1}},
			answer{[]int{2}},
		},

		{
			parameter{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
			answer{[]int{2, 4, 6, 8}},
		},

		{
			parameter{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
			answer{[]int{2, 4, 6, 8, 0, 1}},
		},

		{
			parameter{[]int{1}, []int{9, 9, 9, 9, 9}},
			answer{[]int{0, 0, 0, 0, 0, 1}},
		},

		{
			parameter{[]int{9, 9, 9, 9, 9}, []int{1}},
			answer{[]int{0, 0, 0, 0, 0, 1}},
		},

		{
			parameter{[]int{2, 4, 3}, []int{5, 6, 4}},
			answer{[]int{7, 0, 8}},
		},

		{
			parameter{[]int{1, 8, 3}, []int{7, 1}},
			answer{[]int{8, 9, 3}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2------------------------\n")

	for _, question := range question_list {
		param := question.parameter
		fmt.Printf("【input】:%v       【output】:%v\n", param, List2Ints(addTwoNumbers(Ints2List(param.one), Ints2List(param.another))))
	}
	fmt.Printf("\n\n\n")
}

// 連結數字相加
func addTwoNumbers(FirstNumberListNode *ListNode, SecondNumberListNode *ListNode) *ListNode {
	// AnswerListNode := &ListNode{Val: 0} // 建立答案節點，預設個位數是 0
	AnswerListNode := new(ListNode) // 建立答案節點

	CurrentAnswerListNode := AnswerListNode // 目前答案節點
	number_1 := 0                           // 第 1 個加總數值
	number_2 := 0                           // 第 2 個加總數值
	carry_number := 0                       // 加總後的進位數值
	number_sum := 0                         // 加總數值

	// 若「第 1 數字節點」或「第 2 數字節點」有資料不為 nil 時，或沒有任何進位數字時
	for FirstNumberListNode != nil || SecondNumberListNode != nil || carry_number != 0 {
		// 判斷「第 1 數字節點」
		if FirstNumberListNode == nil {
			// 若「第 1 數字節點」為 nil，表示沒有數字可以做加總了，設定可加總數字為 0
			number_1 = 0
		} else {
			// 若「第 1 數字節點」有值，將節點數值設定為此次第 1 個加總數值
			number_1 = FirstNumberListNode.Val
			FirstNumberListNode = FirstNumberListNode.Next
		}

		// 判斷「第 2 數字節點」
		if SecondNumberListNode == nil {
			// 若「第 2 數字節點」為 nil，表示沒有數字可以做加總了，設定可加總數字為 0
			number_2 = 0
		} else {
			// 若「第 1 數字節點」有值，將節點數值設定為此次第 2 個加總數值
			number_2 = SecondNumberListNode.Val
			SecondNumberListNode = SecondNumberListNode.Next
		}

		// 加總數值
		number_sum = (number_1 + number_2 + carry_number)

		// 設定餘數為目前答案節點數值
		CurrentAnswerListNode.Next = &ListNode{Val: number_sum % 10}

		// 設定答案節點的下一節點為目前節點，繼續往後做加總
		CurrentAnswerListNode = CurrentAnswerListNode.Next

		// 取得加總後的進位數值，繼續往後做加總
		carry_number = number_sum / 10
	}
	return AnswerListNode.Next
}

// 鏈結清單轉換成整數陣列
func List2Ints(HeadListNode *ListNode) []int {
	// 鏈結深度限制，鏈結深度超過限制會出錯
	link_list_depth_limit := 100

	current_link_list_depth := 0

	res := []int{}
	for HeadListNode != nil {
		// 若有連結資料
		current_link_list_depth++
		if current_link_list_depth > link_list_depth_limit {
			msg := fmt.Sprintf("鏈結深度超過%d，可能出现環狀鏈結。檢查錯誤，或者調整 List2Ints 函式中 link_list_depth_limit 的限制。", link_list_depth_limit)
			panic(msg)
		}

		res = append(res, HeadListNode.Val)
		HeadListNode = HeadListNode.Next
	}

	return res
}

// Ints2List convert []int to List
// 整數陣列轉換成鏈結清單
func Ints2List(nums_list []int) *ListNode {
	if len(nums_list) == 0 {
		return nil
	}

	// 初始化鏈結節點
	HeadListNode := &ListNode{}
	// 設定目前鏈結節點
	CurrentListNode := HeadListNode
	for _, number := range nums_list {
		// 設定鏈結節點的下一節點
		CurrentListNode.Next = &ListNode{Val: number}
		// 將下一節點設為目前節點
		CurrentListNode = CurrentListNode.Next
	}
	return HeadListNode.Next
}
