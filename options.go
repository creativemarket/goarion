package goarion

import (
	"errors"
	"strings"
)

type ImageType int

// These must be kept in sync with the defines in arion/carion.h
const (
	JPEG ImageType = iota
	PNG
	JP2
	WEBP
	AVIF
)

type Gravity int

const (
	CENTER Gravity = iota
	NORTH
	SOUTH
	WEST
	EAST
	NORTH_WEST
	NORTH_EAST
	SOUTH_WEST
	SOUTH_EAST
)

type Algo int

const (
	WIDTH Algo = iota
	HEIGHT
	SQUARE
	FILL
)

type WatermarkType int

const (
	STANDARD WatermarkType = iota
	ADAPTIVE
)

func (wt *WatermarkType) UnmarshalText(b []byte) error {
	str := strings.Trim(string(b), `"`)

	switch str {
	case "STANDARD":
		*wt = STANDARD
	case "ADAPTIVE":
		*wt = ADAPTIVE
	default:
		return errors.New("Unknown WatermarkType: " + string(b))
	}

	return nil
}

func (wt WatermarkType) String() string {
	switch wt {
	case STANDARD:
		return "STANDARD"
	case ADAPTIVE:
		return "ADAPTIVE"
	default:
		return "UNKNOWN"
	}
}

type Options struct {
	ImageType       ImageType
	Width           int
	Height          int
	Algo            Algo
	Gravity         Gravity
	Quality         int
	SharpenRadius   float64
	SharpenAmount   int
	WatermarkURL    string
	WatermarkType   WatermarkType
	WatermarkAmount float64
	WatermarkMin    float64
	WatermarkMax    float64
}

func AlgoToString(a Algo) string {
	switch a {
	case WIDTH:
		return "width"
	case HEIGHT:
		return "height"
	case SQUARE:
		return "square"
	case FILL:
		return "fill"
	default:
		return "invalid"
	}
}

func StringToAlgo(s string) Algo {
	switch s {
	case "width":
		return WIDTH
	case "height":
		return HEIGHT
	case "square":
		return SQUARE
	default:
		return FILL
	}
}

func GravtiyToString(g Gravity) string {
	switch g {
	case NORTH:
		return "n"
	case SOUTH:
		return "s"
	case WEST:
		return "w"
	case EAST:
		return "e"
	case NORTH_WEST:
		return "nw"
	case NORTH_EAST:
		return "ne"
	case SOUTH_WEST:
		return "sw"
	case SOUTH_EAST:
		return "se"
	default:
		return "c"
	}
}

func WatermarkTypeToString(w WatermarkType) string {
	switch w {
	case ADAPTIVE:
		return "adaptive"
	default:
		return "standard"
	}
}

func StringToWatermarkType(s string) WatermarkType {
	switch s {
	case "adaptive":
		return ADAPTIVE
	default:
		return STANDARD
	}
}
