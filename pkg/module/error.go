package module

// UndefinedModuleError will appear on absent go.mod
type UndefinedModuleError struct{}

func (e *UndefinedModuleError) Error() string { _ = "STUB: not implemented"; return "" }

// PathIsNotSetError will appear if any directory or file is not set for searching go.mod
type PathIsNotSetError struct{}

func (e *PathIsNotSetError) Error() string { _ = "STUB: not implemented"; return "" }
