package main

import (
	"net/http"
	"os"

	"github.com/snghnaveen/url-shortener/pkg/shortener"
	"github.com/snghnaveen/url-shortener/routers"
	"github.com/snghnaveen/url-shortener/util"
	"go.uber.org/zap"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		panic(err.Error())
	}

	defer util.CloseLoggerOnAppExit()

	// prepare some dummy data for metrics api
	if err := shortener.ForTestCreateTestingData(); err != nil {
		util.Logger().Error("failed to feed some mock data to", zap.Error(err))
		os.Exit(1)
	}

	// configure server
	server := &http.Server{
		Addr:    ":" + config.AppPort,
		Handler: routers.InitRouter(),
	}

	server.ListenAndServe()

	// start server
	if err := server.ListenAndServe(); err != nil {
		util.Logger().Error("failed to run server", zap.Error(err))
		os.Exit(1)
	}
}
