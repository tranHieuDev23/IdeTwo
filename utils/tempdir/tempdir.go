package tempdir

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type TempDir struct {
	path string
}

func New(baseDir string) TempDir {
	path, err := ioutil.TempDir(baseDir, "execute-*")
	if err != nil {
		panic(err)
	}
	abspath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return TempDir{path: abspath}
}

func (dir *TempDir) GetPath() string {
	return dir.path
}

func (dir *TempDir) Close() {
	if err := os.RemoveAll(dir.path); err != nil {
		panic(err)
	}
}
