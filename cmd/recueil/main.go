package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"sort"

	"github.com/theobori/recueil/configuration"
	"github.com/theobori/recueil/image"
	"github.com/theobori/recueil/internal"
	"github.com/theobori/recueil/web"
)

func main() {
	var (
		sourceDirectory      string
		destinationDirectory string
		configurationPath    string
	)

	flag.StringVar(&sourceDirectory, "source-directory", "./recueil", "The source directory which should be your 'recueil' base directory.")
	flag.StringVar(&destinationDirectory, "destination-directory", "./recueil-static", "The destination directory where files are installed.")
	flag.StringVar(&configurationPath, "configuration", "", "The configuration YAML file path.")
	flag.Parse()

	var config *configuration.Configuration
	var err error
	if configurationPath != "" {
		config, err = configuration.NewConfigurationFromPath(configurationPath)
	} else {
		config = configuration.NewConfigurationDefault()
	}

	err = os.MkdirAll(destinationDirectory, 0o755)
	if err != nil {
		log.Fatalln(err)
	}

	baseDestinationDirectory := filepath.Join(destinationDirectory, config.BaseDirectory)
	err = os.MkdirAll(baseDestinationDirectory, 0o755)
	if err != nil {
		log.Fatalln(err)
	}
	tagsDestinationDirectory := filepath.Join(destinationDirectory, config.TagsDirectory)
	err = os.MkdirAll(tagsDestinationDirectory, 0o755)
	if err != nil {
		log.Fatalln(err)
	}

	metadatas, err := image.NewMetadatasFromDirectory(sourceDirectory, destinationDirectory, config.BaseDirectory)
	if err != nil {
		log.Fatalln(err)
	}

	sort.Sort(sort.Reverse(metadatas))

	for _, metadata := range metadatas {
		err := internal.CopyFile(metadata.Source.Path, metadata.Destination.Path.Value)
		if err != nil {
			log.Fatalln(err)
		}
		err = image.CreateThumbnail(metadata.Source.Path, metadata.Destination.ThumbnailPath.Value, 256, 256)
		if err != nil {
			log.Fatalln(err)
		}
	}

	// ^ fonction execute
	err = web.Execute(destinationDirectory, metadatas, config)
	if err != nil {
		log.Fatalln(err)
	}
}
