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
    errNoOutpuData         = errors.New("image data length is zero")
    errInvalidSourceFormat = errors.New("invalid data source format")
    errOperation           = errors.New("error running the operation")
    errEmptyInputUrl       = errors.New("empty intput url")
    errInvalidHeight       = errors.New("provided height is invalid")
    errInvalidWidth        = errors.New("provided width is invalid")
    errInvalidQuality      = errors.New("provided quality is invalid")
)

func ResizeFromFile(inputUrl string, options Options) ([]byte, error) {
    
    // if inputUrl == nil {        
    //     return nil, errEmptyInputUrl
    // }
    
    if options.Height <= 0 {
        return nil, errInvalidHeight
    }
    
    if options.Width <= 0 {
        return nil, errInvalidWidth
    }
    
    if options.Quality <= 0 {
        return nil, errInvalidQuality
    }
    
    cinputUrl := C.CString(inputUrl)
    defer C.free(unsafe.Pointer(cinputUrl))
    
    inputOptions := C.struct_ArionInputOptions{correctOrientation: 1, inputUrl:cinputUrl}

    algo := C.CString(AlgoToString(options.Algo))
    defer C.free(unsafe.Pointer(algo))
    
    gravity := C.CString(GravtiyToString(options.Gravity))
    defer C.free(unsafe.Pointer(gravity))

    // Ability to save to file (to disk) from Arion
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
                                                // watermarkUrl: 
                                                // watermarkAmount: }

    // Run it!
    result := C.ArionResize(inputOptions, resizeOptions)
    
    // Read back results
    outputData := unsafe.Pointer(result.outputData)
    outputJson := unsafe.Pointer(result.resultJson)
    returnCode := int(result.returnCode)
    
    // If we got back output data make sure it gets freed
    if outputData != nil {
        defer C.free(outputData)
    }

    // If we got back json make sure it gets freed
    if outputJson != nil {
        defer C.free(outputJson)
    }

    // Now check the error code
    if returnCode != 0 {
        return nil, errOperation
    }
     
    // We should have data, but we don't
    if outputData == nil {
        return nil, errNoOutpuData
    }

    jpeg := C.GoBytes(outputData, result.outputSize)
    
    return jpeg, nil
}
