package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	fmt.Println("listening...")
	err := http.ListenAndServe(GetPort(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func GetPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		fmt.Println("INFO: No PORT environment variable detected, Starting Server at port " + port)
	}
	return ":" + port
}
