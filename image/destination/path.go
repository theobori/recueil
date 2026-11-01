package destination

import (
	"path/filepath"
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
