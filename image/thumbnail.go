package image

import (
	"fmt"

	"github.com/h2non/bimg"
)

func CreateThumbnail(imagePath string, thumbnailPath string, thumbnailSize int) error {
	if thumbnailSize <= 0 {
		return fmt.Errorf("The width must be greater than zero.")
	}

	buf, err := bimg.Read(imagePath)
	if err != nil {
		return err
	}

	imgSource := bimg.NewImage(buf)
	buf, err = imgSource.Thumbnail(thumbnailSize)
	if err != nil {
		return err
	}

	err = bimg.Write(thumbnailPath, buf)
	if err != nil {
		return err
	}

	return nil
}
