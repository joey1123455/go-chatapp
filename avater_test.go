package main

import (
	"testing"
)

func TesTAuthAvater(t *testing.T) {
	var authAvater AuthAvatar
	testUrl := "http://url-to-gavater"
	client := new(client)
	url, err := authAvater.GetAvaterURL(client)
	if err != nil {
		t.Error("AuthAvater.GetAvaterURL should return ErrNoAvaterURL when no value present")
	}
	if url != testUrl {
		t.Error("AuthAvater.GetAvaterURL should return correct URL")
	}
}
