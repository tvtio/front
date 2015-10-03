package cache

import (
	"fmt"
	"io/ioutil"
	"os"
)

// IsCached check if a hash is cached
func IsCached(hash string) bool {
	filename := "/tmp/tvtio/" + hash
	if _, err := os.Stat(filename); err == nil {
		fmt.Println("Caché exists : " + filename)
		return true
	}
	fmt.Println("Not cached : " + filename)
	return false
}

// Save saves the content to the caché
func Save(filename string, payload string) (success bool, err error) {
	err = ioutil.WriteFile(filename, []byte(payload), 0644)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Get gets the content from the caché
func Get(filename string) (data []byte, err error) {
	data, err = ioutil.ReadFile(filename)
	if err != nil {
		return data, err
	}
	return data, nil
}
