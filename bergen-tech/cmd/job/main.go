package main

import (
	"fmt"
	"net/http"

	"github.com/gravitymir/tests-from-many-companies/bergen-tech/config"

	"github.com/gravitymir/tests-from-many-companies/bergen-tech/internal/handlers"
	log "github.com/sirupsen/logrus"
)

type A struct {
	i   int
	str string
}

func (a *A) Add() *A {
	a.i++
	return a
}
func run() error {

	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetLevel(log.Level(config.Get().LogLevel)) //set log level

	http.HandleFunc("/job/", new(handlers.Handler).Job())

	fmt.Println("http://" + config.Get().HTTPAddr + "/")
	a := new(A)
	fmt.Println(a.Add())
	log.Fatal(http.ListenAndServe(config.Get().HTTPAddr, nil))

	return nil
}

func main() {
	if err := run(); err != nil {
		//fmt.Println(err)
	}
}
