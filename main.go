package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	
	
)

type Obj struct {
	Slack_name      string `json:"slack_name"`
	Current_day     string `json:"current_day"`
	Utc_time        string `json:"utc_time"`
	Track           string `json:"track"`
	Github_file_url string `json:"github_file_url"`
	Github_repo_url string `json:"github_repo_url"`
	Status_code     int    `json:"status_code"`
}


func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract values from URL parameters using keys "slack_name" and "track"
	slackName := r.URL.Query().Get("slack_name")
	track := r.URL.Query().Get("track")

	if slackName == "" || track == "" {
		http.Error(w, "Both parameters 'slack_name' and 'track' are required", http.StatusBadRequest)
		return
	}
	currentDay := time.Now().Format("Monday")
	currentUtcTime := time.Now().UTC().Format("2006-01-02T15:04:05Z")

	response := Obj{
		Slack_name:      slackName,
		Current_day:    currentDay,
		Utc_time:        currentUtcTime,
		Track:           track,
		Github_file_url: "https://github.com/charlesozo/end-point/blob/main/main.go",
		Github_repo_url: "https://github.com/charlesozo/end-point",
		Status_code:     200,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Write(jsonResponse)
}

func main() {
	
	http.HandleFunc("/api", handler)
	fmt.Println("Server is running on :80...")
	log.Fatal(http.ListenAndServe(":80", nil))
}

