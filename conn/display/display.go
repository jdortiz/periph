// Copyright 2016 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package display

import (
	"image"
	"image/color"

	"periph.io/x/periph/conn"
)

// Drawer represents a context to display pixels on an output device. It is a
// write-only interface.
//
// What Drawer represents can be as varied as a 1 bit OLED display or a strip
// of LED lights.
type Drawer interface {
	conn.Resource

	// ColorModel returns the device native color model.
	//
	// It is generally color.NRGBA for a color display or a specialize type for
	// 1-bit or 4-bit display devices.
	ColorModel() color.Model
	// Bounds returns the size of the output device.
	//
	// Generally displays should have Min at {0, 0} but this is not guaranteed in
	// multiple displays setup or when an instance of this interface represents a
	// section of a larger logical display.
	Bounds() image.Rectangle
	// Draw updates the display with this image starting at 'sp' offset into the
	// display into 'r'. The code will likely be faster if the image is in the
	// display's native color format.
	//
	// To be compatible with draw.Drawer, this function doesn't return an error.
	Draw(r image.Rectangle, src image.Image, sp image.Point)
}
