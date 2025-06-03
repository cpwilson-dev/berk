package cmd

import "fmt"

func commandHelp() error {

	fmt.Println("Welcome to berk")
	fmt.Println("A rather foolish attempt to recreate Git that should be used by no one ever.")
	fmt.Println("Available commands:")

	availableCommands := getCommandsMap()
	for _, cmd := range availableCommands {
		fmt.Printf(" -%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Print()
	return nil
}
