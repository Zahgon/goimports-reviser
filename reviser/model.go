package reviser

type groupsImports struct {
	*common
	blanked []string
	dotted  []string
}

type common struct {
	std          []string
	namedStd     []string
	general      []string
	namedGeneral []string
	company      []string
	namedCompany []string
	project      []string
	namedProject []string
}

func (c *common) defaultSorting() [][]string { _ = "STUB: not implemented"; return nil }
