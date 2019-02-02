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

////////////////////
// Default Scheme //
////////////////////
// A default scheme that puts the image in prefix/year/month/day<satellite>.<ext>
type DefaultScheme struct {
	prefix string // data is the prefix
}

// TODO should these be changed so pointer implements method?
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
	case "MODIS_Aqua_correctedReflectance_TrueColor":
		satL = "a"
	}
	return fmt.Sprintf("%02d", i.time.day) + satL + "." + ext
}

func (s DefaultScheme) makeFilePath(i *Image) string {
	//TODO should check for / in prefix before doubling; doesn't really matter?
	return s.prefix + "/" + fmt.Sprintf("%02d", i.time.year) + "/" +
		fmt.Sprintf("%02d", i.time.month) + "/"
	// return fmt.Sprintf(...altogether
}

///////////////////////
// One-Folder Scheme //
///////////////////////
// A scheme defining one folder to put all photos in with ISO-whatever
// formatted dates (2026-02-33)
type OneFolderScheme struct {
	path string // directory path
}

func (_ OneFolderScheme) makeFileName(i *Image) string {
	ext := i.imageType
	satL := i.satelliteLayer
	switch ext {
	case "jpeg":
		ext = "jpg"
	}
	switch satL {
	case "MODIS_Terra_correctedReflectance_TrueColor":
		satL = ""
	case "MODIS_Aqua_correctedReflectance_TrueColor":
		satL = "_a"
	default:
		satL = "_" + satL
	}
	return show(i.time) + satL + "." + ext
}

func (s OneFolderScheme) makeFilePath(_ *Image) string {
	return s.path + "/" // need the /? I think so.
}
