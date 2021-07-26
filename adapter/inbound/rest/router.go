//go:generate go run github.com/swaggo/swag/cmd/swag init -g router.go

package rest

import (
	"blog/adapter/inbound/rest/controller"
	"blog/app/usecase"
	"blog/domain/blog"
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
	"net/http"
	"net/url"
	"path"
)

func NewRouter(repo blog.BlogRepository, shortURL usecase.ShortURLGateway) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/blogs/{id}", controller.NewBlogController(repo, shortURL).Get).Methods(http.MethodGet)
	r.HandleFunc("/blogs/{id}/share", controller.NewBlogController(repo, shortURL).Share).Methods(http.MethodGet)
	attachSwagger(r)
	return r
}

// @title Blog Service Swagger API
// @version 1.0
// @description This is a sample server Blog server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host petstore.swagger.io
func attachSwagger(r *mux.Router) {
	const host = "localhost:8080"
	docPath := path.Join("swagger", "doc.json")
	swaggerURL := url.URL{Scheme: "http", Host: host, Path: docPath}
	r.PathPrefix("/swagger/").HandlerFunc(httpSwagger.Handler(
		httpSwagger.URL(swaggerURL.String()),
	))
}
