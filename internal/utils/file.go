package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// Check if file exist.
func FileExist(file string) bool {
	_, err := os.Stat(file)
	return !errors.Is(err, os.ErrNotExist)
}

// Make a new directory relative to the current directory.
func MkRelDir(dir string) error {
	parent, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error in os.Getwd: %w", err)
	}

	scriptsDir := filepath.Join(parent, dir)

	err = os.MkdirAll(scriptsDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error in os.MkdirAll: %w", err)
	}

	return nil
}
