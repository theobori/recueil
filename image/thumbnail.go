package image

import (
	"fmt"
	"path/filepath"

	"github.com/h2non/bimg"
)

func CreateThumbnail(imagePath string, thumbnailPath string, thumbnailSize int) error {
	if thumbnailSize <= 0 {
		return fmt.Errorf("The width must be greater than zero.")
	}

	buffer, err := bimg.Read(imagePath)
	if err != nil {
		return err
	}

	imgSource := bimg.NewImage(buffer)
	size, err := imgSource.Size()
	if err != nil {
		return err
	}

	newSize := min(size.Width, size.Height)
	midX := size.Width / 2
	midY := size.Height / 2
	bufExtracted, err := imgSource.Extract(midY-(newSize/2), midX-(newSize/2), newSize, newSize)
	if err != nil {
		return err
	}

	bufResize, err := bimg.Resize(bufExtracted, bimg.Options{Width: thumbnailSize, Height: thumbnailSize})
	if err != nil {
		return err
	}

	imagePathExt := filepath.Ext(thumbnailPath)
	img := bimg.NewImage(bufResize)

	var buf []byte
	switch imagePathExt {
	case ".jpeg":
		buf, err = img.Convert(bimg.JPEG)
	case ".png":
		buf, err = img.Convert(bimg.PNG)
	default:
		return fmt.Errorf("The image format associated to the file ext '%s' is not handled.", imagePathExt)
	}
	if err != nil {
		return err
	}

	buf, err = bimg.NewImage(buf).Process(bimg.Options{
		Lossless:    true,
		Compression: 9,
	})
	if err != nil {
		return err
	}

	err = bimg.Write(thumbnailPath, buf)
	if err != nil {
		return err
	}

	return nil
}
