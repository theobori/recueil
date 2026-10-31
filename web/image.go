package web

import (
	"github.com/theobori/recueil/image"
)

type ImageData struct {
	Metadata         *image.Metadata
	PreviousMetadata *image.Metadata
	NextMetadata     *image.Metadata
	TagsDirectory    string
}
