package api

import (
	"../config"
	"fmt"
	"github.com/op/go-logging"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttprouter"
)

var log = logging.MustGetLogger("api")

func Start() {
	log.Info("Start API")

	// Instantiate a new router
	r := fasthttprouter.New()

	// set handlers for api endpoints
	r.GET("/", index)
	r.POST("/job", postJob)
	r.GET("/job/:id", getJob)

	server := &fasthttp.Server{
		Handler:            r.Handler,
		MaxRequestBodySize: 1 * 1024 * 1024,
	}

	// start web server
	host := fmt.Sprintf("%s:%v", config.API_HOST, config.API_PORT)
	log.Infof("Run HTTP server on %s", host)
	log.Fatal(server.ListenAndServe(host))
}
