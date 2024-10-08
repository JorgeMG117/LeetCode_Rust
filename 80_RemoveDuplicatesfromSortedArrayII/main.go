package main

import (
    "fmt"
)

/*
Given an integer array nums sorted in non-decreasing order, remove some duplicates in-place such that each unique element appears at most twice. The relative order of the elements should be kept the same.

Since it is impossible to change the length of the array in some languages, you must instead have the result be placed in the first part of the array nums. More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.

Return k after placing the final result in the first k slots of nums.

Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.

Custom Judge:

The judge will test your solution with the following code:

int[] nums = [...]; // Input array
int[] expectedNums = [...]; // The expected answer with correct length

int k = removeDuplicates(nums); // Calls your implementation

assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}
If all assertions pass, then your solution will be accepted.

 

Example 1:

Input: nums = [1,1,1,2,2,3]
Output: 5, nums = [1,1,2,2,3,_]
Explanation: Your function should return k = 5, with the first five elements of nums being 1, 1, 2, 2 and 3 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).
Example 2:

Input: nums = [0,0,1,1,1,1,2,3,3]
Output: 7, nums = [0,0,1,1,2,3,3,_,_]
Explanation: Your function should return k = 7, with the first seven elements of nums being 0, 0, 1, 1, 2, 3 and 3 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).
 

Constraints:

1 <= nums.length <= 3 * 104
-104 <= nums[i] <= 104
nums is sorted in non-decreasing order.
*/
func removeDuplicates(nums []int) int {
    delay := 0
    val := 105
    times := 0

    for i := 0; i < len(nums); i++ {
        //fmt.Println("Iter: ", i)
        //fmt.Println("Val: ", nums[i])

        v := nums[i]
        if val != v {
            val = v
            times = 1
        } else if times < 2 {
            times++
        } else {
            // Compute delay
            for val == nums[i] {
                i++
                delay++
                if i >= len(nums) {
                    return len(nums) - delay
                }
            }

            times = 1
            val = nums[i]
        }
        nums[i-delay] = nums[i]

    }

    return len(nums) - delay
}



func main() {
    //nums := []int{0,0,1,1,1,1,2,3,3}
    //nums := []int{1,1,1,2,2,2,3,3}
    nums := []int{1,1,1}

    //nums := []int{1, 1, 1, 2, 2, 3}
    res := removeDuplicates(nums)
    fmt.Println(res)
    fmt.Println(nums)
}
