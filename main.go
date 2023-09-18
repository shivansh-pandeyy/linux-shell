package main

import (
	"fmt"
	"bufio"
	"os"
	"os/exec"
	"strings"
	"errors"
)

func main() {
	for {
		fmt.Print("> ")

		startCommands := []string{"whoami", "hostname", "pwd"}

		var hostname string

		for i, command := range startCommands {
			outputBs, err := exec.Command(command).Output()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			output := string(outputBs)

			output = strings.TrimSuffix(output, "\n")

			if i == 1 {
				hostname = hostname + " @ " + output 
			} else if i == 2 {
				hostname = hostname + ":~" + output
			} else {
				hostname = hostname + "" + output
			}
		}
		
		fmt.Print(hostname + "$ ")

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(os.Stderr, err)
		}

		inputError := runCommand(input)
		
		if inputError != nil {
			fmt.Println(inputError)
		}
	}
}

func runCommand(inputString string)  error {
	inputString = strings.TrimSuffix(inputString, "\n")
	
	args := strings.Fields(inputString);

	switch args[0] {
	case "cd": 
		if(len(args) < 2) {
			return errors.New("path required")
		}

		return os.Chdir(args[1])

	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()

}
