package main

import (
	"github.com/gabrielbo1/b2wtest/api"
	"github.com/gabrielbo1/b2wtest/config"
	"github.com/gorilla/handlers"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	"net/http"
)

func init() {
	// Only log the warning severity or above.
	log.SetLevel(log.FatalLevel)

	//Flag parsing
	config.FlagParse()
}

func main() {
	router := api.NewRouter()
	router.Use(handlers.CompressHandler)
	http.Handle("/", router)

	port := config.EnvVal(config.Port)
	n := negroni.Classic()
	n.UseHandler(cors.AllowAll().Handler(router))
	n.Run(":" + port)
}
