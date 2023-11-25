.PHONY: all test normalize-eol update-readme

# Define the python command (Use python if python3 is not available)
PYTHON := $(shell command -v python3 || command -v python)

all: normalize-eol test update-readme
	@echo "Done."

# Target to run all tests
test:
	@echo "Running tests..."
	@find . -type f -name 'test_*.py' -exec $(PYTHON) {} \;
	@echo "All Tests Passed."

# Target to normalize end-of-line characters
normalize-eol:
	@echo "Normalizing end-of-line characters..."
	@find . -type f -exec sed -i 's/\r$$//' {} \;
	@echo "Normalization complete for end-of-line."

# Target to update the README
update-readme:
	@echo "Updating README..."
	@$(PYTHON) update_readme.py
	@echo "README updated."
