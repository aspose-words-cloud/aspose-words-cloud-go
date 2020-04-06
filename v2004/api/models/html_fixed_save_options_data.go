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



// container class for fixed html save options.
type HtmlFixedSaveOptionsData struct {

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

	// Gets or sets value determining whether content of StructuredDocumentTag is updated before saving.
	UpdateSdtContent bool `json:"UpdateSdtContent,omitempty"`

	// Gets or sets a value determining if fields should be updated before saving the document to a fixed page format. Default value for this property is. true
	UpdateFields bool `json:"UpdateFields,omitempty"`

	// Gets or sets a value determining how 3D effects are rendered.
	Dml3DEffectsRenderingMode string `json:"Dml3DEffectsRenderingMode,omitempty"`

	// Gets or sets a value determining whether the Aspose.Words.Properties.BuiltInDocumentProperties.LastPrinted property is updated before saving.
	UpdateLastPrintedProperty bool `json:"UpdateLastPrintedProperty,omitempty"`

	// Gets or sets specifies prefix which is added to all class names in style.css file. Default value is \"aw\".
	CssClassNamesPrefix string `json:"CssClassNamesPrefix,omitempty"`

	// Gets or sets encoding.
	Encoding string `json:"Encoding,omitempty"`

	// Gets or sets specifies whether the CSS (Cascading Style Sheet) should be embedded into Html document.
	ExportEmbeddedCss bool `json:"ExportEmbeddedCss,omitempty"`

	// Gets or sets specifies whether fonts should be embedded into Html document in Base64 format.
	ExportEmbeddedFonts bool `json:"ExportEmbeddedFonts,omitempty"`

	// Gets or sets specifies whether images should be embedded into Html document in Base64 format.
	ExportEmbeddedImages bool `json:"ExportEmbeddedImages,omitempty"`

	// Gets or sets indication of whether form fields are exported as interactive items (as 'input' tag) rather than converted to text or graphics.
	ExportFormFields bool `json:"ExportFormFields,omitempty"`

	// Gets or sets specifies export format of fonts.
	FontFormat string `json:"FontFormat,omitempty"`

	// Gets or sets specifies the horizontal alignment of pages in an HTML document. Default value is HtmlFixedHorizontalPageAlignment.Center.
	PageHorizontalAlignment string `json:"PageHorizontalAlignment,omitempty"`

	// Gets or sets specifies the margins around pages in an HTML document. The margins value is measured in points and should be equal to or greater than 0. Default value is 10 points.
	PageMargins float64 `json:"PageMargins,omitempty"`

	// Gets or sets specifies the physical folder where resources are saved when exporting a document.
	ResourcesFolder string `json:"ResourcesFolder,omitempty"`

	// Gets or sets specifies the name of the folder used to construct resource URIs.
	ResourcesFolderAlias string `json:"ResourcesFolderAlias,omitempty"`

	// Gets or sets flag indicates whether \"@font-face\" CSS rules should be placed into a separate file \"fontFaces.css\" when a document is being saved with external stylesheet (that is, when Aspose.Words.Saving.HtmlFixedSaveOptions.ExportEmbeddedCss is false). Default value is false, all CSS rules are written into single file \"styles.css\".
	SaveFontFaceCssSeparately bool `json:"SaveFontFaceCssSeparately,omitempty"`

	// Gets or sets specifies whether border around pages should be shown.
	ShowPageBorder bool `json:"ShowPageBorder,omitempty"`

	// Gets or sets flag indicates whether fonts from target machine must be used to display the document. If this flag is set to true, Aspose.Words.Saving.HtmlFixedSaveOptions.FontFormat and Aspose.Words.Saving.HtmlFixedSaveOptions.ExportEmbeddedFonts properties do not have effect, also Aspose.Words.Saving.HtmlFixedSaveOptions.ResourceSavingCallback is not fired for fonts. Default is false.
	UseTargetMachineFonts bool `json:"UseTargetMachineFonts,omitempty"`
}

type IHtmlFixedSaveOptionsData interface {
	IsHtmlFixedSaveOptionsData() bool
}
func (HtmlFixedSaveOptionsData) IsHtmlFixedSaveOptionsData() bool {
	return true;
}
func (HtmlFixedSaveOptionsData) IsFixedPageSaveOptionsData() bool {
	return true;
}
