package main

import (
	"fmt"
	"log"
	"os/exec"

	winPS "github.com/cybermind-nick/process-monitor/windwos"
)

type Process struct {
	Pid   string
	Pname string
}

// type ProcessMap map[Process.Pid]Process.Pname

func main() {
	// get the process information
	// if err := syscall.Exec("dir", flag.Args(), os.Environ()); err != nil {
	// 	log.Fatal(err)
	// }
	cmd := exec.Command("ps")
	resp, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	ps := winPS.New()

	psOut, psErr, err := ps.Execute("ls")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(psOut)
	fmt.Println(psErr)

	out := string(resp)
	fmt.Print(out)
}
