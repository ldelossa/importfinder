package importfinder

import (
	"encoding/json"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"sync"
)

const (
	goExt = ".go"
)

// Finder is the struct which will implement our importfinder application
type Finder struct {
	// files found with .go extensions from root of file tree
	FilePaths []string
	// mutex for protecting concurrent access to Imports
	mu sync.Mutex
	// list of discovered imports per file
	Imports map[string][]string
}

// NewFinder is a constructor method for a Finder
func NewFinder() *Finder {
	f := &Finder{
		FilePaths: []string{},
		Imports:   map[string][]string{},
	}

	return f
}

// findFiles is a private helper method for walking the file system tree at our current
// working directory and locating all .go files.
func (f *Finder) findFiles() error {
	// walk current directory. access to finder is provided via closure
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("failed to stat file %s during walk: %v", path, err)
			return err
		}

		// all go files to our finder
		if filepath.Ext(path) == ".go" {
			// log.Printf("found go file %s and adding to inventory", path)
			f.FilePaths = append(f.FilePaths, path)
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// FindImports concurrently parses found go files into AST trees, extracts the found import
// paths and adds them to f.Imports map
func (f *Finder) FindImports() error {
	// recursively enumerate go files
	err := f.findFiles()
	if err != nil {
		return err
	}

	// create wait group
	var wg sync.WaitGroup

	for _, path := range f.FilePaths {
		// add to wait group
		wg.Add(1)

		// launch
		go func(path string) {
			defer wg.Done()

			// parse AST for found file
			file, err := parser.ParseFile(token.NewFileSet(), path, nil, parser.ImportsOnly)
			if err != nil {
				log.Printf("failed to parse AST for go file %s: %v", path, err)
				return
			}

			// extract imports in file
			for _, imp := range file.Imports {
				f.mu.Lock()
				f.Imports[path] = append(f.Imports[path], imp.Path.Value)
				f.mu.Unlock()
			}
		}(path)
	}

	wg.Wait()

	return nil
}

// String json serializes the Imports map and returns it as a string
func (f *Finder) String() string {
	b, _ := json.Marshal(f.Imports)
	return string(b)
}
