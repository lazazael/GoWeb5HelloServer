package main

import (
	"fmt"
	"github.com/lazazael/GoWeb5HelloServer/pkg/handlers"
	"net/http"
)

const portNumber string = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/divide", handlers.Divide)

	/*
		http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
			n, err := fmt.Fprintf(w, "Hello,world!")
			fmt.Println(fmt.Sprintf("Number of bytes written: %d",n))
			if err != nil{
				fmt.Println(err)
			}
		})
	*/

	fmt.Println(fmt.Sprintf("Starting app on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
