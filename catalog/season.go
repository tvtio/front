// Copyright 2016 The front Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package catalog

import (
	"fmt"

	"github.com/repejota/cache"
	"github.com/repejota/logger"
	"github.com/tvtio/tmdb"
)

// Season ...
func Season(id string, snumber string) (result tmdb.Season, err error) {
	l := logger.New("default")

	c, err := cache.NewCache(CachePath)
	if err != nil {
		l.Errorf(err.Error())
	}
	key := c.CreateKey(fmt.Sprintf("tv-%s-season-%s", id, snumber))

	// Check if it is cached
	if c.IsCached(key) {
		err := c.Load(key, &result)
		return result, err
	}

	tmdb := tmdb.NewTMDB()
	result, err = tmdb.GetSeason(id, snumber)
	if err != nil {
		return result, err
	}
	go c.Save(key, result)

	return result, err
}
