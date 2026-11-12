package source

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/theobori/recueil/internal"
	rpath "github.com/theobori/recueil/path"
)

type Metadata struct {
	Title       string      `yaml:"title"`
	Description string      `yaml:"description"`
	Tags        []string    `yaml:"tags"`
	Date        time.Time   `yaml:"date"`
	ImagePath   *rpath.Path `yaml:"-"`
}

func TransformTag(tag string) string {
	tag = strings.TrimSpace(tag)
	tag = strings.ReplaceAll(tag, "/", "")
	tag = strings.ReplaceAll(tag, ".", "")

	return tag
}

func TransformTags(tags []string) []string {
	seen := map[string]any{}
	for _, tag := range tags {
		transformedTag := TransformTag(tag)
		if transformedTag != "" {
			seen[transformedTag] = struct{}{}
		}
	}

	ans := []string{}
	for tag := range seen {
		ans = append(ans, tag)
	}

	slices.Sort(ans)

	return ans
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

	err = internal.TryYamlUnmarshalFromPath(yamlPath, &metadata)
	if err != nil {
		return nil, err
	}

	if metadata.Title == "" {
		metadata.Title = filepath.Base(metadata.ImagePath.Value)
	}

	metadata.Tags = TransformTags(metadata.Tags)

	return &metadata, nil
}
