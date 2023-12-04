package scrapper

import (
	"log"
	"strings"
	"sync"
	"regexp"
	"github.com/gocolly/colly"
)

var client *colly.Collector
var mutex *sync.Mutex
var script string

func ScrapLinks(search string)([]string){

	url := getSearchURL(search)

	client.Visit(url)

	mutex.Lock()
	strings := scriptScrapper(script)
	mutex.Unlock()


	log.Println(strings)
	return strings

}


func getSearchURL(searchStr string) string {

	searchStr = strings.ReplaceAll(searchStr, " ", "+")
	return  "https://www.youtube.com/results?search_query=" + searchStr

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

		links = append(links,key)
	}

	return links
}


func init(){
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
			log.Println(len(data))
			script = data
		}

	})	




}