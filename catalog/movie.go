package catalog

import (
	"encoding/json"

	"github.com/tvtio/front/cache"
	"github.com/tvtio/front/tmdb"
)

// TODO read path from config.json
const path = "/Users/raul/Projects/tvt.io/src/github.com/tvtio/front/.cache"

// SearchMovies ...
func SearchMovies(query string) (result tmdb.SearchMovieResult, err error) {
	// Get a hash
	hash := cache.Hash(query)

	// Check if it is cached
	if cache.IsCached(path, hash) {

		// Get the cached result
		data, err := cache.Get(path, hash)
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
	err = cache.Save(path, hash, string(json))

	return result, err
}
