package internal

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

func YamlUnmarshalFromPath(yamlPath string, out any) error {
	f, err := os.Open(yamlPath)
	if err == nil {
		defer f.Close()

		bytes, err := io.ReadAll(f)
		if err != nil {
			return err
		}

		err = yaml.Unmarshal(bytes, out)
		if err != nil {
			return err
		}
	}

	return nil
}
