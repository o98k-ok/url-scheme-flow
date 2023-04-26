package command

type Commander interface {
	GenURI() string
	IncCount()
	GetCommand() *Command
}

type Command struct {
	Name    string
	Desc    string
	Command string
	Count   uint64
	Direct  bool
	App     string
	Format  string
}

func (c *Command) IncCount() {
	c.Count++
}

func (c *Command) GetCommand() *Command {
	return c
}

var (
	DEFAULT_VSCODE_FORMAT   = "%s://ionutvmi.vscode-commands-executor/runCommands?data=[{\"id\": \"%s\"}]"
	DEFAULT_OBSIDIAN_FORMAT = "%s://advanced-uri?vault=%s&commandid=%s"

	VSCODE   = "vscode"
	OBSIDIAN = "obsidian"
)
