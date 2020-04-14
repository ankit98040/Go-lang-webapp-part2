//mapping news data

package main

import ("fmt"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type Sitemapindex struct {
	Locations []string `xml:"sitemap>loc"`//changing from location type to string type
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword string
	Location string
}

func main() {
	var s Sitemapindex
	var n News
	news_map := make(map[string]NewsMap)
	resp, _ := http.Get("https://www.westernunion.com/sitemap.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)
	//news_map := make(map[string]NewsMap)

	for _, Location := range s.Locations {
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)

		for idx, _ := range n.Keywords {
			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}
	for idx, data := range news_map {
		fmt.Println("\n\n\n\n\n",idx)
		fmt.Println("\n",data.Keyword)
		fmt.Println("\n",data.Location)
	}
}