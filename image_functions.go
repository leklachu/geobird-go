package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
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

// Merging functions to one: because the deferred HTTP close closes the stream
// before a new function can access it atm, and these functions are always
// going to be piped together
func getAndSaveImage(url string, filePath string, fileName string) error {

	var err error
	outputPath := filePath + fileName
	// 0. Check if the file already exists (moved to main)

	// 1. Open the HTTPS stream to get the image
	var resp *http.Response
	resp, err = http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 2. Get the body ready to write
	var imgBytes []byte
	imgBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// 3a. Make the directory if not already there
	err = os.MkdirAll(filePath, 0755)
	if err != nil {
		return err
	}
	// 3b. write the file
	err = ioutil.WriteFile(outputPath, imgBytes, 0644)
	if err != nil {
		return err
	}

	return nil

}
