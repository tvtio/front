package catalog

import (
	"encoding/json"

	"github.com/tvtio/front/cache"
	"github.com/tvtio/front/tmdb"
)

// Person ...
func Person(id string) (result tmdb.Person, err error) {
	// Get a hash
	hash := cache.Hash("person-" + id)

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
	result, err = tmdb.GetPerson(id)
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
