package web

import (
	"html/template"
	"io/fs"
	"os"
)

func ExecuteHelper(sourceFS fs.FS, destinationPath string, data any, patterns ...string) error {
	t, err := template.ParseFS(sourceFS, patterns...)
	if err != nil {
		return err
	}

	f, err := os.Create(destinationPath)
	if err != nil {
		return err
	}

	err = t.Execute(f, data)
	if err != nil {
		return err
	}

	return nil
}
