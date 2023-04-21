package main

import (
	"log"

	"github.com/snghnaveen/url-shortner/db"
	"github.com/snghnaveen/url-shortner/util"
)

func main() {
	_, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err, "cannot load config")
	}

	db.Tmp()

	defer util.CloseLoggerOnAppExit()
}
