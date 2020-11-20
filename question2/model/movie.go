package model

type Movies struct {
	Title    string `json:"Title"`
	Year     string `json:"Year"`
	ImdbID   string `json:"imdbID"`
	Type     string `json:"Type"`
	Poster   string `json:"Poster"`
	Rated    string `json:"Rated"`
	Released string `json:"Released"`
	Runtime  string `json:"Runtime"`
	Error    string `json:"Error"`
}

type ResponseMovies struct {
	Data         []Movies `json:"Search"`
	TotalResults string   `json:"totalResults"`
	Response     string   `json:"Response"`
	Error        string   `json:"Error"`
}
