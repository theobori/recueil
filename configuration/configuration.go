package configuration

import "github.com/theobori/recueil/internal"

type Css struct {
	ColorBg    string `yaml:"colorBg"`
	ColorFg    string `yaml:"colorFg"`
	ColorLink  string `yaml:"colorLink"`
	ColorHover string `yaml:"colorHover"`
}

type Configuration struct {
	// This value define the destination directory within the base destination directory
	BaseDirectory string `yaml:"baseDirectory"`
	TagsDirectory string `yaml:"tagsDirectory"`
	Title         string `yaml:"title"`
	Css           Css    `yaml:"css"`
}

func NewConfigurationDefault() *Configuration {
	return &Configuration{
		BaseDirectory: "/recueil",
		TagsDirectory: "/tags",
		Title:         "recueil",
		Css: Css{
			ColorBg:    "#fffef7",
			ColorFg:    "#394655",
			ColorLink:  "brown",
			ColorHover: "brown",
		},
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
