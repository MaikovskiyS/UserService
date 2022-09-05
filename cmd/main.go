package main

import (
	"practice/internal/controller"
	"practice/internal/repository"
	"practice/internal/service"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	connstr string = "root:Wild54323@tcp(127.0.0.1:3306)/world"
)

func main() {

	log := logrus.New()
	repo, err := repository.NewRepository(connstr, log, time.Second*30)
	if err != nil {
		log.Info("cant connect to sql", err)
		return
	}
	svc := service.NewService(log, repo)
	tr := controller.NewController(log, svc)
	tr = tr.CreateEndPoints()
	tr.Start()
}
