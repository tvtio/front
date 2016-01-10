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

// Movie ...
func Movie(id string) (result tmdb.Movie, err error) {
	l := logger.New("default")

	c, err := cache.NewCache(CachePath)
	if err != nil {
		l.Errorf(err.Error())
	}
	key := c.CreateKey(fmt.Sprintf("movie-%s", id))

	// Check if it is cached
	if c.IsCached(key) {
		err := c.Load(key, &result)
		return result, err
	}

	tmdb := tmdb.NewTMDB()
	result, err = tmdb.GetMovie(id)
	if err != nil {
		return result, err
	}
	go c.Save(key, result)

	return result, err
}

// SearchMovies ...
func SearchMovies(query string) (result tmdb.SearchMovieResult, err error) {
	l := logger.New("default")

	c, err := cache.NewCache(CachePath)
	if err != nil {
		l.Errorf(err.Error())
	}
	key := c.CreateKey(fmt.Sprintf("movie-search-%s", query))

	// Check if it is cached
	if c.IsCached(key) {
		err := c.Load(key, &result)
		return result, err
	}

	tmdb := tmdb.NewTMDB()
	result, err = tmdb.SearchMovie(query)
	if err != nil {
		return result, err
	}
	go c.Save(key, result)

	return result, err
}

// PopularMovies ...
func PopularMovies() (result tmdb.SearchMovieResult, err error) {
	l := logger.New("default")

	c, err := cache.NewCache(CachePath)
	if err != nil {
		l.Errorf(err.Error())
	}
	key := c.CreateKey("movie-popular")

	// Check if it is cached
	if c.IsCached(key) {
		err := c.Load(key, &result)
		return result, err
	}

	tmdb := tmdb.NewTMDB()
	result, err = tmdb.PopularMovie()
	if err != nil {
		return result, err
	}
	go c.Save(key, result)

	return result, err
}
