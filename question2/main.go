package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"stockbit_technical_test/question2/model"
	"stockbit_technical_test/question2/usecase"

	"github.com/gorilla/mux"
)

var (
	address = ":9000"
)

type responseMovies model.ResponseMovies
type movies model.Movies

func main() {
	runtime.GOMAXPROCS(2)

	router := setupRouter()

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt, os.Kill)

	srv := &http.Server{Addr: address, Handler: router}
	fmt.Println("The HTTP starting in address " + address + "...")
	go func() {
		if e := srv.ListenAndServe(); e != nil {
			log.Fatal(http.ListenAndServe(address, router))
		}
	}()

	<-stopChan
	log.Fatal("Shutting Down Server...")
}

func setupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies/{keyword}", usecase.HandlerMovies).Methods("GET")
	router.HandleFunc("/api/movies/{keyword}/page/{page}", usecase.HandlerMovies).Methods("GET")
	router.HandleFunc("/api/movie/{id}", usecase.HandlerMovieByID).Methods("GET")

	return router
}
