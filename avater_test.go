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

func TestGravaterAvater(t *testing.T) {
	var gravaterAvater GravatarAvatar
	client := new(client)
	client.userData = map[string]interface{}{
		"userid": "0bc83cb571cd1c50ba6f3e8a78ef1346",
	}
	url, err := gravaterAvater.GetAvaterURL(client)
	if err != nil {
		t.Error("GravatarAvater.GetAvaterURL shouldnt return an error")
	}
	if url != "//www.gravatar.com/avatar/0bc83cb571cd1c50ba6f3e8a78ef1346" {
		t.Errorf("GravatarAvater.GetAvaterURL shouldnt return %s, which is incorrect", url)
	}
}
