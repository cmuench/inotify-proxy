package util

import "os"

// FileExists checks if a file or directory exists in the filesystem
func FileExists(path string) bool {
	info, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}
