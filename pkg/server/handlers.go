package server

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

const (
	ErrCreatingID            = "failed creating id"
	ErrCouldNotWriteResponse = "could not write response"
)

type ResultList []Result

type Result struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

func (server *Server) index(writer http.ResponseWriter, req *http.Request) {
	list := ResultList{
		Result{ID: "abc123", URL: "something"},
		Result{ID: "zxy456", URL: "another"},
	}

	writer.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(&list)
	if err != nil {
		http.Error(writer, "could not write response", http.StatusInternalServerError)
	}
}

func (server *Server) resultIndex(writer http.ResponseWriter, req *http.Request) {
	if req.ContentLength > server.config.MaxBodySize {
		http.Error(writer, "request body too large", http.StatusExpectationFailed)
		return
	}

	req.Body = http.MaxBytesReader(writer, req.Body, server.config.MaxBodySize)

	id, err := generateID()
	if err != nil {
		http.Error(writer, errors.Wrap(err, ErrCreatingID).Error(), http.StatusInternalServerError)
		return
	}

	list := ResultList{
		Result{ID: id, URL: "something"},
	}

	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(&list)
	if err != nil {
		http.Error(writer, "could not write response", http.StatusInternalServerError)
	}
}

func (server *Server) resultShow(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	// todo: look up data from data file, or in-memory
	list := ResultList{
		Result{ID: id, URL: "something"},
	}

	writer.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(&list)
	if err != nil {
		http.Error(writer, "could not write response", http.StatusInternalServerError)
	}
}

func generateID() (string, error) {
	b := make([]byte, 10)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}
