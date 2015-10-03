package catalog

import (
	"encoding/json"
	"fmt"
	"hash/fnv"

	"github.com/tvtio/front/cache"
	"github.com/tvtio/front/tmdb"
)

func SearchMovies(query string) (result tmdb.SearchMovieResult, err error) {
	h := fnv.New32a()
	h.Write([]byte(query))
	sum := h.Sum32()
	str := fmt.Sprint(sum)
	filename := "/tmp/tvtio/" + str
	if cache.IsCached(str) {
		data, err := cache.Get(filename)
		err = json.Unmarshal(data, &result)
		return result, err
	}
	result, err = tmdb.SearchMovie(query)
	if err != nil {
		return result, err
	}
	json, err := json.Marshal(result)
	if err != nil {
		return result, err
	}
	_, err = cache.Save(filename, string(json))
	return result, err
}
