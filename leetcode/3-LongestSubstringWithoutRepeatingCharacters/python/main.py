class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        temp = ''
        l = 0
        for c in s:
            if c not in temp:
                temp += c
                l = len(temp) if len(temp) > length else length
            else:
                temp = temp.split(c)[-1]
                temp += c
        return l


import unittest


class TestSolution(unittest.TestCase):
    def setUp(self):
        self.solution = Solution()

    def test_case_1(self):
        result = self.solution.lengthOfLongestSubstring('abcabcbb')
        self.assertEqual(result, 3)

    def test_case_2(self):
        result = self.solution.lengthOfLongestSubstring('bbbbb')
        self.assertEqual(result, 1)

    def test_case_3(self):
        result = self.solution.lengthOfLongestSubstring('pwwkew')
        self.assertEqual(result, 3)


if __name__ == '__main__':
    unittest.main()
