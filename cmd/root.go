package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cpwilson-dev/berk/utils"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commandsMap = make(map[string]*cliCommand)

func getCommandsMap() map[string]cliCommand {
	return map[string]cliCommand{
		"add": {
			name:        "add",
			description: "Stages files",
			callback:    commandAdd,
		},
		"cat-file": {
			name:        "cat-file",
			description: "TODO",
			callback:    commandHelp,
		},
		"check-ignore": {
			name:        "check-ignore",
			description: "TODO",
			callback:    commandHelp,
		},
		"checkout": {
			name:        "checkout",
			description: "TODO",
			callback:    commandHelp,
		},
		"commit": {
			name:        "commit",
			description: "Commits staged files",
			callback:    commandCommit,
		},
		"config": {
			name:        "config",
			description: "Configures Berk to your liking",
			callback:    commandConfig,
		},
		"diff": {
			name:        "diff",
			description: "TODO",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the application",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    commandHelp,
		},
		"init": {
			name:        "init",
			description: "Initialises a new Berk repo",
			callback:    commandInit,
		},
		"log": {
			name:        "log",
			description: "TODO",
			callback:    commandLog,
		},
		"ls-files": {
			name:        "ls-files",
			description: "TODO",
			callback:    commandHelp,
		},
		"ls-tree": {
			name:        "ls-tree",
			description: "TODO",
			callback:    commandHelp,
		},
		"rev-parse": {
			name:        "rev-parse",
			description: "TODO",
			callback:    commandHelp,
		},
		"rm": {
			name:        "rm",
			description: "TODO",
			callback:    commandHelp,
		},
		"show-ref": {
			name:        "show-ref",
			description: "TODO",
			callback:    commandHelp,
		},
		"status": {
			name:        "status",
			description: "Gives status of staged & unstaged files",
			callback:    commandStatus,
		},
		"tag": {
			name:        "tag",
			description: "TODO",
			callback:    commandHelp,
		},
		"version": {
			name:        "version",
			description: "TODO",
			callback:    commandHelp,
		},
	}
}

func StartRepl() {
	fmt.Println("Hello from Berk")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")

		scanner.Scan()
		text := scanner.Text()
		cleaned := utils.CleanInput(text)

		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		availableCommands := getCommandsMap()

		command, ok := availableCommands[commandName]

		if !ok {
			fmt.Println("Invalid command")
			continue
		}
		command.callback()

	}
}

func Execute(args []string) {
	commandsMap["help"] = &cliCommand{
		name:        "help",
		description: "help",
		callback:    commandHelp,
	}

	if len(args) < 2 {
		fmt.Println("Starting Repl")
		StartRepl()
	}

	commandName := args[1]
	availableCommands := getCommandsMap()

	command, ok := availableCommands[commandName]

	if !ok {
		fmt.Println("Invalid command")
		os.Exit(1)
	}
	command.callback()

}
