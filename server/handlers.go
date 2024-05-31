package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"ur_shortner/repositories"
)

type CreateRequest struct {
    URL string `json:"url"`
}

type CreateResponse struct {
    ShortURL string `json:"short_url"`
}


func CreateHandler(w http.ResponseWriter, r *http.Request) {
    var req CreateRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    shortURL := generateShortURL()
    err = repositories.SaveURLMapping(shortURL, req.URL)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    response := CreateResponse{ShortURL: "http://localhost:8080/" + shortURL}
    json.NewEncoder(w).Encode(response)
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    shortURL := vars["shortURL"]

    originalURL, err := repositories.GetOriginalURL(shortURL)
    if err != nil {
        http.Error(w, "URL not found", http.StatusNotFound)
        return
    }

    http.Redirect(w, r, originalURL, http.StatusFound)
}
