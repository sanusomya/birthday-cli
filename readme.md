# birthday-cli

A command-line tool for managing birthday entries, built in Go. Easily add, edit, delete, and query birthdays from your terminal, with data stored in MongoDB.

## Features

- Add new birthday entries (name, date, month, mobile)
- Edit existing entries (name, mobile, date)
- Delete entries
- List all birthdays, birthdays today, or birthdays this month
- Input validation and helpful error messages

## Prerequisites

- Go 1.21+
- [birhday-backend](https://github.com/sanusomya/birthday-backend)
- `.env` file

## Setup

1. **Clone the repository:**
   ```sh
   git clone https://github.com/sanusomya/birthday-cli.git
   cd birthday-cli
   ```

2. **Configure environment variables:**
   Create a `.env` file in the root directory with:
   ```
   backend_url=<your-backend-url-including-ports>
   ```

3. **Install dependencies:**
   ```sh
   go mod tidy
   ```

4. **Build the CLI:**
   ```sh
   go build -o birthday-cli main.go
   ```

## Usage

Run the CLI and see available commands:
```sh
./birthday-cli --help
```

### Example Commands

- **Add a birthday:**
  ```sh
  ./birthday-cli add --name "Alice" --date 15 --month 9 --mobile 1234567890
  ```

- **Edit a birthday's mobile:**
  ```sh
  ./birthday-cli edit mobile --name "Alice" --mobile 0987654321
  ```

- **Delete a birthday:**
  ```sh
  ./birthday-cli delete --name "Alice"
  ```

- **List all birthdays:**
  ```sh
  ./birthday-cli get all
  ```

- **List birthdays today:**
  ```sh
  ./birthday-cli get today
  ```

- **List birthdays this month:**
  ```sh
  ./birthday-cli get month
  ```

## Project Structure

- `birthday/` – Birthday model and logic
- `cmd/` – CLI command definitions (add, edit, delete, get)
- `config/` – Configuration and environment loading
- `utils/` – Utility functions and error handling

## Testing

Run unit tests:
```sh
go test ./...
```