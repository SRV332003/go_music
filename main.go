package main

import (
	"bufio"
	"dhvani/handler"
	"log"
	"os"
	"os/exec"
)

func main() {

	bufioReader := bufio.NewReaderSize(os.Stdin, 1000)
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	for {

		os.Stdout.WriteString("dhvani> ")
		str, err := bufioReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		handler.HandleInput(str)
	}

}
