package downloader

import (
	"fmt"
	"log"
	// "os"
	"path"
	"ytdown/scrapper"

	"github.com/Vernacular-ai/godub"
	"github.com/kkdai/youtube/v2"
)

var client youtube.Client

func Getfile(url string) (string, error) {

	video, err := client.GetVideo(url)
	if err != nil {
		panic(err)
	}

	format := video.Formats.FindByItag(140)

	stream, _, err := client.GetStream(video, format)
	if err != nil {
		return "", err
	}
	segment, _ := godub.NewLoader().Load(stream)

	name := fmt.Sprintf("%s.mp3", video.Title)
	// dirname, err := os.UserHomeDir()
	// if err != nil {
	// 	return "", err
	// }
	fileDestination := path.Join(name)

	log.Println("Saving file at", fileDestination)

	err = godub.NewExporter(fileDestination).
		WithDstFormat("mp3").
		WithBitRate(256000).
		Export(segment)

	if err != nil {
		return "", err
	}

	log.Println(video.Title, video.Author, "Downloaded !! ")

	return fileDestination, err
}

func FetchSearch(searchStr string) [][]string {

	links := scrapper.ScrapLinks(searchStr)

	names := [][]string{}

	for i := range links {
		video, err := client.GetVideo(links[i])
		if err != nil {
			panic(err)
		}
		names = append(names, []string{links[i], video.Title, video.Author})
	}

	return names
}

func init() {
	client = youtube.Client{}
}
