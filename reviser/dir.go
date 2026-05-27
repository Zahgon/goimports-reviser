package reviser

import (
	"errors"
	"io/fs"
	"path/filepath"
)

type walkCallbackFunc = func(hasChanged bool, path string, content []byte) error

const (
	goExtension   = ".go"
	recursivePath = "./..."
)

var (
	currentPaths = []string{".", "." + string(filepath.Separator)}
)

var (
	ErrPathIsNotDir = errors.New("path is not a directory")
)

// SourceDir to validate and fix import
type SourceDir struct {
	projectName     string
	dir             string
	isRecursive     bool
	excludePatterns []string // see filepath.Match
}

var defaultExcludes = []string{".git", ".idea", ".vscode"}

func NewSourceDir(projectName string, path string, isRecursive bool, excludes string) *SourceDir {
	_ = "STUB: not implemented"
	return nil
}

// get the absolute path

// if path is recursive, then we need to remove the "/..." suffix

// resolve the absolute path

// Check pattern is well-formed.

func (d *SourceDir) Fix(options ...SourceFileOption) error { _ = "STUB: not implemented"; return nil }

// Find collection of bad formatted paths
func (d *SourceDir) Find(options ...SourceFileOption) (*UnformattedCollection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *SourceDir) walk(callback walkCallbackFunc, options ...SourceFileOption) fs.WalkDirFunc {
	_ = "STUB: not implemented"
	return *new(fs.WalkDirFunc)
}

func (d *SourceDir) isExcluded(path string) bool { _ = "STUB: not implemented"; return false }

type UnformattedCollection struct {
	list []string
}

func newUnformattedCollection(list []string) *UnformattedCollection {
	_ = "STUB: not implemented"
	return nil
}

func (c *UnformattedCollection) List() []string { _ = "STUB: not implemented"; return nil }

func (c *UnformattedCollection) String() string { _ = "STUB: not implemented"; return "" }

func IsDir(path string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func isGoFile(path string) bool { _ = "STUB: not implemented"; return false }
