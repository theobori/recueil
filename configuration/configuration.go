package configuration

import "github.com/theobori/recueil/internal"

type Configuration struct {
	// This value define the destination directory within the base destination directory
	BaseDirectory string `yaml:"baseDirectory"`
	TagsDirectory string `yaml:"tagsDirectory"`
	Title         string `yaml:"title"`
}

func NewConfigurationDefault() *Configuration {
	return &Configuration{
		BaseDirectory: "/recueil",
		TagsDirectory: "/tags",
		Title:         "recueil",
	}
}

func NewConfigurationFromPath(yamlPath string) (*Configuration, error) {
	configuration := NewConfigurationDefault()

	err := internal.YamlUnmarshalFromPath(yamlPath, configuration)
	if err != nil {
		return nil, err
	}

	return configuration, nil
}
