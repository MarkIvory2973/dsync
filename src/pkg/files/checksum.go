package files

import (
	"io"
	"os"

	"github.com/zeebo/blake3"
)

func Checksum(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	hasher := blake3.New()

	_, err = io.Copy(hasher, file)
	if err != nil {
		return nil, err
	}

	checksum := hasher.Sum(nil)

	return checksum, nil
}
