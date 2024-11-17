package transport

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"posts/internal/domain"
)

type Posts interface {
	Create(context.Context, domain.Post) error
	Select() ([]domain.Post, error)
}
type Handler struct {
	posts Posts
}

func NewHandler(posts Posts) *Handler {
	return &Handler{posts: posts}
}
func (h *Handler) InitRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", h.Select).Methods(http.MethodGet)
	r.HandleFunc("/create", h.Create).Methods(http.MethodGet, http.MethodPost)
	return r
}
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		data, _ := ioutil.ReadFile("/Users/air/Desktop/posts/posts/static/html/create_post.html")

		w.Write(data)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Ошибка парсинга формы", http.StatusBadRequest)
		return
	}

	name := r.FormValue("username")
	post := r.FormValue("post")
	h.posts.Create(context.TODO(), domain.Post{Author: name, Post: post})
	http.Redirect(w, r, "/", http.StatusFound)

}
func (h *Handler) Select(w http.ResponseWriter, r *http.Request) {
	posts, err := h.posts.Select()
	if err != nil {
		log.Fatal(err)
	}
	for _, post := range posts {
		fmt.Fprintf(w, "%v\t%v\t%v\t%v\n\n", post.Id, post.Author, post.Post, post.Time)
	}

}
