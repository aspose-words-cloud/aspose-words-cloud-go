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



// Section element.
type HeaderFooter struct {

	// Gets or sets paragraph's text.
	Type_ string `json:"Type,omitempty"`

	Link *WordsApiLink `json:"link,omitempty"`

	// Gets or sets child nodes.
	ChildNodes []NodeLink `json:"ChildNodes,omitempty"`

	DrawingObjects *LinkElement `json:"DrawingObjects,omitempty"`

	Paragraphs *LinkElement `json:"Paragraphs,omitempty"`
}

type IHeaderFooter interface {
	IsHeaderFooter() bool
}
func (HeaderFooter) IsHeaderFooter() bool {
	return true;
}
func (HeaderFooter) IsHeaderFooterLink() bool {
	return true;
}
