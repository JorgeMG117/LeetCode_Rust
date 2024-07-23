
struct Solution {}


use std::collections::BinaryHeap;
use std::cmp::Reverse;


impl Solution {
    pub fn k_smallest_pairs(nums1: Vec<i32>, nums2: Vec<i32>, k: i32) -> Vec<Vec<i32>> {
        let mut heap = BinaryHeap::new();
        let mut k = k as usize;
        let (m, n) = (nums1.len(), nums2.len());
        let mut res = vec![];

        for i in 0..m {
            heap.push(Reverse((nums1[i] + nums2[0], i, 0 as usize)));
        }

        while k > 0 {
            if let Some(Reverse((_sum, i, j))) = heap.pop() {
                res.push(vec![nums1[i], nums2[j]]);

                if j + 1 < n {
                    heap.push(Reverse((nums1[i] + nums2[j + 1], i, j + 1)));
                }
            }

            k -= 1;
        }

        res    
    }
}

/*
You are given two integer arrays nums1 and nums2 sorted in non-decreasing order and an integer k.

Define a pair (u, v) which consists of one element from the first array and one element from the second array.

Return the k pairs (u1, v1), (u2, v2), ..., (uk, vk) with the smallest sums.
*/
fn main() {
    let vec1 = vec![1,2,4,5,6];
    let vec2 = vec![3,5,7,9];
    //[[1,3],[1,5],[1,7]]
    let sol = Solution::k_smallest_pairs(vec1, vec2, 3);

    for inner_vec in &sol {
            if inner_vec.len() == 2 {
                println!("({}, {})", inner_vec[0], inner_vec[1]);
            } else {
                println!("Invalid inner vector length.");
            }
        }
}
