package jar

import (
	"archive/zip"
	"errors"
	"fmt"
)

// may be custom type with Close()

func New(path string) (err error) {
	zipReader, err := zip.OpenReader(path)
	if err != nil {
		return err
	}

	// how to handle error on close
	defer zipReader.Close()
	defer func() {
		err = errors.Join(err, zipReader.Close())
	}()

	for _, f := range zipReader.File {
		fmt.Println(f.Name)
	}

	return err
}
