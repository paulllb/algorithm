package main

// 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
//不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并原地修改输入数组。
//元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
//示例 1: 给定 nums = [3,2,2,3], val = 3, 函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。 你不需要考虑数组中超出新长度后面的元素。
//示例 2: 给定 nums = [0,1,2,2,3,0,4,2], val = 2, 函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。
//你不需要考虑数组中超出新长度后面的元素。

// 思路是双指针，slowIdx 是慢指针，只有遇到了!=val的数据才++
func removeElement(num []int, val int) int {
    slowIdx := 0
    for i := 0; i < len(num); i++ {
        if num[i] != val {
            num[slowIdx] = num[i]
            slowIdx++
        }
    }
    num = num[:slowIdx]
    return slowIdx
}
