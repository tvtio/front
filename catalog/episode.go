package catalog

import (
	"log"

	"github.com/repejota/cache"
	"github.com/tvtio/front/tmdb"
)

// Episode ...
func Episode(id string, snumber string, enumber string) (result tmdb.Episode, err error) {
	c, err := cache.NewCache(CachePath)
	if err != nil {
		log.Fatal(err)
	}
	key := c.CreateKey("tv-" + id + "-season-" + snumber + "-episode-" + enumber)

	// Check if it is cached
	if c.IsCached(key) {
		err := c.Load(key, &result)
		if err != nil {
			return result, err
		}
		return result, err
	}

	result, err = tmdb.GetEpisode(id, snumber, enumber)
	if err != nil {
		return result, err
	}
	err = c.Save(key, result)

	return result, err
}
