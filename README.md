# Goarion
Goarion is a Go wrapper for [Arion](https://github.com/snapwire-media/arion), a fast thumbnail creation and 
image metadata extraction library. 

While Arion provides a powerful CLI and C++ library, Goarion supplements its functionality by allowing for
quick and easy Go integration.  Goarion is intended for creating microservices and batch processing tools, and allows effortless parallelization of Arion operations (see [bench](bench/main.go)).

**Example Applications:**
* On-the-fly image resize/watermarking API
* Batch resize/watermark operations on large datasets
* Existing Go projects that need image manipulation

This project was inspired by the [T-REZ](https://github.com/DAddYE/trez) library.  

## Performance

Below are a few illustrative benchmarks of Goarion from start to finish (reading input image -> generating final JPEG).   (see [raw results](https://raw.githubusercontent.com/wiki/filitchp/goarion/benchmarks/2.5-GHz-Intel-Core-i7-MacBook-Pro-Mid-2015.txt)).

**Hardware:** 2.5 GHz Intel Core i7 MacBook Pro (Mid 2015) <br>
**Cores:** 4 physcial cores, 8 logical cores <br>
**OS:** MacOS X El Capitan <br>

| Stat      | Resize        | Resize + Sharpen |  Resize + Sharpen + Watermark |
|-----------|---------------|------------------|-------------------------------|
| mean      | 54.07 ms      | 65.81 ms         | 70.83 ms |
| min       | 28.94 ms      | 30.66 ms         | 32.73 ms |
| max       | 123.51 ms     | 174.79 ms        | 195.96 ms |
| %99       | 104.92 ms     | 128.32 ms        |  134.39 ms |
| stdDev    | 16.31 ms      | 25.143 ms        | 27.72 ms |
| rate      | 148.2 ops/sec | 119.0 ops/sec    | 113.3 ops/sec |
| count     | 4500          | 3000             | 3000 |


**Hardware:** Intel Xeon CPU E5-2698 v4 @ 2.20GHz <br>
**Cores:** 20 physcial cores, 40 logical cores <br>
**OS:** Ubuntu Xenial 16.04 <br>

| Stat      | Resize        | Resize + Sharpen |  Resize + Sharpen + Watermark |
|-----------|---------------|------------------|-------------------------------|
| mean      | 49.77 ms      | 61.16 ms         | 67.88 ms |
| min       | 17.95 ms      | 18.32 ms         | 20.62 ms |
| max       | 331.81 ms     | 292.40 ms        | 200.48 ms |
| %99       | 112.12 ms     | 143.85 ms        | 138.77 ms |
| stdDev    | 22.72 ms      | 29.00 ms         | 32.37 ms |
| rate      | 810.1 ops/sec | 668.9 ops/sec    | 572.7 ops/sec |
| count     | 90000         | 30000            | 30000 |


## Installation
Installation assumes Ubuntu-based system, but Goarion has been tested on MacOS X as well and installation is very similar.<br> 
Install Go if you haven't already (for other OSs see https://golang.org/doc/install)
```bash
sudo apt install golang-go
```
Make sure your Go path is set
```bash
vim ~/.bashrc
# Add this line to you .bashrc (~/.bash_profile on MacOS X)
export GOPATH=~/code/go
source ~/.bashrc
```
Create the project directory structure and clone the repo
```bash
cd $GOPATH
mkdir -p src/github.com/filitchp
cd src/github.com/filitchp
git clone git@github.com:filitchp/goarion.git
```
NOTE: before building Goarion you have to install the Arion library by following the steps here: https://github.com/snapwire-media/arion#installation

Satisfy Go dependencies (this will recursively install them using the Go standard directory structure)
```bash
cd goarion
go get ./...
```

###Running the benchmark
Build the benchmark (this shows sample usage of the library)
```bash
cd bench
go build
```
The following script benchmarks a few different operations and serves as an example of various Goarion operations
```bash
./benchmark.sh
```

###Running test cases
```bash
cd $GOPATH/src/github.com/filitchp/goarion
go test
```
If Go complains about the package stretchr/testify missing you can install it manually with
```bash
go get github.com/stretchr/testify
```

You should see the following
```
PASS
ok  	github.com/filitchp/goarion	1.131s
```
