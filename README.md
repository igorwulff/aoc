# Advent of Code Helper Tool

This repository provides a framework and helper tools to streamline the process of solving [Advent of Code](https://adventofcode.com/) challenges. It includes utilities for building and running Go plugins, managing input files, and running tests.

## Structure

- `internal/processor.go`: Contains the `PluginProcessor` struct and methods for building plugins, calling functions within those plugins, running tests, and getting input data.
- `internal/input.go`: Contains the `Args` struct and functions for processing command-line arguments and determining file paths.
- `runner.go`: The main entry point that ties everything together by processing arguments, building the plugin, getting input, running tests, and calling the solution function.
- `benchmark.go`: Alternative entry point that benchmarks every part of every day, and displays the result as a table afterwards.
- It is expected to set up a folder structure following the pattern: `/year/day/partnumber.go`. For example, for the first part of the first day's challenge in 2024, the file should be located at `/2024/day1/part1.go`.
- It is expected to add test files for each day. As a minimal it is adviced to add a test case for each part of a day based on the sample input and output provided by `Advent of Code`.

## Usage

1. **Process Arguments**: The tool processes command-line arguments to determine the year, day, part, and input type (sample or input) of the challenge.
2. **Build Plugin**: It builds the Go plugin for the specified challenge part.
3. **Get Input**: Reads the input file for the challenge.
4. **Run Tests**: Executes tests for the specified challenge.
5. **Call Function**: Calls the solution function from the built plugin and prints the output.

## Example

To run the tool, use the following command:

```sh
go run runner.go -year=2024 -day=1 -part=1
```

This command will:
- Process the arguments for year 2024, day 1, part 1, and use the sample input.
- Build the plugin for the specified challenge.
- Read the sample input file.
- Run tests for the specified challenge.
- Call the solution function and print the output.

## Contributing

Feel free to open issues or submit pull requests if you have any improvements or bug fixes.

## License

This project is licensed under the MIT License.