package git

import (
	"log"
	"os"
	"os/exec"
)

func GitCommit(message string, all bool) {
	var cmd *exec.Cmd

	if all {
		cmd = exec.Command("git", "commit", "-a", "-m", string(message), "-e")
	} else {
		cmd = exec.Command("git", "commit", "-m", string(message), "-e")
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil {
		log.Fatalln(err)
	}
}
