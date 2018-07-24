package handler

// HandlerFunc defines a handler function to handle http request.
type HandlerFunc func(*Context) error
