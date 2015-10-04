package tmdb

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	apiKey  = "523587afe262c34af9ee7794c5f8de81"
	baseURL = "http://api.themoviedb.org/3/"
)

// PopularMovie ...
func PopularMovie() (result SearchMovieResult, err error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", baseURL+"movie/popular?api_key="+apiKey, nil)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	return result, nil
}

// PopularTV ...
func PopularTV() (result SearchTVResult, err error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", baseURL+"tv/popular?api_key="+apiKey, nil)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	return result, nil
}

// SearchMulti ...
func SearchMulti(query string) (result SearchMultiResult, err error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", baseURL+"search/multi?api_key="+apiKey+"&query="+query, nil)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	return result, nil
}

// SearchMovie ...
func SearchMovie(query string) (result SearchMovieResult, err error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", baseURL+"search/movie?api_key="+apiKey+"&query="+query, nil)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	return result, nil
}

// GetMovie ...
func GetMovie(id string) (result Movie, err error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", baseURL+"movie/"+id+"?api_key="+apiKey+"&append_to_response=credits,images", nil)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	return result, nil
}

// GetPerson ...
func GetPerson(id string) (result Person, err error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", baseURL+"person/"+id+"?api_key="+apiKey+"&append_to_response=movie_credits,tv_credits", nil)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	return result, nil
}

// GetTV ...
func GetTV(id string) (result TV, err error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", baseURL+"tv/"+id+"?api_key="+apiKey+"&append_to_response=credits", nil)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	return result, nil
}
