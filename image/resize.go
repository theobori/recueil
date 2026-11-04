package image

import (
	"github.com/h2non/bimg"
)

func ResizeIfGtMaxPixelSize(sourcePath string, destinationPath string, maxPixelSize int) error {
	buf, err := bimg.Read(sourcePath)
	if err != nil {
		return err
	}

	size, err := bimg.Size(buf)
	if err != nil {
		return err
	}

	if size.Height > maxPixelSize || size.Width > maxPixelSize {
		var width int
		var height int
		if size.Width > size.Height {
			width = maxPixelSize
			height = (maxPixelSize * size.Height) / size.Width
		} else {
			height = maxPixelSize
			width = (maxPixelSize * size.Width) / size.Height
		}

		buf, err = bimg.Resize(buf, bimg.Options{Width: width, Height: height, Compression: 9, Lossless: true})
		if err != nil {
			return err
		}
	}

	err = bimg.Write(destinationPath, buf)
	if err != nil {
		return err
	}

	return nil
}
