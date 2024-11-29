package internal

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"plugin"
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
	symProcess, err := plug.Lookup("Part" + p.Args.part)
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

func (p PluginProcessor) RunTests() {
	cmd := exec.Command("go", "test", fmt.Sprintf("./%s/day%s/", p.Args.year, p.Args.day))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("\nRunning tests: %s\n", cmd.String())

	if err := cmd.Run(); err != nil {
		fmt.Println("Tests failed.")
		return
	}
}

func (p PluginProcessor) GetInput() (*[]byte, error) {
	// Determine the input file path
	path := filepath.Join("./", p.Args.year, "day"+p.Args.day, p.Args.input+".txt")

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
