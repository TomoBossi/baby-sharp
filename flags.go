package main

import (
	"flag"
	"fmt"
)

type flags struct {
	gaussianBlurDeviation float64
	strength              float64
}

func newFlags() (*flags, error) {
	gaussianBlurDeviation := flag.Float64("gaussian-blur-deviation", 1.0, "DEFAULT 1.0 - Standard Deviation of the gaussian blur (low pass filter) applied to the image as part of the high pass filter calculation.")
	strength := flag.Float64("strength", 1.0, "DEFAULT 1.0 - Strength of the sharpening effect. 0.0 means no sharpening, 1.0 means full sharpening.")
	flag.Parse()

	if *gaussianBlurDeviation < 0 {
		return nil, fmt.Errorf("gaussian-blur-radius must be greater than 0")
	}

	if *strength < 0 || *strength > 1 {
		return nil, fmt.Errorf("strength must be between 0 and 1")
	}

	return &flags{
		gaussianBlurDeviation: *gaussianBlurDeviation,
		strength:              *strength,
	}, nil
}
