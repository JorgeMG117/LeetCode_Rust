//use std::convert::TryInto;

struct Solution {}

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut idx_inf = 0;
        let mut idx_sup = nums.len()-1;
        let mut found = false;
        let mut result: Vec<i32> = Vec::new();
        
        let nums_copy = nums;
        // saving original indices and sorting
        let mut nums_ord: Vec<(usize, i32)> = 
            nums.into_iter()
                .enumerate()
                .collect::<Vec<(usize, i32)>>();

        nums_ord.sort_unstable_by_key(|&(_, n)| n);

        println!("{:?}",nums_copy);

        while !found && idx_inf != idx_sup {
            let val_inf = nums_ord[idx_inf].1;
            let val_sup = nums_ord[idx_sup].1;

            if val_inf + val_sup == target {
                found = true;
                result.push(idx_inf.try_into().unwrap());
                result.push(idx_sup.try_into().unwrap());
            } else if val_inf + val_sup > target {
                idx_sup = idx_sup - 1;
            } else {
                idx_inf = idx_inf + 1;
            }
        }

        //return result.into_iter().map(|a|nums[a]).collect::<Vec<i32>>();
        return result;
    }
}

fn main() {
    let sol = Solution::two_sum(vec![-1, -2, -3, -4, -5], -8);

    for i in &sol {
        println!("{i}");
    }
}
