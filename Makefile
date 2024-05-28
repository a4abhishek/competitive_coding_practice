.PHONY: all test vet update-readme setup

# Define the python command (Use python if python3 is not available)
PYTHON := $(shell command -v python3 || command -v python)

all: vet test update-readme
	@echo "Done."

# Target to run all tests
test:
	@echo "Running Go tests..."
	@go test ./... -v
	@echo "Running Python tests..."
	@find . -type f -name 'test_*.py' -exec $(PYTHON) {} \;
	@echo "All Tests Passed."

# Target to vet the code
vet: normalize-eol

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

# Target to perform initial setup
setup:
	@echo "Setting up the project..."
	@echo "Copying pre-commit hook..."
	@cp scripts/pre-commit .git/hooks/pre-commit
	@chmod +x .git/hooks/pre-commit
	@echo "Pre-commit hook installed."
	@echo "Setup complete."
