package helper

type Option func() (string, error)

func OSGetwdOption() (string, error) { _ = "STUB: not implemented"; return "", nil }

func DetermineProjectName(projectName, filePath string, option Option) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
