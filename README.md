# Statistical Analysis Program

This program is designed to perform basic statistical analysis on a dataset provided in a file named `data.txt`. It calculates the average, median, variance, and standard deviation of the numbers in the dataset.

## Requirements

- Go 1.16 or later installed on your system.

## Installation

1. Clone or download the repository to your local machine.
```bash
git clone https://learn.zone01kisumu.ke/git/nyunja/math-skills.git
```
2. Navigate to the directory containing the program files.
```bash
cd math-skills
```
## Usage

1. Ensure that the dataset file is named `data.txt` and located in the same directory as the program.
2. Open a terminal and navigate to the directory containing the program.
3. Run the program using the following command:

```bash
go run main.go
```

4. The program will display the calculated statistical values in the terminal.

## Notes

- Ensure that the dataset file (`data.txt`) contains numerical values separated by either carriage return and newline (`\r\n`) or just the newline (`\n`) characters.
- The program does not currently support datasets with different delimiters or formats.
- If there are any errors encountered during execution (e.g., incorrect file name, empty file, invalid numbers), appropriate error messages will be displayed in the terminal.

## Contributing

Contributions are welcome! If you encounter any issues or have suggestions for improvements, please open an issue or submit a pull request on GitHub.

## License

This program is licensed under the [MIT License](LICENSE).
