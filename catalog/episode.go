package catalog

import (
	"encoding/json"

	"github.com/tvtio/front/cache"
	"github.com/tvtio/front/tmdb"
)

// Episode ...
func Episode(id string, snumber string, enumber string) (result tmdb.Episode, err error) {
	// Get a hash
	hash := cache.Hash("tv-" + id + "-season-" + snumber + "-episode-" + enumber)

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
	result, err = tmdb.GetEpisode(id, snumber, enumber)
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
