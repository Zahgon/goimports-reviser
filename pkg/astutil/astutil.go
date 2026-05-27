package astutil

import (
	"go/ast"
)

const (
	buildTagPrefix           = "//go:build"
	deprecatedBuildTagPrefix = "//+build"
)

// PackageImports is map of imports with their package names
type PackageImports map[string]string

// UsesImport is for analyze if the import dependency is in use
func UsesImport(f *ast.File, packageImports PackageImports, importPath string) bool {
	_ = "STUB: not implemented"
	return false
}

// LoadPackageDependencies will return all package's imports with it names:
//
//	key - package(ex.: github/pkg/errors), value - name(ex.: errors)
func LoadPackageDependencies(dir, buildTag string) (PackageImports, error) {
	_ = "STUB: not implemented"
	return *new(PackageImports), nil
}

// ParseBuildTag parse `//+build ...` or `//go:build ` on a first line of *ast.File
func ParseBuildTag(f *ast.File) string { _ = "STUB: not implemented"; return "" }

type visitFn func(node ast.Node)

func (f visitFn) Visit(node ast.Node) ast.Visitor {
	_ = "STUB: not implemented"
	return *new(ast.Visitor)
}
