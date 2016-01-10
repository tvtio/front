package models

import "golang.org/x/oauth2"

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
