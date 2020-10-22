package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kmcode-18/task21_10/api"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	startApp()
}

func startApp() {

	r := mux.NewRouter()
	log.Println("server started at port : 8080")
	r.HandleFunc("/image", api.AddImage).Methods("POST")
	r.HandleFunc("/image", api.GetImages).Methods("GET")

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic("error starting server")
	}
	go HandleOSSignals(func() {
		err := api.Stop()
		if err != nil {
			panic(err)
		}
	})
	fmt.Println("closing app")
}

func HandleOSSignals(fn func()) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGUSR1, syscall.SIGINT, syscall.SIGTERM)

	for sig := range signals {
		switch sig {
		case syscall.SIGINT, syscall.SIGUSR1, syscall.SIGTERM:
			fn()
		}
	}
}
