package domain

import "github.com/joaohgf/bigdatacorp/internal/enum"

type File struct {
	Name string
	Type enum.FileName
}

func NewFile() *File {
	return new(File)
}
