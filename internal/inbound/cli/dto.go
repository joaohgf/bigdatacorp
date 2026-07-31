package cli

type (
	CommandInput struct {
		FilePath string
	}
	CommandOutput struct {
		ClubFilePath   string `json:"club_file_path"`
		PlayerFilePath string `json:"player_file_path"`
	}
)

func NewCommandInput() *CommandInput {
	return new(CommandInput)
}

func NewCommandOutput() *CommandOutput {
	return new(CommandOutput)
}
