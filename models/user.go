package models

import (
	"encoding/json"

	"golang.org/x/oauth2"

	"github.com/tvtio/front/cache"
)

const path = "/Users/raul/Projects/tvt.io/src/github.com/tvtio/front/.users"

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

// GetUser a user
func GetUser(id string) (user *User, err error) {
	if id != "<nil>" {
		payload, err := cache.Get(path, id)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(payload, &user)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, err
}
