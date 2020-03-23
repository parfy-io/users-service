package main

import (
	"context"
	"fmt"
	"github.com/ardanlabs/conf"
	"github.com/parfy-io/users-service/internal/storage"
	"github.com/parfy-io/users-service/internal/web"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg, err := newConfig()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to parse config")
	}

	cfgAsString, err := conf.String(&cfg)
	if err != nil {
		logrus.WithError(err).Fatal("Could not build config string")
	}
	fmt.Print(cfgAsString)
	logrus.Infof("Starting service")
	s, err := storage.New(cfg.DB.Host, cfg.DB.Port, cfg.DB.Username, cfg.DB.Password, cfg.DB.Name)
	if err != nil {
		logrus.WithError(err).Fatal("Could not create storage")
	}

	err = s.Migrate(cfg.DB.MigrationsFolderPath)
	if err != nil {
		logrus.WithError(err).Fatal("Could not migrate database")
	}

	server := web.NewServer(nil, cfg.ServerAddress)

	errs := make(chan error)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			errs <- fmt.Errorf("http-server failed to listen: %w", err)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	hasErr := false
	select {
	case sig := <-signals:
		logrus.WithField("signal", sig).Info("Service interrupted")
	case err := <-errs:
		hasErr = true
		logrus.WithError(err).Error("An error occurred")
	}

	shutdownCTX, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	err = server.Shutdown(shutdownCTX)
	if err != nil {
		logrus.WithError(err).Warn("Failed to shutdown http-server")
	}

	cancel()

	if hasErr {
		os.Exit(1)
	}
}
