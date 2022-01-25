package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	abs, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{abs}
}

func (d *DirEntry) readClass(className string) ([]byte, Entry, error) {
	filename := filepath.Join(d.absDir, className)
	bytes, err := ioutil.ReadFile(filename)
	return bytes, d, err
}

func (d *DirEntry) String() string {
	return d.absDir
}
