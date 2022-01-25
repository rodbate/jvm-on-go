package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(path string) CompositeEntry {
	var compositeEntry []Entry
	for _, p := range strings.Split(path, pathListSeparator) {
		compositeEntry = append(compositeEntry, newEntry(p))
	}
	return compositeEntry
}

func (c CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range c {
		class, e, err := entry.readClass(className)
		if err == nil {
			return class, e, err
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (c CompositeEntry) String() string {
	strs := make([]string, len(c))
	for i, entry := range c {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}
