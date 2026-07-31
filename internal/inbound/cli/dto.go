package cli

type (
	CommandInput struct {
		FilePath     string
		ClubOutput   string
		PlayerOutput string
	}
)

const (
	ClubOutputFlag   = "clubs-output"
	PlayerOutputFlag = "players-output"
)

func NewCommandInput() *CommandInput {
	return new(CommandInput)
}
