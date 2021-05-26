package main

import (
	"net/http"

	"github.com/gravitymir/tests-from-many-companies/bergen-tech/config"

	"github.com/gravitymir/tests-from-many-companies/bergen-tech/internal/httphandlers"
	log "github.com/sirupsen/logrus"
)

func run() error {

	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetLevel(log.Level(config.Get().LogLevel)) //set log level

	go http.HandleFunc("/job", httphandlers.Job())
	go http.HandleFunc("/task", httphandlers.Task())

	//fmt.Println("\u001b[34mStarted on \u001b[36mhttp://" + config.Get().HTTPAddr + "/\u001b[0m")
	log.Info("Server is UP on \u001b[36mhttp://" + config.Get().HTTPAddr + "/")

	log.Fatal(http.ListenAndServe(config.Get().HTTPAddr, nil))

	return nil
}

func main() {
	if err := run(); err != nil {
		//fmt.Println(err)
	}
}
