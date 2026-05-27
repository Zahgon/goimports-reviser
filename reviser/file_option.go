package reviser

// SourceFileOption is an int alias for options
type SourceFileOption func(f *SourceFile) error

// SourceFileOptions is a slice of executing options
type SourceFileOptions []SourceFileOption

// WithRemovingUnusedImports is an option to remove unused imports
func WithRemovingUnusedImports(f *SourceFile) error { _ = "STUB: not implemented"; return nil }

// WithUsingAliasForVersionSuffix is an option to set explicit package name in imports
func WithUsingAliasForVersionSuffix(f *SourceFile) error { _ = "STUB: not implemented"; return nil }

// WithCodeFormatting use to format the code
func WithCodeFormatting(f *SourceFile) error { _ = "STUB: not implemented"; return nil }

// WithCompanyPackagePrefixes option for 3d group(by default), like inter-org or company package prefixes
func WithCompanyPackagePrefixes(s string) SourceFileOption {
	_ = "STUB: not implemented"
	return *new(SourceFileOption)
}

// WithImportsOrder will sort by needed order. Default order is "std,general,company,project"
func WithImportsOrder(orders []ImportsOrder) SourceFileOption {
	_ = "STUB: not implemented"
	return *new(SourceFileOption)
}

// WithSkipGeneratedFile will skip formatting and imports sorting for auto-generated file which starts with
// comment on first line: `// Code generated`
func WithSkipGeneratedFile(f *SourceFile) error { _ = "STUB: not implemented"; return nil }

func WithSeparatedNamedImports(f *SourceFile) error { _ = "STUB: not implemented"; return nil }
