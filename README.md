# Expense-Tracker

## Overview

- The Expense Tracker CLI is a command-line tool for managing personal expenses. It allows users to add, update, delete, and view expenses, as well as providing summaries of expenses over time. This tool is built in Go and stores data locally in a JSON file.


## Features

* Add Expense: Add a new expense with a description and amount.
* Update Expense: Update an existing expense by its ID.
* Delete Expense: Delete an expense by its ID.
* List Expenses: View all recorded expenses.
* Expense Summary: View a summary of all expenses or expenses for a specific month.
 * Optional Features:
 * Categorize expenses.
 * Set monthly budgets and get warnings when exceeded.
 * Export expenses to a CSV file.

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/Carl0sL0pez03/Expense-Tracker
   cd task-tracker
   ```

2. **Compile proyect:**

```bash
   make build
```

# Usage

The CLI supports the following commands:

a. **Add an Expense**

```bash
make add description="Lunch" amount=20
```

b. **List All Expenses**

```bash
make list
```

c. **Update an Expense**

```bash
make update --id 1 --description "Lunch at cafe" --amount 25
```

d. **Delete an Expense**

```bash
make delete id=1
```

e. **View Expense Summary**

```bash
make summary
make summary-month month=8
```

# Challange roadmap

- https://roadmap.sh/projects/expense-tracker