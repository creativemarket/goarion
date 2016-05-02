#!/bin/sh

# Basic variation of sizes and algorithms
go run main.go -file file://../testdata/image.jpg -size 100x100,640x480,1024x768 -algo fill,width,height -times 500

# Baseline fill operation
go run main.go -file file://../testdata/image.jpg -size 100x100,640x480,1024x768 -algo fill

# Fill with sharpening
go run main.go -file file://../testdata/image.jpg -size 100x100,640x480,1024x768 -algo fill -sharpen 80

# Fill with sharpening and watermark
go run main.go -file file://../testdata/image.jpg -size 100x100,640x480,1024x768 -algo fill -sharpen 80 -watermark file://../testdata/watermark.png

# Fill with sharpening and watermark, small thumbnail
go run main.go -file file://../testdata/image.jpg -size 100x100 -algo fill -sharpen 80 -watermark file://../testdata/watermark.png

