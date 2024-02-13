# DateTime Validation

This Go application reads a list of date-time strings from an input file, validates the format, and outputs unique normalized date-time values to an output file.

## Usage

1. Prepare your input file (`input.txt`) with a list of date-time strings, each on a new line.
2. Run the application:

    ```bash
    go run main.go helper.go
    ```

3. The unique and normalized date-time values will be written to the output file (`output.txt`).

## Input Format

The application expects date-time strings in the format `YYYY-MM-DDThh:mm:ssZ` or `YYYY-MM-DDThh:mm:ss±hh:mm`.

## Validation Process
The validation process involves validating the length of string and proper connectors between values.

## Normalization Process

The normalization function parses and formats the date-time strings to the desired output format (`YYYY-MM-DDThh:mm:ssZ`).

## Error Handling

The application provides error handling for invalid date-time strings and conversion errors during normalization.

Feel free to customize and extend the application based on your specific requirements.