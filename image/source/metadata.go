package source

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/theobori/recueil/internal"
)

type Metadata struct {
	Title       string   `yaml:"title"`
	Description string   `yaml:"description"`
	Tags        []string `yaml:"tags"`
	Date        time.Time
	Path        string
}

func NewMetadataFromPath(path string) (*Metadata, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	metadata := Metadata{
		Date: fileInfo.ModTime(),
		Path: path,
	}

	yamlPath := strings.TrimSuffix(path, ".") + ".yaml"

	err = internal.YamlUnmarshalFromPath(yamlPath, &metadata)
	if err != nil {
		return nil, err
	}

	if metadata.Title == "" {
		metadata.Title = filepath.Base(metadata.Path)
	}

	return &metadata, nil
}
