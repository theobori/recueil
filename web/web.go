package web

import (
	"embed"
	"path/filepath"

	"github.com/theobori/recueil/configuration"
	"github.com/theobori/recueil/image"
)

//go:embed templates
var templates embed.FS

func Execute(destinationDirectory string, metadatas image.Metadatas, config *configuration.Configuration) error {
	var err error

	byTags := metadatas.ByTags()

	n := len(byTags)
	tags := make([]string, n)
	i := 0
	for tagName := range byTags {
		tags[i] = tagName
		i++
	}

	homeData := HomeData{
		Metadatas:     metadatas,
		Title:         config.Title,
		Tags:          tags,
		TagsDirectory: config.TagsDirectory,
	}
	err = ExecuteHelper(templates, filepath.Join(destinationDirectory, "index.html"), &homeData, "templates/home.html")
	if err != nil {
		return err
	}

	for tagName, boundMetadatas := range byTags {
		tagData := TagData{
			Metadatas:     boundMetadatas,
			Title:         tagName,
			Tags:          tags,
			TagName:       tagName,
			TagsDirectory: config.TagsDirectory,
		}

		err = ExecuteHelper(templates, filepath.Join(destinationDirectory, config.TagsDirectory, tagName+".html"), &tagData, "templates/tag.html")
		if err != nil {
			return err
		}
	}

	metadatas = append(metadatas, nil)
	var previousMetadata *image.Metadata
	var nextMetadata *image.Metadata

	for i := 0; i < len(metadatas)-1; i++ {
		metadata := metadatas[i]
		nextMetadata = metadatas[i+1]

		imageData := ImageData{
			Metadata:         metadata,
			TagsDirectory:    config.TagsDirectory,
			PreviousMetadata: previousMetadata,
			NextMetadata:     nextMetadata,
		}

		err = ExecuteHelper(templates, metadata.Destination.HtmlPath.Value, &imageData, "templates/image.html")
		if err != nil {
			return err
		}

		previousMetadata = metadata
	}

	return nil
}
