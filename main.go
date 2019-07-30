package main

import (
	"fmt"
	GET  "./lib"
	POST "./lib"
	DELETE "./lib"
	PUT "./lib"
)
func main() {

	var opcion int
	fmt.Println("Menu :  ")
	fmt.Println("1 - GET: ")
	fmt.Println("2 - POST: ")
	fmt.Println("3 - DELETE: ")
	fmt.Println("4 - PUT: ")
	fmt.Scan((&opcion))

	switch opcion {
	case 1:	GET.Get()
	case 2: POST.Post()
	case 3: DELETE.Delete()
	case 4: PUT.Put()
	}
}
