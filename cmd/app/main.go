package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/L1irik259/TestForYadro/internal/namescounter"
	"github.com/L1irik259/TestForYadro/internal/output"
)

func main() {
	fileName := flag.String("file", "", "путь к файлу")
	sortByFrequency := flag.Bool("sort", false, "сортировка по частоте")
	workers := flag.Int("workers", runtime.NumCPU(), "количество воркеров")

	flag.Parse()

	if *fileName == "" {
		fmt.Println("Укажи файл через -file")
		os.Exit(1)
	}

	counts, err := namescounter.CountStreaming(*fileName, *workers)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		os.Exit(1)
	}

	if *sortByFrequency {
		output.PrintSorted(counts)
	} else {
		output.PrintAnyOrder(counts)
	}
}
