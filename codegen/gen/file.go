package gen

import (
	"bytes"
	"path/filepath"
	"strings"
)

type FileBuffer struct {
	Out         *bytes.Buffer
	Path        string
	PackageName string
	Imports     []Import
	FileName    string
	IsRoot      bool
}

func (f *FileBuffer) Reset() []byte {
	b := f.Out.Bytes()
	f.Out = bytes.NewBuffer(nil)
	return b
}

func NewFileBuffer(path string, pkg string, n ...string) *FileBuffer {
	f := &FileBuffer{
		Out:         bytes.NewBuffer(nil),
		Path:        strings.TrimSuffix(path, filepath.Ext(path)),
		PackageName: pkg[strings.Index(pkg, "/")+1:],
	}

	if len(n) > 0 {
		f.FileName = n[0]
	}

	f.Imports = make([]Import, len(DefaultImports))
	copy(f.Imports, DefaultImports)
	return f
}
