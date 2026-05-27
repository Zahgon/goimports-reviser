package reviser

// ImportsOrder represents the name of import order
type ImportsOrder string

const (
	// StdImportsOrder is std libs, e.g. fmt, errors, strings...
	StdImportsOrder      ImportsOrder = "std"
	NamedStdImportsOrder ImportsOrder = "namedStd"
	// CompanyImportsOrder is packages that belong to the same organization
	CompanyImportsOrder      ImportsOrder = "company"
	NamedCompanyImportsOrder ImportsOrder = "namedCompany"
	// ProjectImportsOrder is packages that are inside the current project
	ProjectImportsOrder      ImportsOrder = "project"
	NamedProjectImportsOrder ImportsOrder = "namedProject"
	// GeneralImportsOrder is packages that are outside. In other words it is general purpose libraries
	GeneralImportsOrder      ImportsOrder = "general"
	NamedGeneralImportsOrder ImportsOrder = "namedGeneral"
	// BlankedImportsOrder is separate group for "_" imports
	BlankedImportsOrder ImportsOrder = "blanked"
	// DottedImportsOrder is separate group for "." imports
	DottedImportsOrder ImportsOrder = "dotted"
)

const (
	defaultImportsOrder = "std,general,company,project"
)

// ImportsOrders alias to []ImportsOrder
type ImportsOrders []ImportsOrder

func (o ImportsOrders) sortImportsByOrder(importGroups *groupsImports) [][]string {
	_ = "STUB: not implemented"
	return nil
}

func (o ImportsOrders) hasBlankedImportOrder() bool { _ = "STUB: not implemented"; return false }

func (o ImportsOrders) hasDottedImportOrder() bool { _ = "STUB: not implemented"; return false }

func (o ImportsOrders) hasRequiredGroups() bool { _ = "STUB: not implemented"; return false }

// StringToImportsOrders will convert string, like "std,general,company,project" to ImportsOrder array type.
// Default value for empty string is "std,general,company,project"
func StringToImportsOrders(s string) (ImportsOrders, error) {
	_ = "STUB: not implemented"
	return *new(ImportsOrders), nil
}

func unique(s []string) []string { _ = "STUB: not implemented"; return nil }

func appendGroups(input ...[]string) []string { _ = "STUB: not implemented"; return nil }
