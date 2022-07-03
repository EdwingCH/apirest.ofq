package main

import (
	"flag"
	"net/http"
	"os"
	"time"

	"apirest.ofq/controller"
	"github.com/google/logger"
	"github.com/gorilla/mux"
)

const logPath = "apirest.ofq.log"

var verbose = flag.Bool("verbose", false, "print info level logs to stdout")

func main() {
	flag.Parse()
	lf, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		logger.Fatalf("Failed to open log file: %v", err)
	}
	defer lf.Close()

	defer logger.Init("LoggerExample", *verbose, true, lf).Close()

	routerMux := mux.NewRouter().StrictSlash(false)
	routerMux.HandleFunc("/topsecret/", controller.PostTopSecretHandler).Methods("POST")
	routerMux.HandleFunc("/topsecret_split/{satellite_name}", controller.GetTopSecretSplitHandler).Methods("GET")
	routerMux.HandleFunc("/topsecret_split/{satellite_name}", controller.PostTopSecretSplitHandler).Methods("POST")

	server := &http.Server{ //Crea un server parametrizado
		Addr:           ":8080",
		Handler:        routerMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	//http.ListenAndServe(":8080", routerMux) //crea un servidor default, no parametrizado
	logger.Info("Run Server localhost:8080 ...")
	server.ListenAndServe()
}
