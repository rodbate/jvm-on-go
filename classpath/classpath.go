package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath()
	cp.parseUserClasspath(cpOption)
	return cp
}

func (cp *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if bytes, entry, err := cp.bootClasspath.readClass(className); err == nil {
		return bytes, entry, err
	}
	if bytes, entry, err := cp.extClasspath.readClass(className); err == nil {
		return bytes, entry, err
	}
	return cp.userClasspath.readClass(className)
}

func (cp *Classpath) String() string {
	return cp.userClasspath.String()
}

func (cp *Classpath) parseBootAndExtClasspath() {
	jreDir := getJreDir()
	cp.bootClasspath = newWildCardEntry(filepath.Join(jreDir, "lib", "*"))
	cp.extClasspath = newWildCardEntry(filepath.Join(jreDir, "lib", "ext", "*"))
}

func (cp *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	cp.userClasspath = newCompositeEntry(cpOption)
}

func getJreDir() string {
	home := os.Getenv("JAVA_HOME")
	if home == "" {
		panic("please set JAVA_HOME env variable")
	}
	return filepath.Join(home, "jre")
}
