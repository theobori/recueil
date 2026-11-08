package web

type HomeHtmlData GalleryHtmlData

func (e *Executor) executeHome() error {
	homeHtmlData := HomeHtmlData{
		Metadatas:     e.metadatas,
		Title:         e.config.Title,
		Tags:          e.tags,
		TagsDirectory: e.config.TagsDirectory,
		WebringUrl:    e.config.WebringUrl,
	}
	err := e.executeHelper("index.html", &homeHtmlData, "templates/gallery.html")
	if err != nil {
		return err
	}

	return nil
}
