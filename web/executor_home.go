package web

func (e *Executor) executeHome() error {
	var err error

	homeHtmlData := GalleryHtmlData{
		Metadatas:     e.metadatas,
		Title:         e.config.Title,
		Tags:          e.tags,
		TagsDirectory: e.config.TagsDirectory,
		WebringUrl:    e.config.WebringUrl,
	}
	err = e.executeHelper("index.html", &homeHtmlData, "templates/gallery.html")
	if err != nil {
		return err
	}

	err = e.executeHelper(
		"gallery.css",
		&GalleryCssData{
			ColorBg:            e.config.Css.ColorBg,
			ColorFg:            e.config.Css.ColorFg,
			ColorLink:          e.config.Css.ColorLink,
			ColorHover:         e.config.Css.ColorHover,
			ThumbnailPixelSize: e.config.Images.ThumbnailPixelSize,
		},
		"templates/gallery.css",
	)
	if err != nil {
		return err
	}

	return nil
}
