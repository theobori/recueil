package internal

import (
	"path/filepath"
	"strings"
)

func ReplaceBaseDirectory(path string, baseDirectorySource string, baseDirectoryDestination string) string {
	return filepath.Join(baseDirectoryDestination, strings.TrimPrefix(path, baseDirectorySource))
}
