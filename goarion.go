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
	errNoData              = errors.New("image data length is zero")
	errInvalidSourceFormat = errors.New("invalid data source format")
	errEncoding            = errors.New("error during encoding")
    errInvalidHeight       = errors.New("provided height is invalid")
    errInvalidWidth        = errors.New("provided width is invalid")
)

func ResizeFromFile(inputUrl string, options Options) ([]byte, error) {
    
	cinputUrl := C.CString(inputUrl)
	defer C.free(unsafe.Pointer(cinputUrl))
    
    inputOptions := C.struct_ArionInputOptions{correctOrientation: 1, inputUrl:cinputUrl}

    algo := C.CString(AlgoToString(options.Algo))
	defer C.free(unsafe.Pointer(algo))
    
    gravity := C.CString(GravtiyToString(options.Gravity))
    defer C.free(unsafe.Pointer(gravity))

    // Ability to save to file from Arion
	// coutputUrl := C.CString(outputUrl)
	// defer C.free(unsafe.Pointer(coutputUrl))

    resizeOptions := C.struct_ArionResizeOptions{algo: algo, 
                                                height: C.uint(options.Height), 
                                                width:  C.uint(options.Width),
                                                gravity: gravity,
                                                quality: C.uint(options.Quality),
                                                sharpenRadius: C.float(options.SharpenRadius),
                                                sharpenAmount: C.uint(options.SharpenAmount),
                                                preserveMeta: C.uint(0)}
                                                
    result := C.ArionResize(inputOptions, resizeOptions)
    
    outputData := unsafe.Pointer(result.outputData)
    outputJson := unsafe.Pointer(result.resultJson)
    outputError := unsafe.Pointer(result.errorMessage)
    
    if outputData != nil {
        defer C.free(outputData)
    }

    if outputJson != nil {
        defer C.free(outputJson)
    }
    
    if outputError != nil {
        defer C.free(outputError)
        
        // TODO: use the actual output error
        return nil, errEncoding
    }

    jpeg := C.GoBytes(outputData, result.outputSize)
    
    return jpeg, nil
}

// func Resize(data []byte, options Options) ([]byte, error) {
// 	if len(data) == 0 {
// 		return nil, errNoData
// 	}

// 	// return resize(src, options)
//     return nil, errInvalidSourceFormat
// }
