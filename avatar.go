package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"strings"
)

// ErrNoAvater is the error that is returned when the avater instance is unable
// to provide a avater URL.
var ErrNoAvaterURL = errors.New("chat: Unable to get an avatar URL")

// ErrNoEmailAdrrs is the error that is returned when the avater instance is unable
// to provide a email URL.
var ErrNoEmailAdrrs = errors.New("chat: Unable to get an email address")

// Avatar represents types capable of representing user profile pictures
type Avater interface {
	//GetAvaterURL gets the avater for the speciefied client or returns error
	//ErrNoAvater is returned if the object is unable to get a url
	GetAvaterURL(c *client) (string, error)
}

type AuthAvatar struct{}
type GravatarAvatar struct{}

var UseAuthAvater AuthAvatar
var UseGravater GravatarAvatar

func (AuthAvatar) GetAvaterURL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}
	return "", ErrNoAvaterURL
}

func (GravatarAvatar) GetAvaterURL(c *client) (string, error) {
	if email, ok := c.userData["email"]; ok {
		if emailStr, ok := email.(string); ok {
			m := md5.New()
			io.WriteString(m, strings.ToLower(emailStr))
			return fmt.Sprintf("//www.gravatar.com/avatar/%x", m.Sum(nil)), nil
		}
	}
	return "", ErrNoEmailAdrrs
}
