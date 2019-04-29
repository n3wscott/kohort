package kohort

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type KoOptions struct {
	Manifests string
	Verb      string
}

func Run(kopt *KoOptions) (string, error) {
	// want to do a ko resolve -f deployment.yaml

	fmt.Print(kopt.Manifests)

	cmd := Cmd(fmt.Sprintf("ko %s -f -", kopt.Verb))

	cmd.Stdin = strings.NewReader(kopt.Manifests)
	cmdOut, err := cmd.Output()
	return string(cmdOut), err

}

func Cmd(cmdLine string, opts ...string) *exec.Cmd {
	if len(opts) == 0 {
		fmt.Fprintln(os.Stderr, cmdLine)
	}
	cmdSplit := strings.Split(cmdLine, " ")
	cmd := cmdSplit[0]
	args := cmdSplit[1:]

	return exec.Command(cmd, args...)
}
