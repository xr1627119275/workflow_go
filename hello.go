package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func runCommand(command string) io.Writer {
	cmdName := "cmd.exe"
	cmdArgs := []string{"/c", command}
	fmt.Println("Running command: " + command)
	cmd := exec.Command(cmdName, cmdArgs...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	cmd.Run()
	return cmd.Stdout

}

func main() {
	fmt.Println(os.Args)
	fmt.Println(os.Getenv("development"))

	fmt.Println(runCommand("dir C:"))
	fmt.Println(runCommand("dir D:"))
	fmt.Println(runCommand("dir E:"))
	fmt.Println("结束")
}
