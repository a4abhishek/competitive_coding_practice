from unittest import TestCase
from solution import Solution, buildTree


class TestSolution(TestCase):
    def test_getCount(self):
        solution = Solution()

        # Test case: Normal case 1
        # Tree structure: 5 4 6 3 N N 7 1
        #     5
        #    / \
        #   4   6
        #  /     \
        # 3       7
        #  \
        #   1
        root = buildTree("5 4 6 3 N N 7 1")
        self.assertEqual(solution.getCount(root, 1, 5), 4)
