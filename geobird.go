/* The G - O Bird: a bird's-eye look at the geoid.

Geobird is a tool for batch-downloading images from Nasa's earthdata.
*/

package main

import (
	"fmt"
	"net/url"
	"os"
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
	// dryErr := errors.New("we're only pretending")
	var exists bool
	for img := range iChannel {

		// arrange URL and file paths
		iURL = img.getImageURL()
		iDir, iName := img.getImagePath()

		// Check if the file already exists
		_, err = os.Stat(iDir + iName)
		exists = !os.IsNotExist(err)

		// verbosity
		switch options.verbosity {
		case 1:
			if !exists {
				fmt.Println("Producing", iDir+iName)
			}
		case 2:
			fmt.Println("getting", iURL.String())
			fmt.Println("and putting at", iDir+iName)
			if exists {
				fmt.Println("...but it's already there")
			}
			if options.dryRun {
				fmt.Println("...but we're only pretending")
			}
		}

		// If it's not already there, and we're not a dry-run,
		// do that funky stuff!
		if !exists && !options.dryRun {
			err = getAndSaveImage(iURL.String(), iDir, iName)
			if err != nil {
				fmt.Println("however...", err, "\n")
			}
		}
	}

	if options.verbosity > 0 {
		fmt.Println("birdy!")
	}
}
