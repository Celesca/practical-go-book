package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func handleUserAPI(w http.ResponseWriter, r *http.Request) {
	log.Println("I started processing the request")
	time.Sleep(15 * time.Second)
	fmt.Println(w, "Hello World!")
	log.Println("I finished processing the request")
}

func main() {
	listenAddr := os.Getenv("LISTEN_ADDR")
	if len(listenAddr) == 0 {
		listenAddr = ":8080"
	}

	timeoutDuration := 14 * time.Second

	userHandler := http.HandlerFunc(handleUserAPI)
	hTimeout := http.TimeoutHandler(userHandler, timeoutDuration, "Timeout!")

	mux := http.NewServeMux()
	mux.Handle("/user", hTimeout)

	log.Fatal(http.ListenAndServe(listenAddr, mux))
}
