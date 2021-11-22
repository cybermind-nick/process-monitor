// library to invoke powershell instead of cmd
package windows

import (
	"bytes"
	"log"
	"os/exec"
)

type Powershell struct {
	powershell string
}

func New() *Powershell {
	ps, _ := exec.LookPath("powershell.exe")
	return &Powershell{
		powershell: ps,
	}
}

func (p *Powershell) Execute(args ...string) (stdOut string, stdErr string, err error) {
	// arguments := append([]string, args...)
	cmd := exec.Command(p.powershell, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	stdOut, stdErr = stdout.String(), stderr.String()
	return
}
