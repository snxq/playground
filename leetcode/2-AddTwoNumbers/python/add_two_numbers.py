# coding: utf-8
# athor: oliverch

# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        # 进位标识
        flag = False
        # 链表指针
        p1, p2 = l1, l2
        p3 = l3 = ListNode(0)

        while True:
            s = sum([p.val for p in [p1, p2] if p])
            if flag:
                s += 1
            if s >= 10:
                flag = True
                s %= 10
            else:
                flag = False
            
            p3.val = s

            p1 = p1.next if p1 else None
            p2 = p2.next if p2 else None

            if any([p1, p2]):
                p3.next = ListNode(0)
                p3 = p3.next
            else:
                if flag:
                    p3.next = ListNode(1)
                else:
                    p3 = None
                break
        
        return l3

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