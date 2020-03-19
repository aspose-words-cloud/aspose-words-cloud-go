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



// Words document property DTO.
type DocumentProperty struct {

	Link *WordsApiLink `json:"link,omitempty"`

	// Gets or sets a value indicating whether flag indicates whether the property is built-in or not. If true the property is built-in, if false the property is custom.
	BuiltIn bool `json:"BuiltIn,omitempty"`

	// Gets or sets name of the document property.
	Name string `json:"Name,omitempty"`

	// Gets or sets string value of the document property.
	Value string `json:"Value,omitempty"`
}

type IDocumentProperty interface {
	IsDocumentProperty() bool
}
func (DocumentProperty) IsDocumentProperty() bool {
	return true;
}
func (DocumentProperty) IsLinkElement() bool {
	return true;
}
