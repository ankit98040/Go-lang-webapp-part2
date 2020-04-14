//webscraper
//Accessing the internet

package main

import ("fmt"
"net/http"
"io/ioutil")

func main(){
	resp,_ := http.Get("https://www.westernunion.com/sitemap.xml")

		//https://www.washingtonpost.com/news-sitemap-index.xml
		//as http.Get returns 2 variables so we have declared two variables. same goes for ioutil
		//for web parsing. sending request and getting data
		//get is for getting request

	bytes,_ := ioutil.ReadAll(resp.Body) 
	// we are going to grap all the information from the body
	string_body := string(bytes)
	//lastly we need to convert the data into string
	fmt.Println(string_body)
	resp.Body.Close()
}