package root

import "fmt"

const (
	allCommands string = "*"
	cmdHelp     string = "help"
	cmdExit     string = "exit"
	cmdSearch   string = "cmd"
	cmdGlue     string = "/"
)

var (
	errInvalidCommand    = fmt.Errorf("invalid command")
	errInvalidParamCount = fmt.Errorf("invalid param count")
	getHelp              = fmt.Sprintf("Type `%s` for list of commands, `%s <keyword>` to search for command", cmdHelp, cmdSearch)
	helpSkipCommands     = []string{cmdHelp, cmdExit, cmdSearch}
)

// CmdHandler takes in a list of string parameters
type CmdHandler = func([]string)

type CmdConfig struct {
	Command   string
	MinParams int
	Docs      string
	Handler   CmdHandler
}

// NewCommand creates a new CmdConfig
func NewCommand(command string, minParams int, docs string, handler CmdHandler) *CmdConfig {
	return new(CmdConfig{command, minParams, docs, handler})
}

// NewCommandMap creates a new map of command to CmdConfigs
func NewCommandMap(cfgs ...*CmdConfig) map[string]*CmdConfig {
	commands := make(map[string]*CmdConfig)
	for _, cfg := range cfgs {
		commands[cfg.Command] = cfg
	}
	return commands
}
