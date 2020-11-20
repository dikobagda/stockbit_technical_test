package repository

import "stockbit_technical_test/question2/model"

type Service interface {
	GetMovies(param1, param2 string) (string, []model.Movies)
	GetMovieByID(param1 string) (model.Movies, string)
}
