package catalog

import (
	"encoding/json"

	"github.com/tvtio/front/cache"
	"github.com/tvtio/front/tmdb"
)

// Movie ...
func Movie(id string) (result tmdb.Movie, err error) {
	// Get a hash
	hash := cache.Hash("movie-" + id)

	// Check if it is cached
	if cache.IsCached(CachePath, hash) {

		// Get the cached result
		data, err := cache.Get(CachePath, hash)
		if err != nil {
			return result, err
		}
		err = json.Unmarshal(data, &result)
		return result, err
	}

	// Query to the backend
	result, err = tmdb.GetMovie(id)
	if err != nil {
		return result, err
	}

	// Cache the result
	json, err := json.Marshal(result)
	if err != nil {
		return result, err
	}
	err = cache.Save(CachePath, hash, string(json))

	return result, err
}

// SearchMovies ...
func SearchMovies(query string) (result tmdb.SearchMovieResult, err error) {
	// Get a hash
	hash := cache.Hash("movie-search-" + query)

	// Check if it is cached
	if cache.IsCached(CachePath, hash) {

		// Get the cached result
		data, err := cache.Get(CachePath, hash)
		if err != nil {
			return result, err
		}
		err = json.Unmarshal(data, &result)
		return result, err
	}

	// Query to the backend
	result, err = tmdb.SearchMovie(query)
	if err != nil {
		return result, err
	}

	// Cache the result
	json, err := json.Marshal(result)
	if err != nil {
		return result, err
	}
	err = cache.Save(CachePath, hash, string(json))

	return result, err
}

// PopularMovies ...
func PopularMovies() (result tmdb.SearchMovieResult, err error) {
	// Get a hash
	hash := cache.Hash("movie-popular")

	// Check if it is cached
	if cache.IsCached(CachePath, hash) {
		// Get the cached result
		data, err := cache.Get(CachePath, hash)
		if err != nil {
			return result, err
		}
		err = json.Unmarshal(data, &result)
		return result, err
	}

	// Query to the backend
	result, err = tmdb.PopularMovie()
	if err != nil {
		return result, err
	}

	// Cache the result
	json, err := json.Marshal(result)
	if err != nil {
		return result, err
	}
	err = cache.Save(CachePath, hash, string(json))
	return result, err
}
