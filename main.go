package main

import (
	"fmt"

	"igorwulff.com/aoc/internal"
)

func main() {
	args := internal.ProcessArgs()

	plugin := internal.PluginProcessor{Args: args}
	if err := plugin.Build(); err != nil {
		fmt.Println(err)
		return
	}

	input, err := plugin.GetInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	plugin.RunTests()

	output, err := plugin.CallFunc(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nSolution:\n", output)
}
