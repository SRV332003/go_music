package downloader

import (
	"fmt"
	"go_music/scrapper"
	"log"
	"os"
	"path"

	"github.com/Vernacular-ai/godub"
	"github.com/kkdai/youtube/v2"
)

var client youtube.Client

func Getfile(url string) (string, string, error) {

	// log.Println("Fetching", url)

	video, err := client.GetVideo(url)
	if err != nil {
		return "", "", err
	}

	// log.Println("Downloading", video.Title, video.Author)

	format := video.Formats.Itag(140)[0]

	stream, _, err := client.GetStream(video, &format)
	if err != nil {
		return "", "", err
	}
	segment, _ := godub.NewLoader().Load(stream)

	name := fmt.Sprintf("%s.mp3", video.Title)
	dirname, err := os.UserHomeDir()
	if err != nil {
		return "", "", err
	}
	fileDestination := path.Join(dirname, "Music", name)

	log.Println("Saving file at", fileDestination)

	err = godub.NewExporter(fileDestination).
		WithDstFormat("mp3").
		WithBitRate(256000).
		Export(segment)

	if err != nil {
		return "", "", err
	}

	log.Println(video.Title, video.Author, "Downloaded !! ")

	return name, fileDestination, err
}

func FetchSearch(searchStr string) [][]string {

	// log.Println("Searching for", searchStr)
	links := scrapper.ScrapLinks(searchStr)

	log.Println("Found", links)

	names := [][]string{}

	for i := range links {
		video, err := client.GetVideo(links[i])
		if err != nil {

			log.Panic("Downloader:", err)
			return names
		}
		names = append(names, []string{links[i], video.Title, video.Author})
	}

	return names
}

func init() {
	client = youtube.Client{}
}
