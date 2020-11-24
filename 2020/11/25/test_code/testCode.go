package main

import (
	"fmt"
	"os/exec"
)

func main() {

	out, _ := exec.Command("pwd").Output()
	fmt.Println(string(out))

	// cmd := exec.Command("cp", "-r", dir1, dir2)
	// stdoutStderr, _ := cmd.CombinedOutput()
	// fmt.Println(string(stdoutStderr))
}
