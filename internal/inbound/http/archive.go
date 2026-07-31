package http

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/joaohgf/bigdatacorp/internal/usecase/domain"
)

func createArchive(path string, files []*domain.File) error {
	archive, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create archive: %w", err)
	}
	writer := zip.NewWriter(archive)
	for _, file := range files {
		if err := addToArchive(writer, file.Name); err != nil {
			writer.Close()
			archive.Close()
			return err
		}
	}
	if err := writer.Close(); err != nil {
		archive.Close()
		return fmt.Errorf("finalize archive: %w", err)
	}
	if err := archive.Close(); err != nil {
		return fmt.Errorf("close archive: %w", err)
	}
	return nil
}

func addToArchive(archive *zip.Writer, path string) error {
	source, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open generated file: %w", err)
	}
	defer source.Close()
	target, err := archive.Create(filepath.Base(path))
	if err != nil {
		return fmt.Errorf("create archive entry: %w", err)
	}
	if _, err := io.Copy(target, source); err != nil {
		return fmt.Errorf("write archive entry: %w", err)
	}
	return nil
}
