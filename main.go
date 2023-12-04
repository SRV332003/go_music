package main

import (
	"bufio"
	"dhvani/handler"
	"log"
	"os"
)

func main() {

	// filemanager.Getfiles()
	// fmt.Println(filemanager.Search("tum ho"))

	bufioReader := bufio.NewReaderSize(os.Stdin, 1000)

	for {
		str, err := bufioReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		go handler.HandleInput(str)
	}

}
