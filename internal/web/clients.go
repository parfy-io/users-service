package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/parfy-io/users-service/internal"
	"github.com/sirupsen/logrus"
	"net/http"
	"regexp"
)

func (s *Server) createClient(w http.ResponseWriter, r *http.Request) {
	reqBody := struct {
		ID string `json:"id"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	match, err := regexp.MatchString("^([a-z]|[0-9]|[A-Z])+$", reqBody.ID)
	if err != nil {
		logrus.WithError(err).Error("failed to compile create-client id-check regex")
		writeInternalServerError(w)
		return
	}
	if !match {
		writeError(w, http.StatusBadRequest, "client>id must only contain a-zA-Z0-9")
		return
	}

	err = s.service.CreateClient(reqBody.ID)
	if err != nil {
		if errors.Is(err, internal.ErrClientAlreadyExists) {
			writeError(w, http.StatusConflict, "client already exists")
			return
		}

		logrus.WithError(err).Error("failed to create client")
		writeInternalServerError(w)
		return
	}
	w.Header().Add("Location", fmt.Sprintf("/v1/clients/%s", reqBody.ID))
	w.WriteHeader(http.StatusCreated)
}
