package jar

import (
	"archive/zip"
	"errors"
	"fmt"
	"path/filepath"
	"strings"
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

func HasSuffix(path string, suffix string) bool {
	ext := filepath.Ext(path)
	if ext == "" {
		return strings.HasSuffix(path, suffix)
	}

	pathWoExt, _ := strings.CutSuffix(path, ext)
	return strings.HasSuffix(pathWoExt, suffix)
}
