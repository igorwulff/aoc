package aoc

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	internal "github.com/igorwulff/aoc/internal"
)

type BenchmarkResult struct {
	totalTimeInMs int64
}

func Benchmark() {
	args := internal.ProcessArgs()

	processors, err := internal.GetAllPluginProcessors(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	results := make([]BenchmarkResult, len(processors))
	fmt.Println("Starting benchmark!")
	for i, processor := range processors {
		input, err := processor.GetInput()
		if err != nil {
			fmt.Printf("Could not find input.txt for Day %s Part %s\n", processor.Args.Day, processor.Args.Part)
			return
		}

		fmt.Printf("Building day %s part %s... ", processor.Args.Day, processor.Args.Part)
		if err := processor.Build(); err != nil {
			fmt.Print("ERROR!\n")
			fmt.Println(err)
			return
		}
		fmt.Print("OK!\n")

		fmt.Print("Executing tests... ")
		if err := processor.RunTests(); err != nil {
			fmt.Print("ERROR!\n")
			fmt.Println(err)
			return
		}

		fmt.Print("Executing main function... ")
		timeBeforeExec := time.Now().UnixMilli()
		_, err = processor.CallFunc(input)
		if err != nil {
			fmt.Print("ERROR!\n")
			fmt.Println(err)
			return
		}
		timeAfterExec := time.Now().UnixMilli()
		results[i] = BenchmarkResult{totalTimeInMs: timeAfterExec - timeBeforeExec}
	}

	fmt.Println("Benchmark finished!")
	w := tabwriter.NewWriter(os.Stdout, 0, 4, 4, ' ', 0)
	fmt.Fprintln(w, "Day\tPart\tTime (ms)")
	for i, processor := range processors {
		// @TODO: would be nice to join day & parts in a single row, but for now it's already nice they're alphabetically sorted
		fmt.Fprintf(w, "%s\t%s\t%d\n", processor.Args.Day, processor.Args.Part, results[i].totalTimeInMs)
	}
	w.Flush()
}
