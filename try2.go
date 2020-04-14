package main

import ("fmt"
"net/http")

func index_handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<h1>hey there</h1>")
	fmt.Fprintf(w, "<p>Go is awesome</p>")
	fmt.Fprintf(w, "<p>You %s can even add %s </p>","can","<strong>variables</strong")
}

func main(){
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8000", nil)// nil is the return type
}