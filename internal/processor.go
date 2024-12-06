package internal

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"plugin"
	"strings"
)

type PluginProcessor struct {
	Args      Args
	Benchmark *BenchmarkResult
}

func (p PluginProcessor) Build() error {
	cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", p.Args.path+".so", p.Args.path+".go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("\nBuilding: %s", cmd.String())
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error running command: %v", err)
	}

	return nil
}

func (p PluginProcessor) CallFunc(input *[]byte) (string, error) {
	plug, err := plugin.Open(p.Args.path + ".so")
	if err != nil {
		return "", fmt.Errorf("error loading plugin: %v", err)
	}

	// Look up the Process function
	symProcess, err := plug.Lookup("Part" + p.Args.Part)
	if err != nil {
		return "", fmt.Errorf("error looking up process: %v", err)
	}

	// Assert that the symbol is a function with the expected signature
	processFunc, ok := symProcess.(func(string) string)
	if !ok {
		return "", fmt.Errorf("expected func(string) string, got %T", symProcess)
	}

	// Call the function with an input variable
	return processFunc(string(*input)), nil
}

func (p PluginProcessor) RunTests() error {
	cmd := exec.Command("go", "test", fmt.Sprintf("./%s/day%s/", p.Args.Year, p.Args.Day))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("\nRunning tests: %s\n", cmd.String())

	if err := cmd.Run(); err != nil {
		fmt.Println("Tests failed.")
		return err
	}

	return nil
}

func (p PluginProcessor) GetInput() (*[]byte, error) {
	// Determine the input file path
	path := filepath.Join("./", p.Args.Year, "day"+p.Args.Day, "input.txt")

	// Read the input file
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading input file: %v", err)
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("input file is empty: %v", err)
	}

	return &data, nil
}

type Path int

const (
	year Path = iota
	day
	part
)

func GetProcessors(args Args) ([]PluginProcessor, error) {
	dir := filepath.Join("./", args.Year)
	processors := make([]PluginProcessor, 0)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		parts := strings.Split(path, string(os.PathSeparator))
		if len(parts) != 3 {
			return nil
		}

		if !strings.HasPrefix(parts[day], "day") || (args.Day != "0" && parts[day] != "day"+args.Day) {
			return nil
		}

		if !strings.HasPrefix(parts[part], "part") || strings.Contains(parts[part], "test") || !strings.HasSuffix(parts[part], ".go") {
			return nil
		}

		if args.Part != "0" && parts[part] != "part"+args.Part+".go" {
			return nil
		}

		args := Args{
			Year: parts[year],
			Day:  parts[day][3:],
			Part: parts[part][4 : len(parts[part])-3],
		}
		args.path, err = args.getPath()
		if err != nil {
			return err
		}

		processors = append(processors, PluginProcessor{
			Args:      args,
			Benchmark: &BenchmarkResult{},
		})

		return nil
	})

	if err != nil {
		fmt.Println("Error scanning directories:", err)
		return nil, err
	}

	return processors, nil
}
