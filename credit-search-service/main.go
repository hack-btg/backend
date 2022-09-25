package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/hack-btg/backend/credit-search-service/client"
)

var history []interface{}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/credit/search", getBankRates)
	http.HandleFunc("/credit/history", getBankRates)

	err := http.ListenAndServe(":5002", nil)
	log.Println(err)
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got '/' request")
	io.WriteString(w, "ok")
}
func getBankRates(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got '/credit/search' request")
	banks, err := client.GetBanks()
	if err != nil {
		log.Println("[getBankRates] err ", err.Error())
		return
	}
	rand.Seed(time.Now().UnixNano())

	fmt.Println(rand.Intn(len(banks)))
	io.WriteString(w, "Hello, HTTP!\n")
}
