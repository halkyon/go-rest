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

func (server *Server) index(writer http.ResponseWriter, req *http.Request) error {
	// todo: get data from data source
	list := []Resource{{"Bob"}, {"Joe"}}

	writer.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(&list)
	if err != nil {
		return StatusError{http.StatusInternalServerError, errors.Wrap(err, "could not write response")}
	}

	return nil
}

func (server *Server) resourceIndex(writer http.ResponseWriter, req *http.Request) error {
	if req.ContentLength > server.config.MaxBodySize {
		return StatusError{http.StatusExpectationFailed, errors.New("request body too large")}
	}

	req.Body = http.MaxBytesReader(writer, req.Body, server.config.MaxBodySize)

	var resource Resource
	err := json.NewDecoder(req.Body).Decode(&resource)
	if err != nil {
		return StatusError{http.StatusBadRequest, errors.Wrap(err, "failed to parse request body")}
	}

	// todo: validate data, and write resource data

	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(&resource)
	if err != nil {
		return StatusError{http.StatusInternalServerError, errors.Wrap(err, "could not write response")}
	}

	return nil
}

func (server *Server) resourceShow(writer http.ResponseWriter, req *http.Request) error {
	vars := mux.Vars(req)
	id := vars["id"]

	// todo: get data, output 404 if not found
	resource := Resource{id}

	writer.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(&resource)
	if err != nil {
		return StatusError{http.StatusInternalServerError, errors.Wrap(err, "could not write response")}
	}

	return nil
}
