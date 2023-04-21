package main

import (
	"log"
	"net/http"
	"os"

	"github.com/snghnaveen/url-shortner/routers"
	"github.com/snghnaveen/url-shortner/util"
	"go.uber.org/zap"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err, "cannot load config")
	}

	defer util.CloseLoggerOnAppExit()

	server := &http.Server{
		Addr:    ":" + config.AppPort,
		Handler: routers.InitRouter(),
	}

	server.ListenAndServe()

	if err := server.ListenAndServe(); err != nil {
		util.Logger().Error("failed to run server", zap.Error(err))
		os.Exit(1)
	}
}
