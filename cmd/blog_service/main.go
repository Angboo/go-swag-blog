package main

import (
	"blog/adapter/inbound/rest"
	"blog/adapter/outbound"
	"log"
	"net/http"
)

func main() {
	config := NewConfig()
	db := outbound.NewMemoryDB()
	blogRepo := outbound.NewBlogRepositoryImpl(db)
	shortURLGateway := outbound.NewShortURLGatewayImpl(config.ShortURLServiceHost)
	router := rest.NewRouter(blogRepo, shortURLGateway)
	log.Fatal(http.ListenAndServe(":8080", router))
}
