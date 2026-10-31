package image

import (
	"time"

	"github.com/theobori/recueil/image/destination"
	"github.com/theobori/recueil/image/source"
)

type Metadata struct {
	Source      *source.Metadata
	Destination *destination.Metadata
}

func (m *Metadata) Title() string {
	return m.Source.Title
}

func (m *Metadata) Description() string {
	return m.Source.Description
}

func (m *Metadata) Tags() []string {
	return m.Source.Tags
}

func (m *Metadata) Date() time.Time {
	return m.Source.Date
}
