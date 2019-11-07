package cmd

import (
	"bytes"
	"log"
	"os/exec"
)

func ExecShell(s string) string {
	cmd := exec.Command("/bin/bash", "-c", s)
	var out bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &out
	//err := cmd.Start()
	err := cmd.Run()
	if err != nil {
		log.Print(err, cmd.Stderr)
		return ""
	}
	//fmt.Printf("id=%s", out.String())
	return out.String()
}
