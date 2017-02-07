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

## Benchmark
Goarion comes with a handy benchmarking tool
```
cd bench
go build
./benchmark.sh
```

Below are a few illustrative benchmarks of Goarion from start to finish (reading input image -> generating final JPEG) on a 2.5 GHz Intel Core i7 MacBook Pro (Mid 2015) (see [raw results](https://raw.githubusercontent.com/wiki/filitchp/goarion/benchmarks/2.5-GHz-Intel-Core-i7-MacBook-Pro-Mid-2015.txt)).

**Basic resizing**

| Stat | Time |
|-----------|---------|
| mean      | 54.07 ms |
| min       | 28.94 ms |
| max       | 123.51 ms |
| %99       | 104.92 ms |
| stdDev    | 16.31 ms |
| rate      | 148.2 ops/second |
| count     | 4500 |

**Resizing + Sharpening**

| Stat | Time |
|-----------|---------|
| mean      | 65.81 ms |
| min       | 30.66 ms |
| max       | 174.79 ms |
| %99       | 128.32 ms |
| stdDev    | 25.143 ms |
| rate      | 119.0 ops/second |
| count     | 3000 |

**Resizing + Sharpening + Adaptive Watermark**
     
| Stat | Time |
|-----------|---------|
| mean      | 70.83 ms |
| min       | 32.73 ms |
| max       | 195.96 ms |
| %99       | 134.39 ms |
| stdDev    | 27.72 ms |
| rate      | 113.3 ops/second |
| count     | 3000 |

## Installation
Install Go if you haven't already
```bash
sudo apt install golang-go
```
Make sure your Go path is set
```bash
vim ~/.bashrc
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
