package catalog

import (
	"log"

	"github.com/repejota/cache"
	"github.com/tvtio/front/tmdb"
)

// Person ...
func Person(id string) (result tmdb.Person, err error) {
	c, err := cache.NewCache(CachePath)
	if err != nil {
		log.Fatal(err)
	}
	key := c.CreateKey("person-" + id)

	// Check if it is cached
	if c.IsCached(key) {
		err := c.Load(key, &result)
		if err != nil {
			return result, err
		}
		return result, err
	}

	// Query to the backend
	result, err = tmdb.GetPerson(id)
	if err != nil {
		return result, err
	}
	err = c.Save(key, result)

	return result, err
}
