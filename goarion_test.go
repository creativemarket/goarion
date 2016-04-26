package goarion

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

const tmpdir = "./testoutput"

// Setup the temp directory for a new test run
func cleanup() {
	os.RemoveAll(tmpdir)
	os.MkdirAll(tmpdir, 0775)
}

// Write jpeg encoded byte array to disk
func dumpImage(filename string, data []byte) {
	err := ioutil.WriteFile(path.Join(tmpdir, filename+".jpg"), data, 0665)
	if err != nil {
		panic(err)
	}
}

func imageTestHelper(t *testing.T, srcPath string, outputPrefix string, opts Options) {
	assert := assert.New(t)

	// Perform the resize operation and make sure there are no errors
	// Data will be jpeg encoded
	data, err := ResizeFromFile(srcPath, opts)
    
    assert.NoError(err)

	destName := fmt.Sprintf("%s%dx%d_%s", outputPrefix, opts.Width, opts.Height, AlgoToString(opts.Algo))

	// Write the resized image back to disk
	dumpImage(destName, data)

	// Read back the resized image
	im, _, err := image.DecodeConfig(bytes.NewReader(data))

	// Make sure the resized image looks good
	assert.NoError(err)
	assert.Equal(opts.Width, im.Width)
	assert.Equal(opts.Height, im.Height)
}

func TestJpg(t *testing.T) {
	cleanup()
	assert := assert.New(t)
 
    srcPath := "file://testdata/image.jpg"
 
	opts := []Options{
		{Algo: WIDTH,  Quality: 92, Height: 2000,  Width: 150,  SharpenRadius:1.0, SharpenAmount:200},
		{Algo: WIDTH,  Quality: 20, Height: 2000,  Width: 300,  SharpenRadius:0.5, SharpenAmount:80},
		{Algo: HEIGHT, Quality: 92, Height: 150,   Width: 2000, SharpenRadius:0.5, SharpenAmount:80},
		{Algo: HEIGHT, Quality: 92, Height: 300,   Width: 200, SharpenRadius:0.5, SharpenAmount:80},
		{Algo: SQUARE, Quality: 92, Height: 100,   Width: 100,  SharpenRadius:0.5, SharpenAmount:80},
		{Algo: SQUARE, Quality: 92, Height: 300,   Width: 300,  SharpenRadius:0.5, SharpenAmount:80},
	}

	for _, opt := range opts {
        data, err := ResizeFromFile(srcPath, opt)
        
        assert.NoError(err)
        
        outputName := fmt.Sprintf("image_jpg_to_%dx%d_%s", opt.Width, opt.Height, AlgoToString(opt.Algo))
        
        dumpImage(outputName, data)
        _, _, err = image.DecodeConfig(bytes.NewReader(data))
        assert.NoError(err)
	}
}

func Test100x100(t *testing.T) {
    
	srcPath := "file://testdata/100x100_square.png"

	outputPrefix := "100x100_square_to_"

	opts := Options{
		Algo:    FILL,
		Gravity: WEST,
		Width:   50,
		Height:  50,
        Quality: 92,
	}
    
	imageTestHelper(t, srcPath, outputPrefix, opts)

}

// Here we have a tall source image and we are always cropping a tall portion at the center of the image
func Test100x200TallCenter(t *testing.T) {

	srcPath := "file://testdata/100x200_tall_center.png"
	outputPrefix := "100x200_tall_center_to_"

	// Just a crop, take the center
	opts := Options{
		Algo:    FILL,
		Gravity: CENTER,
		Width:   50,
		Height:  200,
        Quality: 92,
	}
	imageTestHelper(t, srcPath, outputPrefix, opts)

	// Shrink, take the center
	opts = Options{
		Algo:    FILL,
		Gravity: NORTH,
		Width:   25,
		Height:  100,
        Quality: 92,
	}
	imageTestHelper(t, srcPath, outputPrefix, opts)

	// Enlarge, take the center
	opts = Options{
		Algo:    FILL,
		Gravity: SOUTH,
		Width:   100,
		Height:  400,
        Quality: 92,
	}
	imageTestHelper(t, srcPath, outputPrefix, opts)

}

// Here we have a tall source image and we are always cropping a tall portion at the left of the image
func Test100x200TallLeft(t *testing.T) {

	srcPath := "file://testdata/100x200_tall_left.png"
	outputPrefix := "100x200_tall_left_to_"

	// Just a crop, take the left
	opts := Options{
		Algo:    FILL,
		Gravity: WEST,
		Width:   50,
		Height:  200,
        Quality: 92,
	}
	imageTestHelper(t, srcPath, outputPrefix, opts)

	// Shrink, take the left
	opts = Options{
		Algo:    FILL,
		Gravity: NORTH_WEST,
		Width:   25,
		Height:  100,
        Quality: 92,
	}
	imageTestHelper(t, srcPath, outputPrefix, opts)

	// Enlarge, take the left
	opts = Options{
		Algo:    FILL,
		Gravity: SOUTH_WEST,
		Width:   100,
		Height:  400,
        Quality: 92,
	}
	imageTestHelper(t, srcPath, outputPrefix, opts)

}

