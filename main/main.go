package main

import (
	"fmt"
	"net/http"

	"github.com/vexth1/handler"
	"github.com/vexth1/middleware"

	"github.com/vexth1/content"
)

func main() {
	config := &handler.Configuration{":8100"}
	ep := handler.NewEntrypoint(config, nil)

	ep.AttachMiddleware(middleware.NegroniRecoverHandler())
	ep.AttachMiddleware(middleware.NegroniLoggerHandler())
	ep.AttachMiddleware(middleware.NegroniCorsAllowAll())
	ep.AttachMiddleware(middleware.NegroniJwtHandler(content.PrivateTokenKey, skipper, nil, jwtErrHandler))

	router := handler.NewRouter()
	router.Get("/", content.Hello)
	router.Post("/test", content.Hello)

	router.Get("/skipper", content.GetToken)

	ep.Start(router.Handler())

	ep.Run()
}

func skipper(path string) bool {
	if path == "/skipper" {
		return true
	}
	return false
}

func jwtErrHandler(w http.ResponseWriter, r *http.Request, err string) {
	fmt.Println(err)
	http.Error(w, err, 401)
}
