package json

type _Status int

const (
	statusNone _Status = iota
	statusPackageDefine
	statusPackageDefining
	statusImportDefine
	statusImportDefining
	statusEnumDefine
	statusEnumDefining
	statusStructDefine
	statusStructDefining
	statusFieldsDefine
	statusFieldsDefining
)
