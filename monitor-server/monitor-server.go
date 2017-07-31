package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
)

type Summary struct {
	Elapsed int `json:"Elapsed"`
	Ghs5s float32 `json:"GHS 5s"`
	GhsAv float32 `json:"GHS av"`
	FoundBlocks int `json:"Found Blocks"`
	GetWorks int `json:"Getworks"`
	Accepted int `json:"Accepted"`
	Rejected int `json:"Rejected"`
	HardwareErrors int `json:"Hardware Errors"`
	Utility float32 `json:"Utility"`
	Discarded int `json:"Discarded"`
	Stale int `json:"Stale"`
	GetFailures int `json:"Get Failures"`
	LocalWork int `json:"Local Work"`
	RemoteFailures int `json:"Remote Failures"`
	NetworkBlocks int `json:"Network Blocks"`
	TotalMH float64 `json:"Total MH"`
	WorkUtility float32 `json:"Work Utility"`
	DifficultyAccepted float64 `json:"Difficulty Accepted"`
	DifficultyRejected float64 `json:"Difficulty Rejected"`
	DifficultyStale float64 `json:"Difficulty Stale"`
	BestShare int `json:"Best Share"`
	HardwareErrorRate float32 `json:"Device Hardware%"`
	DeviceRejected float32 `json:"Device Rejected%"`
	PoolRejected float32 `json:"Pool Rejected%"`
	PoolStale float32 `json:"Pool Stale%"`
	LastGetwork int `json:"Last getwork"`
}


func main() {
	fmt.Println("Hello, world!")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/miner_status/{miner_id}", StatusReceived).Methods("POST")
	log.Fatal(http.ListenAndServe(":8888", router))
}

func StatusReceived (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var status Summary
	_ = json.NewDecoder(r.Body).Decode(&status)

	log.Printf("Received status from miner %s: %f GH/s\n", params["miner_id"], status.GhsAv)
}

func Index (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
