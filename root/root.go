package root

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/zeroibot/pack/dict"
	"github.com/zeroibot/pack/fail"
	"github.com/zeroibot/pack/lang"
	"github.com/zeroibot/pack/str"
	"golang.org/x/term"
)

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

// ParamsMap gets the key=value map from the parameters list
func ParamsMap(params []string, required []string, optional []string) (dict.Strings, error) {
	if required == nil {
		required = make([]string, 0)
	}
	if optional == nil {
		optional = make([]string, 0)
	}
	paramsMap := make(dict.Strings)
	for _, param := range params {
		parts := str.CleanSplitN(param, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key, value := parts[0], parts[1]
		if !slices.Contains(required, key) && !slices.Contains(optional, key) {
			continue
		}
		paramsMap[key] = value
	}
	for _, key := range required {
		if _, ok := paramsMap[key]; !ok {
			return nil, fail.MissingParams
		}
	}
	return paramsMap, nil
}

// Authenticate performs authentication for the Root account in the command-line app
func Authenticate(authFn func(string) error) error {
	fmt.Print("Enter password: ")
	pwd, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return err
	}
	fmt.Println()
	password := strings.TrimSpace(string(pwd))
	err = authFn(password)
	if err != nil {
		return fmt.Errorf("root authentication failed: %w", err)
	}
	return nil
}

// validateCommandParams checks if the command exists and the parameters meet the min parameter count
func validateCommandParams(cmdMap map[string]*CmdConfig, command string, params []string) error {
	if command == cmdExit || command == cmdHelp {
		return nil
	}
	if command == cmdSearch {
		return lang.Ternary(len(params) < 1, errInvalidParamCount, nil)
	}
	cfg, ok := cmdMap[command]
	if !ok {
		return errInvalidCommand
	}
	if len(params) < cfg.MinParams {
		return errInvalidParamCount
	}
	return nil
}

// getCommandParams gets the command and parameters from the line
func getCommandParams(cmdMap map[string]*CmdConfig, line string) (string, []string) {
	if strings.TrimSpace(line) == "" {
		fmt.Println(getHelp)
		return "", nil
	}
	args := str.SpaceSplit(line)
	command, params := strings.ToLower(args[0]), args[1:]
	err := validateCommandParams(cmdMap, command, params)
	if err != nil {
		fmt.Println("Error:", err)
		if errors.Is(err, errInvalidCommand) {
			fmt.Println(getHelp)
		} else if errors.Is(err, errInvalidParamCount) {
			displayHelp(cmdMap, command)
		}
		return "", nil
	}
	return command, params
}

// displayHelp displays the help list
func displayHelp(cmdMap map[string]*CmdConfig, command string) {

}
