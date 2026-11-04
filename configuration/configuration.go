package configuration

import "github.com/theobori/recueil/internal"

type ConfigurationCss struct {
	ColorBg    string `yaml:"colorBg"`
	ColorFg    string `yaml:"colorFg"`
	ColorLink  string `yaml:"colorLink"`
	ColorHover string `yaml:"colorHover"`
}

type ConfigurationImages struct {
	AutoResize         bool `yaml:"autoResize"`
	MaxPixelSize       int  `yaml:"maxPixelSize"`
	ThumbnailPixelSize int  `yaml:"thumbnailPixelSize"`
}

type Configuration struct {
	// This value define the destination directory within the base destination directory
	BaseDirectory         string              `yaml:"baseDirectory"`
	TagsDirectory         string              `yaml:"tagsDirectory"`
	Title                 string              `yaml:"title"`
	Description           string              `yaml:"description"`
	SiteUrl               string              `yaml:"siteUrl"`
	Css                   ConfigurationCss    `yaml:"css"`
	PublicSourceDirectory string              `yaml:"publicSourceDirectory"`
	WebringUrl            string              `yaml:"webringUrl"`
	Images                ConfigurationImages `yaml:"images"`
}

func NewConfigurationDefault() *Configuration {
	return &Configuration{
		BaseDirectory: "/recueil",
		TagsDirectory: "/tags",
		Title:         "recueil",
		Css: ConfigurationCss{
			ColorBg:    "#fffef7",
			ColorFg:    "#394655",
			ColorLink:  "brown",
			ColorHover: "brown",
		},
		SiteUrl:               "http://localhost:8080",
		PublicSourceDirectory: "./public",
		Images: ConfigurationImages{
			AutoResize:         true,
			MaxPixelSize:       960,
			ThumbnailPixelSize: 256,
		},
		WebringUrl: "",
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
