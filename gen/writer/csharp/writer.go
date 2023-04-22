package csharp

import (
	"fpbs/gen/parser"
	"fpbs/gen/writer"
)

type csharpWriter struct {
}

func NewWriter() writer.Writer {
	return &csharpWriter{}
}

func (this *csharpWriter) Write(outDir string, packages map[string]*parser.Package) error {
	return nil
}
