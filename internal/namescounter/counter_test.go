package namescounter

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestCountStreamingCountsNames(t *testing.T) {
	t.Parallel()

	file := writeTempFile(t, []string{
		"Alice",
		"",
		"Bob",
		"Alice",
		"Carol",
		"Bob",
		"Bob",
		"Alice",
	})

	counts, err := CountStreaming(file, 3)
	if err != nil {
		t.Fatalf("CountStreaming returned error: %v", err)
	}

	want := map[string]int{
		"Alice": 3,
		"Bob":   3,
		"Carol": 1,
	}

	if !reflect.DeepEqual(counts, want) {
		t.Fatalf("unexpected counts: got %#v want %#v", counts, want)
	}
}

func TestCountStreamingMissingFile(t *testing.T) {
	t.Parallel()

	_, err := CountStreaming(filepath.Join(t.TempDir(), "missing.txt"), 2)
	if err == nil {
		t.Fatal("expected error for missing file")
	}
}

func writeTempFile(t *testing.T, lines []string) string {
	t.Helper()

	file := filepath.Join(t.TempDir(), "input.txt")
	content := ""
	for i, line := range lines {
		content += line
		if i < len(lines)-1 {
			content += "\n"
		}
	}

	if err := os.WriteFile(file, []byte(content), 0o600); err != nil {
		t.Fatalf("write temp file: %v", err)
	}

	return file
}
