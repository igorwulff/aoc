package internal

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Args struct {
	Day  string
	Part string
	Year string
	path string // plugin path
}

func ProcessArgs() Args {
	year := flag.String("year", fmt.Sprint(time.Now().Year()), "Year of the challenge")
	day := flag.String("day", "1", "Day of the challenge")
	part := flag.String("part", "1", "Part of the challenge")
	flag.Parse()

	args := Args{
		Day:  *day,
		Part: *part,
		Year: *year,
	}

	path, err := args.getPath()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	args.path = path

	return args
}

func (args Args) getPath() (string, error) {
	path := filepath.Join("./", args.Year, "day"+args.Day, "part"+args.Part)
	if _, err := os.Stat(filepath.Dir(path)); os.IsNotExist(err) {
		return "", fmt.Errorf("the directory %s does not exist", filepath.Dir(path))
	}

	return path, nil
}
