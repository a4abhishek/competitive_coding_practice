#!/usr/bin/env python3

import os
import re
from collections import OrderedDict


def extract_problem_info(directory):
    """Extract the problem name and the Markdown link from the README.md file."""
    readme_path = f'{directory}/README.md'
    with open(readme_path, 'r') as file:
        for line in file:
            if line.startswith('# '):
                clean_line = line.strip('# \n').strip()
                link_match = re.search(r'\[([^\]]+)]\(([^\)]+)\)', line)
                if link_match:
                    # Extract both the link text and the link
                    return link_match.group(1), os.path.relpath(directory, '.')
                else:
                    # Extract the header text and use the directory name as link
                    return clean_line, os.path.relpath(directory, '.')
    return None, None


def update_problems_list(main_readme, directory_links):
    """Update the 'Problems List' section in the main README.md."""
    problems_info = OrderedDict()

    problems_section_pattern = re.compile(r"(## Problems List\n)(.*?)(\n## )", re.DOTALL)
    problems_list_match = problems_section_pattern.search(main_readme)

    if problems_list_match:
        problems_list = problems_list_match.group(2)
        problem_lines = problems_list.split('\n')

        for line in problem_lines:
            if line.strip() == '':
                continue
            link_match = re.search(r'\[([^\]]+)\]\(([^)]+)\)', line)
            if link_match:
                problem_name = link_match.group(1)

                if problem_name not in directory_links:
                    continue

                directory = directory_links[problem_name]
                if os.path.isfile(directory):
                    directory = os.path.dirname(directory)

                # Keep only those problems in the list which are present in the directory
                if os.path.isdir(directory):
                    problem_name, directory_link = extract_problem_info(directory)
                    problems_info[problem_name] = directory_link

    # Add new problem at the end
    for problem_name, directory_link in directory_links.items():
        if problem_name not in problems_info:
            problems_info[problem_name] = directory_link

    # Generate the updated problems list
    updated_list = '\n'.join([f"- [{problem_name}]({directory})" for problem_name, directory in problems_info.items()])

    # Replace the old problems list with the updated one
    return problems_section_pattern.sub(r"\1" + updated_list + r"\3", main_readme)


def main():
    """Main function to update the README.md with the problems list."""
    directory_links = {}

    # Gather information about each problem
    for directory in os.listdir('.'):
        if os.path.isdir(directory) and not directory.startswith('.'):
            problem_name, directory_link = extract_problem_info(directory)
            if problem_name:
                directory_links[problem_name] = directory_link

    # Read the main README.md
    with open('README.md', 'r') as file:
        main_readme_contents = file.read()

    # Update the main README.md
    updated_readme = update_problems_list(main_readme_contents, directory_links)

    # Write the updated content back to README.md
    with open('README.md', 'w') as file:
        file.write(updated_readme)


if __name__ == "__main__":
    main()
