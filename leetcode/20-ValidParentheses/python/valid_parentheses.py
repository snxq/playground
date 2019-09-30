class Solution:
    def isValid(self, s: str) -> bool:
        stack = [s[0]]
        couples = {
            "(": ")",
            "[": "]",
            "{": "}"
        }

        for char in s[1:]:
            if not stack:
                stack.append(char)
            elif stack[0] not in couples:
                return False
            elif stack[-1] in couples and char == couples[stack[-1]]:
                stack.pop()
            else:
                stack.append(char)
        
        return not bool(stack)


import logging
import unittest
logger = logging.getLogger()


class TestSolution(unittest.TestCase):
    def setUp(self):
        self.solution = Solution()
    
    def __test(self, string, result):
        test_result = self.solution.isValid(string)
        self.assertEqual(test_result, result)
    
    def test(self):
        cases = {
            r'()': True,
            r'()[]{}': True,
            r'(]': False, r'([)]': False,
            r'{[]}': True
        }
        for key, value in cases.items():
            try:
                self.__test(key, value)
                logger.info(f'{key}\t Test Pass')
            except AssertionError as exc:
                logger.error(f'{key}\t Test Failed')


if __name__ == '__main__':
    unittest.main()