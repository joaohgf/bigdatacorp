package enum

type (
	FileName string
	FileType string
)

const (
	PlayerFileName  FileName = "players"
	ClubFileName    FileName = "clubs"
	UploadFileName  FileName = "upload"
	ArchiveFileName FileName = "result"

	CSVType   FileType = "csv"
	JSONLType FileType = "jsonl"
	ZIPType   FileType = "zip"
)
