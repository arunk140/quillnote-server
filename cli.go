package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SetupCLIHandler() {
	r := bufio.NewReader(os.Stdin)
	go func() {
		for {
			line, err := r.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading from stdin:", err)
				return
			}
			done := ProcessCommand(line)
			if !done {
				fmt.Println("Invalid Command")
			} else {
				fmt.Println("Done")
			}
		}
	}()
}

func ProcessCommand(command string) bool {
	arguments := strings.Split(command, " ")
	if len(arguments) < 1 {
		return false
	}
	for agrId := 0; agrId < len(arguments); agrId++ {
		arguments[0] = strings.TrimSpace(arguments[0])
	}
	switch arguments[0] {
	case "user":
		switch arguments[1] {
		case "add":
			if len(arguments) != 4 {
				return false
			}
			AddUser(arguments[2], arguments[3])
			return true
		}
	case "exit":
		os.Exit(0)
		return true
	case "migrate":
		Migrate()
		return true
	}

	return false
}
