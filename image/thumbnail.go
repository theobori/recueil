package image

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"golang.org/x/image/draw"
)

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func CreateThumbnail(imagePath string, thumbnailPath string, thumbnailSize int) error {
	if thumbnailSize <= 0 {
		return fmt.Errorf("The width must be greater than zero.")
	}

	f, err := os.Open(imagePath)
	if err != nil {
		return err
	}
	defer f.Close()

	imgSource, _, err := image.Decode(f)
	if err != nil {
		return err
	}

	rectSource := imgSource.Bounds()
	rectSize := rectSource.Size()

	size := min(rectSize.X, rectSize.Y)
	midX := rectSize.X / 2
	midY := rectSize.Y / 2
	r := image.Rect(midX-(size/2), midY-(size/2), midX+(size/2), midY+(size/2))

	imgSquare := imgSource.(SubImager).SubImage(r)
	fw, err := os.Create(thumbnailPath)
	if err != nil {
		return err
	}
	defer fw.Close()

	imgDestination := image.NewRGBA(image.Rect(0, 0, thumbnailSize, thumbnailSize))

	draw.NearestNeighbor.Scale(imgDestination, imgDestination.Rect, imgSquare, imgSquare.Bounds(), draw.Over, nil)

	imagePathExt := filepath.Ext(imagePath)
	switch imagePathExt {
	case ".jpeg":
		err = jpeg.Encode(fw, imgDestination, nil)
	case ".png":
		err = png.Encode(fw, imgDestination)
	default:
		return fmt.Errorf("The image format associated to the file ext '%s' is not handled.", imagePathExt)
	}
	if err != nil {
		return err
	}

	return nil
}
