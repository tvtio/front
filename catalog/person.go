// Copyright 2016 The front Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package catalog

import (
	"fmt"

	"github.com/repejota/cache"
	"github.com/tvtio/tmdb"
)

// Person ...
func Person(id string) (result tmdb.Person, err error) {
	c, err := cache.NewCache(CachePath)
	if err != nil {
		return result, err
	}
	key := c.CreateKey(fmt.Sprintf("person-%s", id))

	// Check if it is cached
	if c.IsCached(key) {
		err := c.Load(key, &result)
		return result, err
	}

	tmdb := tmdb.New()
	result, err = tmdb.GetPerson(id)
	if err != nil {
		return result, err
	}
	go c.Save(key, result)

	return result, err
}
