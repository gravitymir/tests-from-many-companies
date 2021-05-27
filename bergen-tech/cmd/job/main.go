package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gravitymir/tests-from-many-companies/bergen-tech/config"

	"github.com/gravitymir/tests-from-many-companies/bergen-tech/internal/httphandlers"
	postgres "github.com/gravitymir/tests-from-many-companies/bergen-tech/internal/pg"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetLevel(log.Level(config.Get().LogLevel)) //set log level

	if err := postgres.Init(); err != nil { //Postgres initialisation
		log.Error(err)
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func run() error {
	http.HandleFunc("/job", httphandlers.Job())
	http.HandleFunc("/task", httphandlers.Task())

	log.Info("Bergen-tech server is running on http://" + config.Get().HTTPHost + config.Get().HTTPPort + "/")

	go func() {
		log.Warn(http.ListenAndServe(config.Get().HTTPPort, nil))
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	fmt.Println("closing")

	return nil
}
