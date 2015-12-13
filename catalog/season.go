package catalog

import (
	"encoding/json"

	"github.com/tvtio/front/cache"
	"github.com/tvtio/front/tmdb"
)

// Season ...
func Season(id string, snumber string) (result tmdb.Season, err error) {
	// Get a hash
	hash := cache.Hash("tv-" + id + "-season-" + snumber)

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
	result, err = tmdb.GetSeason(id, snumber)
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
