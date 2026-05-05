package usage

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestIncrementNormalizesEquivalentEmojiVariants(t *testing.T) {
	t.Setenv("alfred_workflow_data", t.TempDir())

	if err := Increment("🖐️"); err != nil {
		t.Fatalf("increment variant with VS16: %v", err)
	}
	if err := Increment("🖐"); err != nil {
		t.Fatalf("increment variant without VS16: %v", err)
	}

	stats, err := Load()
	if err != nil {
		t.Fatalf("load stats: %v", err)
	}

	if got := stats.Count("🖐️"); got != 2 {
		t.Fatalf("expected normalized count 2, got %d", got)
	}

	if got := stats[NormalizeEmoji("🖐️")]; got != 2 {
		t.Fatalf("expected stored normalized count 2, got %d", got)
	}
}

func TestLoadMissingFileReturnsEmptyStats(t *testing.T) {
	t.Setenv("alfred_workflow_data", t.TempDir())

	stats, err := Load()
	if err != nil {
		t.Fatalf("load stats: %v", err)
	}

	if len(stats) != 0 {
		t.Fatalf("expected empty stats, got %d entries", len(stats))
	}
}

func TestSaveCreatesExpectedFile(t *testing.T) {
	dataDir := t.TempDir()
	t.Setenv("alfred_workflow_data", dataDir)

	if err := Save(Stats{"🙂": 3}); err != nil {
		t.Fatalf("save stats: %v", err)
	}

	if _, err := os.Stat(filepath.Join(dataDir, statsFileName)); err != nil {
		t.Fatalf("stat saved file: %v", err)
	}
}

func TestResetRemovesStatsFile(t *testing.T) {
	dataDir := t.TempDir()
	t.Setenv("alfred_workflow_data", dataDir)

	if err := Save(Stats{"🙂": 3}); err != nil {
		t.Fatalf("save stats: %v", err)
	}

	if err := Reset(); err != nil {
		t.Fatalf("reset stats: %v", err)
	}

	if _, err := os.Stat(filepath.Join(dataDir, statsFileName)); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("expected stats file to be removed, got err=%v", err)
	}
}
