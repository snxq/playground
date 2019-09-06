use std::collections::HashMap;

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut dict = HashMap::with_capacity(nums.len());

        for (index, num) in nums.iter().enumerate() {
            match dict.get(&(target-num)) {
                Some(result) => { return vec![index as i32, *result as i32] },
                None => { dict.insert(num, index); }
            }
        }
        vec![]
    }
}