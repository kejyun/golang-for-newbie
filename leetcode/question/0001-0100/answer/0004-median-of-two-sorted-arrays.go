package main

import (
    "fmt"
)

type QuestionList struct {
    Parameter
    Answer
}

// para 是参数
// one 代表第一个参数
type Parameter struct {
    nums1 []int
    nums2 []int
}

// ans 是答案
// one 代表第一个答案
type Answer struct {
    one float64
}

func main() {

    question_list := []QuestionList{
        {
            Parameter{[]int{400}, []int{111, 222, 333, 444, 555, 666, 777}},
            Answer{422},
        },
        {
            Parameter{[]int{700}, []int{111, 222, 333, 444, 555, 666, 777}},
            Answer{499.5},
        },
        {
            Parameter{[]int{600}, []int{111, 222, 333, 444, 555, 666, 777}},
            Answer{499.5},
        },
        {
            Parameter{[]int{500}, []int{111, 222, 333, 444, 555, 666, 777}},
            Answer{472},
        },
        {
            Parameter{[]int{1, 3, 5, 7, 9}, []int{2, 3, 4, 7, 11, 13}},
            Answer{3.5},
        },
        {
            Parameter{[]int{1, 3}, []int{2}},
            Answer{2.0},
        },

        {
            Parameter{[]int{1, 2}, []int{3, 4}},
            Answer{2.5},
        },
    }

    fmt.Printf("------------------------Leetcode Problem 4------------------------\n")

    for _, question := range question_list {
        // 問題參數
        Ans, Param := question.Answer, question.Parameter
        fmt.Printf("【input】:%+v    answer:%+v    【output】:%+v\n", Param, Ans, findMedianSortedArrays(Param.nums1, Param.nums2))
    }
    fmt.Printf("\n\n\n")
}

func findMedianSortedArrays(nums_list_1 []int, nums_list_2 []int) float64 {
    // 假设 nums1 的长度小
    if len(nums_list_1) > len(nums_list_2) {
        // 若數字清單 1 長度大於數字清單 2，將清單較短的數字清單 2 優先傳入第一個參數
        return findMedianSortedArrays(nums_list_2, nums_list_1)
    }
    // 「數字清單 1」左方數字索引
    num1_left_num_index := 0
    // 「數字清單 1」右方數字索引
    num1_right_nums_index := len(nums_list_1)
    // 所有數字中位數可能位置，全部長度+1，bit 往右位移除 2
    all_nums_list_divide_position := (len(nums_list_1) + len(nums_list_2) + 1) >> 1

    // 「數字清單 1」中位數位置
    nums1_median_index := 0
    // 「數字清單 2」中位數位置
    nums2_median_index := 0

    // === 中位數區塊劃分 ===
    for num1_left_num_index <= num1_right_nums_index {
        // 「當左方數字索引」比「右方數字索引」還要小，表示搜尋的數字還沒交叉重複到

        // 數字清單 1 中位數位置 = 目前左方最小數字位置 + 右方剩餘數字取中位數 bit 往右位移(除 2)
        // 中位數分界點，左側是 median_index -1，右側是 median_index，
        // nums1:  ……………… nums1[nums1_median_index-1] | nums1[nums1_median_index] ……………………
        // nums2:  ……………… nums2[nums2_median_index-1] | nums2[nums2_median_index] ……………………
        nums1_median_index = num1_left_num_index + (num1_right_nums_index-num1_left_num_index)>>1
        // 數字清單 2 中位數位置 =  所有數字中位數可能位置 - 數字清單 1 中位數位置
        nums2_median_index = all_nums_list_divide_position - nums1_median_index

        if nums1_median_index > 0 && nums_list_1[nums1_median_index-1] > nums_list_2[nums2_median_index] { // nums1 中的分界线划多了，要向左边移动
            // 「數字清單 1」中位數位置不是第一個數字
            // 「數字清單 1」中位數左側數字比「數字清單 2」右側中位數還大（排序過的數字左側應比右側小）
            // 表示整個數字列表的中位數在「數字清單 1」目前中位數之前
            // 「數字清單 1」右方數字索引往左側移動，找「數字清單 1」中位數前面小一點的數字
            num1_right_nums_index = nums1_median_index - 1
        } else if nums1_median_index != len(nums_list_1) && nums_list_1[nums1_median_index] < nums_list_2[nums2_median_index-1] { // nums1 中的分界线划少了，要向右边移动
            // 「數字清單 1」中位數位置不是最後一個數字
            // 「數字清單 1」中位數右側數字比「數字清單 2」左側中位數還小（排序過的數字右側應比左側大）
            // 「數字清單 1」中位數位置數字 比 「數字清單 2」中位數位位置的前一個數字還小
            // 「數字清單 1」左方數字索引往右側移動，找「數字清單 1」中位數後面大一點的數字
            num1_left_num_index = nums1_median_index + 1
        } else {
            // 無法再劃分左右側中位數位置
            break
        }
    }

    // === 找出中位數 ===
    // 中位數
    var median_num float64 = 0.0
    // 左側中位數
    median_num_left := 0
    // 右側中位數
    median_num_right := 0

    if nums1_median_index == 0 {
        // 「數字清單 1」中位數位置為第一個數字
        // 中位數左側數字 = 「數字清單 2」 左側第一個數字」
        median_num_left = nums_list_2[nums2_median_index-1]
    } else if nums2_median_index == 0 {
        // 「數字清單 2」中位數位置為第一個數字
        // 中位數左側數字 = 「數字清單 1」 左側第一個數字」
        median_num_left = nums_list_1[nums1_median_index-1]
    } else {
        // 中位數左側數字 = 「數字清單 1 左側位置第一個數字」與「數字清單 2 左側位置第一個數字」取最大的數字
        // 左側的數字比右側小，所以要找比較大的數字才會接近右側數字
        median_num_left = max(nums_list_1[nums1_median_index-1], nums_list_2[nums2_median_index-1])
    }

    // 判斷數字數量是奇數還是偶數
    if (len(nums_list_1)+len(nums_list_2))&1 == 1 {
        // 若數字總數量為奇數，直接回傳中位數數字
        median_num = float64(median_num_left)
        return median_num
    }

    if nums1_median_index == len(nums_list_1) {
        // 「數字清單 1」中位數位置在不在清單中，索引超出清單長度（索引從 0 開始算）
        // 中位數右側數字 = 「數字清單 2」中位數右側第一個數字
        median_num_right = nums_list_2[nums2_median_index]
    } else if nums2_median_index == len(nums_list_2) {
        // 「數字清單 2」中位數位置在不在清單中，索引超出清單長度（索引從 0 開始算）
        // 中位數右側數字 = 「數字清單 1」中位數右側第一個數字
        median_num_right = nums_list_1[nums1_median_index]
        fmt.Println("<median_num_right (2)>: num1_medium =>  nums_list_1[nums1_median_index]")
    } else {
        // 中位數右側數字 = 「數字清單 1 右側第一個數字」與「數字清單 2 右側第一個數字」取最小的數字
        // 右側的數字比左側大，所以要找比較小的數字才會接近左側數字
        median_num_right = min(nums_list_1[nums1_median_index], nums_list_2[nums2_median_index])
    }


    median_num = float64(median_num_left+median_num_right) / 2

    return median_num
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a int, b int) int {
    if a > b {
        return b
    }
    return a
}
