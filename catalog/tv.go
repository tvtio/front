// Copyright 2016 The front Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package catalog

import (
	"fmt"

	"github.com/repejota/cache"
	"github.com/tvtio/tmdb"
)

// TV ...
func TV(id string) (result tmdb.TV, err error) {
	c, err := cache.NewCache(CachePath)
	if err != nil {
		return result, err
	}
	key := c.CreateKey(fmt.Sprintf("tv-%s", id))

	// Check if it is cached
	if c.IsCached(key) {
		err := c.Load(key, &result)
		return result, err
	}

	tmdb := tmdb.New()
	result, err = tmdb.GetTV(id)
	if err != nil {
		return result, err
	}
	go c.Save(key, result)

	return result, err
}

// PopularTV ...
func PopularTV() (result tmdb.SearchTVResult, err error) {
	c, err := cache.NewCache(CachePath)
	if err != nil {
		return result, err
	}
	key := c.CreateKey("tv-popular")

	// Check if it is cached
	if c.IsCached(key) {
		err := c.Load(key, &result)
		return result, err
	}

	tmdb := tmdb.New()
	result, err = tmdb.PopularTV()
	if err != nil {
		return result, err
	}
	go c.Save(key, result)

	return result, err
}
