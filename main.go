package main

import (
	"GoGinExample/models"
	"GoGinExample/pkg/logging"
	"GoGinExample/pkg/setting"
	"GoGinExample/routers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()

	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Println("server listenServer err", err)
	}
}
