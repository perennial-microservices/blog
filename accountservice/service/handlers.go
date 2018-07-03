package service

import (
	"github.com/perennial-microservices/blog/dbclient"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"strconv"
	"github.com/perennial-microservices/blog/accountservice/model"
	blogService "github.com/perennial-microservices/blog/service"
)

var DBClient dbclient.IBoltClient

func CreateAccount(w http.ResponseWriter, r *http.Request) {

}

func GetAccounts(w http.ResponseWriter, r *http.Request) {

}

func GetAccount(w http.ResponseWriter, r *http.Request) {
	var accountId = mux.Vars(r)["accountId"]

	result, err := DBClient.GetOne(accountId)

	account, _ := result.(model.Account)

	account.ServedBy = blogService.GetIP()

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

func UpdateAccount(w http.ResponseWriter, r *http.Request) {

}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {

}