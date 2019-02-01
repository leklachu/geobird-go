/* The G - O Bird: a bird's-eye look at the geoid.

Geobird is a tool for batch-downloading images from Nasa's earthdata.
*/

package main

import (
	"fmt"
	// "net/http"
)

func main() {
	// Work out the opts and choose images
	// Construct urls to fetch, and paths to save
	// Fetch urls and save as files

	u := i1.getImageURL()
	fmt.Println(u.String())
	f := i1.getImagePath()

	// image := getImage(u.String())
	// saveImage(image, f)
	getAndSaveImage(u.String(), f)

	fmt.Println("birdy!")
}
