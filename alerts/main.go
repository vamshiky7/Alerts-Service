package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type Alert struct {
	AlertID     string `json:"alert_id"`
	ServiceID   string `json:"service_id"`
	ServiceName string `json:"service_name"`
	Model       string `json:"model"`
	AlertType   string `json:"alert_type"`
	AlertTS     string `json:"alert_ts"`
	Severity    string `json:"severity"`
	TeamSlack   string `json:"team_slack"`
}

var alerts []Alert

func main() {
	// Router setup
	r := chi.NewRouter()
	// Route requests
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from home"))
	})
	r.Post("/alerts", WriteAlert)
	r.Get("/alerts", ReadAlerts)
	// Server start
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", "8080"),
		Handler: r,
	}
	log.Println("Server started...")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(fmt.Sprintf("%+v", err))
	}
}

// POST Request Handler (Write Alert)
func WriteAlert(w http.ResponseWriter, r *http.Request) {
	// 1. Parse the JSON request body
	// 2. Validate the input data
	// 3. Store the alert data
	// 4. Handle errors
	// 5. Respond with an appropriate HTTP status code and JSON response }
	// GET Request Handler (Read Alerts)
}
func ReadAlerts(w http.ResponseWriter, r *http.Request) {
	// 1. Parse and validate query parameters
	// 2. Query data storage to retrieve alerts
	// 3. Create a response JSON object
	// 4. Handle errors
	// 5. Respond with an appropriate HTTP status code and JSON response }
}
