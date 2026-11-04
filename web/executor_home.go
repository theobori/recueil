package web

import (
	"github.com/theobori/recueil/image"
)

type HomeHtmlData struct {
	Metadatas     image.Metadatas
	Tags          []string
	Title         string
	TagsDirectory string
	WebringUrl    string
}

type HomeCssData struct {
	ColorBg            string
	ColorFg            string
	ColorLink          string
	ColorHover         string
	ThumbnailPixelSize int
}

func (e *Executor) executeHome() error {
	var err error

	homeHtmlData := HomeHtmlData{
		Metadatas:     e.metadatas,
		Title:         e.config.Title,
		Tags:          e.tags,
		TagsDirectory: e.config.TagsDirectory,
		WebringUrl:    e.config.WebringUrl,
	}
	err = e.executeHelper("index.html", &homeHtmlData, "templates/home.html")
	if err != nil {
		return err
	}

	err = e.executeHelper(
		"home.css",
		&HomeCssData{
			ColorBg:            e.config.Css.ColorBg,
			ColorFg:            e.config.Css.ColorFg,
			ColorLink:          e.config.Css.ColorLink,
			ColorHover:         e.config.Css.ColorHover,
			ThumbnailPixelSize: e.config.Images.ThumbnailPixelSize,
		},
		"templates/home.css",
	)
	if err != nil {
		return err
	}

	return nil
}
