package path

import (
	"path/filepath"
)

type Path struct {
	// Relative from where the process is execute
	// or
	// an absolute path
	Value string
	// Relative from where it will be available publicly
	PublicValue string
}

func NewPath(directoryPath string, path string) *Path {
	return &Path{
		Value:       filepath.Join(directoryPath, path),
		PublicValue: path,
	}
}

func (p *Path) SetDirectory(directoryPath string) {
	p.Value = filepath.Join(directoryPath, p.PublicValue)
}
