# Alerts Service

The Alerts Service is a Go application to read and write alerts. 
Alerts are stored in memory and can be queried using service ID, start timestamp and end timestamp as the parameters.


### Prerequisites

- Golang installed on your system.
- A web browser or API client to interact with the service.

### Running the application

1. Clone this repository to your local machine:
   ```bash
   git clone https://github.com/vamshiky7/Alerts-Service
2. Navigate to the project directory
    ```bash
    cd alerts
3. Run the application
    ```bash
    go run .

### Reading alerts
    ```bash
    curl --location --request GET 'http://localhost:8080/alerts?service_id=my_test_service_id&start_ts=1695644190&end_ts=1695644192'

### Writing alerts
```bash
      curl --location --request POST 'http://localhost:8080/alerts' \
--header 'Content-Type: application/json' \
--data-raw '{
    "alert_id": "b950482e9911ec7e41f7ca5e5d9a424",
    "service_id": "my_test_service_id",
    "service_name": "my_test_service",
    "model": "my_test_model",
    "alert_type": "anomaly",
    "alert_ts": "1695644190",
    "severity": "warning",
    "team_slack": "slack_ch"
}'








  

  







   
  
