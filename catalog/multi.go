package catalog

import (
	"log"

	"github.com/repejota/cache"
	"github.com/tvtio/front/tmdb"
)

// SearchMulti ...
func SearchMulti(query string) (result tmdb.SearchMultiResult, err error) {
	c, err := cache.NewCache(CachePath)
	if err != nil {
		log.Fatal(err)
	}
	key := c.CreateKey("multi-search-" + query)

	// Check if it is cached
	if c.IsCached(key) {
		err := c.Load(key, &result)
		if err != nil {
			return result, err
		}
		return result, err
	}

	// Query to the backend
	result, err = tmdb.SearchMulti(query)
	if err != nil {
		return result, err
	}
	err = c.Save(key, result)

	return result, err
}
