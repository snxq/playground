/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
    let dict = {}
    for (let i=0; i<nums.length; i++) {
        if(dict[target-nums[i]] !== undefined){
            return [i, dict[target-nums[i]]]
        }
        dict[nums[i]] = i
    }
    return []
};