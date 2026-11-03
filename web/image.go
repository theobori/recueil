package web

import (
	"github.com/theobori/recueil/image"
)

type ImageHtmlData struct {
	Metadata         *image.Metadata
	PreviousMetadata *image.Metadata
	NextMetadata     *image.Metadata
	TagsDirectory    string
	WebringUrl       string
}

type ImageCssData struct {
	ColorBg    string
	ColorFg    string
	ColorLink  string
	ColorHover string
}
