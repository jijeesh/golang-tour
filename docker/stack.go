package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Docker Shell")
	fmt.Println("---------------------")
	for {
		fmt.Print("DockerShell-> ")
		// Read the keyboad input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		// Handle the execution of the input
		if err = execInput(input); err != nil {
			fmt.Println("-----------")
			fmt.Fprintln(os.Stderr, err)

		}

	}
}

func execInput(input string) error {
	// convert CRLF to LF
	// input = strings.Replace(input, "\n", "", -1)
	// Remove the newline character
	input = strings.TrimSuffix(input, "\n")

	if strings.Compare("hi", input) == 0 {
		fmt.Println("hello, Thank you for using Docker")
	}
	// Split the input seperate the command and the arguments
	args := strings.Split(input, " ")
	switch args[0] {
	case "exit":
		os.Exit(0)

	}
	docker := "docker"
	// args := input
	cmd := exec.Command(docker, args[0:]...)
	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and return the error.
	return cmd.Run()
	// cmdReader, err := cmd.StdoutPipe()
	//
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
	// 	os.Exit(1)
	// }
	// scanner := bufio.NewScanner(cmdReader)
	// go func() {
	// 	for scanner.Scan() {
	// 		fmt.Printf("docker %s  | %s\n", input, scanner.Text())
	// 	}
	// }()
	// err = cmd.Start()
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, "Error starting for Cmd", err)
	// 	os.Exit(1)
	// }
	// err = cmd.Wait()
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
	// 	os.Exit(1)
	// }
	// return err
}
