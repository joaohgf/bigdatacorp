package enum

type (
	FileName string
	FileType string
)

const (
	PlayerFileName FileName = "players"
	ClubFileName   FileName = "clubs"
	UploadFileName FileName = "upload"

	CSVType   FileType = "csv"
	JSONLType FileType = "jsonl"
)
