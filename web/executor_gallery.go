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
	ColorSelectionFg   string
	ColorSelectionBg   string
	ThumbnailPixelSize int
}

func (e *Executor) executeGallery() error {
	err := e.executeHelper(
		"gallery.css",
		&GalleryCssData{
			ColorBg:            e.config.Css.ColorBg,
			ColorFg:            e.config.Css.ColorFg,
			ColorLink:          e.config.Css.ColorLink,
			ColorHover:         e.config.Css.ColorHover,
			ColorSelectionFg:   e.config.Css.ColorSelectionFg,
			ColorSelectionBg:   e.config.Css.ColorSelectionBg,
			ThumbnailPixelSize: e.config.Images.ThumbnailPixelSize,
		},
		"templates/gallery.css",
	)
	if err != nil {
		return err
	}

	return nil
}
