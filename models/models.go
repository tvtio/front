package models

import (
	"encoding/json"

	"github.com/tvtio/front/cache"
)

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
