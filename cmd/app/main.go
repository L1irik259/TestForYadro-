package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/L1irik259/TestForYadro/internal/namescounter"
	"github.com/L1irik259/TestForYadro/internal/output"
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout io.Writer, stderr io.Writer) int {
	flagSet := flag.NewFlagSet("app", flag.ContinueOnError)
	flagSet.SetOutput(stderr)

	fileName := flagSet.String("file", "", "путь к файлу")
	sortByFrequency := flagSet.Bool("sort", false, "сортировка по частоте")
	workers := flagSet.Int("workers", runtime.NumCPU(), "количество воркеров")

	if err := flagSet.Parse(args); err != nil {
		return 1
	}

	if *fileName == "" {
		fmt.Fprintln(stdout, "Укажи файл через -file")
		return 1
	}

	counts, err := namescounter.CountStreaming(*fileName, *workers)
	if err != nil {
		fmt.Fprintf(stdout, "Ошибка: %v\n", err)
		return 1
	}

	if *sortByFrequency {
		output.PrintSortedTo(stdout, counts)
	} else {
		output.PrintAnyOrderTo(stdout, counts)
	}

	return 0
}
