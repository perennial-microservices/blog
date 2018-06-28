package service

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"strconv"
	"fmt"
	"github.com/perennial-microservices/blog/accountservice/dbclient"
)

//var DBClient dbclient.IBoltClient
var DBClient dbclient.IRedisClient

func GetAccount(w http.ResponseWriter, r *http.Request) {

	var accountId = mux.Vars(r)["accountId"]

	account, err := DBClient.QueryAccount(accountId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, _ := json.Marshal(account)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	dbUp := DBClient.Check()
	if dbUp && isHealthy {
		data, _ := json.Marshal(healthCheckResponse{Status: "UP"})
		writeJsonResponse(w, http.StatusOK, data)
	} else {
		data, _ := json.Marshal(healthCheckResponse{Status: "Database unaccessible"})
		writeJsonResponse(w, http.StatusServiceUnavailable, data)
	}
}

var isHealthy = true

func SetHealthyState(w http.ResponseWriter, r *http.Request) {
	var state, err = strconv.ParseBool(mux.Vars(r)["state"])

	if err != nil {
		fmt.Println("Invalid request to SetHealthyState, allowed values are true or false")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	isHealthy = state
	w.WriteHeader(http.StatusOK)
}

func writeJsonResponse(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(status)
	w.Write(data)
}

type healthCheckResponse struct {
	Status string `json:"status"`
}