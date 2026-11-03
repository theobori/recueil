package destination

import (
	"path/filepath"
	"strings"

	"github.com/theobori/recueil/path"
)

type Metadata struct {
	DirectoryPath *path.Path
	ImagePath     *path.Path
	ThumbnailPath *path.Path
	HtmlPath      *path.Path
}

func NewMetadata(directoryPath string, publicDirectoryPath string, filePath string) *Metadata {
	base := filepath.Base(filePath)
	publicFilePath := filepath.Join(publicDirectoryPath, base)
	ext := filepath.Ext(publicFilePath)
	publicFilePathWithoutExt := strings.TrimSuffix(publicFilePath, ext)
	publicThumbnailPath := publicFilePathWithoutExt + ".thumbnail.jpeg"
	htmlPath := publicFilePathWithoutExt + ".html"

	return &Metadata{
		DirectoryPath: path.NewPath(directoryPath, publicDirectoryPath),
		ImagePath:     path.NewPath(directoryPath, publicFilePath),
		ThumbnailPath: path.NewPath(directoryPath, publicThumbnailPath),
		HtmlPath:      path.NewPath(directoryPath, htmlPath),
	}
}

func (m *Metadata) SetDirectory(directoryPath string) {
	m.DirectoryPath.SetDirectory(directoryPath)
	m.ImagePath.SetDirectory(directoryPath)
	m.ThumbnailPath.SetDirectory(directoryPath)
	m.HtmlPath.SetDirectory(directoryPath)
}
