package cache

import (
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"os"
	"path/filepath"
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
		fmt.Println("Caché exists : " + filename)
		return true
	}
	fmt.Println("Not cached : " + filename)
	return false
}

// Save saves the content to the caché
func Save(path string, hash string, payload string) (err error) {
	filename := filepath.Join(path, hash)
	err = ioutil.WriteFile(filename, []byte(payload), 0644)
	if err != nil {
		return err
	}
	return nil
}

// Get gets the content from the caché
func Get(path string, hash string) (data []byte, err error) {
	filename := filepath.Join(path, hash)
	data, err = ioutil.ReadFile(filename)
	if err != nil {
		return data, err
	}
	return data, nil
}
