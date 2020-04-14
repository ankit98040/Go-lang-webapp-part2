//using of maps

package main

import "fmt"

func main(){
	grades := make(map[string]float32)
	grades["ankit"]= 92
	grades["amit"]=67
	grades["bikash"]=62

	fmt.Println(grades)
	Timsgrade := grades["amit"]
	fmt.Println(Timsgrade)

	delete(grades, "ankit")
	fmt.Println(grades)

	for k,v := range grades {
		fmt.Println(k, ":" ,v)// iterating through loops
	}
}