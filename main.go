package main

import (
	"github.com/MinatoNamikaze02/GoBlockchain/src"
	"time"
	// "container/list"
	"net/http"
	"io"
	"fmt"
	"encoding/json"
)

var bc = blockchain.Blockchain{}


func getChain(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		io.WriteString(w, "Method not allowed!\n")
		return
	}	
	chain := bc.Chain
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chain)

}
type Transaction struct {
	sender string
	content string
	timestamp time.Time
}
func newTransaction(w http.ResponseWriter, r *http.Request) {
	var body Transaction
	if r.Method != "POST" {
		io.WriteString(w, "Method not allowed!\n")
		return
	}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	body.timestamp = time.Now()
	bc.UnconfirmedTransactions.PushBack(body)
	w.WriteHeader(http.StatusCreated)

}

func pendingTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		io.WriteString(w, "Method not allowed!\n")
		return
	}	
	chain := bc.UnconfirmedTransactions
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chain)
}

type ErrorResponseData struct {
	detail string
}

func mine(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		io.WriteString(w, "Method not allowed!\n")
		return
	}	
	var data ErrorResponseData
	x := bc.Mine()
	w.Header().Set("Content-Type", "application/json")
	if x == 0{
		data.detail = "No transaction to mine"
		json.NewEncoder(w).Encode(data)
		return
	}
	w.WriteHeader(http.StatusCreated)
	data.detail = fmt.Sprintf("Block %d mined successfully", x)
	json.NewEncoder(w).Encode(data)
}

func main() {
	// Create a new blockchain
	http.HandleFunc("/chain", getChain)
	http.HandleFunc("/new_transaction", newTransaction)
	http.HandleFunc("/pending_transaction", pendingTransaction)
	http.HandleFunc("/mine", mine)
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println(err)
	}
}
