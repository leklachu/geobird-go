package main

import (
	"io"
	"io/ioutil"
	"net/http"
)

// Given an image URL, return a pointer to the image
func getImage(URL string) io.ReadCloser {
	resp, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return resp.Body
}

func saveImage(img io.ReadCloser, fileName string) {
	imgBytes, err := ioutil.ReadAll(img)
	if err != nil {
		panic(err)
	}

	outputPath := fileName //filepath.Join(outputDir, fileName)
	err = ioutil.WriteFile(outputPath, imgBytes, 0644)
	if err != nil {
		panic(err)
	}
}
