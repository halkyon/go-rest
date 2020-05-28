package server

import (
	"fmt"
	"io"
	"net/http"
)

type Error interface {
	error
	Status() int
}

type StatusError struct {
	Code int
	Err  error
}

func (se StatusError) Error() string {
	return se.Err.Error()
}

func (se StatusError) Status() int {
	return se.Code
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

type Handler struct {
	stderr        io.Writer
	nestedHandler HandlerFunc
}

func (handler Handler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	err := handler.nestedHandler(writer, req)
	if err != nil {
		switch e := err.(type) {
		case Error:
			if e.Status() >= http.StatusInternalServerError {
				fmt.Fprintf(handler.stderr, "%s\n", e)
			}
			http.Error(writer, e.Error(), e.Status())
		default:
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}
