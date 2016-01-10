// Copyright 2016 The front Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package catalog

import (
	"log"

	"github.com/repejota/cache"
	"github.com/tvtio/tmdb"
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
