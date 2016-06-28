package goarion

import "testing"
import "github.com/stretchr/testify/assert"

var watermarkTypeDeserializationTests = []struct {
	text        []byte
	expected    WatermarkType
	expectError bool
}{
	{[]byte("STANDARD"), STANDARD, false},
	{[]byte("ADAPTIVE"), ADAPTIVE, false},
	{[]byte("BANANA"), -1, true},
}

func TestWatermarkTypeDeserialization(t *testing.T) {
	for _, test := range watermarkTypeDeserializationTests {
		// Need weird indirection to be able to test
		// a pointer to a WatermarkType
		var wmt *WatermarkType
		std := STANDARD
		wmt = &std
		err := wmt.UnmarshalText(test.text)
		if err != nil {
			if !test.expectError {
				t.Error(err)
			}
		} else {
			assert.Equal(t, test.expected, *wmt)
		}
	}
}

var watermarkTypeStringTests = []struct {
	wmt      WatermarkType
	expected string
}{
	{STANDARD, "STANDARD"},
	{ADAPTIVE, "ADAPTIVE"},
	{-1, "UNKNOWN"},
}

func TestWatermarkTypeString(t *testing.T) {
	for _, test := range watermarkTypeStringTests {
		assert.Equal(t, test.expected, test.wmt.String())
	}
}
