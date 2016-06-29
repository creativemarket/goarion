package goarion

//#cgo CFLAGS: -Wall -Wextra -Os -Wno-unused-function -Wno-unused-parameter
//#cgo LDFLAGS: -lcarion
//#include <stdio.h>
//#include <stdlib.h>
//#include <carion.h>
//
import "C"
import (
	"errors"
	"unsafe"
)

var (
	errNoOutpuData         = errors.New("Image data length is zero")
	errInvalidSourceFormat = errors.New("Invalid data source format")
	errOperation           = errors.New("Error running the operation")
	errInvalidHeight       = errors.New("Provided height is invalid")
	errInvalidWidth        = errors.New("Provided width is invalid")
	errInvalidQuality      = errors.New("Provided quality is invalid")
)

// ResizeFromFile Performs a resize operation given an input url
// On success this will return JPEG data in a byte array
func ResizeFromFile(inputURL string, options Options) (jpeg []byte, json string, err error) {

	// Set a default JSON response...
	json = `{"result":false,"error_message":"Unknown error"}`

	if options.Height <= 0 {
		return nil, json, errInvalidHeight
	}

	if options.Width <= 0 {
		return nil, json, errInvalidWidth
	}

	if options.Quality <= 0 {
		return nil, json, errInvalidQuality
	}

	cinputURL := C.CString(inputURL)
	inputOptions := C.struct_ArionInputOptions{correctOrientation: 1, inputUrl: cinputURL}
	algo := C.CString(AlgoToString(options.Algo))
	gravity := C.CString(GravtiyToString(options.Gravity))
	watermarkURL := C.CString(options.WatermarkURL)
	watermarkType := C.CString(WatermarkTypeToString(options.WatermarkType))

	// Ability to save to file (to disk) from Arion
	// coutputUrl := C.CString(outputUrl)
	// If used, put this after call to resize!
	// defer C.free(unsafe.Pointer(coutputUrl))

	resizeOptions := C.struct_ArionResizeOptions{algo: algo,
		height:          C.uint(options.Height),
		width:           C.uint(options.Width),
		gravity:         gravity,
		quality:         C.uint(options.Quality),
		sharpenRadius:   C.float(options.SharpenRadius),
		sharpenAmount:   C.uint(options.SharpenAmount),
		preserveMeta:    C.uint(0),
		watermarkUrl:    watermarkURL,
		watermarkType:   watermarkType,
		watermarkMin:    C.float(options.WatermarkMin),
		watermarkMax:    C.float(options.WatermarkMax),
		watermarkAmount: C.float(options.WatermarkAmount)}

	// Run it!
	result := C.ArionResize(inputOptions, resizeOptions)

	// Cleanup
	defer C.free(unsafe.Pointer(cinputURL))
	defer C.free(unsafe.Pointer(algo))
	defer C.free(unsafe.Pointer(gravity))
	defer C.free(unsafe.Pointer(watermarkURL))
	defer C.free(unsafe.Pointer(watermarkType))

	// Read back results
	outputData := unsafe.Pointer(result.outputData)
	outputSize := int(result.outputSize)
	outputJSON := unsafe.Pointer(result.resultJson)
	returnCode := int(result.returnCode)

	// If we got back json make sure it gets freed
	if outputJSON != nil {
		json = C.GoString(result.resultJson)
		defer C.free(outputJSON)
	}

	// Now check the error code
	if returnCode != 0 {

		// If we got back output data make sure it gets freed
		if outputData != nil {
			defer C.free(outputData)
		}

		return nil, json, errOperation
	}

	// We should have data, but we don't
	if outputData == nil {
		return nil, json, errNoOutpuData
	}

	// This works, but creates an extra copy...
	// jpeg := C.GoBytes(outputData, result.outputSize)
	// If we got back output data make sure it gets freed
	// if outputData != nil {
	//   defer C.free(outputData)
	// }

	// Avoid the extra copy
	// See https://github.com/golang/go/wiki/cgo#turning-c-arrays-into-go-slices
	// TODO: does this need to be freed or will Go use garbage collection?
	jpeg = (*[1 << 30]byte)(outputData)[:outputSize:outputSize]

	return jpeg, json, nil
}
