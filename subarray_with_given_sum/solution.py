import math


class Solution:
    @staticmethod
    def subArraySum(arr, n, s):
        start = 0
        end = 0
        current_sum = arr[0]
        while True:
            if current_sum == s:
                if start > end:
                    break
                return [start + 1, end + 1]
            elif start > end or current_sum < s:
                end += 1
                if end >= n:
                    break
                current_sum += arr[end]
            else:
                current_sum -= arr[start]
                start += 1
        return [-1]

    @staticmethod
    def subArraySum2(arr, n, s):
        start = 0
        end = 0
        current_sum = 0

        while end < n:
            current_sum += arr[end]

            while current_sum > s and start < end:
                current_sum -= arr[start]
                start += 1

            if current_sum == s:
                return [start + 1, end + 1]

            end += 1

        return [-1]


def main():
    T = int(input())
    while (T > 0):
        NS = input().strip().split()
        N = int(NS[0])
        S = int(NS[1])

        A = list(map(int, input().split()))
        ob = Solution()
        ans = ob.subArraySum(A, N, S)

        for i in ans:
            print(i, end=" ")

        print()

        T -= 1


if __name__ == '__main__':
    main()
