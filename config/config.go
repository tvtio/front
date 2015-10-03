package config

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

// Configuration ...
type Configuration struct {
	Host      string `json:"host"`
	Port      int    `json:"port"`
	Templates string `json:"templates"`
	Cache     string `json:"cache"`
}

// ServerAddress gets server address
func (c *Configuration) ServerAddress() string {
	return c.Host + ":" + strconv.Itoa(c.Port)
}

// ServerURL gets the server URL
func (c *Configuration) ServerURL() string {
	return "http://" + c.ServerAddress()
}

// Load loads configuration onfiguration
func Load(filename string) (configuration Configuration, err error) {
	data, err := ioutil.ReadFile(filename)
	err = json.Unmarshal(data, &configuration)
	return configuration, err
}
