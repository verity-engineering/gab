package main

import (
  "fmt"
	"log"
)

func main() {
	log.Println("Tidying up")
	Initialize()
  storage := GetRepositories(true)
  
  //&mt.Println(&storage)
  fmt.Printf("%+v\n", storage) 
 

}
