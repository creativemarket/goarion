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

func ResizeFromFile(inputUrl string, options Options) ([]byte, string, error) {
  
//   if inputUrl == nil {    
//     return nil, errEmptyInputUrl
//   }
  
  // Set a default JSON response...
  json := "{\"result\":false,\"error_message\":\"Unknown error\"}"
  
  if options.Height <= 0 {
    return nil, json, errInvalidHeight
  }
  
  if options.Width <= 0 {
    return nil, json, errInvalidWidth
  }
  
  if options.Quality <= 0 {
    return nil, json, errInvalidQuality
  }
  
  cinputUrl       := C.CString(inputUrl)
  inputOptions    := C.struct_ArionInputOptions{correctOrientation: 1, inputUrl:cinputUrl}
  algo            := C.CString(AlgoToString(options.Algo))
  gravity         := C.CString(GravtiyToString(options.Gravity))
  watermarkUrl    := C.CString(options.WatermarkUrl)
  watermarkType   := C.CString(WatermarkTypeToString(options.WatermarkType))
  
  // Ability to save to file (to disk) from Arion
  // coutputUrl := C.CString(outputUrl)
  // defer C.free(unsafe.Pointer(coutputUrl))
  
  resizeOptions := C.struct_ArionResizeOptions{algo:            algo, 
                                               height:          C.uint(options.Height), 
                                               width:           C.uint(options.Width),
                                               gravity:         gravity,
                                               quality:         C.uint(options.Quality),
                                               sharpenRadius:   C.float(options.SharpenRadius),
                                               sharpenAmount:   C.uint(options.SharpenAmount),
                                               preserveMeta:    C.uint(0),
                                               watermarkUrl:    watermarkUrl,
                                               watermarkType:   watermarkType,
                                               watermarkMin:    C.float(options.WatermarkMin),
                                               watermarkMax:    C.float(options.WatermarkMax),
                                               watermarkAmount: C.float(options.WatermarkAmount)}

  // Run it!
  result := C.ArionResize(inputOptions, resizeOptions)
  
  // Cleanup
  defer C.free(unsafe.Pointer(cinputUrl))
  defer C.free(unsafe.Pointer(algo))
  defer C.free(unsafe.Pointer(gravity))
  defer C.free(unsafe.Pointer(watermarkUrl))
  defer C.free(unsafe.Pointer(watermarkType))
  
  // Read back results
  outputData := unsafe.Pointer(result.outputData)
  outputSize := int(result.outputSize)
  outputJson := unsafe.Pointer(result.resultJson)
  returnCode := int(result.returnCode)
  
  // If we got back json make sure it gets freed
  if outputJson != nil {
    json = C.GoString(result.resultJson)
    defer C.free(outputJson)
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
  jpeg := (*[1 << 30]byte)(outputData)[:outputSize:outputSize]
  
  return jpeg, json, nil
}
