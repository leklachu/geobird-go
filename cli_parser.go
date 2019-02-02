// Parses the commandline arguments and builds an ImageSet (metadata)
package main

import (
	"path/filepath"
	"strings"

	"github.com/spf13/pflag"
)

type OptionGroup struct {
	outputDir       string
	satelliteLayers string
	// dates           string
	startDate  string
	endDate    string
	interval   string
	satellites string
	// latMin, latMax, lonMin, lonMax string
	coordinates   string // latMin,lonMin,latMax,lonMax: as per earthview
	width, height string
}

func parseCommandLineArguments() *OptionGroup {

	options := OptionGroup{} // do this as a global or return it?

	pflag.StringVarP(&options.outputDir, "output-dir", "o", ".", "The directory to output the image files into")
	pflag.Parse()

	// change relative paths to absolute ones
	options.outputDir, _ = filepath.Abs(options.outputDir)
	return &options
}

func prepImageSet(o *OptionGroup) *ImageSet {
	imageSet := defaultImageSet
	if o.outputDir != "" {
		imageSet.fileSchema = DefaultScheme{o.outputDir}
	}
	if o.startDate != "" && o.endDate != "" {
		if o.interval == "" {
			o.interval = "0-0-1"
		}
		imageSet.dates = dateRange(
			readDate(o.startDate), readDate(o.endDate), readPeriod(o.interval))
	}
	if o.coordinates != "" {
		imageSet.latMin, imageSet.lonMin, imageSet.latMax, imageSet.lonMax =
			readCoordinates(o.coordinates)
	}
	// imagetype
	// w+h (would be cool to calculate this!)
	// other?
	return &imageSet
}

// Split a "x1,y1,x2,y2" into "x1","y1","x2","y2"
func readCoordinates(acme string) (a, c, m, e string) {
	ss := strings.Split(acme, ",")
	if len(ss) != 4 {
		panic("wrong number of coordinates")
	}
	return ss[0], ss[1], ss[2], ss[3]
}
