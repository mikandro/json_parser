# JSON Validator CLI

A simple command-line tool written in Go that checks if an input file contains valid JSON.

## Features

- Validates JSON files and reports errors.
- Outputs a success message for valid JSON.
- Easy to use with simple command-line arguments.

## Installation

1. Ensure you have Go installed on your machine. If not, you can download it from the [official Go website](https://golang.org/dl/).

2. Clone this repository:

   ```bash
   git clone https://github.com/yourusername/json-validator-cli.git
   cd json-validator-cli
   ```

3. Run the application:

   ```bash
   go run main.go <input_file.json>
   ```

## Usage

To check if a JSON file is valid, run the command:

```bash
go run main.go path/to/your/file.json
```

### Example

For a valid JSON file (`valid.json`):

```json
{
    "name": "John Doe",
    "age": 30,
    "email": "john.doe@example.com"
}
```

Run the command:

```bash
go run main.go valid.json
```

Output:

```
Valid JSON!
```

For an invalid JSON file (`invalid.json`):

```json
{
    "name": "John Doe",
    "age": 30,
    "email": "john.doe@example.com"
```

Run the command:

```bash
go run main.go invalid.json
```

Output:

```
invalid JSON: unexpected end of JSON input
```

## Running Tests

To ensure the tool is functioning correctly, you can run the tests using:

```bash
go test
```

## Contributing

Feel free to open issues and submit pull requests. Contributions are welcome!

## License

This project is licensed under the Personal License - see the [LICENSE](LICENSE) file for details.

