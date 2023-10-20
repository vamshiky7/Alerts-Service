package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadAlerts_Success(t *testing.T) {
	alerts = []Alert{{
		AlertID:     "1",
		ServiceID:   "my_test_service_id",
		ServiceName: "my_test_service",
		Model:       "model",
		AlertType:   "anamoly",
		AlertTS:     "1695644188",
		Severity:    "anamoly",
		TeamSlack:   "test_slack",
	}}
	req := httptest.NewRequest("GET", "/alerts?service_id=my_test_service_id&start_ts=1695643160&end_ts=16956443190", nil)
	w := httptest.NewRecorder()

	ReadAlerts(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestReadAlerts_MissingServiceID(t *testing.T) {
	r := httptest.NewRequest("GET", "/alerts", nil)
	w := httptest.NewRecorder()

	ReadAlerts(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)

}

func TestWriteAlert_Success(t *testing.T) {
	request := `{
        "alert_id": "b950482e9911ec7e41f7ca5e5d9a424",
        "service_id": "my_test_service_id",
        "service_name": "my_test_service",
        "model": "my_test_model",
        "alert_type": "anomaly",
        "alert_ts": "1695644190",
        "severity": "warning",
        "team_slack": "slack_ch"
    }`

	req := httptest.NewRequest("POST", "/alerts", strings.NewReader(request))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	WriteAlert(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestWriteAlert_DuplicateAlertID(t *testing.T) {
	request := `{
        "alert_id": "b950482e9911ec7e41f7ca5e5d9a424",
        "service_id": "my_test_service_id",
        "service_name": "my_test_service",
        "model": "my_test_model",
        "alert_type": "anomaly",
        "alert_ts": "1695644190",
        "severity": "warning",
        "team_slack": "slack_ch"
    }`
	alerts = []Alert{{AlertID: "b950482e9911ec7e41f7ca5e5d9a424"}}

	req := httptest.NewRequest("POST", "/alerts", strings.NewReader(request))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	WriteAlert(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestWriteAlert_EmptyAlertID(t *testing.T) {
	request := `{
        "alert_id": "",
        "service_id": "my_test_service_id",
        "service_name": "my_test_service",
        "model": "my_test_model",
        "alert_type": "anomaly",
        "alert_ts": "1695644190",
        "severity": "warning",
        "team_slack": "slack_ch"
    }`

	req := httptest.NewRequest("POST", "/alerts", strings.NewReader(request))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	WriteAlert(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
