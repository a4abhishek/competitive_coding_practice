from unittest import TestCase, main
from solution import Solution


class TestSolution(TestCase):
    def test_smallest_sub_with_sum(self):
        solution = Solution()

        # Test case 1: Normal case
        self.assertEqual(solution.smallestSubWithSum([1, 4, 45, 6, 0, 19], 6, 51), 3,
                         "Case [1,4,45,6,0,19] with sum 51")


if __name__ == '__main__':
    main()
