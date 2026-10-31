package web

import (
	"github.com/theobori/recueil/image"
)

type HomeData struct {
	Metadatas     image.Metadatas
	Tags          []string
	Title         string
	TagsDirectory string
}
