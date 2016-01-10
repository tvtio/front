// Copyright 2016 The front Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package catalog

import (
	"log"

	"github.com/repejota/cache"
	"github.com/tvtio/tmdb"
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

	tmdb := tmdb.NewTMDB()
	result, err = tmdb.GetEpisode(id, snumber, enumber)
	if err != nil {
		return result, err
	}
	err = c.Save(key, result)

	return result, err
}
