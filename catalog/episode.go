// Copyright 2016 The front Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package catalog

import (
	"github.com/repejota/cache"
	"github.com/repejota/logger"
	"github.com/tvtio/tmdb"
)

// Episode ...
func Episode(id string, snumber string, enumber string) (result tmdb.Episode, err error) {
	l := logger.New("default")

	c, err := cache.NewCache(CachePath)
	if err != nil {
		l.Errorf(err.Error())
	}
	key := c.CreateKey("tv-" + id + "-season-" + snumber + "-episode-" + enumber)

	// Check if it is cached
	if c.IsCached(key) {
		err := c.Load(key, &result)
		return result, err
	}

	tmdb := tmdb.NewTMDB()
	result, err = tmdb.GetEpisode(id, snumber, enumber)
	if err != nil {
		return result, err
	}
	go c.Save(key, result)

	return result, err
}
