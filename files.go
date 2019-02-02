// Schematics for file naming
package main

import (
	"fmt"
	// "strconv"
)

// Interface for choosing a filename from an Image
type Schema interface {
	makeFileName(*Image) string
	makeFilePath(*Image) string
}

// A default scheme that puts the image in prefix/year/month/day<satellite>.<ext>
type DefaultScheme struct {
	prefix string // data is the prefix
}

func (_ DefaultScheme) makeFileName(i *Image) string {
	ext := i.imageType
	satL := i.satelliteLayer
	switch ext {
	case "jpeg":
		ext = "jpg"
	}
	switch satL {
	case "MODIS_Terra_correctedReflectance_TrueColor":
		satL = ""
	}
	return fmt.Sprintf("%02d", i.time.day) + satL + "." + ext
}

func (s DefaultScheme) makeFilePath(i *Image) string {
	//TODO should check for / in prefix before doubling; doesn't really matter?
	return s.prefix + "/" + fmt.Sprintf("%02d", i.time.year) + "/" +
		fmt.Sprintf("%02d", i.time.month) + "/"
}
