package service

import (
	"net/http"
	"strconv"
	"net"
)

func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	//dbUp := DBClient.Check()
	//if dbUp {
	//	data, _ := json.Marshal(healthCheckResponse{Status: "UP"})
	//	writeJsonResponse(w, http.StatusOK, data)
	//} else {
	//	data, _ := json.Marshal(healthCheckResponse{Status: "Database unaccessible"})
	//	writeJsonResponse(w, http.StatusServiceUnavailable, data)
	//}
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

func GetIP() string {
	addr, err := net.InterfaceAddrs()
	if err != nil {
		return "error"
	}
	for _, address := range addr {
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return  ipNet.IP.String()
			}
		}
	}
	panic("Unable to determine local IP address (non loopback). Exiting.")
}