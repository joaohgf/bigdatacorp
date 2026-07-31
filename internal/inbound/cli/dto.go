package cli

type (
	// CommandInput contains the CLI input and output paths.
	CommandInput struct {
		FilePath     string
		ClubOutput   string
		PlayerOutput string
	}
)

const (
	// ClubOutputFlag configures the club CSV output path.
	ClubOutputFlag = "clubs-output"
	// PlayerOutputFlag configures the player CSV output path.
	PlayerOutputFlag = "players-output"
)

// NewCommandInput creates an empty CommandInput.
func NewCommandInput() *CommandInput {
	return new(CommandInput)
}
