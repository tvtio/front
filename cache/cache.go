package cache

import (
	"fmt"
	"hash/fnv"
	"log"
	"os"
	"path/filepath"

	"github.com/mgutz/ansi"
	"github.com/repejota/kvson"
)

// Hash returns the hash for an string
func Hash(input string) (hash string) {
	h := fnv.New32a()
	h.Write([]byte(input))
	return fmt.Sprint(h.Sum32())
}

// IsCached check if a hash is cached on a path
func IsCached(path string, hash string) bool {
	filename := filepath.Join(path, hash)
	if _, err := os.Stat(filename); err == nil {
		log.Println(ansi.Color("DEBUG: ", "blue"), "Caché exists : "+filename)
		return true
	}
	log.Println(ansi.Color("DEBUG: ", "blue"), "Not cached : "+filename)
	return false
}

// Save saves the content to the caché
func Save(path string, hash string, payload string) (err error) {
	el := kvson.Element{
		ID:      hash,
		Payload: []byte(payload),
	}
	err = el.Save(path)
	if err != nil {
		return err
	}
	return nil
}

// Get gets the content from the caché
func Get(path string, hash string) (data []byte, err error) {
	el := kvson.Element{}
	filename := filepath.Join(path, hash)
	el, err = el.Get(filename)
	if err != nil {
		return data, err
	}
	return el.Payload, nil
}
