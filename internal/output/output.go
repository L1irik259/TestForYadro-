package output

import (
	"fmt"
	"io"
	"sort"
)

func PrintAnyOrder(counts map[string]int) {
	PrintAnyOrderTo(io.Discard, counts)
}

func PrintAnyOrderTo(w io.Writer, counts map[string]int) {
	for name, count := range counts {
		fmt.Fprintf(w, "%q: %d\n", name, count)
	}
}

func PrintSorted(counts map[string]int) {
	PrintSortedTo(io.Discard, counts)
}

func PrintSortedTo(w io.Writer, counts map[string]int) {
	type pair struct {
		Name  string
		Count int
	}

	pairs := make([]pair, 0, len(counts))
	for name, count := range counts {
		pairs = append(pairs, pair{Name: name, Count: count})
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].Count == pairs[j].Count {
			return pairs[i].Name < pairs[j].Name
		}

		return pairs[i].Count > pairs[j].Count
	})

	for _, p := range pairs {
		fmt.Fprintf(w, "%q: %d\n", p.Name, p.Count)
	}
}
