# Goarion
Goarion is a Go wrapper for [Arion](https://github.com/snapwire-media/arion), a fast thumbnail creation and 
image metadata extraction library. 

While Arion provides a powerful CLI and C++ library, Goarion supplements its functionality by allowing for
quick and easy Go integration.  Goarion is intended for microservice or batch processing tools, and can easily parallelize Arion operations (see [bench](bench/main.go)).

**Example use cases**
* On-the-fly image resize/watermarking API
* Batch resize/watermark operations on large datasets
* Existing Go projects that need image manipulation

This project was inspired by the [T-REZ](https://github.com/DAddYE/trez) library.  
