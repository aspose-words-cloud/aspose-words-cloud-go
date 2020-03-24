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



// Container class for text save options.
type TextSaveOptionsData struct {

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

	// Gets or sets specifies whether to add bi-directional marks before each BiDi run when exporting in plain text format. The default value is true.
	AddBidiMarks bool `json:"AddBidiMarks,omitempty"`

	// Gets or sets specifies the encoding to use when exporting in plain text format.
	Encoding string `json:"Encoding,omitempty"`

	// Gets or sets specifies whether to output headers and footers when exporting in plain text format. default value is TxtExportHeadersFootersMode.PrimaryOnly.
	ExportHeadersFootersMode string `json:"ExportHeadersFootersMode,omitempty"`

	// Gets or sets allows to specify whether the page breaks should be preserved during export. The default value is false.
	ForcePageBreaks bool `json:"ForcePageBreaks,omitempty"`

	// Gets or sets specifies the string to use as a paragraph break when exporting in plain text format.
	ParagraphBreak string `json:"ParagraphBreak,omitempty"`

	// Gets or sets specifies whether the program should attempt to preserve layout of tables when saving in the plain text format.
	PreserveTableLayout bool `json:"PreserveTableLayout,omitempty"`

	// Gets or sets specifies whether the program should simplify list labels in case of complex label formatting not being adequately represented by plain text.
	SimplifyListLabels bool `json:"SimplifyListLabels,omitempty"`
}

type ITextSaveOptionsData interface {
	IsTextSaveOptionsData() bool
}
func (TextSaveOptionsData) IsTextSaveOptionsData() bool {
	return true;
}
func (TextSaveOptionsData) IsSaveOptionsData() bool {
	return true;
}
