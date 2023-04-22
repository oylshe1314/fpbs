package writer

import "fpbs/gen/parser"

type Writer interface {
	Write(outDir string, packages map[string]*parser.Package) error
}
