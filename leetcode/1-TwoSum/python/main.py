# coding: utf-8
# athor: oliverch

from typing import List

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        adict = {}
        for index, num in enumerate(nums):
            if (target - num) in adict:
                return [adict[target-num], index]
            else:
                adict[num] = index


if __name__ == "__main__":
    solution = Solution()
    print(solution.twoSum([2, 7, 11, 15], 9))
