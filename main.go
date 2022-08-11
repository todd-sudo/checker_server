package main

import (
	"context"
	"errors"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/todd-sudo/checker_server/pkg/handler"
	"github.com/todd-sudo/checker_server/pkg/logging"
	"github.com/todd-sudo/checker_server/pkg/server"
)

const (
	timeout  = 10 * time.Second
	logLevel = "trace"
)

func main() {
	log := logging.GetLogger(logLevel)

	portFlag := flag.String("port", "", "server")

	flag.Parse()
	port := *portFlag
	if port == "" {
		log.Errorln("The 'port' flag cannot be empty. For example: -port=8000")
		return
	}

	log.Infoln("Connect logger successfully!")

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	router := mux.NewRouter().StrictSlash(true)
	handler := handler.NewHandler(log)
	handler.InitRoutes(router)

	srv := server.NewServer(port, router)

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			panic("error occurred while running http server: " + err.Error())
		}
	}()
	log.Info("Server started on http://127.0.0.1:" + port)

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	if err := srv.Stop(ctx); err != nil {
		log.Info("Server stopped :(")
		return
	}
	log.Info("Server stopped")
}
