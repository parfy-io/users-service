package web

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Service interface {
}

type Server struct {
	server  *http.Server
	service Service
}

func NewServer(service Service, bindAddress string) *Server {
	s := &Server{
		service: service,
	}

	r := mux.NewRouter()
	v1 := r.PathPrefix("/v1").Subrouter()
	v1.Path("/internal/alive").Methods(http.MethodGet).HandlerFunc(s.aliveHandler)

	s.server = &http.Server{
		Addr:    bindAddress,
		Handler: r,
	}

	return s
}

func (s *Server) ListenAndServe() error {
	err := s.server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (s Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func writeInternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	_, err := w.Write([]byte(`{"message":"internal server error"}`))
	if err != nil {
		logrus.WithError(err).Error("Failed to write error response")
	}
}

func writeError(w http.ResponseWriter, statusCode int, message string) {
	b, err := json.Marshal(struct {
		Message string `json:"message"`
	}{
		Message: message,
	})
	if err != nil {
		logrus.WithError(err).Error("Failed to marshal json error response")
		writeInternalServerError(w)
		return
	}

	w.WriteHeader(statusCode)
	_, err = w.Write(b)
	if err != nil {
		logrus.WithError(err).Error("Failed to write error response")
		writeInternalServerError(w)
		return
	}
}
