# Goarion
Goarion is a Go wrapper for [Arion](https://github.com/snapwire-media/arion), a fast thumbnail creation and 
image metadata extraction library. 

While Arion library provides a powerful CLI and C++ library, Goarion supplements that functionality by allowing for
quick and easy service creation.  Goarion is built with performance in mind by providing multi-threaded support
(see [bench](blob/master/bench/main.go)).

**Example use cases**
* On-the-fly image resize/watermarking API
* Batch resize/watermark operations on large datasets
* Existing Go projects that need image manipulation

This project was inspired by the [T-REZ](https://github.com/DAddYE/trez) library.  
