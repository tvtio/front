package catalog

import (
	"encoding/json"

	"github.com/tvtio/front/cache"
	"github.com/tvtio/front/tmdb"
)

// TV ...
func TV(id string) (result tmdb.TV, err error) {
	// Get a hash
	hash := cache.Hash("tv-" + id)

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
	result, err = tmdb.GetTV(id)
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

// PopularTV ...
func PopularTV() (result tmdb.SearchTVResult, err error) {
	// Get a hash
	hash := cache.Hash("tv-popular")

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
	result, err = tmdb.PopularTV()
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
