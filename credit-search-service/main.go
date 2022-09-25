package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/hack-btg/backend/credit-search-service/client"
)

var history [][]Rate

type Rate struct {
	BankName     string  `json:"bank_name"`
	AnualTaxRate float64 `json:"anual_tax_rate"`
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/credit/search", getBankRates)
	http.HandleFunc("/credit/history", getHistory)

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
	resp := []Rate{}

	for i := 0; i < 3; i++ {
		rand.Seed(time.Now().UnixNano())
		bank := rand.Intn(len(banks) - 1)

		rand.Seed(time.Now().UnixNano())
		bankRate := rand.Intn(40)

		resp = append(resp, Rate{
			BankName:     banks[bank].OrganisationName,
			AnualTaxRate: (float64(bankRate) / 100),
		})
	}
	history = append(history, resp)

	fmt.Printf("history: %+v\n", history)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func getHistory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got '/credit/search' request")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(history)
}
