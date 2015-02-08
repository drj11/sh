// I'm really just messing around

package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd1 := exec.Command("ls", "-s")
	cmd2 := exec.Command("sort", "-n")

	cmd2.Stdin, _ = cmd1.StdoutPipe()
	cmd2.Stdout = os.Stdout

	cmd1.Start()
	cmd2.Start()
	cmd2.Wait()
}
