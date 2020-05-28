package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ashmilhussain/GO-RestWithUnitTest/pkg/models"
	"github.com/ashmilhussain/GO-RestWithUnitTest/pkg/responses"
	"github.com/gorilla/mux"
)

func (server *Server) CreatePost(w http.ResponseWriter, r *http.Request) {

	post := models.Post{}
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	post.Prepare()
	err = post.SavePost(server.DB)

	if err != nil {

		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, "Post Created Successfuly")
}

func (server *Server) GetPosts(w http.ResponseWriter, r *http.Request) {

	var post models.Post
	posts, err := post.FindAllPosts(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, posts)
}

func (server *Server) UpdatePost(w http.ResponseWriter, r *http.Request) {

	var post models.Post
	json.NewDecoder(r.Body).Decode(&post)
	err := post.UpdatePost(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(w, http.StatusCreated, "Post Updated Successfuly")

}

func (server *Server) DeletePost(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	postID := vars["id"]
	var post models.Post
	post.PostID, _ = strconv.Atoi(postID)
	err := post.DeletePost(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(w, http.StatusCreated, "Post Deleted Successfuly")

}
