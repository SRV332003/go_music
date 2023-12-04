package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	// "time"
	"ytdown/downloader"
	
	"ytdown/handler"
)

func timepass() {

	// take input using bufio

	bufioReader := bufio.NewReaderSize(os.Stdin, 1000)
	str, err := bufioReader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	data := downloader.FetchSearch(str)

	for i := range data {
		fmt.Printf("%d. %s\t %20s\n", i+1, data[i][1][:30], data[i][2])
	}

	fmt.Print("Enter the number of the song -> ")

	var n int
	fmt.Scan(&n)

	if n > len(data) {
		log.Fatal("Invalid Input")
	}

	_, err = downloader.Getfile(data[n-1][0])
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	
	bufioReader := bufio.NewReaderSize(os.Stdin, 1000)
	

	for{
		str, err := bufioReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		go handler.HandleInput(str)
	}
	


	

}
