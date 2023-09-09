package main

import (
	"encoding/json"
	"fmt"
	
	"net/http"
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

	response := Obj{
		Slack_name:      slackName,
		Current_day:     "Monday",
		Utc_time:        "2023-08-21T15:04:05Z",
		Track:           track,
		Github_file_url: "https://github.com/username/repo/blob/main/file_name.ext",
		Github_repo_url: "https://github.com/username/repo",
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
	http.HandleFunc("/charlesozo.com/api", handler)
	fmt.Println("Server is running on :8080...")
	http.ListenAndServe(":8080", nil)
}
