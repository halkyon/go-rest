package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type Resource struct {
	Name string `json:"name"`
}

func (server *Server) index(writer http.ResponseWriter, req *http.Request) {
	// todo: get data from data source
	list := []Resource{
		{Name: "Bob"},
		{Name: "Joe"},
	}

	writer.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(&list)
	if err != nil {
		http.Error(writer, errors.Wrap(err, "could not write response").Error(), http.StatusInternalServerError)
	}
}

func (server *Server) resourceIndex(writer http.ResponseWriter, req *http.Request) {
	if req.ContentLength > server.config.MaxBodySize {
		http.Error(writer, "request body too large", http.StatusExpectationFailed)
		return
	}

	req.Body = http.MaxBytesReader(writer, req.Body, server.config.MaxBodySize)

	var resource Resource
	err := json.NewDecoder(req.Body).Decode(&resource)
	if err != nil {
		http.Error(writer, errors.Wrap(err, "failed to parse request body").Error(), http.StatusBadRequest)
		return
	}

	// todo: validate data, and write resource data

	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(&resource)
	if err != nil {
		http.Error(writer, errors.Wrap(err, "could not write response").Error(), http.StatusInternalServerError)
	}
}

func (server *Server) resourceShow(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	// todo: get data, output 404 if not found
	resource := Resource{Name: id}

	writer.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(&resource)
	if err != nil {
		http.Error(writer, errors.Wrap(err, "could not write response").Error(), http.StatusInternalServerError)
	}
}
