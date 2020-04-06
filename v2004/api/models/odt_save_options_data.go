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



// container class for odt/ott save options.
type OdtSaveOptionsData struct {

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

	// Gets or sets specifies whether export should correspond to ODT specification 1.1 strictly.
	IsStrictSchema11 bool `json:"IsStrictSchema11,omitempty"`

	// Gets or sets allows to specify units of measure to apply to document content. The default value is Aspose.Words.Saving.OdtSaveMeasureUnit.Centimeters.  Open Office uses centimeters when specifying lengths, widths and other measurable formatting and content properties in documents whereas MS Office uses inches.
	MeasureUnit string `json:"MeasureUnit,omitempty"`

	// Gets or sets a password to encrypt document.
	Password string `json:"Password,omitempty"`

	// Gets or sets specifies whether or not use pretty formats output.
	PrettyFormat bool `json:"PrettyFormat,omitempty"`
}

type IOdtSaveOptionsData interface {
	IsOdtSaveOptionsData() bool
}
func (OdtSaveOptionsData) IsOdtSaveOptionsData() bool {
	return true;
}
func (OdtSaveOptionsData) IsSaveOptionsData() bool {
	return true;
}
