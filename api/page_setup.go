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

// Represents the page setup properties of a section.             
type PageSetup struct {

	// Gets or sets link to the document.
	Link *WordsApiLink `json:"link,omitempty"`

	// Gets or sets specifies that this section contains bidirectional (complex scripts) text.             
	Bidi bool `json:"Bidi,omitempty"`

	// Gets or sets specifies where the page border is positioned relative to intersecting texts and objects.             
	BorderAlwaysInFront bool `json:"BorderAlwaysInFront,omitempty"`

	// Gets or sets specifies which pages the page border is printed on.             
	BorderAppliesTo string `json:"BorderAppliesTo,omitempty"`

	// Gets or sets a value that indicates whether the specified page border is measured from the edge of the page or from the text it surrounds.             
	BorderDistanceFrom string `json:"BorderDistanceFrom,omitempty"`

	// Gets or sets returns or sets the distance (in points) between the bottom edge of the page and the bottom boundary of the body text.             
	BottomMargin float64 `json:"BottomMargin,omitempty"`

	// Gets or sets true if a different header or footer is used on the first page.             
	DifferentFirstPageHeaderFooter bool `json:"DifferentFirstPageHeaderFooter,omitempty"`

	// Gets or sets the paper tray (bin) to use for the first page of a section. The value is implementation (printer) specific.             
	FirstPageTray int32 `json:"FirstPageTray,omitempty"`

	// Gets or sets returns or sets the distance (in points) between the footer and the bottom of the page.             
	FooterDistance float64 `json:"FooterDistance,omitempty"`

	// Gets or sets the amount of extra space added to the margin for document binding.             
	Gutter float64 `json:"Gutter,omitempty"`

	// Gets or sets returns or sets the distance (in points) between the header and the top of the page.             
	HeaderDistance float64 `json:"HeaderDistance,omitempty"`

	// Gets or sets returns or sets the distance (in points) between the left edge of the page and the left boundary of the body text.             
	LeftMargin float64 `json:"LeftMargin,omitempty"`

	// Gets or sets returns or sets the numeric increment for line numbers.             
	LineNumberCountBy int32 `json:"LineNumberCountBy,omitempty"`

	// Gets or sets distance between the right edge of line numbers and the left edge of the document.             
	LineNumberDistanceFromText float64 `json:"LineNumberDistanceFromText,omitempty"`

	// Gets or sets the way line numbering runs  that is, whether it starts over at the beginning of a new page or section or runs continuously.             
	LineNumberRestartMode string `json:"LineNumberRestartMode,omitempty"`

	// Gets or sets the starting line number.             
	LineStartingNumber int32 `json:"LineStartingNumber,omitempty"`

	// Gets or sets returns or sets the orientation of the page.             
	Orientation string `json:"Orientation,omitempty"`

	// Gets or sets the paper tray (bin) to be used for all but the first page of a section. The value is implementation (printer) specific.             
	OtherPagesTray int32 `json:"OtherPagesTray,omitempty"`

	// Gets or sets returns or sets the height of the page in points.             
	PageHeight float64 `json:"PageHeight,omitempty"`

	// Gets or sets the page number format.             
	PageNumberStyle string `json:"PageNumberStyle,omitempty"`

	// Gets or sets the starting page number of the section.             
	PageStartingNumber int32 `json:"PageStartingNumber,omitempty"`

	// Gets or sets returns or sets the width of the page in points.             
	PageWidth float64 `json:"PageWidth,omitempty"`

	// Gets or sets returns or sets the paper size.             
	PaperSize string `json:"PaperSize,omitempty"`

	// Gets or sets true if page numbering restarts at the beginning of the section.             
	RestartPageNumbering bool `json:"RestartPageNumbering,omitempty"`

	// Gets or sets returns or sets the distance (in points) between the right edge of the page and the right boundary of the body text.             
	RightMargin float64 `json:"RightMargin,omitempty"`

	// Gets or sets whether Microsoft Word uses gutters for the section based on a right-to-left language or a left-to-right language.             
	RtlGutter bool `json:"RtlGutter,omitempty"`

	// Gets or sets returns or sets the type of section break for the specified object.             
	SectionStart string `json:"SectionStart,omitempty"`

	// Gets or sets true if endnotes are printed at the end of the next section that doesn't suppress endnotes.                 Suppressed endnotes are printed before the endnotes in that section.             
	SuppressEndnotes bool `json:"SuppressEndnotes,omitempty"`

	// Gets or sets returns or sets the distance (in points) between the top edge of the page and the top boundary of the body text.             
	TopMargin float64 `json:"TopMargin,omitempty"`

	// Gets or sets returns or sets the vertical alignment of text on each page in a document or section.             
	VerticalAlignment string `json:"VerticalAlignment,omitempty"`
}
