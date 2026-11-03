package image

import (
	"log"
	"os"
	"path/filepath"

	"github.com/theobori/recueil/image/destination"
	"github.com/theobori/recueil/image/source"
	"github.com/theobori/recueil/image/validate"
)

type Metadatas []*Metadata

func NewMetadatasFromDirectory(sourceDirectoryPath string, destinationDirectoryPath string, publicDestinationDirectoryPath string, verbose bool) (Metadatas, error) {
	entries, err := os.ReadDir(sourceDirectoryPath)
	if err != nil {
		return nil, err
	}

	metadatas := Metadatas{}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		path := filepath.Join(sourceDirectoryPath, entry.Name())
		isImage, err := validate.IsImageFromPath(path)
		if err != nil {
			if verbose {
				log.Println(err)
			}
			continue
		}
		if !isImage {
			continue
		}

		if verbose {
			log.Printf("Fetching source metadatas for the file '%s'.\n", path)
		}
		sourceMetadata, err := source.NewMetadataFromPath(path)
		if err != nil {
			return nil, err
		}

		destinationMetadata := destination.NewMetadata(destinationDirectoryPath, publicDestinationDirectoryPath, path)
		metadata := Metadata{
			Source:      sourceMetadata,
			Destination: destinationMetadata,
		}

		metadatas = append(metadatas, &metadata)
	}

	return metadatas, nil
}

func (m Metadatas) Len() int           { return len(m) }
func (m Metadatas) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m Metadatas) Less(i, j int) bool { return m[i].Source.Date.Before(m[j].Source.Date) }

func (m Metadatas) ByTags() map[string][]*Metadata {
	ans := map[string][]*Metadata{}

	for _, metadata := range m {
		for _, tag := range metadata.Tags() {
			ans[tag] = append(ans[tag], metadata)
		}
	}

	return ans
}
