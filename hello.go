package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

// import "fmt"

// import "rsc.io/quote"

// func main() {
// 	fmt.Println("Hello, World!")
// }
func main(){
	log.Printf(get("https://jsonplaceholder.typicode.com/posts/1"))
}

func get(url string) string {
	resp, err := http.Get(url)

	if err != nil {
	log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	log.Fatalln(err)
	}
	//Convert the body to type string
	bodyString := string(body)
	return bodyString
}

func constructVoyagerUrl(){
	baseUrl := "https://www.linkedin.com/voyager/api/
}