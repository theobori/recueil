package web

import (
	"path/filepath"

	"github.com/theobori/recueil/image"
)

type TagHtmlData GalleryHtmlData

func (e *Executor) executeTag(tagName string, boundMetadatas []*image.Metadata) error {
	return e.executeHelper(
		filepath.Join(e.config.TagsDirectory, tagName+".html"),
		&TagHtmlData{
			Metadatas:     boundMetadatas,
			Title:         tagName,
			Tags:          e.tags,
			TagsDirectory: e.config.TagsDirectory,
			Webrings:      e.config.Webrings,
		},
		"templates/gallery.html",
	)
}

func (e *Executor) executeTags() error {
	for tagName, boundMetadatas := range e.byTags {
		err := e.executeTag(tagName, boundMetadatas)
		if err != nil {
			return err
		}
	}

	return nil
}
