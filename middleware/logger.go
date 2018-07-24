package middleware

import "github.com/urfave/negroni"

// NegroniLoggerHandler 返回处理日志的方法.
func NegroniLoggerHandler() negroni.Handler {
	return negroni.NewLogger()
}
