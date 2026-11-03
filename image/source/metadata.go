package source

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/theobori/recueil/internal"
	rpath "github.com/theobori/recueil/path"
)

type Metadata struct {
	Title       string   `yaml:"title"`
	Description string   `yaml:"description"`
	Tags        []string `yaml:"tags"`
	Date        time.Time
	ImagePath   *rpath.Path
}

func NewMetadataFromPath(path string) (*Metadata, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	metadata := Metadata{
		Date:      fileInfo.ModTime(),
		ImagePath: rpath.NewPath("", path),
	}

	yamlPath := strings.TrimSuffix(path, ".") + ".yaml"

	err = internal.YamlUnmarshalFromPath(yamlPath, &metadata)
	if err != nil {
		return nil, err
	}

	if metadata.Title == "" {
		metadata.Title = filepath.Base(metadata.ImagePath.Value)
	}

	return &metadata, nil
}
