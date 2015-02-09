// I'm really just messing around

package main

import (
	"os"
	"os/exec"
)

func main() {
	pipeline := [][]string{[]string{"ls", "-s"}, []string{"sort", "-n"}}
	var cmds []*exec.Cmd
	for i, command := range pipeline {
		cmds = append(cmds, exec.Command(command[0], command[1:]...))
		if i > 0 {
			cmds[i].Stdin, _ = cmds[i-1].StdoutPipe()
		}
	}
	cmds[len(cmds)-1].Stdout = os.Stdout

	for _, cmd := range cmds {
		cmd.Start()
	}
	for _, cmd := range cmds {
		cmd.Wait()
	}
}
