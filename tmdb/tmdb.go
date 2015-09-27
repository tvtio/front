package tmdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// SearchMulti ...
func SearchMulti(query string) (result SearchMultiResult, err error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://api.themoviedb.org/3/search/multi?api_key=523587afe262c34af9ee7794c5f8de81&query="+query, nil)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return result, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return result, err
	}
	return result, nil
}

// SearchMovie ...
func SearchMovie(query string) (result SearchMovieResult, err error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://api.themoviedb.org/3/search/movie?api_key=523587afe262c34af9ee7794c5f8de81&query="+query, nil)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return result, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return result, err
	}
	return result, nil
}
