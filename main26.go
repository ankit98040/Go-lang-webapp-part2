//Html templates
//basictemplating.html

package main

import ("fmt"
"net/http"
"html/template")

type NewsAggPage struct {
    Title string
    News string
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {

	p := NewsAggPage{Title: "Amazing News Aggregator", News: "and here whatever you write will be shown up"}
	t, _ := template.ParseFiles("basictemplating.html")
	t.Execute(w, p)
}


func index_handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<h1>hey there. go is really nice</h1>")
}

func main(){
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000", nil)
}