// Here we have a tall source image and we are always cropping a tall portion at the right of the image
func Test100x200TallRight(t *testing.T) {

	srcPath := "file://testdata/100x200_tall_right.png"
	outputPrefix := "100x200_tall_right_to_"

	// Just a crop, take the right
	opts := Options{
		Algo:    FILL,
		Gravity: EAST,
		Width:   50,
		Height:  200,
        Quality: 92,
	}
	imageTestHelper(t, srcPath, outputPrefix, opts)

	// Shrink, take the right
	opts = Options{
		Algo:    FILL,
		Gravity: NORTH_EAST,
		Width:   25,
		Height:  100,
        Quality: 92,
	}
	imageTestHelper(t, srcPath, outputPrefix, opts)

	// Enlarge, take the right
	opts = Options{
		Algo:    FILL,
		Gravity: SOUTH_EAST,
		Width:   100,
		Height:  400,
        Quality: 92,
	}
	imageTestHelper(t, srcPath, outputPrefix, opts)

}

// Here we have a tall source image and we are always cropping a wide portion at the bottom of the image
func Test100x200WideBottom(t *testing.T) {

	srcPath := "file://testdata/100x200_wide_bottom.png"
	outputPrefix := "100x200_wide_bottom_to_"

	// Just a crop, take the bottom
	opts := Options{
		Algo:    FILL,
		Gravity: SOUTH,
		Width:   100,
		Height:  50,
        Quality: 92,
	}
	imageTestHelper(t, srcPath, outputPrefix, opts)

	// Shrink, take the bottom
	opts = Options{
		Algo:    FILL,
		Gravity: SOUTH_EAST,
		Width:   50,
		Height:  25,
        Quality: 92,
	}
	imageTestHelper(t, srcPath, outputPrefix, opts)

	// Enlarge, take the bottom
	opts = Options{
		Algo:    FILL,
		Gravity: SOUTH_WEST,
		Width:   200,
		Height:  100,
        Quality: 92,
	}
	imageTestHelper(t, srcPath, outputPrefix, opts)

}

// Here we have a tall source image and we are always cropping a wide portion at the bottom of the image
func Test100x200WideCenter(t *testing.T) {

	srcPath := "file://testdata/100x200_wide_center.png"
	outputPrefix := "100x200_wide_center_to_"

	// Just a crop, take the bottom
	opts := Options{
		Algo:    FILL,
		Gravity: CENTER,
		Width:   100,
		Height:  50,
        Quality: 92,
	}
	imageTestHelper(t, srcPath, outputPrefix, opts)

	// Shrink, take the bottom
	opts = Options{
		Algo:    FILL,
		Gravity: EAST,
		Width:   50,
		Height:  25,
        Quality: 92,
	}
	imageTestHelper(t, srcPath, outputPrefix, opts)

	// Enlarge, take the bottom
	opts = Options{
		Algo:    FILL,
		Gravity: WEST,
		Width:   200,
		Height:  100,
        Quality: 92,
	}
	imageTestHelper(t, srcPath, outputPrefix, opts)

}

// Here we have a tall source image and we are always cropping a wide portion at the top of the image
func Test100x200WideTop(t *testing.T) {

	srcPath := "file://testdata/100x200_wide_top.png"
	outputPrefix := "100x200_wide_top_to_"

	// Just a crop, take the top
	opts := Options{
		Algo:    FILL,
		Gravity: NORTH,
		Width:   100,
		Height:  50,
        Quality: 92,
	}
	imageTestHelper(t, srcPath, outputPrefix, opts)

	// Shrink, take the top
	opts = Options{
		Algo:    FILL,
		Gravity: NORTH_EAST,
		Width:   50,
		Height:  25,
        Quality: 92,
	}
	imageTestHelper(t, srcPath, outputPrefix, opts)

	// Enlarge, take the top
	opts = Options{
		Algo:    FILL,
		Gravity: NORTH_WEST,
		Width:   200,
		Height:  100,
        Quality: 92,
	}
	imageTestHelper(t, srcPath, outputPrefix, opts)

}

// Here we have a wide source image and we are always cropping a tall portion at the center of the image
func Test200x100TallCenter(t *testing.T) {

	srcPath := "file://testdata/200x100_tall_center.png"
	outputPrefix := "200x100_tall_center_to_"

	// Just a crop, take the center
	opts := Options{
		Algo:    FILL,
		Gravity: CENTER,
		Width:   50,
		Height:  100,
        Quality: 92,
	}
	imageTestHelper(t, srcPath, outputPrefix, opts)

	// Shrink, take the center
	opts = Options{
		Algo:    FILL,
		Gravity: NORTH,
		Width:   25,
		Height:  50,
        Quality: 92,
	}
	imageTestHelper(t, srcPath, outputPrefix, opts)

	// Enlarge, take the center
	opts = Options{
		Algo:    FILL,
		Gravity: SOUTH,
		Width:   100,
		Height:  200,
        Quality: 92,
	}
	imageTestHelper(t, srcPath, outputPrefix, opts)

}

