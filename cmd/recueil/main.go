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
		verbose              bool
	)

	flag.StringVar(&sourceDirectory, "source-directory", "./recueil", "The source directory which should be your 'recueil' base directory.")
	flag.StringVar(&destinationDirectory, "destination-directory", "./recueil-static", "The destination directory where files are installed.")
	flag.StringVar(&configurationPath, "configuration", "", "The configuration YAML file path.")
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose.")
	flag.Parse()

	var config *configuration.Configuration
	var err error
	if configurationPath != "" {
		config, err = configuration.NewConfigurationFromPath(configurationPath)
	} else {
		config = configuration.NewConfigurationDefault()
	}

	log.Printf("Creating the directory '%s'.\n", destinationDirectory)
	err = os.MkdirAll(destinationDirectory, 0o755)
	if err != nil {
		log.Fatalln(err)
	}

	baseDestinationDirectory := filepath.Join(destinationDirectory, config.BaseDirectory)
	log.Printf("Creating the directory '%s'.\n", baseDestinationDirectory)
	err = os.MkdirAll(baseDestinationDirectory, 0o755)
	if err != nil {
		log.Fatalln(err)
	}

	tagsDestinationDirectory := filepath.Join(destinationDirectory, config.TagsDirectory)
	log.Printf("Creating the directory '%s'.\n", tagsDestinationDirectory)
	err = os.MkdirAll(tagsDestinationDirectory, 0o755)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Fetching the metadatas.")
	metadatas, err := image.NewMetadatasFromDirectory(sourceDirectory, destinationDirectory, config.BaseDirectory, verbose)
	if err != nil {
		log.Fatalln(err)
	}

	sort.Sort(sort.Reverse(metadatas))

	log.Printf("Copying images and creating the associated thumbnails.\n")
	for _, metadata := range metadatas {
		if verbose {
			log.Printf("Copying from '%s' to '%s'.\n", metadata.Source.ImagePath.Value, metadata.Destination.ImagePath.Value)
		}
		err := internal.CopyFile(metadata.Source.ImagePath.Value, metadata.Destination.ImagePath.Value)
		if err != nil {
			log.Fatalln(err)
		}

		if verbose {
			log.Printf("Creating the thumbnail at '%s'.\n", metadata.Destination.ThumbnailPath.Value)
		}
		err = image.CreateThumbnail(metadata.Source.ImagePath.Value, metadata.Destination.ThumbnailPath.Value, config.ThumbnailPixelSize)
		if err != nil {
			log.Fatalln(err)
		}
	}

	log.Printf("Copying the files from '%s' into '%s'.\n", config.PublicSourceDirectory, destinationDirectory)
	err = internal.CopyDir(destinationDirectory, config.PublicSourceDirectory)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Evaluating then creating the web files.\n")
	webExecutor := web.NewExecutor(destinationDirectory, metadatas, config, verbose)
	err = webExecutor.Execute()
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("The static website is ready at '%s'.\n", destinationDirectory)
}
