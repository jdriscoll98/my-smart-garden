package main

import (
	"encoding/json"
	"fmt"
	"io"
	"my-smart-garden/common"
	"net/http"
)

func dataHandler(w http.ResponseWriter, r *http.Request) {
	var data common.SensorData
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "Error parsing JSON data", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Received data: %+v\n", data)
	fmt.Fprintf(w, "Data received successfully")
}

func main() {
	http.HandleFunc("/data", dataHandler)
	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}
