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
	fmt.Printf("images available at %s\n", tmpdir)
}

// Write jpeg encoded byte array to disk
func dumpImage(filename string, data []byte) {
	err := ioutil.WriteFile(path.Join(tmpdir, filename+".jpg"), data, 0665)
	if err != nil {
		panic(err)
	}
}

func imageTestHelper(t *testing.T, src_path string, output_prefix string, opts Options) {
	assert := assert.New(t)

	// Perform the resize operation and make sure there are no errors
	// Data will be jpeg encoded
	data, err := ResizeFromFile(src_path, opts)
    
    assert.NoError(err)

	dest_name := fmt.Sprintf("%s%dx%d_%s", output_prefix, opts.Width, opts.Height, AlgoToString(opts.Algo))

	// Write the resized image back to disk
	dumpImage(dest_name, data)

	// Read back the resized image
	im, _, err := image.DecodeConfig(bytes.NewReader(data))

	// Make sure the resized image looks good
	assert.NoError(err)
	assert.Equal(opts.Width, im.Width)
	assert.Equal(opts.Height, im.Height)
}

func Test100x100(t *testing.T) {
    cleanup()
    
	src_path := "file://testdata/100x100_square.png"

	output_prefix := "100x100_square_to_"

	opts := Options{
		Algo:    FILL,
		Gravity: WEST,
		Width:   50,
		Height:  50,
        Quality: 92,
	}
    
	imageTestHelper(t, src_path, output_prefix, opts)

}