package controller

import (
	"blog/app/usecase"
	"blog/domain/blog"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type BlogController struct {
	repo        blog.BlogRepository
	shortURL    usecase.ShortURLGateway
	blogUseCase *usecase.BlogUseCase
}

func NewBlogController(repo blog.BlogRepository, shortURL usecase.ShortURLGateway) BlogController {
	return BlogController{repo: repo, shortURL: shortURL,
		blogUseCase: usecase.NewBlogUseCase(repo, shortURL)}
}

type blogResponse struct {
	// blog identify
	Id    int64  `json:"id" example:"10086"`
	// title of blog
	Title string `json:"title" example:"title"`
	// a rich body for blog
	Body  string `json:"body" example:"body"`
}

func newBlogResponse(blog *blog.Blog) blogResponse {
	return blogResponse{Id: blog.Id, Title: blog.Title, Body: blog.Body}
}

// ShowBlog godoc
// @summary Show a blog
// @description get blog by ID
// @id get-blog-by-id
// @accept json
// @produce json
// @param id path int true "Blog ID"
// @success 200 {object} blogResponse
// @failure 404
// @failure 500
// @router /blogs/{id} [get]
func (c BlogController) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	found, err := c.blogUseCase.FindBlog(int64(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if found == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resp, _ := json.Marshal(newBlogResponse(found))
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(resp)
}

func (c BlogController) Share(_ http.ResponseWriter, _ *http.Request) {

}
