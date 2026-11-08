package web

import "github.com/theobori/recueil/image"

type GalleryHtmlData struct {
	Metadatas     image.Metadatas
	Tags          []string
	Title         string
	TagsDirectory string
	WebringUrl    string
}

type GalleryCssData struct {
	ColorBg            string
	ColorFg            string
	ColorLink          string
	ColorHover         string
	ThumbnailPixelSize int
}
