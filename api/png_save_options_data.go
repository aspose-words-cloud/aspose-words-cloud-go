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

// container class for png save options.
type PngSaveOptionsData struct {

	// Gets or sets format of save.
	SaveFormat string `json:"SaveFormat,omitempty"`

	// Gets or sets name of destination file.
	FileName string `json:"FileName,omitempty"`

	// Gets or sets a value determining how DrawingML shapes are rendered. { Fallback | DrawingML }.
	DmlRenderingMode string `json:"DmlRenderingMode,omitempty"`

	// Gets or sets a value determining how DrawingML effects are rendered. { Simplified | None | Fine }.
	DmlEffectsRenderingMode string `json:"DmlEffectsRenderingMode,omitempty"`

	// Gets or sets controls zip output or not. Default value is false.
	ZipOutput bool `json:"ZipOutput,omitempty"`

	// Gets or sets a value determining whether the Aspose.Words.Properties.BuiltInDocumentProperties.LastSavedTime property is updated before saving.
	UpdateLastSavedTimeProperty bool `json:"UpdateLastSavedTimeProperty,omitempty"`

	// Gets or sets value determining whether content of  is updated before saving.
	UpdateSdtContent bool `json:"UpdateSdtContent,omitempty"`

	// Gets or sets a value determining if fields should be updated before saving the document to a fixed page format. Default value for this property is. true
	UpdateFields bool `json:"UpdateFields,omitempty"`

	// Gets or sets a value determining how colors are rendered. { Normal | Grayscale}.
	ColorMode string `json:"ColorMode,omitempty"`

	// Gets or sets determines the quality of the JPEG images inside PDF document.
	JpegQuality int32 `json:"JpegQuality,omitempty"`

	// Gets or sets allows to specify metafile rendering options.
	MetafileRenderingOptions *MetafileRenderingOptionsData `json:"MetafileRenderingOptions,omitempty"`

	// Gets or sets indicates the symbol set that is used to represent numbers while rendering to fixed page formats.
	NumeralFormat string `json:"NumeralFormat,omitempty"`

	// Gets or sets flag indicates whether it is required to optimize output of XPS. If this flag is set redundant nested canvases and empty canvases are removed, also neighbor glyphs with the same formatting are concatenated. Note: The accuracy of the content display may be affected if this property is set to true.  Default is false.
	OptimizeOutput bool `json:"OptimizeOutput,omitempty"`

	// Gets or sets determines number of pages to render.
	PageCount int32 `json:"PageCount,omitempty"`

	// Gets or sets determines 0-based index of the first page to render.
	PageIndex int32 `json:"PageIndex,omitempty"`

	// Gets or sets allows to specify additional System.Drawing.Graphics quality options.
	GraphicsQualityOptions *GraphicsQualityOptionsData `json:"GraphicsQualityOptions,omitempty"`

	// Gets or sets the horizontal resolution for the generated images, in dots per inch.  This property has effect only when saving to raster image formats. The default value is 96.
	HorizontalResolution float64 `json:"HorizontalResolution,omitempty"`

	// Gets or sets brightness of image.
	ImageBrightness float64 `json:"ImageBrightness,omitempty"`

	// Gets or sets color mode of image.
	ImageColorMode string `json:"ImageColorMode,omitempty"`

	// Gets or sets contrast of image.
	ImageContrast float64 `json:"ImageContrast,omitempty"`

	// Gets or sets background (paper) color of image.
	PaperColor string `json:"PaperColor,omitempty"`

	// Gets or sets pixel format of image.
	PixelFormat string `json:"PixelFormat,omitempty"`

	// Gets or sets both horizontal and vertical resolution for the generated images, in dots per inch.  This property has effect only when saving to raster image formats. The default value is 96.
	Resolution float64 `json:"Resolution,omitempty"`

	// Gets or sets zoom factor of image.
	Scale float64 `json:"Scale,omitempty"`

	// Gets or sets determine whether or not to use anti-aliasing for rendering.
	UseAntiAliasing bool `json:"UseAntiAliasing,omitempty"`

	// Gets or sets a value determining whether to use GDI+ or Aspose.Words metafile renderer when saving to EMF.
	UseGdiEmfRenderer bool `json:"UseGdiEmfRenderer,omitempty"`

	// Gets or sets determine whether or not to use high quality (i.e. slow) rendering algorithms.
	UseHighQualityRendering bool `json:"UseHighQualityRendering,omitempty"`

	// Gets or sets the vertical resolution for the generated images, in dots per inch.  This property has effect only when saving to raster image formats. The default value is 96.
	VerticalResolution float64 `json:"VerticalResolution,omitempty"`
}
type IPngSaveOptionsData interface {
	IsPngSaveOptionsData() bool
}
func (PngSaveOptionsData) IsPngSaveOptionsData() bool {
	return true;
}
func (PngSaveOptionsData) IsImageSaveOptionsData() bool {
	return true;
}

