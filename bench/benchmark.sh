#!/bin/sh

# Basic variation of sizes and algorithms
echo
echo '------------------------------------------------------------------------------------------'
echo
echo ' Benchmark 1: Three resize types on 3 output sizes. No sharpening / no watermark'
echo
echo '------------------------------------------------------------------------------------------'

go run main.go -file file://../testdata/image.jpg -size 100x100,640x480,1024x768 -algo fill,width,height -times 500


# Baseline fill operation
echo
echo '------------------------------------------------------------------------------------------'
echo
echo ' Benchmark 2: "Fill" resize type with 3 output sizes. No sharpening / no watermark'
echo
echo '------------------------------------------------------------------------------------------'

go run main.go -file file://../testdata/image.jpg -size 100x100,640x480,1024x768 -algo fill


# Fill with sharpening
echo
echo '------------------------------------------------------------------------------------------'
echo
echo ' Benchmark 3: "Fill" resize type with 3 output sizes. Sharpening / no watermark'
echo
echo '------------------------------------------------------------------------------------------'

go run main.go -file file://../testdata/image.jpg -size 100x100,640x480,1024x768 -algo fill -sharpen 80


# Fill with sharpening and watermark
echo
echo '------------------------------------------------------------------------------------------'
echo
echo ' Benchmark 4: Fill resize type with 3 output sizes. Sharpening and watermark'
echo
echo '------------------------------------------------------------------------------------------'

go run main.go -file file://../testdata/image.jpg -size 100x100,640x480,1024x768 -algo fill -sharpen 80 -watermark file://../testdata/watermark.png


# Fill with sharpening and adaptive watermark
echo
echo '------------------------------------------------------------------------------------------'
echo
echo ' Benchmark 5: Fill resize type with 3 output sizes. Sharpening and adaptive watermark'
echo
echo '------------------------------------------------------------------------------------------'

go run main.go -file file://../testdata/image.jpg -size 100x100,640x480,1024x768 -algo fill -sharpen 80 -watermark file://../testdata/watermark.png -watermarkType adaptive


# Fill with sharpening and watermark, small thumbnail
echo
echo '------------------------------------------------------------------------------------------'
echo
echo ' Benchmark 6: Fill resize type with one small output size. Sharpening and watermark'
echo
echo '------------------------------------------------------------------------------------------'

go run main.go -file file://../testdata/image.jpg -size 100x100 -algo fill -sharpen 80 -watermark file://../testdata/watermark.png -times 2000


