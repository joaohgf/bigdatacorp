package domain

import "github.com/joaohgf/bigdatacorp/internal/enum"

// File describes a generated output artifact.
type File struct {
	Name string
	Type enum.FileName
}

// NewFile creates an empty File descriptor.
func NewFile() *File {
	return new(File)
}
