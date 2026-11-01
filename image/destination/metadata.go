package destination

import (
	"path/filepath"
	"strings"
)

type Metadata struct {
	DirectoryPath *Path
	Path          *Path
	ThumbnailPath *Path
	HtmlPath      *Path
}

func NewMetadata(directoryPath string, publicDirectoryPath string, filePath string) *Metadata {
	base := filepath.Base(filePath)
	publicFilePath := filepath.Join(publicDirectoryPath, base)
	ext := filepath.Ext(publicFilePath)
	publicThumbnailPath := strings.TrimSuffix(publicFilePath, ext) + ".thumbnail.jpeg"
	htmlPath := strings.TrimSuffix(publicFilePath, ext) + ".html"

	return &Metadata{
		DirectoryPath: NewPath(directoryPath, publicDirectoryPath),
		Path:          NewPath(directoryPath, publicFilePath),
		ThumbnailPath: NewPath(directoryPath, publicThumbnailPath),
		HtmlPath:      NewPath(directoryPath, htmlPath),
	}
}

func (m *Metadata) SetDirectory(directoryPath string) {
	m.DirectoryPath.SetDirectory(directoryPath)
	m.Path.SetDirectory(directoryPath)
	m.ThumbnailPath.SetDirectory(directoryPath)
	m.HtmlPath.SetDirectory(directoryPath)
}
