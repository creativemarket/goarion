package main

import (
    "flag"
    "fmt"
    "log"
    "runtime"
    "strconv"
    "strings"
    "sync"
    "time"

    "github.com/filitchp/goarion"
    "github.com/pkg/profile"
    "github.com/rcrowley/go-metrics"
)

var images = metrics.NewTimer()

func printStats() {
    fmt.Printf("mean: % 12s, min: % 12s, max: % 12s, %%99: % 12s, stdDev: % 12s, rate: % 8.1f, count: % 8d\n",
        time.Duration(images.Mean()),
        time.Duration(images.Min()),
        time.Duration(images.Max()),
        time.Duration(images.Percentile(0.99)),
        time.Duration(images.StdDev()),
        images.RateMean(),
        images.Count(),
    )
}

func main() {
    p := profile.Start(profile.MemProfile, profile.ProfilePath("."))
    defer p.Stop()
    times := 1000
    workers := runtime.GOMAXPROCS(0)
    filename := "file://../testdata/image.jpg"
    watermark := ""
    watermarkType := "standard"
    watermarkAmount := 0.10
    watermarkMin := 0.1
    watermarkMax := 0.5
    sharpenRadius := 0.5
    sharpen := 0
    size := "200x200,400x180,800x600"
    algo := "fill"
    quality := 92

    flag.IntVar(&times, "times", times, "number of resizes")
    flag.IntVar(&workers, "workers", workers, "number of workers")
    flag.StringVar(&filename, "file", filename, "input image")
    flag.StringVar(&watermark, "watermark", watermark, "watermark image")
    flag.StringVar(&watermarkType, "watermarkType", watermarkType, "watermark type")
    flag.StringVar(&size, "size", size, "comma separated list of sizes")
    flag.StringVar(&algo, "algo", algo, "comma separated list of algos")
    flag.IntVar(&sharpen, "sharpen", sharpen, "amount to sharpen")
    flag.Parse()

    log.Printf("GOMAXPROCS: %d, WORKERS: %d", runtime.GOMAXPROCS(0), workers)

    if filename == "" {
        fmt.Println("usage: bench filename [options]")
        flag.PrintDefaults()
        log.Fatal("please specify a filename")
    }

    opts := []goarion.Options{}
    
    // Parse size / algo
    for _, sizeString := range strings.Split(size, ",") {
        
        size := strings.Split(sizeString, "x")
        
        w, err := strconv.Atoi(size[0])
        if err != nil {
            log.Fatal(err)
        }
        
        h, err := strconv.Atoi(size[1])
        if err != nil {
            log.Fatal(err)
        }
        
        for _, algoString := range strings.Split(algo, ",") {
            
            // Set a default value
            a := goarion.StringToAlgo(algoString)
            wmt := goarion.StringToWatermarkType(watermarkType)

            opts = append(opts, goarion.Options{Algo: a, 
                                                Quality: quality, 
                                                Width: w, 
                                                Height: h, 
                                                SharpenRadius: sharpenRadius,
                                                SharpenAmount: sharpen,
                                                WatermarkUrl: watermark,
                                                WatermarkType: wmt,
                                                WatermarkMin: watermarkMin,
                                                WatermarkMax: watermarkMax,
                                                WatermarkAmount: watermarkAmount})
        }
    }

    optionsCount := len(opts)
    if optionsCount == 0 {
        log.Fatal("you must provide at least one size")
    }

    log.Printf("%d different options, run each %d times for a total of %d operations", optionsCount, times, optionsCount*times)
    for _, opt := range opts {
        log.Printf("%+v", opt)
    }

    go func() {
        for _ = range time.Tick(time.Second) {
            printStats()
        }
    }()

    wg := new(sync.WaitGroup)
    ch := make(chan goarion.Options, workers)
    wg.Add(workers)

    for i := 0; i < workers; i++ {
        go func() {
            defer wg.Done()
            for opt := range ch {
                images.Time(func() {
                    _, _, err := goarion.ResizeFromFile(filename, opt)
                    if err != nil {
                        log.Fatal(err)
                    }
                })
            }
        }()
    }

    for i := 0; i < times; i++ {
        for _, opt := range opts {
            ch <- opt
        }
    }

    close(ch)
    wg.Wait()
    printStats()
}
