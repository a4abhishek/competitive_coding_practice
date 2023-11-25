from collections import deque


def isMirror(left, right):
    """Helper function to check if two trees are mirror images"""
    # If both are None, they are mirrors
    if left is None and right is None:
        return True
    # If one of them is None, they are not mirrors
    if left is None or right is None:
        return False
    # The two nodes must have the same value and their respective children must be mirrors
    return (left.data == right.data) and isMirror(left.left, right.right) and isMirror(left.right, right.left)


class Solution:
    def isSymmetric(self, root):
        # If tree is empty, it is symmetric
        if root is None:
            return True
        return isMirror(root.left, root.right)


# Tree Node
class Node:
    def __init__(self, val):
        self.right = None
        self.data = val
        self.left = None


# Function to Build Tree
def buildTree(s):
    # Corner Case
    if (len(s) == 0 or s[0] == "N"):
        return None

    # Creating list of strings from input
    # string after spliting by space
    ip = list(map(str, s.split()))

    # Create the root of the tree
    root = Node(int(ip[0]))
    size = 0
    q = deque()

    # Push the root to the queue
    q.append(root)
    size = size + 1

    # Starting from the second element
    i = 1
    while (size > 0 and i < len(ip)):
        # Get and remove the front of the queue
        currNode = q[0]
        q.popleft()
        size = size - 1

        # Get the current node's value from the string
        currVal = ip[i]

        # If the left child is not null
        if (currVal != "N"):
            # Create the left child for the current node
            currNode.left = Node(int(currVal))

            # Push it to the queue
            q.append(currNode.left)
            size = size + 1
        # For the right child
        i = i + 1
        if (i >= len(ip)):
            break
        currVal = ip[i]

        # If the right child is not null
        if (currVal != "N"):
            # Create the right child for the current node
            currNode.right = Node(int(currVal))

            # Push it to the queue
            q.append(currNode.right)
            size = size + 1
        i = i + 1
    return root


if __name__ == "__main__":
    t = int(input())
    for _ in range(0, t):
        s = input()
        root = buildTree(s)
        ob = Solution()
        if ob.isSymmetric(root):
            print("True")
        else:
            print("False")
