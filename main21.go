//Parsing XML (extensible markup language)

package main

import ("fmt"
"net/http"
"io/ioutil"
"encoding/xml")

type Sitemapindex struct {
  Locations []Location `xml:"sitemap"`
}

type Location struct {
  Loc string `xml:"loc"`
}

//final part to convert slices to list
func (l Location) String () string {
  return fmt.Sprintf(l.Loc)
}
//for converting the given data into strings
//sprintf for string method


func main() {
  resp, _ := http.Get("https://www.westernunion.com/sitemap.xml")
  	//https://www.westernunion.com/sitemap.xml
  	//we now get all the urls of the site
  bytes, _ := ioutil.ReadAll(resp.Body)
  var s Sitemapindex
  xml.Unmarshal(bytes, &s)
  fmt.Println(s.Locations)
}


// [5]int == array

// []int == slice.     here location is a slice means its size can be changed