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
	// TODO how to denote multilayer within image? currently each layer is
	// separate image
	// TODO have rawlayer; also easy layer selection
	startDate string
	endDate   string
	interval  string
	// TODO want option to do specific dates regardless of regularity
	satellites  string
	coordinates string // latMin,lonMin,latMax,lonMax: as per earthview
	size        string
	imageType   string
	// TODO would be neat to be able to take a minimal set of
	// lat+long+size[+res] or lat/long*2[+res] and calculate
	schema string

	// TODO also need a dry-run option
}

// Parse the command line and put details into an OptionGroup
func parseCommandLineArguments() *OptionGroup {

	options := OptionGroup{} // do this as a global or return it?

	// Output dir and schema (not implemented)
	pflag.StringVarP(&options.outputDir,
		"output-dir", "o", ".", "The directory to output the image files into")
	pflag.StringVarP(&options.schema,
		"schema", "k", "default",
		"Schematic for filenaming, currently 'default' (y/m/d[layer].ext) or 'one' (y-m-d[l].ext")

	// Satellite image layers
	pflag.StringVarP(&options.satelliteLayers,
		"layers", "l", "", "The satellite layers from which to get images")

	// Dates
	pflag.StringVarP(&options.startDate,
		"start", "s", "", "Start date, format y-m-d")
	pflag.StringVarP(&options.endDate,
		"end", "e", "", "End date, format y-m-d")
	pflag.StringVarP(&options.interval,
		"interval", "i", "0-0-1", "Interval between image dates, format y-m-d")

	// Coordinates and size
	pflag.StringVarP(&options.coordinates,
		"coordinates", "c", "",
		"Bounding coordinates, in lat,lon,lat,lon (bottom-left - top-right)")
	pflag.StringVarP(&options.size,
		"size", "z", "", "Size of image in pixels: x,y")

	// image type
	pflag.StringVarP(&options.imageType,
		"format", "f", "jpeg", "Image format, accepts jpeg or png")

	pflag.Parse()

	// change relative paths to absolute ones
	options.outputDir, _ = filepath.Abs(options.outputDir)
	return &options
}

// Read an OptionGroup to create an ImageSet. (Gives possibility in future of
// making multiple ImageSets in one run, if that's easier than making ImageSet
// more versatile as a single thing.)
func prepImageSet(o *OptionGroup) ImageSet {
	imageSet := defaultImageSet

	switch o.schema {
	case "default":
		imageSet.fileSchema = DefaultScheme{o.outputDir}
	case "one":
		imageSet.fileSchema = OneFolderScheme{o.outputDir}
	default:
		panic("invalid file schema")
	}

	if o.satelliteLayers != "" {
		imageSet.satelliteLayers = strings.Split(o.satelliteLayers, ",")
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
	if o.size != "" {
		imageSet.width, imageSet.height = readSize(o.size)
	}

	switch o.imageType {
	case "jpeg":
		imageSet.imageType = "jpeg"
	case "jpg":
		imageSet.imageType = "jpeg"
	case "png":
		imageSet.imageType = "png"
	default:
		panic("invalid image type")
	}

	return imageSet
}

// Split a "x1,y1,x2,y2" into "x1","y1","x2","y2"
func readCoordinates(acme string) (a, c, m, e string) {
	ss := strings.Split(acme, ",")
	if len(ss) != 4 {
		panic("wrong number of coordinates")
	}
	return ss[0], ss[1], ss[2], ss[3]
}

// Split "x,y" to "x","y"
func readSize(xy string) (x, y string) {
	ss := strings.Split(xy, ",")
	if len(ss) != 2 {
		panic("size-parse didn't work")
	}
	return ss[0], ss[1]
}
