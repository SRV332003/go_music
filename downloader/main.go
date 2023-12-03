package main

import (
	"fmt"
	"path"
	
	"log"
	"os"

	"github.com/Vernacular-ai/godub"
	"github.com/kkdai/youtube/v2"
)

func Getfile(url string) {
// https://www.youtube.com/results?search_query=
	client := youtube.Client{}

	video, err := client.GetVideo(url)
	if err != nil {
		panic(err)
	}

	format := video.Formats.FindByItag(140)
	// fmt.Println(format)

	stream, _, err := client.GetStream(video, format)
	if err != nil {
		panic(err)
	}
	segment, _ := godub.NewLoader().Load(stream)

	name := fmt.Sprintf("%s.mp3", video.Title)
	dirname, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	fileDestination := path.Join(dirname, "Music", name)

	log.Println("Saving file at",fileDestination)

	err = godub.NewExporter(fileDestination).
		WithDstFormat("mp3").
		WithBitRate(256000).
		Export(segment)

	if err != nil {
		log.Fatal(err) 
	}

	log.Println(video.Title, video.Author, "Downloaded !! ")


}
