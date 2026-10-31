package destination

import (
	"path/filepath"
	"strings"
)

type Path struct {
	// Relative from where the process is executed or absolute path
	Value string
	// Relative from where it will be available publicly
	PublicValue string
}

func NewPath(directoryPath string, value string) *Path {
	return &Path{
		Value:       filepath.Join(directoryPath, value),
		PublicValue: value,
	}
}

func (p *Path) SetDirectory(directoryPath string) {
	p.Value = filepath.Join(directoryPath, p.PublicValue)
}

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
	HtmlPath := strings.TrimSuffix(publicFilePath, ext) + ".html"

	return &Metadata{
		DirectoryPath: NewPath(directoryPath, publicDirectoryPath),
		Path:          NewPath(directoryPath, publicFilePath),
		ThumbnailPath: NewPath(directoryPath, publicThumbnailPath),
		HtmlPath:      NewPath(directoryPath, HtmlPath),
	}
}

func (m *Metadata) SetDirectory(directoryPath string) {
	m.DirectoryPath.SetDirectory(directoryPath)
	m.Path.SetDirectory(directoryPath)
	m.ThumbnailPath.SetDirectory(directoryPath)
	m.HtmlPath.SetDirectory(directoryPath)
}
