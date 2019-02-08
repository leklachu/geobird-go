/* The G - O Bird: a bird's-eye look at the geoid.

Geobird is a tool for batch-downloading images from Nasa's earthdata.
*/

package main

import (
	"fmt"
	// "net/http"
	"net/url"
)

func main() {
	// Work out the opts and choose images
	options := parseCommandLineArguments()
	imageSet := prepImageSet(options)

	// Construct urls to fetch, and paths to save
	iChannel := make(chan Image)
	go imageSet.getImages(iChannel)

	// Fetch urls and save as files
	var iURL url.URL
	var err error
	for img := range iChannel {
		iURL = img.getImageURL()
		// fmt.Println("getting", iURL.String())
		iDir, iName := img.getImagePath()
		fmt.Println("and putting at", iDir+iName)
		err = getAndSaveImage(iURL.String(), iDir, iName)
		if err != nil {
			fmt.Println("however...", err, "\n")
		}
	}

	fmt.Println("birdy!")
}
