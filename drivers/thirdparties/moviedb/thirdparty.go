package moviedb

import (
	"context"
	"encoding/json"
	"net/http"
	"ticketing/business/moviedb"
)

type FetchMovies struct {
	httpClient http.Client
}

func NewFetchMovies() moviedb.Repository {
	return &FetchMovies{
		httpClient: http.Client{},
	}
}

func (fm *FetchMovies) GetMovies(ctx context.Context, search string) ([]moviedb.Domain, error) {
	req, _ := http.NewRequest("GET", "https://api.themoviedb.org/3/search/movie?api_key=824a63eb0b94410090afccb2a7398fac&query="+ search +"&page=1", nil)
	resp, err := fm.httpClient.Do(req)
	if err != nil {
		return []moviedb.Domain{}, err
	}

	defer resp.Body.Close()

	var data Response
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return []moviedb.Domain{}, err
	}
	respDomain := toDomain(data)
	return respDomain, nil
}
