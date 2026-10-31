package web

import (
	"github.com/theobori/recueil/image"
)

type TagData struct {
	Metadatas     image.Metadatas
	Tags          []string
	Title         string
	TagName       string
	TagsDirectory string
}
