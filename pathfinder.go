// Works out URLs and filepaths
package main

import (
	// "net/http"
	"net/url"
	// "strconv"
)

// a struct with all the parameters
type ImageSet struct {
	_ int
}

// struct that represents individual image
type Image struct {
	satelliteLayer                 string // full layer name? or satellite name?
	latMin, latMax, lonMin, lonMax string // string easier to crib from worldview
	width, height                  string // by which point... eh, just string
	time                           string // or use d/m/y?
	imageType                      string // jpeg/png
}

// get Images from ImageSet; pipe them back through a channel
func (iSet ImageSet) getImages(ichannel chan<- Image) {
	for i := 0; i < 10; i++ {
		ichannel <- Image{}
	}
	close(ichannel)
}

// Defaults for image fetching
const (
	iurl = "https://gibs.earthdata.nasa.gov.wms.epsg4326/best/wms.cgi"
)

var i1 = Image{
	satelliteLayer: "MODIS_Terra_correctedReflectance_TrueColor",
	latMin:         "28.505126953125",
	latMax:         "29.700439453125",
	lonMin:         "82.00634765625",
	lonMax:         "83.6982421875",
	width:          "385",
	height:         "272",
	time:           "2018-01-29",
	imageType:      "jpeg"}

// Method to take image struct and give URL
func (i Image) getImageURL() url.URL {

	query := make(url.Values)
	query.Add("SERVICE", "WMS")
	query.Add("REQUEST", "GetMap")
	query.Add("VERSION", "1.3.0")
	query.Add("LAYERS", i.satelliteLayer)
	query.Add("FORMAT", "image/"+i.imageType)
	query.Add("HEIGHT", i.height)
	query.Add("WIDTH", i.width)
	query.Add("TIME", i.time)
	query.Add("CRS", "EPSG:4326")
	query.Add("BBox", i.latMin+","+i.lonMin+","+i.latMax+","+i.lonMax)
	return url.URL{
		Scheme:   "https",
		Host:     "gibs.earthdata.nasa.gov",
		Path:     "/wms/epsg4326/best/wms.cgi",
		RawQuery: query.Encode()}
}

// Method to take image struct and give filepath
func (i Image) getImagePath() string {
	return "testimg.jpg"
}
