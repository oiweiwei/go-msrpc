package midl

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

var (
	fileStore   = make(map[string]*File)
	fileStoreMu = new(sync.Mutex)
)

// Load function loads the file content.
func (f *File) Load() (*File, error) {

	fmt.Fprintln(os.Stderr, "loading...", f.Path)

	var path, base string
	var err error

	if filepath.IsAbs(f.Path) {
		path, base = f.Path, filepath.Dir(f.Path)
	} else if path, base, err = FindFile(f.Path); err != nil {
		if err != nil {
			return f, err
		}
	}

	fmt.Println("[BASE]", base)

	if rf, ok := FileLoad(path); ok {
		return rf, nil
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return f, err
	}
	rf, err := Parse(string(b))
	if err != nil {
		return f, err
	}

	rf.FullPath, rf.GoPkg, rf.Path = path, filepath.Join(base, f.GoPkg), f.Path
	fmt.Println("[FILE]", "full_path", rf.FullPath)
	fmt.Println("[FILE]", "go_pkg", rf.GoPkg)
	fmt.Println("[FILE]", "path", f.Path)

	*f = *rf

	FileStore(path, f)

	return f, nil
}

// FileLoad ...
func FileLoad(p string) (*File, bool) {
	fileStoreMu.Lock()
	defer fileStoreMu.Unlock()
	file, ok := fileStore[p]
	return file, ok
}

func Files() []*File {
	fileStoreMu.Lock()
	defer fileStoreMu.Unlock()

	keys := make([]string, 0, len(fileStore))
	for k := range fileStore {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	ret := make([]*File, len(fileStore))
	for i := range keys {
		ret[i] = fileStore[keys[i]]
	}

	return ret
}

func FileStore(p string, f *File) {
	fileStoreMu.Lock()
	defer fileStoreMu.Unlock()
	fileStore[p] = f
}

func FindPath() []string {
	return strings.Split(os.Getenv("MSIDLPATH"), ":")
}

func LookupType(n string) *Type {
	for _, f := range Files() {
		t, ok := f.exportSyms[n]
		if !ok || t.Type == nil {
			continue
		}
		return t.Type
	}
	return nil
}

func FindFile(f string) (string, string, error) {
	for _, dir := range FindPath() {

		reldirs := []string{""}

		if idx := strings.LastIndex(f, "/"); idx > 0 {
			reldirs[0], f = f[:idx], f[idx+1:]
		}

		ps, err := filepath.Glob(filepath.Join(dir, "*"))
		if err != nil {
			return "", "", err
		}

		for _, p := range ps {
			if info, _ := os.Stat(p); info != nil && info.IsDir() {
				reldirs = append(reldirs, filepath.Base(p))
			}
		}
		for _, reldir := range reldirs {
			dir := dir
			if reldir != "" {
				dir = filepath.Join(dir, reldir)
			}
			for file := range map[string]struct{}{
				strings.TrimPrefix(f, "ms-"): {},
				strings.TrimPrefix(f, "mc-"): {},
				f:                            {},
			} {
				path := filepath.Join(dir, file)
				fmt.Fprintln(os.Stderr, "checking", path, "...")
				if _, err := os.Stat(path); err != nil {
					continue
				}
				return path, reldir, nil
			}
		}
	}
	return f, "", fmt.Errorf("cannot find file %q: path: %s", f, strings.Join(FindPath(), ":"))
}

func NewFile(p, pkg string) *File {

	if pkg == "" {
		pkg = /* "go-" + */ strings.TrimPrefix(strings.TrimPrefix(strings.TrimSuffix(filepath.Base(p), filepath.Ext(p)), "ms-"), "mc-")
	}

	pkg = strings.ReplaceAll(pkg, "-", "_")

	return &File{Path: p, GoPkg: pkg}
}

// File structure represents the parsed file contents.
type File struct {
	FullPath string `json:"full_path"`
	// Path is a path to file.
	Path string `json:"path"`
	// GoPkg ...
	GoPkg string `json:"go_pkg"`
	// The list of imports.
	Imports []string `json:"imports,omitempty"`
	// Export is a map of exported symbols on file level.
	Export map[string]*Export `json:"exports,omitempty"`
	// Interfaces is a list of interfaces.
	Interfaces []*Interface `json:"interfaces,omitempty"`
	// ComClasses is a list of COM classes.
	ComClasses []*ComClass `json:"com_classes,omitempty"`
	// DispatchInterfaces ...
	DispatchInterfaces []*DispatchInterface `json:"dispatch_interfaces,omitempty"`
	// Libraries ...
	Libraries []*Library `json:"libraries,omitempty"`
	// exportSyms ...
	exportSyms map[string]*Export `json:"-"`
}

// Exports function returns the exported symbols as a sorted array.
func (f *File) Exports() []*Export {

	exportSyms := make([]*Export, 0, len(f.Export))
	for _, sym := range f.Export {
		exportSyms = append(exportSyms, sym)
	}

	sort.Slice(exportSyms, func(i, j int) bool {
		return exportSyms[i].Position < exportSyms[j].Position
	})

	return exportSyms
}

// LookupType function lookups the type inside the file.
func (f *File) LookupType(n string) (*Type, bool) {

	if typ, ok := f.Export[n]; ok && typ.Type != nil {
		return typ.Type, true
	}

	for i := range f.Interfaces {
		if typ, ok := f.Interfaces[i].Body.Export[n]; ok && typ.Type != nil {
			return typ.Type, true
		}
	}

	return nil, false
}

func (f *File) LookupAlias(n string) []string {

	if typ, ok := f.Export[n]; ok {
		return typ.Aliases
	}

	for i := range f.Interfaces {
		if typ, ok := f.Interfaces[i].Body.Export[n]; ok {
			return typ.Aliases
		}
	}
	return nil
}

func (f *File) IsLocal(tn string) bool {
	if _, ok := f.exportSyms[tn]; ok {
		return true
	}
	return false
}

// GoPackage function returns the go package name for the type.
func (f *File) GoPackage(tn string) (string, bool) {

	if _, ok := f.exportSyms[tn]; ok {
		return "", true
	}

	for _, file := range Files() {
		if _, ok := file.exportSyms[tn]; ok {
			return file.GoPkg, true
		}
	}

	return "", false
}

func (f *File) IsDCOM() bool {
	for _, iff := range f.Interfaces {
		if iff.IsObject() {
			return true
		}
	}
	return false
}

func (f *File) PkgIs(pkg string) bool {
	return f.GoPkg == pkg
}
