package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunRequiresFile(t *testing.T) {
	t.Parallel()

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	code := run(nil, &stdout, &stderr)
	if code != 1 {
		t.Fatalf("unexpected exit code: got %d want 1", code)
	}

	if !strings.Contains(stdout.String(), "Укажи файл через -file") {
		t.Fatalf("unexpected stdout: %q", stdout.String())
	}
}

func TestRunProducesSortedOutput(t *testing.T) {
	t.Parallel()

	file := filepath.Join(t.TempDir(), "input.txt")
	content := strings.Join([]string{"Bob", "Alice", "Bob"}, "\n")
	if err := os.WriteFile(file, []byte(content), 0o600); err != nil {
		t.Fatalf("write temp file: %v", err)
	}

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	code := run([]string{"-file", file, "-sort"}, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("unexpected exit code: got %d want 0", code)
	}

	got := strings.TrimSpace(stdout.String())
	want := []string{"\"Bob\": 2", "\"Alice\": 1"}
	gotLines := strings.Split(got, "\n")
	if len(gotLines) != len(want) {
		t.Fatalf("unexpected output: %q", got)
	}

	for i, line := range want {
		if gotLines[i] != line {
			t.Fatalf("unexpected line %d: got %q want %q", i, gotLines[i], line)
		}
	}
}
