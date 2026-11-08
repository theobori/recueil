package web

import (
	"embed"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"sort"

	"github.com/theobori/recueil/configuration"
	"github.com/theobori/recueil/image"
)

//go:embed templates
var templates embed.FS

type Executor struct {
	destinationDirectory string
	metadatas            image.Metadatas
	config               *configuration.Configuration
	verbose              bool
	// cached data
	tags []string
	// cached data
	byTags map[string][]*image.Metadata
}

func NewExecutor(destinationDirectory string, metadatas image.Metadatas, config *configuration.Configuration, verbose bool) *Executor {
	byTags := metadatas.ByTags()

	n := len(byTags)
	tags := make([]string, n)
	i := 0
	for tagName := range byTags {
		tags[i] = tagName
		i++
	}

	sort.Strings(tags)

	return &Executor{
		destinationDirectory: destinationDirectory,
		metadatas:            metadatas,
		config:               config,
		verbose:              verbose,
		tags:                 tags,
		byTags:               byTags,
	}
}

func (e *Executor) join(elem ...string) string {
	elem = append([]string{e.destinationDirectory}, elem...)

	return filepath.Join(elem...)
}

func (e *Executor) executeHelper(path string, data any, patterns ...string) error {
	t, err := template.ParseFS(templates, patterns...)
	if err != nil {
		return err
	}

	destinationPath := e.join(path)

	f, err := os.Create(destinationPath)
	if err != nil {
		return err
	}

	if e.verbose {
		log.Printf("Evaluating template and writing to '%s'\n", destinationPath)
	}
	err = t.Execute(f, data)
	if err != nil {
		return err
	}

	return nil
}

func (e *Executor) Execute() error {
	var err error

	err = e.executeGallery()
	if err != nil {
		return err
	}
	err = e.executeHome()
	if err != nil {
		return err
	}
	err = e.executeTags()
	if err != nil {
		return err
	}
	err = e.executeImages()
	if err != nil {
		return err
	}
	err = e.executeRss()
	if err != nil {
		return err
	}

	return nil
}
