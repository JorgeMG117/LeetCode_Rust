use std::convert::TryInto;

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut idx_inf = 0;
        let mut idx_sup = nums.len()-1;
        let mut found = false;
        let mut result: Vec<i32> = Vec::new();

        while !found {
            let val_inf = nums[idx_inf];
            let val_sup = nums[idx_sup];

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
        return result;
    }
}
