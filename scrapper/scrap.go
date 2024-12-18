package scrapper

import (
	"log"
	"regexp"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

var client *colly.Collector
var mutex *sync.Mutex
var script string

func ScrapLinks(search string) []string {

	script = ""
	url := getSearchURL(search)

	log.Println(url)
	err := client.Visit(url)
	if err != nil {
		log.Fatal(err)
		return []string{}
	}

	mutex.Lock()
	strings := scriptScrapper(script)
	mutex.Unlock()

	// log.Println(strings)
	return strings

}

func getSearchURL(searchStr string) string {

	searchStr = strings.ReplaceAll(searchStr, " ", "+")

	return "https://www.youtube.com/results?search_query=" + searchStr

}

func scriptScrapper(s string) []string {

	re := regexp.MustCompile(`\{&#34;videoRenderer&#34;:\{&#34;videoId&#34;:&#34;[\-_a-zA-Z0-9]{11}`)
	match := re.FindAllStringSubmatch(s, 10)

	links := []string{}
	for _, j := range match {

		j[0] = "https://www.youtube.com/watch?v=" + strings.Split(j[0], ";")[5]

		// log.Println(j[0])

		links = append(links, j[0])
	}

	return links
}

func init() {
	client = colly.NewCollector()
	client.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	mutex = &sync.Mutex{}

	script = ""

	client.OnHTML("script", func(e *colly.HTMLElement) {

		data, err := e.DOM.Html()
		if err != nil {
			log.Fatal(err)
		}

		if len(data) < 300 {
			return
		}

		mutex.Lock()
		defer mutex.Unlock()
		if len(script) <= len(data) {
			// log.Println(len(data))

			// store the script tag with the largest length in file
			// file, err := os.Create("script.js")
			// if err != nil {
			// 	log.Fatal(err)
			// }
			// defer file.Close()

			// _, err = file.WriteString(data)
			// if err != nil {
			// 	log.Fatal(err)
			// }

			script = data
		}

	})

}
