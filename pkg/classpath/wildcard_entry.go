package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

func newWildCardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1]
	var compositeEntry []Entry
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			compositeEntry = append(compositeEntry, newZipEntry(path))
		}
		return nil
	}
	filepath.Walk(baseDir, walkFn)
	return compositeEntry
}
