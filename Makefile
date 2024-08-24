# Variables
BINARY_NAME = expense-tracker
BUILD_DIR = bin
SOURCE_DIR = ./cmd/expense-tracker

# Build the binary
build:
	@echo "Building the application..."
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(SOURCE_DIR)

# Add an expense
add:
	@echo "Adding an expense..."
	@$(BUILD_DIR)/$(BINARY_NAME) add --description "$(description)" --amount $(amount)

# List all expenses
list:
	@echo "Listing all expenses..."
	@$(BUILD_DIR)/$(BINARY_NAME) list

# Update an expense
update:
	@$(BINARY) update --id $(id) --description "$(description)" --amount $(amount)

# Delete an expense by ID
delete:
	@echo "Deleting an expense..."
	@$(BUILD_DIR)/$(BINARY_NAME) delete --id $(id)

# Show a summary of all expenses
summary:
	@echo "Showing the summary of all expenses..."
	@$(BUILD_DIR)/$(BINARY_NAME) summary

# Show a summary of expenses for a specific month
summary-month:
	@echo "Showing the summary of expenses for month $(month)..."
	@$(BUILD_DIR)/$(BINARY_NAME) summary --month $(month)

# Clean the build artifacts
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)