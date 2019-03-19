package main

import (
	"../routers"
	"github.com/codegangsta/negroni"
	"log"
	"net/http"
	"../common"
)

func main()  {
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	server := &http.Server{
		Addr:		common.AppConfig.Server,
		Handler:	n,
	}
	log.Println("Listening at: " + common.AppConfig.Server)
	server.ListenAndServe()
}