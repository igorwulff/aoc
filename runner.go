package aoc

import (
	"fmt"
	"os"
	"text/tabwriter"

	internal "github.com/igorwulff/aoc/internal"
)

func Run() {
	args := internal.ProcessArgs()
	plugins, err := internal.GetProcessors(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, p := range plugins {
		input, err := p.GetInput()
		if err != nil {
			fmt.Printf("Could not find input.txt for Day %s Part %s\n", p.Args.Day, p.Args.Part)
			return
		}

		fmt.Printf("Running day %s, part %s", p.Args.Day, p.Args.Part)

		if err := p.Build(); err != nil {
			fmt.Print("ERROR!\n", err)
			return
		}

		if err := p.RunTests(); err != nil {
			fmt.Print("ERROR!\n", err)
			return
		}

		fmt.Println("Trying to solve puzzle...")

		p.Benchmark.StartTimer()
		output, err := p.CallFunc(input)
		if err != nil {
			fmt.Print("ERROR!\n", err)
			return
		}
		p.Benchmark.StopTimer()

		fmt.Printf("Solution: %s\n\n", output)
	}

	fmt.Println("Execution finished!")
	w := tabwriter.NewWriter(os.Stdout, 0, 4, 4, ' ', 0)
	fmt.Fprintln(w, "\nDay\tPart\tTime (ms)")

	var sum float64 = 0.0
	for _, p := range plugins {
		sum += p.Benchmark.GetTotalTime()
		// @TODO: would be nice to join day & parts in a single row, but for now it's already nice they're alphabetically sorted
		fmt.Fprintf(w, "%s\t%s\t%8.3f\n", p.Args.Day, p.Args.Part, p.Benchmark.GetTotalTime())
	}
	w.Flush()

	fmt.Printf("\nTotal time: %.3f ms\n", sum)
}
