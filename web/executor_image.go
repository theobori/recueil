package web

import (
	"github.com/theobori/recueil/configuration"
	"github.com/theobori/recueil/image"
)

type ImageHtmlData struct {
	Metadata         *image.Metadata
	PreviousMetadata *image.Metadata
	NextMetadata     *image.Metadata
	TagsDirectory    string
	Webrings         []configuration.Webring
}

type ImageCssData struct {
	ColorBg          string
	ColorFg          string
	ColorLink        string
	ColorHover       string
	ColorSelectionFg string
	ColorSelectionBg string
}

func (e *Executor) executeImages() error {
	n := len(e.metadatas)
	metadatas := append(e.metadatas, nil)

	var previousMetadata *image.Metadata
	var nextMetadata *image.Metadata

	for i := range n {
		metadata := metadatas[i]
		nextMetadata = metadatas[i+1]

		imageHtmlData := ImageHtmlData{
			Metadata:         metadata,
			TagsDirectory:    e.config.TagsDirectory,
			PreviousMetadata: previousMetadata,
			NextMetadata:     nextMetadata,
			Webrings:         e.config.Webrings,
		}

		err := e.executeHelper(metadata.Destination.HtmlPath.PublicValue, &imageHtmlData, "templates/image.html")
		if err != nil {
			return err
		}

		previousMetadata = metadata
	}
	metadatas = metadatas[0:n]

	err := e.executeHelper(
		"image.css",
		&ImageCssData{
			ColorBg:          e.config.Css.ColorBg,
			ColorFg:          e.config.Css.ColorFg,
			ColorLink:        e.config.Css.ColorLink,
			ColorHover:       e.config.Css.ColorHover,
			ColorSelectionFg: e.config.Css.ColorSelectionFg,
			ColorSelectionBg: e.config.Css.ColorSelectionBg,
		},
		"templates/image.css",
	)
	if err != nil {
		return err
	}

	return nil
}
