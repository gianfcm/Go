package main

import (
	"fmt"
	"os"
	GET  "./lib"
	POST "./lib"
	DELETE "./lib"
	PUT "./lib"
)

func main() {

	os.Setenv("PHOTOS_API_URL","https://jsonplaceholder.typicode.com/photos/")
	os.Setenv("SLACK_API_URL","https://hooks.slack.com/services/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")
	var option int
	
	fmt.Println("Menu :  ")
	fmt.Println("1 - GET: ")
	fmt.Println("2 - POST: ")
	fmt.Println("3 - DELETE: ")
	fmt.Println("4 - PUT: ")

	fmt.Scan((&option))
	switch option {
	case 1:	GET.Get()
	case 2: POST.Post()
	case 3: DELETE.Delete()
	case 4: PUT.Put()
	}

}
