//Looping in go

package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Sitemapindex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
	resp, _ := http.Get("https://www.westernunion.com/sitemap.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	var s Sitemapindex
	xml.Unmarshal(bytes, &s)
	for _, Location := range s.Locations {
		fmt.Printf("\n%s", Location)
	}
	//fmt.Println(s.Locations)
}