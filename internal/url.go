package internal

import (
	"path/filepath"
	"strings"
)

func UrlJoin(url string, elem ...string) string {
	return strings.TrimSuffix(url, "/") + "/" + strings.TrimPrefix(filepath.Join(elem...), "/")
}
