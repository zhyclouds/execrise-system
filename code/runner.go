package main

import (
	"bytes"
	"io"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("go", "run", "code-user/main.go")
	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	stdinPipe, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	io.WriteString(stdinPipe, "1 2\n")
	if err := cmd.Run(); err != nil {
		log.Fatalln(err, stderr.String())
	}

	log.Println(out.String())

	println(out.String() == "3\n")

}
