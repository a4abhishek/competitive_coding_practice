from unittest import TestCase, main
from solution import Solution


class TestSolution(TestCase):
    def test_areIsomorphic(self):
        solution = Solution()

        self.assertEqual(solution.areIsomorphic('aab', 'xxy'), True)
        self.assertEqual(solution.areIsomorphic('aba', 'xxy'), False)
        self.assertEqual(solution.areIsomorphic('aab', 'xyz'), False)
        self.assertEqual(solution.areIsomorphic('rfkqyuqf', 'jkxyqvnr'), False)


if __name__ == '__main__':
    main()
