package models

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type WordFilter struct {
	mu          sync.RWMutex
	bannedWords map[string]struct{}
	filePath    string
}

// Initialize Word Filter
func NewWordFilter(filePath string) (*WordFilter, error) {
	filter := &WordFilter{
		bannedWords: make(map[string]struct{}),
		filePath:    filePath,
	}
	err := filter.loadFromFile()
	if err != nil {
		return nil, err
	}
	return filter, nil
}

// Load Banned Words from File
func (wf *WordFilter) loadFromFile() error {
	file, err := os.Open(wf.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word != "" {
			wf.bannedWords[word] = struct{}{}
		}
	}
	return scanner.Err()
}

// Testfor Banned Word
func (wf *WordFilter) IsBanned(word string) bool {
	wf.mu.RLock()
	defer wf.mu.RUnlock()
	_, exists := wf.bannedWords[strings.ToLower(word)]
	return exists
}

// Add New Banned Word
func (wf *WordFilter) AddBannedWord(word string) error {
	word = strings.ToLower(strings.TrimSpace(word))
	if word == "" {
		return fmt.Errorf("cannot add an empty word")
	}

	wf.mu.Lock()
	defer wf.mu.Unlock()

	if _, exists := wf.bannedWords[word]; exists {
		return fmt.Errorf("word already banned: %s", word)
	}
	wf.bannedWords[word] = struct{}{}

	// write to file
	file, err := os.OpenFile(wf.filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to update banned words file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(word + "\n")
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}

// Delete Banned Word
func (wf *WordFilter) DeleteBannedWord(word string) error {
	word = strings.ToLower(strings.TrimSpace(word))
	if word == "" {
		return fmt.Errorf("cannot delete an empty word")
	}
	wf.mu.Lock()
	defer wf.mu.Unlock()
	if _, exists := wf.bannedWords[word]; !exists {
		return fmt.Errorf("word not banned: %s", word)
	}
	delete(wf.bannedWords, word)

	// write to file
	file, err := os.OpenFile(wf.filePath, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("failed to update banned words file: %w", err)
	}
	defer file.Close()
	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != word {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}
	file.Truncate(0)
	file.Seek(0, 0)
	for _, line := range lines {
		_, err = file.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("failed to write to file: %w", err)
		}
	}
	return nil
}
