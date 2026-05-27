package module

const goModFilename = "go.mod"

// Name reads module value from ./go.mod
func Name(goModRootPath string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// GoModRootPath in case of any directory or file of the project will return root dir of the project where go.mod file
// is exist
func GoModRootPath(path string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func DetermineProjectName(projectName, filePath string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
