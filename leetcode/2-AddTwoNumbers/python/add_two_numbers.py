# coding: utf-8
# athor: oliverch

# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        p3 = l3 = ListNode(0)
        carry = 0

        while l1 != None or l2 != None or carry > 0:
            sum = carry

            if l1:
                sum += l1.val
                l1 = l1.next
            if l2:
                sum += l2.val
                l2 = l2.next
            
            carry = sum // 10
            p3.next = ListNode(sum % 10)
            p3 = p3.next
        return l3.next

# unittest

import unittest


class TestSolution(unittest.TestCase):
    def test_case_1(self):
        """
        81 + 0 = 81
        """
        l1 = ListNode(1)
        l1.next = ListNode(8)

        l2 = ListNode(0)

        solution = Solution()
        result = solution.addTwoNumbers(l1, l2)
        self.assertEqual(result.val, 1)
        self.assertEqual(result.next.val, 8)
        self.assertEqual(result.next.next, None)
    
    def test_case_2(self):
        """
        342 + 465 = 807
        """
        l1 = ListNode(2)
        l1.next = ListNode(4)
        l1.next.next = ListNode(3)

        l2 = ListNode(5)
        l2.next = ListNode(6)
        l2.next.next = ListNode(4)

        solution = Solution()
        result = solution.addTwoNumbers(l1, l2)
        self.assertEqual(result.val, 7)
        self.assertEqual(result.next.val, 0)
        self.assertEqual(result.next.next.val, 8)
        self.assertEqual(result.next.next.next, None)
    

if __name__ == '__main__':
    unittest.main()