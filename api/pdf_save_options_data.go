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

// container class for pdf save options.
type PdfSaveOptionsData struct {

	// Gets or sets a value determining how colors are rendered. { Normal | Grayscale}.
	ColorMode string `json:"ColorMode,omitempty"`

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

	// Gets or sets specifies the PDF standards compliance level for output documents.
	Compliance string `json:"Compliance,omitempty"`

	// Gets or sets specifies whether to convert footnote/endnote references in main text story into active hyperlinks. When clicked the hyperlink will lead to the corresponding footnote/endnote. Default is false.
	CreateNoteHyperlinks bool `json:"CreateNoteHyperlinks,omitempty"`

	// Gets or sets a value determining the way  are exported to PDF file. Default value is .
	CustomPropertiesExport string `json:"CustomPropertiesExport,omitempty"`

	// Gets or sets specifies the details for signing the output PDF document.
	DigitalSignatureDetails *PdfDigitalSignatureDetailsData `json:"DigitalSignatureDetails,omitempty"`

	// Gets or sets a flag specifying whether the window�s title bar should display the document title taken from the Title entry of the document information dictionary.
	DisplayDocTitle bool `json:"DisplayDocTitle,omitempty"`

	// Gets or sets allows to specify downsample options.
	DownsampleOptions *DownsampleOptionsData `json:"DownsampleOptions,omitempty"`

	// Gets or sets controls how fonts are embedded into the resulting PDF documents.
	EmbedFullFonts bool `json:"EmbedFullFonts,omitempty"`

	// Gets or sets specifies the details for encrypting the output PDF document.
	EncryptionDetails *PdfEncryptionDetailsData `json:"EncryptionDetails,omitempty"`

	// Gets or sets a flag specifying whether URI should be escaped before writing.             
	EscapeUri bool `json:"EscapeUri,omitempty"`

	// Gets or sets determines whether or not to export document structure.
	ExportDocumentStructure bool `json:"ExportDocumentStructure,omitempty"`

	// Gets or sets specifies the font embedding mode.
	FontEmbeddingMode string `json:"FontEmbeddingMode,omitempty"`

	// Gets or sets determines how bookmarks in headers/footers are exported. The default value is Aspose.Words.Saving.HeaderFooterBookmarksExportMode.All.
	HeaderFooterBookmarksExportMode string `json:"HeaderFooterBookmarksExportMode,omitempty"`

	// Gets or sets specifies how the color space will be selected for the images in PDF document.
	ImageColorSpaceExportMode string `json:"ImageColorSpaceExportMode,omitempty"`

	// Gets or sets specifies compression type to be used for all images in the document.
	ImageCompression string `json:"ImageCompression,omitempty"`

	// Gets or sets determines whether hyperlinks in the output Pdf document are forced to be opened in a new window (or tab) of a browser.
	OpenHyperlinksInNewWindow bool `json:"OpenHyperlinksInNewWindow,omitempty"`

	// Gets or sets allows to specify outline options.
	OutlineOptions *OutlineOptionsData `json:"OutlineOptions,omitempty"`

	// Gets or sets specifies how the PDF document should be displayed when opened in the PDF reader.
	PageMode string `json:"PageMode,omitempty"`

	// Gets or sets a value determining whether or not to preblend transparent images with black background color.
	PreblendImages bool `json:"PreblendImages,omitempty"`

	// Gets or sets specifies whether to preserve Microsoft Word form fields as form fields in PDF or convert them to text.
	PreserveFormFields bool `json:"PreserveFormFields,omitempty"`

	// Gets or sets specifies compression type to be used for all textual content in the document.
	TextCompression string `json:"TextCompression,omitempty"`

	// Gets or sets determines whether the document should be saved using a booklet printing layout.
	UseBookFoldPrintingSettings bool `json:"UseBookFoldPrintingSettings,omitempty"`

	// Gets or sets determines whether or not to substitute TrueType fonts Arial, Times New Roman, Courier New and Symbol with core PDF Type 1 fonts.
	UseCoreFonts bool `json:"UseCoreFonts,omitempty"`

	// Gets or sets determines what type of zoom should be applied when a document is opened with a PDF viewer.
	ZoomBehavior string `json:"ZoomBehavior,omitempty"`

	// Gets or sets determines zoom factor (in percentages) for a document.
	ZoomFactor int32 `json:"ZoomFactor,omitempty"`
}
