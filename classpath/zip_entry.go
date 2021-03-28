package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	abs, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{abs}
}

func (z *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	reader, err := zip.OpenReader(z.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer reader.Close()

	for _, f := range reader.File {
		if f.Name == className {
			open, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer open.Close()
			bytes, err := ioutil.ReadAll(open)
			if err != nil {
				return nil, nil, err
			}
			return bytes, z, err
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (z *ZipEntry) String() string {
	return z.absPath
}
