// Works out URLs and filepaths
package main

import (
	"net/url"
	// "strconv"
)

// a struct with all the parameters
type ImageSet struct {
	satelliteLayers                []string
	dates                          []Date
	latMin, latMax, lonMin, lonMax string
	width, height                  string
	imageType                      string // jpeg/png
	fileSchema                     Schema
}

var defaultImageSet = ImageSet{
	satelliteLayers: []string{"MODIS_Terra_correctedReflectance_TrueColor"},
	latMin:          "28.505126953125",
	latMax:          "29.700439453125",
	lonMin:          "82.00634765625",
	lonMax:          "83.6982421875",
	width:           "385",
	height:          "272",
	dates:           []Date{Date{2018, 01, 29}},
	imageType:       "jpeg",
	fileSchema:      DefaultScheme{"./img/"}}

// struct that represents individual image
type Image struct {
	satelliteLayer                 string // full layer name? or satellite name?
	latMin, latMax, lonMin, lonMax string // string easier to crib from worldview
	width, height                  string // by which point... eh, just string
	time                           Date   // for easy file handling
	imageType                      string // jpeg/png
	parent                         *ImageSet
}

// get Images from ImageSet; pipe them back through a channel
func (iSet *ImageSet) getImages(ichannel chan<- Image) {
	img := Image{parent: iSet}

	// not sure whether to fit these in here, or iterate over them, or just call
	// them from the parent pointer later. Copying them in for now
	img.latMin, img.latMax, img.lonMin, img.lonMax =
		iSet.latMin, iSet.latMax, iSet.lonMin, iSet.lonMax
	img.width, img.height = iSet.width, iSet.height
	img.imageType = iSet.imageType

	for _, d := range iSet.dates {
		for _, s := range iSet.satelliteLayers {
			img.time = d
			img.satelliteLayer = s
			ichannel <- img
		}
	}
	close(ichannel)
}

var i1 = Image{
	satelliteLayer: "MODIS_Terra_correctedReflectance_TrueColor",
	latMin:         "28.505126953125",
	latMax:         "29.700439453125",
	lonMin:         "82.00634765625",
	lonMax:         "83.6982421875",
	width:          "385",
	height:         "272",
	time:           Date{2018, 01, 29},
	imageType:      "jpeg"}

// Method to take image struct and give URL
func (i *Image) getImageURL() url.URL {

	query := make(url.Values)
	query.Add("SERVICE", "WMS")
	query.Add("REQUEST", "GetMap")
	query.Add("VERSION", "1.3.0")
	query.Add("LAYERS", i.satelliteLayer)
	query.Add("FORMAT", "image/"+i.imageType)
	query.Add("HEIGHT", i.height)
	query.Add("WIDTH", i.width)
	query.Add("TIME", show(i.time))
	query.Add("CRS", "EPSG:4326")
	query.Add("BBox", i.latMin+","+i.lonMin+","+i.latMax+","+i.lonMax)
	return url.URL{
		Scheme:   "https",
		Host:     "gibs.earthdata.nasa.gov",
		Path:     "/wms/epsg4326/best/wms.cgi",
		RawQuery: query.Encode()}
}

// Method to take image struct and give filepath
func (i *Image) getImagePath() string {
	return i.parent.fileSchema.makeFilePath(i) +
		i.parent.fileSchema.makeFileName(i)
}
