# Competitive Coding Practices

This repository contains my solutions and practices for various competitive coding problems. Each problem is organized in its own directory, complete with a detailed explanation, source code, and unit tests.

## Initial Setup

After cloning the repository, it's important to perform an initial setup to configure the environment correctly. Currently, this setup includes installing the pre-commit hook. In the future, it may be expanded to include additional setup tasks as needed. To complete the initial setup, simply run the following command in the root of the repository:

```bash
make setup
````

## Repository Structure

- **Problem Directories**: Each directory is named after a specific problem and contains the following:
  - `README.md`: A detailed explanation of the problem, including its context and constraints.
  - `solution*.*`: The solution code file (the extension depends on the programming language used).
  - `test_*.*`: Unit test code for the solution (the extension depends on the testing framework).

## Problems List
- [Count BST nodes that lie in a given range](count_bst_nodes_that_lie_in_a_given_range)
- [Smallest subarray with sum greater than x](smallest_subarray_with_sum_greater_than_x)
- [Subarray with given sum](subarray_with_given_sum)
- [Symmetric Tree](symmetric_tree)
- [Isomorphic Strings](isomorphic_strings)
- [Best Time to Buy and Sell Stock](best_time_to_buy_and_sell_stock)
- [LeetCode - 108. Convert Sorted Array to Binary Search Tree](convert_sorted_array_to_binary_search_tree)
- [Problem 114: Flatten Binary Tree to Linked List](flatten_binary_tree_to_linked_list)
- [Maximum Number of Vowels in a Substring of Given Length](maximum_number_of_vowels_in_a_substring_of_given_length)
- [Leetcode 1220: Count Vowels Permutation](count_vowels_permutation)
- [LeetCode 334: Increasing Triplet Subsequence](increasing_triplet_subsequence)
- [LeetCode 5: Longest Palindromic Substring](longest-palindromic-substring)
## How to Use This Repository

- **Viewing a Problem**: Navigate to the respective problem directory and open the `README.md` file for a detailed problem description.
- **Viewing the Solution**: In the problem directory, the solution file (`solution*.*`) contains the code.
- **Running Unit Tests**: Unit tests can be found in the `test_*.*` file within each problem directory.

## Contributing

This repository primarily documents my personal journey through competitive coding. While I deeply appreciate interest and engagement, I humbly request that direct contributions to this repository be limited. However, you are more than welcome to:

- **Fork the Repository**: Feel free to fork the repository and make your own changes. You can use it as a template or as a starting point for your own coding journey.

- **Share Feedback**: If you have suggestions or feedback, please feel free to open an issue or contact me directly. I value your input in improving this resource.

- **Collaborate**: If you're interested in collaborative projects or discussing coding problems, I'm open to such opportunities. Please reach out via my contact details below.

### Updating the Problem List

The **Problems List** section of this README is automatically updated daily using a GitHub Action, ensuring it stays current with the latest problems added to the repository. 

However, it is recommended to update the problem list manually along with the addition of a new solution. This ensures immediate reflection of your latest contributions in the repository. To manually update the list, follow these steps:

1. **Run the Script**: Execute `update_readme.py` from the root of the repository. It will scan each problem's README.md and update the list accordingly.

    ```bash
    python update_readme.py
    ```

2. **Commit the Changes**: After running the script, if there are any changes, the main README.md file will be updated. You can then commit and push these changes to the repository.

    ```bash
    git add README.md
    git commit -m "Updated Problems List"
    git push
    ```

## Finalizing Changes

Before committing your changes, it is recommended to run the following command:

```bash
make all
```

This will perform a series of actions including normalizing end-of-line characters, running unit tests, and updating the problem list. It ensures that your changes are properly vetted and the repository remains consistent.

## Writing Unit Tests

When adding a new solution or modifying an existing one, it's important to include corresponding unit tests. A typical test file should look like this:

```python
from unittest import TestCase, main
from solution import Solution

class TestSolution(TestCase):
    def test_your_function(self):
        solution = Solution()

        # Replace with your test cases
        self.assertEqual(solution.your_function(arguments), expected_result)

# This section allows the test to be run by the Makefile
if __name__ == '__main__':
    main()
```

### Key Points for Writing Tests:

- Import `TestCase` and `main` from the `unittest` module.
- Import the class or functions you want to test (e.g., `Solution` and `buildTree`).
- Define a class that inherits from `TestCase`.
- Write test methods for each aspect of your solution you want to test. These methods should start with `test_`.
- Use assertions like `assertEqual`, `assertTrue`, etc., to validate the behavior of your solution.
- Include the `if __name__ == '__main__': main()` block to enable running tests via the Makefile.

Adhering to this structure ensures that your tests are consistent and can be automatically run using the Makefile.

## License

[MIT License](LICENSE.md)

## Contact

- **Abhishek Kashyap**
- Email: [avskksyp@gmail.com](mailto:avskksyp@gmail.com)
- LinkedIn: https://www.linkedin.com/in/abkashyap/
