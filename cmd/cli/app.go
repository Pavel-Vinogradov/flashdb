package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type App struct {
	commands map[string]CommandFunc
}

type CommandFunc func(args []string) error

func NewApp() *App {
	return &App{
		commands: make(map[string]CommandFunc),
	}
}

func (a *App) RegisterCommand(name string, fn CommandFunc) {
	a.commands[name] = fn
}

func (a *App) Run() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("flashdb> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		parts := strings.Fields(input)
		command := parts[0]
		args := parts[1:]

		if command == "exit" || command == "quit" {
			fmt.Println("Goodbye!")
			break
		}

		if fn, exists := a.commands[command]; exists {
			if err := fn(args); err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} else {
			fmt.Printf("Unknown command: %s\n", command)
			a.printHelp()
		}
	}
}

func (a *App) printHelp() {
	fmt.Println("Available commands:")
	for cmd := range a.commands {
		fmt.Printf("  %s\n", cmd)
	}
	fmt.Println("  exit, quit - Exit the CLI")
}
