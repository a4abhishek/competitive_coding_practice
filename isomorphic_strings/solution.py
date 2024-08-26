# User function Template for python3

class Bidict:
    def __init__(self):
        self.forward = {}
        self.backward = {}

    def insert(self, key, value):
        if key in self.forward:
            old_value = self.forward[key]
            del self.backward[old_value]
        if value in self.backward:
            old_key = self.backward[value]
            del self.forward[old_key]

        self.forward[key] = value
        self.backward[value] = key

    def get_key(self, value):
        return self.backward.get(value, None)

    def get_value(self, key):
        return self.forward.get(key, None)


class Solution:

    # Function to check if two strings are isomorphic.
    def areIsomorphic(self, str1, str2):
        if len(str1) != len(str2):
            return False

        bd = Bidict()
        for i in range(len(str1)):
            v = bd.get_value(str1[i])
            k = bd.get_key(str2[i])
            if v is None and k is None:
                bd.insert(str1[i], str2[i])
            elif k != str1[i] or v != str2[i]:
                return False

        return True


# {
# Driver Code Starts
# Initial Template for Python 3

# import atexit
# import io
# import sys
# from collections import defaultdict

# _INPUT_LINES = sys.stdin.read().splitlines()
# input = iter(_INPUT_LINES).__next__
# _OUTPUT_BUFFER = io.StringIO()
# sys.stdout = _OUTPUT_BUFFER


# @atexit.register
# def write():
#     sys.__stdout__.write(_OUTPUT_BUFFER.getvalue())


# if __name__ == '__main__':
#     t = int(input())
#     for i in range(t):
#         s = str(input())
#         p = str(input())
#         ob = Solution()
#         if (ob.areIsomorphic(s, p)):
#             print(1)
#         else:
#             print(0)
# } Driver Code Ends
