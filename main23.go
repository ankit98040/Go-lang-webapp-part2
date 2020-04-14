//continuing with our go web application

//now we want to visit those sitemaps and pull the information from there,
//the articles themselves and their related data

//Looping in go

package main

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type Sitemapindex struct {
	Locations []string `xml:"sitemap>loc"`//changing from location type to string type
}

//this whole work is now shortened using the above code
/*type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}
*/

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}


func main() {
	var s Sitemapindex
	var n News
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)
	news_map := make(map[string]NewsMap)

	for _, Location := range s.Locations {
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)
	}
}