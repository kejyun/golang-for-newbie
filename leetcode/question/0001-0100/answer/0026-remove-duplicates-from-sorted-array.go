package main

import (
    "fmt"
)

type QuestionList struct {
    Parameter
    Answer
}

// Parameter 是参数
// nums_list 代表第一个参数
type Parameter struct {
    nums_list []int
}

// Answer 是答案
// nums 代表第一个答案
type Answer struct {
    nums int
}


func main() {

    question_list := []QuestionList{
        {
            Parameter{[]int{1, 1, 2}},
            Answer{2},
        },

        {
            Parameter{[]int{0, 0, 1, 1, 1, 1, 2, 3, 4, 4}},
            Answer{5},
        },

        {
            Parameter{[]int{0, 0, 0, 0, 0}},
            Answer{1},
        },
        {
            Parameter{[]int{1}},
            Answer{1},
        },
    }

    fmt.Printf("------------------------Leetcode Problem 26------------------------\n")

    for _, question := range question_list {
        Ans, Param := question.Answer, question.Parameter
        fmt.Printf("【input】:%v  answer:%+v  【output】:%v\n", Param.nums_list, Ans.nums, removeDuplicates(Param.nums_list))
    }

}


// 解法一
func removeDuplicates(nums_list []int) int {
    if len(nums_list) == 0 {
        // 若沒有數字資料
        return 0
    }
    // 數字清單長度
    num_list_length := len(nums_list)
    // 最大數字所在索引
    max_num_index := num_list_length -1
    // 最後一個非重複數字索引
    last_none_duplicate_num_index := 0
    // 非重複數字搜尋索引
    non_duplicate_num_finder_index := 0


    for last_none_duplicate_num_index < max_num_index {
        // 最後一個非重複數字索引不是最後一個項目，繼續往後找
        for nums_list[non_duplicate_num_finder_index] == nums_list[last_none_duplicate_num_index] {
            // 若持續找到相同的數字，繼續往後找，直到數字不同在跳出
            non_duplicate_num_finder_index++

            if non_duplicate_num_finder_index == num_list_length {
                // 找到最後一個項目，非重複的長度為最後一個項目索引 +1
                return last_none_duplicate_num_index + 1
            }
        }

        // 將非重複元素索引往後移動
        last_none_duplicate_num_index++
        // 找到不同的元素了，將找到的不同元素複製到前方
        nums_list[last_none_duplicate_num_index] = nums_list[non_duplicate_num_finder_index]
    }

    return last_none_duplicate_num_index + 1
}
