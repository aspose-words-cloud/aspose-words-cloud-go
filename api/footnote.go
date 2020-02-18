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

// Footnote.
type Footnote struct {

	// Gets or sets link to the document.
	Link *WordsApiLink `json:"link,omitempty"`

	// Gets or sets node id.
	NodeId string `json:"NodeId,omitempty"`

	// Gets or sets content of footnote.
	Content *StoryChildNodes `json:"Content,omitempty"`

	// Gets or sets returns a value that specifies whether this is a footnote or endnote.
	FootnoteType string `json:"FootnoteType,omitempty"`

	// Gets or sets link to comment range start node.
	Position *DocumentPosition `json:"Position,omitempty"`

	// Gets or sets /sets custom reference mark to be used for this footnote. Default value is , meaning auto-numbered footnotes are used.
	ReferenceMark string `json:"ReferenceMark,omitempty"`

	// Gets or sets this is a convenience property that allows to easily get or set text of the footnote.
	Text string `json:"Text,omitempty"`
}
type IFootnote interface {
	IsFootnote() bool
}
func (Footnote) IsFootnote() bool {
	return true;
}
func (Footnote) IsFootnoteLink() bool {
	return true;
}

