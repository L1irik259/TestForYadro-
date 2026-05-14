package output

import (
	"bytes"
	"strings"
	"testing"
)

func TestPrintAnyOrderTo(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer
	PrintAnyOrderTo(&buf, map[string]int{"Alice": 2, "Bob": 1})

	got := buf.String()
	if !strings.Contains(got, "\"Alice\": 2") || !strings.Contains(got, "\"Bob\": 1") {
		t.Fatalf("unexpected output: %q", got)
	}
}

func TestPrintSortedTo(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer
	PrintSortedTo(&buf, map[string]int{"Bob": 2, "Alice": 2, "Carol": 1})

	want := []string{
		"\"Alice\": 2",
		"\"Bob\": 2",
		"\"Carol\": 1",
	}

	gotLines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(gotLines) != len(want) {
		t.Fatalf("unexpected number of lines: got %d want %d (%q)", len(gotLines), len(want), buf.String())
	}

	for i, line := range want {
		if gotLines[i] != line {
			t.Fatalf("unexpected line %d: got %q want %q", i, gotLines[i], line)
		}
	}
}
