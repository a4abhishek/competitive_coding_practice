class Solution:
    def smallestSubWithSum(self, a, n, x):
        start = 0
        end = 0
        min_len = n + 1
        s = 0

        while end < n:
            # Add elements to the s until it is greater than x
            while s <= x and end < n:
                s += a[end]
                end += 1

            # Move the start pointer to find the smallest subarray
            while s > x and start < end:
                min_len = min(min_len, end - start)
                s -= a[start]
                start += 1

        # If no such subarray is found, return 0
        return 0 if min_len == n + 1 else min_len


def main():
    T = int(input())

    while (T > 0):
        sz = [int(x) for x in input().strip().split()]
        n, m = sz[0], sz[1]
        a = [int(x) for x in input().strip().split()]
        print(Solution().smallestSubWithSum(a, n, m))

        T -= 1


if __name__ == "__main__":
    main()
