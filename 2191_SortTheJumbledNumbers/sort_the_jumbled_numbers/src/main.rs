use std::collections::BinaryHeap;
use std::cmp::Reverse;

struct Solution {}


impl Solution {

    pub fn sort_jumbled(mapping: Vec<i32>, nums: Vec<i32>) -> Vec<i32> {
        let mut heap = BinaryHeap::new();

        for (idx, &num) in nums.iter().enumerate() {
            // Transfor number
            let mut mapped_num = mapping[(num % 10) as usize];
            let mut aux_num = num / 10;
            let mut i = 1;
            let base: i32 = 10;
            while aux_num > 0 {
                mapped_num += mapping[(aux_num % 10) as usize] * (base.pow(i)); 
                aux_num = aux_num / 10;
                i += 1;
            }
            //println!("Num: {}", mapped_num);


            heap.push(Reverse((mapped_num, idx, num)));
        }

        //return heap.into_sorted_vec();

        let mut sorted_vec = Vec::new();
        while let Some(Reverse((_, _, num))) = heap.pop() {
            sorted_vec.push(num);
        }

        sorted_vec


        //return heap.into_sorted_vec();
        // Create a map of digits out of mapping
        // for each number in nums
        //   transfor to correct number
        //   add to ordered vector
        //
        // return vector

                


        // Transform number to mapped number
        // Order vector
        //
        // Can I do it all at once??
        //
        // Easy solution:
        // Transform each number of nums by mapping
        // Order hole vector
        //
        // Ascending order, heap???

    }
}

/*
You are given a 0-indexed integer array mapping which represents the mapping rule of a shuffled decimal system. mapping[i] = j means digit i should be mapped to digit j in this system.

The mapped value of an integer is the new integer obtained by replacing each occurrence of digit i in the integer with mapping[i] for all 0 <= i <= 9.

You are also given another integer array nums. Return the array nums sorted in non-decreasing order based on the mapped values of its elements.

Notes:

Elements with the same mapped values should appear in the same relative order as in the input.
The elements of nums should only be sorted based on their mapped values and not be replaced by them.

Input: mapping = [8,9,4,0,2,1,3,5,7,6], nums = [991,338,38]
Output: [338,38,991]
Explanation: 
Map the number 991 as follows:
1. mapping[9] = 6, so all occurrences of the digit 9 will become 6.
2. mapping[1] = 9, so all occurrences of the digit 1 will become 9.
Therefore, the mapped value of 991 is 669.
338 maps to 007, or 7 after removing the leading zeros.
38 maps to 07, which is also 7 after removing leading zeros.
Since 338 and 38 share the same mapped value, they should remain in the same relative order, so 338 comes before 38.
Thus, the sorted array is [338,38,991].

Input: mapping = [0,1,2,3,4,5,6,7,8,9], nums = [789,456,123]
Output: [123,456,789]
Explanation: 789 maps to 789, 456 maps to 456, and 123 maps to 123. Thus, the sorted array is [123,456,789].
*/
fn main() {
    let mapping = vec![8,9,4,0,2,1,3,5,7,6];
    let nums = vec![991,338,38];

    let res = Solution::sort_jumbled(mapping, nums);
    println!("Vector: {:?}", res);
}
