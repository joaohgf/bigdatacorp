package jsonl

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/joaohgf/bigdatacorp/internal/enum"
	"github.com/joaohgf/bigdatacorp/internal/port"
)

// JSONL streams newline-delimited JSON records and maps them to output values.
type JSONL[I, O any] struct {
	mapper port.To[I, O]
}

// NewJSONL creates a JSONL decoder backed by mapper.
func NewJSONL[I, O any](mapper port.To[I, O]) *JSONL[I, O] {
	target := new(JSONL[I, O])
	target.mapper = mapper
	return target
}

// Decode returns a lazy sequence over records in fileName.
func (j *JSONL[I, O]) Decode(fileName string) port.Sequence[O] {
	return func(yield func(O, error) bool) {
		file, err := os.Open(j.getFileName(fileName))
		if err != nil {
			var zero O
			yield(zero, fmt.Errorf("error opening file %s: %w", fileName, err))
			return
		}
		defer file.Close()
		reader := bufio.NewReader(file)
		for lineNumber := 1; ; lineNumber++ {
			line, readErr := reader.ReadBytes('\n')
			line = bytes.TrimSpace(line)
			if len(line) > 0 {
				var source I
				if err := json.Unmarshal(line, &source); err == nil {
					if !yield(j.mapper.To(source), nil) {
						return
					}
				}
			}
			if errors.Is(readErr, io.EOF) {
				return
			}
			if readErr != nil {
				var zero O
				yield(zero, fmt.Errorf("read JSONL file %q at line %d: %w", fileName, lineNumber, readErr))
				return
			}
		}
	}
}

// getFileName ensures the input path has a JSONL extension.
func (j *JSONL[I, O]) getFileName(fileName string) string {
	if !strings.HasSuffix(fileName, string(enum.JSONLType)) {
		fileName = fmt.Sprintf("%s.%s", fileName, enum.JSONLType)
	}
	return fileName
}