// Here we have a tall source image and we are always cropping a tall portion at the left of the image
func Test200x100TallLeft(t *testing.T) {

	srcPath := "file://testdata/200x100_tall_left.png"
	outputPrefix := "200x100_tall_left_to_"

	// Just a crop, take the left
	opts := Options{
		Algo:    FILL,
		Gravity: WEST,
		Width:   50,
		Height:  100,
        Quality: 92,
	}
	imageTestHelper(t, srcPath, outputPrefix, opts)

	// Shrink, take the left
	opts = Options{
		Algo:    FILL,
		Gravity: NORTH_WEST,
		Width:   25,
		Height:  50,
        Quality: 92,
	}
	imageTestHelper(t, srcPath, outputPrefix, opts)

	// Enlarge, take the left
	opts = Options{
		Algo:    FILL,
		Gravity: SOUTH_WEST,
		Width:   100,
		Height:  200,
        Quality: 92,
	}
	imageTestHelper(t, srcPath, outputPrefix, opts)

}

// Here we have a tall source image and we are always cropping a tall portion at the right of the image
func Test200x100TallRight(t *testing.T) {

	srcPath := "file://testdata/200x100_tall_right.png"
	outputPrefix := "200x100_tall_right_to_"

	// Just a crop, take the right
	opts := Options{
		Algo:    FILL,
		Gravity: EAST,
		Width:   50,
		Height:  100,
        Quality: 92,
	}
	imageTestHelper(t, srcPath, outputPrefix, opts)

	// Shrink, take the right
	opts = Options{
		Algo:    FILL,
		Gravity: NORTH_EAST,
		Width:   25,
		Height:  50,
        Quality: 92,
	}
	imageTestHelper(t, srcPath, outputPrefix, opts)

	// Enlarge, take the right
	opts = Options{
		Algo:    FILL,
		Gravity: SOUTH_EAST,
		Width:   100,
		Height:  200,
        Quality: 92,
	}
	imageTestHelper(t, srcPath, outputPrefix, opts)

}

// Here we have a tall source image and we are always cropping a wide portion at the bottom of the image
func Test200x100WideBottom(t *testing.T) {

	srcPath := "file://testdata/200x100_wide_bottom.png"
	outputPrefix := "200x100_wide_bottom_to_"

	// Just a crop, take the bottom
	opts := Options{
		Algo:    FILL,
		Gravity: SOUTH,
		Width:   200,
		Height:  50,
        Quality: 92,
	}
	imageTestHelper(t, srcPath, outputPrefix, opts)

	// Shrink, take the bottom
	opts = Options{
		Algo:    FILL,
		Gravity: SOUTH_EAST,
		Width:   100,
		Height:  25,
        Quality: 92,
	}
	imageTestHelper(t, srcPath, outputPrefix, opts)

	// Enlarge, take the bottom
	opts = Options{
		Algo:    FILL,
		Gravity: SOUTH_WEST,
		Width:   400,
		Height:  100,
        Quality: 92,
	}
	imageTestHelper(t, srcPath, outputPrefix, opts)

}

// Here we have a tall source image and we are always cropping a wide portion at the bottom of the image
func Test200x100WideCenter(t *testing.T) {

	srcPath := "file://testdata/200x100_wide_center.png"
	outputPrefix := "200x100_wide_center_to_"

	// Just a crop, take the bottom
	opts := Options{
		Algo:    FILL,
		Gravity: CENTER,
		Width:   200,
		Height:  50,
        Quality: 92,
	}
	imageTestHelper(t, srcPath, outputPrefix, opts)

	// Shrink, take the bottom
	opts = Options{
		Algo:    FILL,
		Gravity: EAST,
		Width:   100,
		Height:  25,
        Quality: 92,
	}
	imageTestHelper(t, srcPath, outputPrefix, opts)

	// Enlarge, take the bottom
	opts = Options{
		Algo:    FILL,
		Gravity: WEST,
		Width:   400,
		Height:  100,
        Quality: 92,
	}
	imageTestHelper(t, srcPath, outputPrefix, opts)

}

// Here we have a tall source image and we are always cropping a wide portion at the top of the image
func Test200x100WideTop(t *testing.T) {

	srcPath := "file://testdata/200x100_wide_top.png"
	outputPrefix := "200x100_wide_top_to_"

	// Just a crop, take the top
	opts := Options{
		Algo:    FILL,
		Gravity: NORTH,
		Width:   200,
		Height:  50,
        Quality: 92,
	}
	imageTestHelper(t, srcPath, outputPrefix, opts)

	// Shrink, take the top
	opts = Options{
		Algo:    FILL,
		Gravity: NORTH_EAST,
		Width:   100,
		Height:  25,
        Quality: 92,
	}
	imageTestHelper(t, srcPath, outputPrefix, opts)

	// Enlarge, take the top
	opts = Options{
		Algo:    FILL,
		Gravity: NORTH_WEST,
		Width:   400,
		Height:  100,
        Quality: 92,
	}
	imageTestHelper(t, srcPath, outputPrefix, opts)

}