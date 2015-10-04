package catalog

import (
	"encoding/json"

	"github.com/tvtio/front/cache"
	"github.com/tvtio/front/tmdb"
)

// PopularTV ...
func PopularTV() (result tmdb.SearchTVResult, err error) {
	// Get a hash
	hash := cache.Hash("populartv")

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
	result, err = tmdb.PopularTV()
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
