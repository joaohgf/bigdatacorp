package http

import (
	"archive/zip"
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joaohgf/bigdatacorp/internal/enum"
	"github.com/joaohgf/bigdatacorp/internal/usecase/domain"
)

func TestOutputNameUsesDefaultsExtensionsAndBaseName(t *testing.T) {
	t.Parallel()
	tests := []struct {
		query    string
		fallback enum.FileName
		fileType enum.FileType
		want     string
	}{
		{"", enum.ClubFileName, enum.CSVType, "clubs.csv"},
		{"clubs-output=report", enum.ClubFileName, enum.CSVType, "report.csv"},
		{"clubs-output=report.CSV", enum.ClubFileName, enum.CSVType, "report.CSV"},
		{"clubs-output=../outside", enum.ClubFileName, enum.CSVType, "outside.csv"},
	}
	for _, test := range tests {
		test := test
		t.Run(test.want, func(t *testing.T) {
			t.Parallel()
			request := httptest.NewRequest(http.MethodGet, "/?"+test.query, nil)
			context, _ := gin.CreateTestContext(httptest.NewRecorder())
			context.Request = request
			if got := outputName(context, clubsOutputParam, test.fallback, test.fileType); got != test.want {
				t.Fatalf("outputName() = %q, want %q", got, test.want)
			}
		})
	}
}

func TestCreateArchiveIncludesOnlyFileBaseNames(t *testing.T) {
	t.Parallel()
	directory := t.TempDir()
	first := filepath.Join(directory, "clubs.csv")
	second := filepath.Join(directory, "players.csv")
	if err := os.WriteFile(first, []byte("clubs"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(second, []byte("players"), 0o600); err != nil {
		t.Fatal(err)
	}
	archivePath := filepath.Join(directory, "result.zip")
	if err := createArchive(archivePath, []*domain.File{{Name: first}, {Name: second}}); err != nil {
		t.Fatalf("createArchive() error = %v", err)
	}
	archive, err := zip.OpenReader(archivePath)
	if err != nil {
		t.Fatal(err)
	}
	defer archive.Close()
	if len(archive.File) != 2 || archive.File[0].Name != "clubs.csv" || archive.File[1].Name != "players.csv" {
		t.Fatalf("archive entries = %#v", archive.File)
	}
}

func TestSaveUpload(t *testing.T) {
	t.Parallel()
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	ignored, err := writer.CreateFormField("description")
	if err != nil {
		t.Fatal(err)
	}
	_, _ = io.WriteString(ignored, "ignored")
	file, err := writer.CreateFormFile("file", "input.jsonl")
	if err != nil {
		t.Fatal(err)
	}
	_, _ = io.WriteString(file, "{\"club_id\":\"A\"}\n")
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodPost, "/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	context.Request = request
	path := filepath.Join(t.TempDir(), "input.jsonl")

	status, err := new(Handler).saveUpload(context, path)
	if err != nil || status != 0 {
		t.Fatalf("saveUpload() = (%d, %v)", status, err)
	}
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if string(content) != "{\"club_id\":\"A\"}\n" {
		t.Fatalf("uploaded content = %q", content)
	}
}

func TestSaveUploadRequiresFileField(t *testing.T) {
	t.Parallel()
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	_ = writer.Close()
	request := httptest.NewRequest(http.MethodPost, "/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	context.Request = request

	status, err := new(Handler).saveUpload(context, filepath.Join(t.TempDir(), "input.jsonl"))
	if status != http.StatusBadRequest || err == nil {
		t.Fatalf("saveUpload() = (%d, %v), want bad request", status, err)
	}
}
