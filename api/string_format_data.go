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

// Allows to specify System.Drawing.StringFormat options.
type StringFormatData struct {

	// Gets or sets horizontal alignment of the string.
	Alignment string `json:"Alignment,omitempty"`

	// Gets or sets a System.Drawing.StringFormatFlags enumeration that contains formatting information.
	FormatFlags string `json:"FormatFlags,omitempty"`

	// Gets or sets the System.Drawing.Text.HotkeyPrefix object for this System.Drawing.StringFormat object.
	HotkeyPrefix string `json:"HotkeyPrefix,omitempty"`

	// Gets or sets the vertical alignment of the string.
	LineAlignment string `json:"LineAlignment,omitempty"`

	// Gets or sets the System.Drawing.StringTrimming enumeration for this System.Drawing.StringFormat object.
	Trimming string `json:"Trimming,omitempty"`
}
type IStringFormatData interface {
	IsStringFormatData() bool
}
func (StringFormatData) IsStringFormatData() bool {
	return true;
}

