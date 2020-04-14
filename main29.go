//The idea of the defer statement is to put off (defer) the execution of something until the surrounding function is done.

package main

import "fmt"

func foo() {
  defer fmt.Println("Done!")
  fmt.Println("Doing some stuff, who knows what?")
}

func main() {
  foo()
}