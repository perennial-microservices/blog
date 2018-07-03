package service

import (
	"github.com/perennial-microservices/blog/dbclient"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"strconv"
	"github.com/perennial-microservices/blog/postservice/model"
	blogService "github.com/perennial-microservices/blog/service"
)

var DBClient dbclient.IBoltClient

func CreatePost(w http.ResponseWriter, r *http.Request) {

}

func GetPosts(w http.ResponseWriter, r *http.Request) {

}

func GetPost(w http.ResponseWriter, r *http.Request) {
	var postId = mux.Vars(r)["postId"]

	result, err := DBClient.GetOne(postId)

	post, _ := result.(model.Post)

	post.ServedBy = blogService.GetIP()

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, _ := json.Marshal(post)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {

}

func DeletePost(w http.ResponseWriter, r *http.Request) {

}