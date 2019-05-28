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
		input, _ := reader.ReadString('\n')
		// convert CRLF to LF
		input = strings.Replace(input, "\n", "", -1)

		if strings.Compare("hi", input) == 0 {
			fmt.Println("hello, Yourself")
		}

		docker := "docker"
		args := input
		cmd := exec.Command(docker, args)
		cmdReader, err := cmd.StdoutPipe()

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
			os.Exit(1)
		}
		scanner := bufio.NewScanner(cmdReader)
		go func() {
			for scanner.Scan() {
				fmt.Printf("docker %s out | %s\n", input, scanner.Text())
			}
		}()
		err = cmd.Start()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error starting for Cmd", err)
			os.Exit(1)
		}
		err = cmd.Wait()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
			os.Exit(1)
		}

	}
}
