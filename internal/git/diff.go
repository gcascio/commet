package git

import (
	"log"
	"os/exec"
)

func GitDiff(all bool) string {
	var cmd *exec.Cmd

	if all {
		cmd = exec.Command("git", "diff")
	} else {
		cmd = exec.Command("git", "diff", "--cached")
	}

	diff, err := cmd.Output()

	if err != nil {
		log.Fatalln(err)
	}

	return string(diff)
}
