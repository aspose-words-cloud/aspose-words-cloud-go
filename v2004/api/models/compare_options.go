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



// Container class for compare documents options.
type CompareOptions struct {

	// Gets or sets a value indicating whether true indicates that documents comparison is case insensitive. By default comparison is case sensitive.             
	IgnoreCaseChanges bool `json:"IgnoreCaseChanges,omitempty"`

	// Gets or sets a value indicating whether specifies whether to compare the differences in data contained in tables. By default tables are not ignored.             
	IgnoreTables bool `json:"IgnoreTables,omitempty"`

	// Gets or sets a value indicating whether specifies whether to compare differences in fields. By default fields are not ignored.             
	IgnoreFields bool `json:"IgnoreFields,omitempty"`

	// Gets or sets a value indicating whether specifies whether to compare differences in footnotes and endnotes. By default footnotes are not ignored.             
	IgnoreFootnotes bool `json:"IgnoreFootnotes,omitempty"`

	// Gets or sets a value indicating whether specifies whether to compare differences in comments. By default comments are not ignored.             
	IgnoreComments bool `json:"IgnoreComments,omitempty"`

	// Gets or sets a value indicating whether specifies whether to compare differences in the data contained within text boxes. By default textboxes are not ignored.             
	IgnoreTextboxes bool `json:"IgnoreTextboxes,omitempty"`

	// Gets or sets a value indicating whether true indicates that formatting is ignored. By default document formatting is not ignored.             
	IgnoreFormatting bool `json:"IgnoreFormatting,omitempty"`

	// Gets or sets a value indicating whether true indicates that headers and footers content is ignored. By default headers and footers are not ignored.             
	IgnoreHeadersAndFooters bool `json:"IgnoreHeadersAndFooters,omitempty"`

	// Gets or sets specifies which document shall be used as a target during comparison.             
	Target string `json:"Target,omitempty"`
}

type ICompareOptions interface {
	IsCompareOptions() bool
}
func (CompareOptions) IsCompareOptions() bool {
	return true;
}
