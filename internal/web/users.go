package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/parfy-io/users-service/internal"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {
	var reqBody struct {
		ID       int64    `json:"id"`
		FullName string   `json:"full_name"`
		Names    []string `json:"names"`
	}

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	if reqBody.ID != 0 {
		writeError(w, http.StatusBadRequest, "user>id must not be set")
		return
	}

	if reqBody.FullName == "" {
		writeError(w, http.StatusBadRequest, "user>full_name must not be set")
		return
	}

	if len(reqBody.Names) < 1 {
		writeError(w, http.StatusBadRequest, "user>names must contain at least one name")
		return
	}

	clientID := mux.Vars(r)["clientID"]
	userID, err := s.service.CreateUser(clientID, internal.User{
		FullName: reqBody.FullName,
		Names:    reqBody.Names,
	})
	if err != nil {
		if errors.Is(err, internal.ErrClientDoesntExists) {
			http.NotFound(w, r)
			return
		}

		logrus.WithError(err).Error("failed to create user")
		writeInternalServerError(w)
		return
	}
	w.Header().Add("Location", fmt.Sprintf("/v1/clients/%s/users/%d", clientID, userID))
	w.WriteHeader(http.StatusCreated)
}
