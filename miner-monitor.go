package main

import (
	"fmt"
	"net"
	"bufio"
	"encoding/json"
	"log"
	"strings"
	"time"
	"net/http"
	"bytes"
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

type SummaryResponse struct {
	SummaryData []Summary `json:"SUMMARY"`
}

func execute_cmd(command string) string {
	conn, _ := net.Dial("tcp", "127.0.0.1:4028")
	fmt.Fprintf(conn,"{\"command\" : \"%s\", \"parameter\" : 0}", command)
	message, _ := bufio.NewReader(conn).ReadString('\x00')


	return strings.Replace(message,"\x00", "", -1)
}

func get_miner_summary() Summary {
	message := []byte(execute_cmd("summary"))

	summary_output := SummaryResponse{}
	err := json.Unmarshal(message, &summary_output)

	if err != nil {

		log.Fatal("Fatal error while parsing JSON response for summary command: {}", err)
	}

	return summary_output.SummaryData[0]
}

func get_miner_id() string {
	ifs, _ := net.Interfaces()

	for _, v := range ifs {
		if (v.Name != "lo") {
			return v.HardwareAddr.String();
		}
	}

	return "Unknown";
}

func report_status() {
	url := "http://10.90.113.195/minermon/report.php"
	jsonStr, _ := json.Marshal(get_miner_summary())

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Miner-Id", get_miner_id())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Failed to report hashrate back to controller.")
	}

	defer resp.Body.Close()
}

func main() {

	log.Println("Starting miner-monitor // @friedkiwi 2017")

	report_status()

	for {
		output := get_miner_summary()
		report_status()
		log.Printf("hash rate: %f\n", output.GhsAv)
		time.Sleep(5 * time.Second)
	}
}
