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
	Args Args
}

func (p PluginProcessor) Build() error {
	cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", p.Args.path+".so", p.Args.path+".go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("\nCreating plugin: %s\n", cmd.String())
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
		return "", fmt.Errorf("unexpected function signature")
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

func GetAllPluginProcessors(args Args) ([]PluginProcessor, error) {
	var baseDir = filepath.Join("./", args.Year)
	var processors []PluginProcessor

	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
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

		if !strings.HasPrefix(parts[1], "day") {
			return nil
		}

		if !strings.HasPrefix(parts[2], "part") || strings.Contains(parts[2], "test") || !strings.HasSuffix(parts[2], ".go") {
			return nil
		}

		year := parts[0]
		day := parts[1][3:]
		part := parts[2][4 : len(parts[2])-3]

		args := Args{
			Year: year,
			Day:  day,
			Part: part,
		}
		args.path, err = args.getPath()
		if err != nil {
			return err
		}

		processors = append(processors, PluginProcessor{
			Args: args,
		})

		return nil
	})

	if err != nil {
		fmt.Println("Error scanning directories:", err)

		return nil, err
	}

	return processors, nil
}
