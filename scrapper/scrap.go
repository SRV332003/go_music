package scrapper

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"regexp"

	"github.com/gocolly/colly"
)

func Scrap() []string {

	searchurl := getSearchURL()
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	mutex := &sync.Mutex{}

	script :=""

	max := 0

	// select string for a tag with name ytd-video-renderer
	c.OnHTML("script", func(e *colly.HTMLElement) {
		// get html of the page
		data, err := e.DOM.Html()
		if err != nil {
			log.Fatal(err)
		}

		if len(data) < 300 {
			return
		}

		mutex.Lock()
		defer mutex.Unlock()
		if max <= len(data) {
			log.Println(len(data))
			max = len(data)
			script = data
		}

	})

	
	c.Visit(searchurl)

	mutex.Lock()

	strings := scriptScrapper(script)
	mutex.Unlock()
	fmt.Println(strings)
	return strings
}

func getSearchURL() string {

	reader := bufio.NewReader(os.Stdin)
	song, err := reader.ReadString('\n')
	// song := "tum ho"

	if err != nil {
		log.Fatal(err)
	}
	song = strings.Split(song, "\n")[0]
	song = strings.ReplaceAll(song, " ", "+")
	song = "https://www.youtube.com/results?search_query=" + song
	fmt.Println(song)

	return song
}

func scriptScrapper(s string) []string {
	re := regexp.MustCompile(`https://i.ytimg.com/vi/[A-Za-z]{11}/`)
	match := re.FindAllStringSubmatch(s,1000)

	mp := make(map[string]int)


	for _,j:=range match{

		mp[j[0]]=1
	}

	links := []string{}

	for key := range mp {

		key = "https://www.youtube.com/watch?v="+strings.Split(key, "/")[4]

		links = append(links,key )
	}

	return links
}