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
	// fmt.Println(options)

	// Construct urls to fetch, and paths to save
	imageSet := defaultImageSet
	imageSet.fileSchema = DefaultScheme{options.outputDir}
	imageSet.dates = dateRange(Date{2018, 1, 1}, Date{2018, 1, 31}, Period{0, 0, 1})

	// Fetch urls and save as files
	iChannel := make(chan Image)
	go imageSet.getImages(iChannel)
	var iURL url.URL
	for img := range iChannel {
		iURL = img.getImageURL()
		fmt.Println("getting", iURL.String())
		iDir, iName := img.getImagePath()
		fmt.Println("and putting at", iDir+iName)
		getAndSaveImage(iURL.String(), iDir, iName)
	}

	// u := i1.getImageURL()
	// fmt.Println(u.String())
	// f := i1.getImagePath()
	// getAndSaveImage(u.String(), f)

	// d1 := Date{2012, 2, 1}
	// d2 := Date{2018, 3, 2}
	// p1 := Period{0, 0, 365}
	// fmt.Println(dateRange(d1, d2, p1))
	// fmt.Println(show(d1))

	// s1 := DefaultScheme{"./"}
	// fmt.Println(s1.makeFilePath(i1))
	// fmt.Println(s1.makeFileName(i1))

	fmt.Println("birdy!")
}
