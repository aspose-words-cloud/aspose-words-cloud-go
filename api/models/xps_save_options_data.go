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



// Container class for xps save options.
type XpsSaveOptionsData struct {

	// Gets or sets a value determining how colors are rendered. { Normal | Grayscale}.
	ColorMode string `json:"ColorMode,omitempty"`

	// Gets or sets determines the quality of the JPEG images inside PDF document.
	JpegQuality int32 `json:"JpegQuality,omitempty"`

	MetafileRenderingOptions *MetafileRenderingOptionsData `json:"MetafileRenderingOptions,omitempty"`

	// Gets or sets indicates the symbol set that is used to represent numbers while rendering to fixed page formats.
	NumeralFormat string `json:"NumeralFormat,omitempty"`

	// Gets or sets flag indicates whether it is required to optimize output of XPS. If this flag is set redundant nested canvases and empty canvases are removed, also neighbor glyphs with the same formatting are concatenated. Note: The accuracy of the content display may be affected if this property is set to true.  Default is false.
	OptimizeOutput bool `json:"OptimizeOutput,omitempty"`

	// Gets or sets determines number of pages to render.
	PageCount int32 `json:"PageCount,omitempty"`

	// Gets or sets determines 0-based index of the first page to render.
	PageIndex int32 `json:"PageIndex,omitempty"`

	// Gets or sets specifies the level in the XPS document outline at which to display Word bookmarks.
	BookmarksOutlineLevel int32 `json:"BookmarksOutlineLevel,omitempty"`

	// Gets or sets specifies how many levels of headings (paragraphs formatted with the Heading styles) to include in the XPS document outline.
	HeadingsOutlineLevels int32 `json:"HeadingsOutlineLevels,omitempty"`

	OutlineOptions *OutlineOptionsData `json:"OutlineOptions,omitempty"`

	// Gets or sets determines whether the document should be saved using a booklet printing layout.
	UseBookFoldPrintingSettings bool `json:"UseBookFoldPrintingSettings,omitempty"`
}

type IXpsSaveOptionsData interface {
	IsXpsSaveOptionsData() bool
}
func (XpsSaveOptionsData) IsXpsSaveOptionsData() bool {
	return true;
}
func (XpsSaveOptionsData) IsFixedPageSaveOptionsData() bool {
	return true;
}
