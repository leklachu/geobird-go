// Schematics for file naming
package main

import "strconv"

// Interface for choosing a filename from an Image
type Schema interface {
	makeFileName(Image) string
	makeFilePath(Image) string
}

// A default scheme that puts the image in prefix/year/month/day<satellite>.<ext>
type DefaultScheme struct {
	prefix string // data is the prefix
}

func (_ DefaultScheme) makeFileName(i Image) string {
	ext := i.imageType
	switch ext {
	case "jpeg":
		ext = "jpg"

	}
	return strconv.Itoa(i.time.day) + i.satelliteLayer + "." + ext
}

func (s DefaultScheme) makeFilePath(i Image) string {
	//TODO should check for / in prefix before doubling; doesn't really matter?
	return s.prefix + "/" + strconv.Itoa(i.time.year) + "/" +
		strconv.Itoa(i.time.month) + "/"
}
