package aoc

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	internal "github.com/igorwulff/aoc/internal"
)

type BenchmarkResult struct {
	TotalTimeInMs int64
}

func Run() ([]BenchmarkResult, error) {
	args := internal.ProcessArgs()
	var plugins []internal.PluginProcessor
	// Benchmarking mode
	if args.Day == "0" {
		allPlugins, err := internal.GetAllPluginProcessors(args)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		plugins = append(plugins, allPlugins...)
	} else {
		plugins = append(plugins, internal.PluginProcessor{Args: args})
	}

	results := make([]BenchmarkResult, len(plugins))
	for i, plugin := range plugins {
		input, err := plugin.GetInput()
		if err != nil {
			fmt.Printf("Could not find input.txt for Day %s Part %s\n", plugin.Args.Day, plugin.Args.Part)
			return nil, err
		}

		fmt.Printf("Building day %s part %s... ", plugin.Args.Day, plugin.Args.Part)
		if err := plugin.Build(); err != nil {
			fmt.Print("ERROR!\n")
			fmt.Println(err)
			return nil, err
		}
		fmt.Print("OK!\n")

		fmt.Print("Executing tests... ")
		if err := plugin.RunTests(); err != nil {
			fmt.Print("ERROR!\n")
			fmt.Println(err)
			return nil, err
		}

		fmt.Print("Executing main function... ")
		timeBeforeExec := time.Now().UnixMilli()
		output, err := plugin.CallFunc(input)
		if err != nil {
			fmt.Print("ERROR!\n")
			fmt.Println(err)
			return nil, err
		}
		timeAfterExec := time.Now().UnixMilli()
		results[i] = BenchmarkResult{TotalTimeInMs: timeAfterExec - timeBeforeExec}

		fmt.Printf("Solution: %s\n", output)
	}

	fmt.Println("Execution finished! Total time:")
	w := tabwriter.NewWriter(os.Stdout, 0, 4, 4, ' ', 0)
	fmt.Fprintln(w, "Day\tPart\tTime (ms)")
	for i, plugin := range plugins {
		// @TODO: would be nice to join day & parts in a single row, but for now it's already nice they're alphabetically sorted
		fmt.Fprintf(w, "%s\t%s\t%d\n", plugin.Args.Day, plugin.Args.Part, results[i].TotalTimeInMs)
	}
	w.Flush()

	return results, nil
}
