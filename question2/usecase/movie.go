package usecase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"stockbit_technical_test/question2/model"

	"github.com/gorilla/mux"
)

const (
	BASEURL = "http://www.omdbapi.com/"
	APIKEY  = "faf7e5bb"
)

type responseMovies model.ResponseMovies
type movies model.Movies

func (res responseMovies) GetMovies(param1, param2 string) ([]model.Movies, string) {
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
		return nil, res.Error
	}

	movies := []model.Movies{}
	for _, val := range res.Data {
		p := model.Movies{
			Title:  val.Title,
			Year:   val.Year,
			ImdbID: val.ImdbID,
			Type:   val.Type,
			Poster: val.Poster,
		}
		movies = append(movies, p)
	}

	return movies, res.Error
}

func (res movies) GetMovieByID(param1 string) (model.Movies, string) {
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

	movies := model.Movies{
		Title:    res.Title,
		Year:     res.Year,
		ImdbID:   res.ImdbID,
		Type:     res.Type,
		Poster:   res.Poster,
		Rated:    res.Rated,
		Released: res.Released,
		Runtime:  res.Runtime,
	}

	return movies, res.Error
}

// function router handler
func HandlerMovies(w http.ResponseWriter, r *http.Request) {
	var (
		params = mux.Vars(r)
		param1 = params["keyword"]
		param2 = params["page"]
		status = http.StatusOK
	)
	api := &responseMovies{}
	movies, err := api.GetMovies(param1, param2)

	if err != "" {
		// if response return False
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  status,
			"message": err,
		})
		fmt.Println("💔💔💔 API [GetMovies] Succesfully load with error. errorMessage : ", err)
	} else {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  status,
			"message": movies,
		})
		fmt.Println("💚💚💚 API [GetMovies] Succesfully load . status code : ", status)
	}
}

func HandlerMovieByID(w http.ResponseWriter, r *http.Request) {
	var (
		params = mux.Vars(r)
		param1 = params["id"]
		status = http.StatusOK
	)
	api := &movies{}
	movie, err := api.GetMovieByID(param1)
	if err != "" {
		// if response return False
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  status,
			"message": err,
		})
		fmt.Println("💔💔💔 API [GetMovieByID] Succesfully load with error. errorMessage : ", err)
	} else {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  status,
			"message": movie,
		})
		fmt.Println("💚💚💚 API [GetMovieByID] Succesfully load . status code : ", status)
	}

}
