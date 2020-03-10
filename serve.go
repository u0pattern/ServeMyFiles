package main

import (
	"net/http"
	"os"
	"log"
	"fmt"
)

func main() {
	var dir string
	var err error
	if len(os.Args) == 1 { // Get The Current Working Directory
		dir, err = os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
	}else{
		dir = os.Args[1]
	}
	fmt.Println("Serving "+dir+" files 127.0.0.1:3000 ....")
	http.Handle("/", http.FileServer(http.Dir(dir)))
	http.ListenAndServe(":3000", nil)
}
