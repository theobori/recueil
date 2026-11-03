package web

import (
	"embed"
	"log"
	"path/filepath"
	"sort"

	"github.com/theobori/recueil/configuration"
	"github.com/theobori/recueil/image"
)

//go:embed templates
var templates embed.FS

func Execute(destinationDirectory string, metadatas image.Metadatas, config *configuration.Configuration, verbose bool) error {
	var err error

	byTags := metadatas.ByTags()

	n := len(byTags)
	tags := make([]string, n)
	i := 0
	for tagName := range byTags {
		tags[i] = tagName
		i++
	}

	sort.Strings(tags)

	homeHtmlData := HomeHtmlData{
		Metadatas:     metadatas,
		Title:         config.Title,
		Tags:          tags,
		TagsDirectory: config.TagsDirectory,
		WebringUrl:    config.WebringUrl,
	}
	err = ExecuteHelper(templates, filepath.Join(destinationDirectory, "index.html"), &homeHtmlData, verbose, "templates/home.html")
	if err != nil {
		return err
	}
	homeCssData := HomeCssData{
		ColorBg:    config.Css.ColorBg,
		ColorFg:    config.Css.ColorFg,
		ColorLink:  config.Css.ColorLink,
		ColorHover: config.Css.ColorHover,
	}
	err = ExecuteHelper(templates, filepath.Join(destinationDirectory, "home.css"), &homeCssData, verbose, "templates/home.css")
	if err != nil {
		return err
	}

	for tagName, boundMetadatas := range byTags {
		tagHtmlData := TagHtmlData{
			Metadatas:     boundMetadatas,
			Title:         tagName,
			Tags:          tags,
			TagsDirectory: config.TagsDirectory,
		}

		err = ExecuteHelper(templates, filepath.Join(destinationDirectory, config.TagsDirectory, tagName+".html"), &tagHtmlData, verbose, "templates/home.html")
		if err != nil {
			return err
		}
	}

	n = len(metadatas)
	metadatas = append(metadatas, nil)
	var previousMetadata *image.Metadata
	var nextMetadata *image.Metadata

	for i := 0; i < n; i++ {
		metadata := metadatas[i]
		nextMetadata = metadatas[i+1]

		imageHtmlData := ImageHtmlData{
			Metadata:         metadata,
			TagsDirectory:    config.TagsDirectory,
			PreviousMetadata: previousMetadata,
			NextMetadata:     nextMetadata,
			WebringUrl:       config.WebringUrl,
		}

		err = ExecuteHelper(templates, metadata.Destination.HtmlPath.Value, &imageHtmlData, verbose, "templates/image.html")
		if err != nil {
			return err
		}

		previousMetadata = metadata
	}
	metadatas = metadatas[0:n]

	imageCssData := ImageCssData{
		ColorBg:    config.Css.ColorBg,
		ColorFg:    config.Css.ColorFg,
		ColorLink:  config.Css.ColorLink,
		ColorHover: config.Css.ColorHover,
	}
	err = ExecuteHelper(templates, filepath.Join(destinationDirectory, "image.css"), &imageCssData, verbose, "templates/image.css")
	if err != nil {
		return err
	}

	destinationPath := filepath.Join(destinationDirectory, "rss.xml")
	log.Printf("Evaluating template and writing to '%s'\n", destinationPath)
	err = ExecuteRss(metadatas, destinationPath, config.Title, config.SiteUrl)
	if err != nil {
		return err
	}

	return nil
}
