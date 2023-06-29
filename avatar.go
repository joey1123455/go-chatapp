package main

import (
	"errors"
)

// ErrNoAvater is the error that is returned when the avater instance is unable
// to provide a avater URL.
var ErrNoAvaterURL = errors.New("chat: Unable to get an avatar URL")

// Avatar represents types capable of representing user profile pictures
type Avater interface {
	//GetAvaterURL gets the avater for the speciefied client or returns error
	//ErrNoAvater is returned if the object is unable to get a url
	GetAvaterURL(c *client) (string, error)
}

type AuthAvatar struct{}

var UseAuthAvater AuthAvatar

func (AuthAvatar) GetAvaterURL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}
	return "", ErrNoAvaterURL
}
