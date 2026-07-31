package jsonl

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/joaohgf/bigdatacorp/internal/enum"
	"github.com/joaohgf/bigdatacorp/internal/port"
)

type JSONL[I, O any] struct {
	mapper port.ToMany[I, O]
}

func NewJSONL[I, O any](mapper port.ToMany[I, O]) *JSONL[I, O] {
	target := new(JSONL[I, O])
	target.mapper = mapper
	return target
}

func (j *JSONL[I, O]) Decode(fileName string) ([]O, error) {
	file, err := os.Open(j.getFileName(fileName))
	if err != nil {
		return nil, fmt.Errorf("error opening file %s: %w", fileName, err)
	}
	decoder := json.NewDecoder(file)
	targets := []I{}
	for {
		var target I
		err := decoder.Decode(&target)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, fmt.Errorf("error decoding jsonl file %s: %w", fileName, err)
		}
		targets = append(targets, target)
	}
	mapped := j.mapper.ToMany(targets...)
	return mapped, nil
}

func (j *JSONL[I, O]) getFileName(fileName string) string {
	if !strings.HasSuffix(fileName, string(enum.JSONLType)) {
		fileName = fmt.Sprintf("%s.%s", fileName, enum.JSONLType)
	}
	return fileName
}
