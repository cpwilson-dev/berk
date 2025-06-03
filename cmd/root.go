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

func getCommandsMap() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the application",
			callback:    commandExit,
		},
		"add": {
			name:        "add",
			description: "Stages files",
			callback:    commandAdd,
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
		"init": {
			name:        "init",
			description: "Initialises a new Berk repo",
			callback:    commandInit,
		},
		"status": {
			name:        "status",
			description: "Gives status of staged & unstaged files",
			callback:    commandStatus,
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
