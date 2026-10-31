package validate

import (
	"os"

	"github.com/h2non/filetype"
)

func IsImageFromPath(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}

	head := make([]byte, 261)
	_, err = f.Read(head)
	if err != nil {
		return false, nil
	}

	isImage := filetype.IsImage(head)

	return isImage, nil
}
