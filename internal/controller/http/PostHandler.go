package http

import (
	"fmt"
	dto "forum/internal/DTO"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != urlPostCreate {
		h.notFound(w)
		return
	}

	if r.Method == "GET" {
		ts, err := template.ParseFiles("./ui/templates/post/createPost.html")
		if err != nil {
			log.Printf("Create Post: Execute:%v", err)
			return
		}

		categories, err := h.services.GetAllCategories()
		if err != nil {
			h.errorLog.Println(err.Error())
			h.serverError(w, err)
			return
		}

		err = ts.Execute(w, categories)
		if err != nil {
			h.serverError(w, err)
			return
		}
	} else if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			log.Println("error parse form :", err)
			return
		}
		title := r.PostFormValue("title")
		text := r.PostFormValue("text")
		categories := r.Form["categories"]

		if title == "" || text == "" || categories == nil {
			h.clientError(w, 400)
			return
		}
		var userID int64 = 60 /// DEL

		err = h.services.CreatePost(dto.PostDto{Title: title, Text: text, User: dto.UserDto{ID: userID}}, categories)

		if err != nil {
			h.serverError(w, err)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/"), http.StatusSeeOther)
	}
}

func (h *Handler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdatePost из page"))
}

func (h *Handler) GetPost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		h.notFound(w)
		return
	}

	fmt.Fprintf(w, "GetPost ID %d...", id)
}

func (h *Handler) ListPosts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ListPosts из page"))
}

func (h *Handler) DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeletePost из page"))
}