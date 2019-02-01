package main

import (
	"path/filepath"

	"github.com/spf13/pflag"
)

type OptionGroup struct {
	outputDir string
}

func parseCommandLineArguments() OptionGroup {

	options := OptionGroup{} // do this as a global or return it?

	pflag.StringVarP(&options.outputDir, "output-dir", "o", ".", "The directory to output the image files into")
	pflag.Parse()

	// change relative paths to absolute ones
	options.outputDir, _ = filepath.Abs(options.outputDir)
	return options
}
