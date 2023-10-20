package main

import (
	"encoding/json"
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

type CustomAlert struct {
	AlertID   string `json:"alert_id"`
	Model     string `json:"model"`
	AlertType string `json:"alert_type"`
	AlertTS   string `json:"alert_ts"`
	Severity  string `json:"severity"`
	TeamSlack string `json:"team_slack"`
}

type WriteResponse struct {
	AlertID string `json:"alert_id"`
	Error   string `json:"error"`
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
	var alert Alert
	if err := json.NewDecoder(r.Body).Decode(&alert); err != nil {
		writeAlertError(w, alert.AlertID, http.StatusInternalServerError, err.Error())
		return
	}

	if alert.AlertID == "" || alert.ServiceName == "" || alert.ServiceID == "" {
		writeAlertError(w, alert.AlertID, http.StatusBadRequest, "Alert ID, Service Name and Service ID cannot be empty")
		return
	}
	for _, existingAlert := range alerts {
		if existingAlert.AlertID == alert.AlertID {
			writeAlertError(w, alert.AlertID, http.StatusBadRequest, "Alert ID already present")
			return
		}
	}
	alerts = append(alerts, alert)

	response := WriteResponse{
		AlertID: alert.AlertID,
		Error:   "",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// GET Request Handler (Read Alerts)
func ReadAlerts(w http.ResponseWriter, r *http.Request) {
	serviceID := r.URL.Query().Get("service_id")
	startTS := r.URL.Query().Get("start_ts")
	endTS := r.URL.Query().Get("end_ts")

	if serviceID == "" {
		readAlertsError(w, http.StatusBadRequest, "Service ID cannot be empty")
		return
	}

	responseAlerts, err := readAlertsResponse(serviceID, startTS, endTS)
	if err != nil {
		readAlertsError(w, http.StatusNotFound, err.Error())
		return
	}
	response := struct {
		ServiceID   string        `json:"service_id"`
		ServiceName string        `json:"service_name"`
		Alerts      []CustomAlert `json:"alerts"`
	}{
		ServiceID:   serviceID,
		ServiceName: responseAlerts[0].ServiceName,
	}

	//To omit the repetition of ServiceID and ServiceName for each alertID in the response
	for _, alert := range responseAlerts {
		customAlert := CustomAlert{
			AlertID:   alert.AlertID,
			Model:     alert.Model,
			AlertType: alert.AlertType,
			AlertTS:   alert.AlertTS,
			Severity:  alert.Severity,
			TeamSlack: alert.TeamSlack,
		}
		response.Alerts = append(response.Alerts, customAlert)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

//readAlertsResponse returns the list of alerts mtaching to the request paramters sent
func readAlertsResponse(serviceID, startTS, endTS string) ([]Alert, error) {
	readAlertsResponse := []Alert{}
	for _, alert := range alerts {
		if (alert.ServiceID == serviceID) &&
			(startTS == "" || alert.AlertTS >= startTS) &&
			(endTS == "" || alert.AlertTS <= endTS) {
			readAlertsResponse = append(readAlertsResponse, alert)
		}
	}
	if len(readAlertsResponse) == 0 {
		return nil, fmt.Errorf("No alerts found")
	}

	return readAlertsResponse, nil
}

//readAlertsError handles the errors while retrieving the alerts
func readAlertsError(w http.ResponseWriter, httpStatusCode int, errorMessage string) {
	response := struct {
		Success bool   `json:"success"`
		Error   string `json:"error"`
	}{
		Success: false,
		Error:   errorMessage,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	json.NewEncoder(w).Encode(response)
}

//writeAlertError handles the errors while storing the alerts
func writeAlertError(w http.ResponseWriter, alertID string, httpStatusCode int, errorMessage string) {
	response := WriteResponse{
		AlertID: alertID,
		Error:   errorMessage,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	json.NewEncoder(w).Encode(response)
}
