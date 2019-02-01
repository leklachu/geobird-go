package main

import (
	"path/filepath"

	"github.com/spf13/pflag"
)

var outputDir string

func parseCommandLineArguments() {

	pflag.StringVarP(&outputDir, "output-dir", "o", ".", "The directory to output the image files into")
	pflag.Parse()

	// change relative paths to absolute ones
	outputDir, _ = filepath.Abs(outputDir)
}
