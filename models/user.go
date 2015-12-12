package models

import (
	"encoding/json"

	"golang.org/x/oauth2"

	"github.com/tvtio/front/cache"
)

const path = "/users"

// User ...
type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Picture struct {
		Data struct {
			URL string `json:"url"`
		} `json:"data"`
	} `json:"picture"`
	Credentials *oauth2.Token `json:"credentials"`
}

// Save user
func (u *User) Save() error {
	b, err := json.Marshal(u)
	if err != nil {
		return err
	}
	return cache.Save(path, u.ID, string(b))
}
