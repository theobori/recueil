package web

import (
	"html/template"
	"io/fs"
	"log"
	"os"
)

func ExecuteHelper(sourceFS fs.FS, destinationPath string, data any, verbose bool, patterns ...string) error {
	t, err := template.ParseFS(sourceFS, patterns...)
	if err != nil {
		return err
	}

	f, err := os.Create(destinationPath)
	if err != nil {
		return err
	}

	if verbose {
		log.Printf("Evaluating template and writing to '%s'\n", destinationPath)
	}
	err = t.Execute(f, data)
	if err != nil {
		return err
	}

	return nil
}
