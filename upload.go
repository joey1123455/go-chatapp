package main

import (
	"io"
	"net/http"
	"os"
	"path"
)

func uploadHandler(w http.ResponseWriter, req *http.Request) {
	userId := req.FormValue("userid")
	file, header, err := req.FormFile("avatarFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	baseDir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fileName := path.Join(baseDir, "avatars", userId+path.Ext(header.Filename))
	f, err := os.Create(fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = os.WriteFile(fileName, data, 0777)
	// err = os.Write(data, 0777)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, "Successful")
	defer f.Close()
}
