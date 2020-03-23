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
type Section struct {

	// Gets or sets link to the document.
	Link *WordsApiLink `json:"link,omitempty"`

	// Gets or sets child nodes.
	ChildNodes []NodeLink `json:"ChildNodes,omitempty"`

	// Gets or sets link to HeaderFooters resource.
	HeaderFooters *LinkElement `json:"HeaderFooters,omitempty"`

	// Gets or sets link to PageSetup resource.
	PageSetup *LinkElement `json:"PageSetup,omitempty"`

	// Gets or sets link to Paragraphs resource.
	Paragraphs *LinkElement `json:"Paragraphs,omitempty"`

	// Gets or sets link to Tables resource.
	Tables *LinkElement `json:"Tables,omitempty"`
}
type ISection interface {
	IsSection() bool
}
func (Section) IsSection() bool {
	return true;
}
func (Section) IsLinkElement() bool {
	return true;
}
