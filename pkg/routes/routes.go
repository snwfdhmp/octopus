package routes

import (
	"../handler"
	"github.com/julienschmidt/httprouter"
	"log"
)

func Init() *httprouter.Router {
	r := httprouter.New()

	log.Println("Will listen to:")

	log.Println("GET /tap")
	r.GET("/tap", handler.Tap)
	log.Println("GET /claim/:ltoken/:rtoken/:port")
	r.GET("/claim/:ltoken/:rtoken/:port", handler.Claim)
	log.Println("GET /do/tap/:ip/:port")
	r.GET("/do/tap/:ip/:port", handler.DoTap)

	return r
}
