package bwutil

import (
	"os"
	"path/filepath"
	"strings"
)

func ExpandPath(path string) (string) {
	if strings.HasPrefix(path, "~") {
		home, _ := os.UserHomeDir()
		return filepath.Join(home, path[1:])
	}
	return path
}
