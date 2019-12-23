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


package api

// container class for options of metafile rendering.
type MetafileRenderingOptionsData struct {

	// Gets or sets determines how EMF+ Dual metafiles should be rendered.
	EmfPlusDualRenderingMode string `json:"EmfPlusDualRenderingMode,omitempty"`

	// Gets or sets a value determining whether or not the raster operations should be emulated.             
	EmulateRasterOperations bool `json:"EmulateRasterOperations,omitempty"`

	// Gets or sets determines how metafile images should be rendered.
	RenderingMode string `json:"RenderingMode,omitempty"`

	// Gets or sets determines how WMF metafiles with embedded EMF metafiles should be rendered.
	UseEmfEmbeddedToWmf bool `json:"UseEmfEmbeddedToWmf,omitempty"`

	// Gets or sets a value determining whether or not to scale fonts in WMF metafile according to metafile size on the page. The default value is true.
	ScaleWmfFontsToMetafileSize bool `json:"ScaleWmfFontsToMetafileSize,omitempty"`
}
