package usage

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

const statsFileName = "emoji-usage.json"

type Stats map[string]int

func Load() (Stats, error) {
	path, err := statsFilePath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return Stats{}, nil
	}
	if err != nil {
		return nil, err
	}

	stats := Stats{}
	if err := json.Unmarshal(data, &stats); err != nil {
		return nil, err
	}

	return stats, nil
}

func Increment(emojiChar string) error {
	stats, err := Load()
	if err != nil {
		return err
	}

	stats[NormalizeEmoji(emojiChar)]++
	return Save(stats)
}

func Save(stats Stats) error {
	path, err := statsFilePath()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(stats, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0o644)
}

func (s Stats) Count(emojiChar string) int {
	if s == nil {
		return 0
	}

	return s[NormalizeEmoji(emojiChar)]
}

func statsFilePath() (string, error) {
	dataDir := os.Getenv("alfred_workflow_data")
	if dataDir == "" {
		return "", errors.New("alfred_workflow_data is not set")
	}
	return filepath.Join(dataDir, statsFileName), nil
}

func NormalizeEmoji(e string) string {
	return strings.ReplaceAll(e, "\uFE0F", "")
}
