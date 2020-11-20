package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"

	"github.com/gorilla/mux"
)

const (
	BASEURL = "http://www.omdbapi.com/"
	APIKEY  = "faf7e5bb"
)

var (
	address = ":9000"
)

type Repository interface {
	GetMovies(param1, param2 string)
	GetMovieByID(param1 string)
}

type Movies struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
	Error  string `json:"Error"`
}

type ResponseMovies struct {
	Data         []Movies `json:"Search"`
	TotalResults string   `json:"totalResults"`
	Response     string   `json:"Response"`
	Error        string   `json:"Error"`
}

func (res ResponseMovies) GetMovies(param1, param2 string) (string, []Movies) {
	var (
		url = BASEURL + "?apikey=" + APIKEY + "&s=" + param1 + "&page=" + param2
	)
	// call request to omdbAPI
	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)

	e := json.Unmarshal(data, &res)
	if e != nil {
		log.Fatal(err)
	}

	if res.Error != "" {
		return res.Error, nil
	}

	movies := []Movies{}
	for _, val := range res.Data {
		p := Movies{
			Title:  val.Title,
			Year:   val.Year,
			ImdbID: val.ImdbID,
			Type:   val.Type,
			Poster: val.Poster,
		}
		movies = append(movies, p)
	}

	return res.Error, movies
}

func (res Movies) GetMovieByID(param1 string) (Movies, string) {
	var (
		url = BASEURL + "?apikey=" + APIKEY + "&i=" + param1
	)

	// call request to omdbAPI to get single detail of movie
	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)

	e := json.Unmarshal(data, &res)
	if e != nil {
		log.Fatal(err)
	}

	movies := Movies{
		Title:  res.Title,
		Year:   res.Year,
		ImdbID: res.ImdbID,
		Type:   res.Type,
		Poster: res.Poster,
	}

	return movies, res.Error
}

func handlermMovies(w http.ResponseWriter, r *http.Request) {
	var (
		params = mux.Vars(r)
		param1 = params["keyword"]
		param2 = params["page"]
		status = http.StatusOK
	)
	api := &ResponseMovies{}
	movies, err := api.GetMovies(param1, param2)

	if err != nil {
		// if response return False
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  status,
			"message": err,
		})
	} else {

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  status,
			"message": movies,
		})
	}
}

func handlermMovieByID(w http.ResponseWriter, r *http.Request) {
	var (
		params = mux.Vars(r)
		param1 = params["id"]
		status = http.StatusOK
	)
	api := &Movies{}
	movie, err := api.GetMovieByID(param1)
	if err != "" {
		// if response return False
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  status,
			"message": err,
		})
	} else {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  status,
			"message": movie,
		})
	}

}

func main() {
	runtime.GOMAXPROCS(2)

	router := setupServer()

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

func setupServer() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies/{keyword}", handlermMovies).Methods("GET")
	router.HandleFunc("/api/movies/{keyword}/page/{page}", handlermMovies).Methods("GET")
	router.HandleFunc("/api/movie/{id}", handlermMovieByID).Methods("GET")

	return router
}
