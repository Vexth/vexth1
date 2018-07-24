package middleware

import (
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

func NegroniCorsAllowAll() negroni.Handler {
	return cors.AllowAll()
}

func NegroniCorsNew(opt cors.Options) negroni.Handler {
	return cors.New(opt)
}
