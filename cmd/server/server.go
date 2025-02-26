package server

import (
	"base-app/cmd/routes"
	"base-app/config"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

func GetServer() *http.Server {
	config.LoadConfig()
	r := routes.InitRouting()

	readTimeout, err := time.ParseDuration(config.AppConfig.ReadTimeout)
	if err != nil {
		panic(fmt.Sprintf("failed parse server read timeout %v", err))
	}

	writeTimeout, err := time.ParseDuration(config.AppConfig.WriteTimeout)
	if err != nil {
		panic(fmt.Sprintf("failed parse server read timeout %v", err))
	}

	server := &http.Server{
		Addr:         config.AppConfig.Port,
		Handler:      r,
		TLSConfig:    &tls.Config{},
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	return server
}
