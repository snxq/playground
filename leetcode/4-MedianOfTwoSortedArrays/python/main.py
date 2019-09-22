

class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: float
        """
        pointer1 = 0
        pointer2 = 0
        length = len(nums1) + len(nums2)

        nums3 = []
        for index in range((length//2)+1):
            if pointer1 >= len(nums1):
                nums3.append(nums2[pointer2])
                pointer2 += 1
            elif pointer2 >= len(nums2):
                nums3.append(nums1[pointer1])
                pointer1 += 1
            elif nums1[pointer1] <= nums2[pointer2]:
                nums3.append(nums1[pointer1])
                pointer1 += 1
            elif nums1[pointer1] > nums2[pointer2]:
                nums3.append(nums2[pointer2])
                pointer2 += 1

        if length % 2:
            return nums3[-1]
        else:
            return (nums3[-2] + nums3[-1]) / 2


import unittest


class TestSolution(unittest.TestCase):
    def setUp(self):
        self.solution = Solution()

    def test_case_1(self):
        result = self.solution.findMedianSortedArrays([1, 2], [3, 4])
        self.assertEqual(result, 2.5)

    def test_case_2(self):
        result = self.solution.findMedianSortedArrays([3, 4], [1, 2])
        self.assertEqual(result, 2,5)
    
    def test_case_3(self):
        result = self.solution.findMedianSortedArrays([1, 3], [2])
        self.assertEqual(result, 2)


if __name__ == '__main__':
    unittest.main()