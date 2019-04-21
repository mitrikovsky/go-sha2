package main

import (
	"./api"
	"./core"
	"./db"
	"github.com/op/go-logging"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
)

var log = logging.MustGetLogger("main")

func main() {
	setupCloseHandler()

	// start http api
	go api.Start()

	log.Info("All done")

	// don't stop keep the moving
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}

func init() {
	// set logs
	format := logging.MustStringFormatter("SHA2.%{module}.%{shortfile}.%{shortfunc}() > %{level:.8s} : %{message}")
	b := logging.NewBackendFormatter(logging.NewLogBackend(os.Stderr, "", 0), format)
	logging.SetBackend(b)

	log.Info("Logs ok")
	log.Infof("GOMAXPROCS: %v", runtime.GOMAXPROCS(0))
}

func setupCloseHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Info("program interrupt: waiting for active jobs to finish...")
		core.WG.Wait()
		db.Close()
		os.Exit(0)
	}()
}
