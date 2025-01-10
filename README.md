# Expense Tracker Go-Implementation

A simple command-line expense tracking application built in Go. This project is designed for beginners and is part of the backend projects on [roadmap.sh](https://roadmap.sh/projects/expense-tracker).

## Features

- Add expenses with a description and amount
- Delete expenses by ID
- List all expenses
- Show a summary of total expenses
- Data is stored in a JSON file

## Project Structure
```bash
Expense-Tracker-Go-Implementation/
├── cmd/
│ └── main.go
├── expense/
│ └── expense.go
├── storage/
│ └── storage.go
└── data/
└── expenses.json (automatically created)
```


## Getting Started

### Prerequisites

- Go (version 1.15 or later)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/alielmi98/Expense-Tracker-Go-Implementation.git

2. Run the application:
    ```bash
    go run cmd/main.go

### Usage
You can use the following commands to manage your expenses:

* Add an expense:
```bash
go run cmd/main.go add --description "Lunch" --amount 20
```
* List all expenses:
```bash
go run cmd/main.go list
```
* Show total expenses:
```bash
go run cmd/main.go summary
```
* Delete an expense by ID:
```bash
go run cmd/main.go delete --id 1
```

## Contributing
Contributions are welcome! Please feel free to open issues or submit pull requests.

## License
This project is licensed under the MIT License - see the LICENSE file for details.