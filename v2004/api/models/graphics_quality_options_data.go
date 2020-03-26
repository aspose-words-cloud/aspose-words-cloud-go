/*
* MIT License

* Copyright (c) 2019 Aspose Pty Ltd

* Permission is hereby granted, free of charge, to any person obtaining a copy
* of this software and associated documentation files (the "Software"), to deal
* in the Software without restriction, including without limitation the rights
* to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
* copies of the Software, and to permit persons to whom the Software is
* furnished to do so, subject to the following conditions:

* The above copyright notice and this permission notice shall be included in all
* copies or substantial portions of the Software.

* THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
* IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
* FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
* AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
* LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
* OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
* SOFTWARE.
*/


package models



// Allows to specify additional System.Drawing.Graphics quality options.
type GraphicsQualityOptionsData struct {

	// Gets or sets a value that specifies how composited images are drawn to this Graphics.
	CompositingMode string `json:"CompositingMode,omitempty"`

	// Gets or sets the rendering quality of composited images drawn to this Graphics.
	CompositingQuality string `json:"CompositingQuality,omitempty"`

	// Gets or sets the interpolation mode associated with this Graphics.
	InterpolationMode string `json:"InterpolationMode,omitempty"`

	// Gets or sets the rendering quality for this Graphics.
	SmoothingMode string `json:"SmoothingMode,omitempty"`

	StringFormat *StringFormatData `json:"StringFormat,omitempty"`

	// Gets or sets the rendering mode for text associated with this Graphics.
	TextRenderingHint string `json:"TextRenderingHint,omitempty"`
}

type IGraphicsQualityOptionsData interface {
	IsGraphicsQualityOptionsData() bool
}
func (GraphicsQualityOptionsData) IsGraphicsQualityOptionsData() bool {
	return true;
}
