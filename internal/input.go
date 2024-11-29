package internal

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Args struct {
	day   string
	part  string
	input string // sample or input
	year  string
	path  string // plugin path
}

func ProcessArgs() Args {
	year := flag.String("year", fmt.Sprint(time.Now().Year()), "Year of the challenge")
	day := flag.String("day", "1", "Day of the challenge")
	part := flag.String("part", "1", "Part of the challenge")
	input := flag.String("input", "sample", "Puzzle input type (input or sample)")
	flag.Parse()

	args := Args{
		day:   *day,
		part:  *part,
		input: *input,
		year:  *year,
	}

	path, err := args.getPath()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	args.path = path

	fmt.Println("Using input type:", args.input)

	return args
}

func (args Args) getPath() (string, error) {
	path := filepath.Join("./", args.year, "day"+args.day, "part"+args.part)
	if _, err := os.Stat(filepath.Dir(path)); os.IsNotExist(err) {
		return "", fmt.Errorf("the directory %s does not exist", filepath.Dir(path))
	}

	return path, nil
}
