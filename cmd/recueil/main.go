package main

import (
	"log"
	"os"
	"path/filepath"
	"sort"
	"sync"

	"bufio"
	"fmt"
	"strings"

	"github.com/alecthomas/kingpin/v2"
	"github.com/theobori/recueil/configuration"
	"github.com/theobori/recueil/image"
	"github.com/theobori/recueil/image/source"
	"github.com/theobori/recueil/internal"
	"github.com/theobori/recueil/web"
	"gopkg.in/yaml.v3"
)

var (
	app = kingpin.New("recueil", "My opinionated static gallery generator.")

	build                     = app.Command("build", "Build the static gallery web site.")
	buildSourceDirectory      = build.Flag("source-directory", "The source directory which should be your 'recueil' base directory.").Default("./recueil").String()
	buildDestinationDirectory = build.Flag("destination-directory", "The destination directory where files are installed.").Default("./recueil-static").String()
	buildConfigurationPath    = build.Flag("configuration", "The configuration YAML file path.").Default("").String()
	buildVerbose              = build.Flag("verbose", "Enable verbose.").Bool()

	add          = app.Command("add", "Add a new image.")
	addImagePath = add.Arg("image", "Image file path.").Required().String()
	addDirectory = add.Arg("directory", "Image directory.").Required().String()

	delete          = app.Command("delete", "Delete image.")
	deleteImagePath = delete.Arg("image", "Image file path.").Required().String()
)

func Delete() error {
	log.Printf("Deleting the file '%s'", *deleteImagePath)
	err := os.Remove(*deleteImagePath)
	if err != nil {
		return err
	}

	metadataDestinationPath := *deleteImagePath + ".yaml"
	if _, err := os.Stat(metadataDestinationPath); err == nil {
		log.Printf("Deleting the metadata file '%s'", *&metadataDestinationPath)
		err := os.Remove(metadataDestinationPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func Add() error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter title: ")
	title, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	fmt.Print("Enter tags splitted by ',': ")
	tagsExpr, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	tags := strings.Split(tagsExpr, ",")
	filteredTags := []string{}
	for _, tag := range tags {
		tag := strings.TrimSpace(tag)
		if tag == "" {
			continue
		}

		filteredTags = append(filteredTags, tag)
	}

	fmt.Print("Enter the description: ")
	description, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	fileInfo, err := os.Stat(*addImagePath)
	if err != nil {
		return err
	}

	// with YAML fields only
	metadata := source.Metadata{
		Title:       title,
		Description: description,
		Tags:        filteredTags,
		Date:        fileInfo.ModTime(),
	}

	imageDestinationPath := filepath.Join(*addDirectory, filepath.Base(*addImagePath))
	log.Printf("Copying the file from '%s' to '%s'", *addImagePath, imageDestinationPath)
	err = internal.CopyFile(*addImagePath, imageDestinationPath)
	if err != nil {
		return err
	}

	metadataDestinationPath := imageDestinationPath + ".yaml"
	log.Printf("Creating the file '%s'", metadataDestinationPath)
	f, err := os.Create(metadataDestinationPath)
	if err != nil {
		return err
	}
	defer f.Close()

	yamlEncoder := yaml.NewEncoder(f)
	err = yamlEncoder.Encode(&metadata)
	if err != nil {
		return err
	}
	defer yamlEncoder.Close()

	return nil
}

func Build() error {
	var config *configuration.Configuration
	var err error
	if *buildConfigurationPath != "" {
		config, err = configuration.NewConfigurationFromPath(*buildConfigurationPath)
	} else {
		config = configuration.NewConfigurationDefault()
	}

	log.Printf("Creating the directory '%s'.\n", *buildDestinationDirectory)
	err = os.MkdirAll(*buildDestinationDirectory, 0o755)
	if err != nil {
		return err
	}

	baseDestinationDirectory := filepath.Join(*buildDestinationDirectory, config.BaseDirectory)
	log.Printf("Creating the directory '%s'.\n", baseDestinationDirectory)
	err = os.MkdirAll(baseDestinationDirectory, 0o755)
	if err != nil {
		return err
	}

	tagsDestinationDirectory := filepath.Join(*buildDestinationDirectory, config.TagsDirectory)
	log.Printf("Creating the directory '%s'.\n", tagsDestinationDirectory)
	err = os.MkdirAll(tagsDestinationDirectory, 0o755)
	if err != nil {
		return err
	}

	log.Printf("Fetching the metadatas.")
	metadatas, err := image.NewMetadatasFromDirectory(*buildSourceDirectory, *buildDestinationDirectory, config.BaseDirectory, *buildVerbose)
	if err != nil {
		return err
	}

	sort.Sort(sort.Reverse(metadatas))

	log.Printf("Copying images and creating the associated thumbnails.\n")
	var wg sync.WaitGroup

	for _, metadata := range metadatas {
		wg.Go(func() {
			var err error
			if config.Images.AutoResize {
				if *buildVerbose {
					log.Printf("Resizing from '%s' to '%s'.\n", metadata.Source.ImagePath.Value, metadata.Destination.ImagePath.Value)
				}
				err = image.ResizeIfGtMaxPixelSize(metadata.Source.ImagePath.Value, metadata.Destination.ImagePath.Value, config.Images.MaxPixelSize)
			} else {
				if *buildVerbose {
					log.Printf("Copying from '%s' to '%s'.\n", metadata.Source.ImagePath.Value, metadata.Destination.ImagePath.Value)
				}
				err = internal.CopyFile(metadata.Source.ImagePath.Value, metadata.Destination.ImagePath.Value)
			}
			if err != nil {
				log.Fatalln(err)
			}

			if *buildVerbose {
				log.Printf("Creating the thumbnail at '%s'.\n", metadata.Destination.ThumbnailPath.Value)
			}
			err = image.CreateThumbnail(metadata.Source.ImagePath.Value, metadata.Destination.ThumbnailPath.Value, config.Images.ThumbnailPixelSize)
			if err != nil {
				log.Fatalln(err)
			}
		})
	}

	wg.Wait()

	log.Printf("Evaluating then creating the web files.\n")
	webExecutor := web.NewExecutor(*buildDestinationDirectory, metadatas, config, *buildVerbose)
	err = webExecutor.Execute()
	if err != nil {
		return err
	}

	log.Printf("Copying the files from '%s' into '%s'.\n", config.PublicSourceDirectory, *buildDestinationDirectory)
	err = internal.CopyDir(*buildDestinationDirectory, config.PublicSourceDirectory)
	if err != nil {
		return err
	}

	log.Printf("The static website is ready at '%s'.\n", *buildDestinationDirectory)

	return nil
}

func runCommand() error {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case build.FullCommand():
		return Build()
	case add.FullCommand():
		return Add()
	case delete.FullCommand():
		return Delete()
	default:
		return fmt.Errorf("invalid command.")
	}
}

func main() {
	err := runCommand()
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}
