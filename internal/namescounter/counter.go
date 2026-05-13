package namescounter

import (
	"bufio"
	"os"
	"sync"
)

const chunkSize = 5

func CountStreaming(fileName string, numWorkers int) (map[string]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	chunks := make(chan []string, numWorkers)
	results := make(chan map[string]int, numWorkers)

	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(chunks, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	scanner := bufio.NewScanner(file)
	buffer := make([]byte, 64*1024)
	scanner.Buffer(buffer, 1024*1024)

	chunk := make([]string, 0, chunkSize)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		chunk = append(chunk, line)

		if len(chunk) >= chunkSize {
			safeChunk := make([]string, len(chunk))
			copy(safeChunk, chunk)
			chunks <- safeChunk
			chunk = chunk[:0]
		}
	}

	if err := scanner.Err(); err != nil {
		close(chunks)
		return nil, err
	}

	if len(chunk) > 0 {
		safeChunk := make([]string, len(chunk))
		copy(safeChunk, chunk)
		chunks <- safeChunk
	}

	close(chunks)

	finalCounts := make(map[string]int)
	for partial := range results {
		for name, count := range partial {
			finalCounts[name] += count
		}
	}

	return finalCounts, nil
}

func worker(chunks <-chan []string, results chan<- map[string]int, wg *sync.WaitGroup) {
	defer wg.Done()

	localCounts := make(map[string]int)

	for chunk := range chunks {
		for _, name := range chunk {
			localCounts[name]++
		}
	}

	results <- localCounts
}
