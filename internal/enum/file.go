package enum

type (
	// FileName identifies a default generated artifact name.
	FileName string
	// FileType identifies a supported generated artifact extension.
	FileType string
)

const (
	// PlayerFileName is the default player CSV base name.
	PlayerFileName FileName = "players"
	// ClubFileName is the default club CSV base name.
	ClubFileName FileName = "clubs"
	// UploadFileName is the default uploaded input base name.
	UploadFileName FileName = "upload"
	// ArchiveFileName is the default ZIP archive base name.
	ArchiveFileName FileName = "result"

	// CSVType identifies CSV artifacts.
	CSVType FileType = "csv"
	// JSONLType identifies JSONL artifacts.
	JSONLType FileType = "jsonl"
	// ZIPType identifies ZIP artifacts.
	ZIPType FileType = "zip"
)
