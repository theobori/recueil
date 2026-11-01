package web

import (
	"github.com/theobori/recueil/image"
)

type HomeHtmlData struct {
	Metadatas     image.Metadatas
	Tags          []string
	Title         string
	TagsDirectory string
}

type HomeCssData struct {
}
