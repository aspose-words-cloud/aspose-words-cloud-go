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

// Represents Words document DTO.
type Document struct {

	// Gets or sets a list of links that originate from this document.
	Links []Link `json:"Links,omitempty"`

	// Gets or sets the name of the file.
	FileName string `json:"FileName,omitempty"`

	// Gets or sets the original format of the document.
	SourceFormat string `json:"SourceFormat"`

	// Gets or sets a value indicating whether returns true if the document is encrypted and requires a password to open.
	IsEncrypted bool `json:"IsEncrypted"`

	// Gets or sets a value indicating whether returns true if the document contains a digital signature. This property merely informs that a digital signature is present on a document, but it does not specify whether the signature is valid or not.
	IsSigned bool `json:"IsSigned"`

	// Gets or sets returns document properties.
	DocumentProperties *DocumentProperties `json:"DocumentProperties,omitempty"`
}
type IDocument interface {
	IsDocument() bool
}
func (Document) IsDocument() bool {
	return true;
}

