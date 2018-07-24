package main

import (
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
	ep.AttachMiddleware(middleware.NegroniJwtHandler(content.PrivateTokenKey, content.Skipper, nil, content.JwtErrHandler))

	router := handler.NewRouter()
	router.Get("/", content.Hello)
	router.Post("/test", content.Hello)

	router.Get("/skipper", content.GetToken)

	ep.Start(router.Handler())

	ep.Run()
}
