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

// Represents all formatting for a table row.
type TableCellFormat struct {

	// Gets or sets link to the document.
	Link *WordsApiLink `json:"link,omitempty"`

	// Gets or sets returns or sets the amount of space (in points) to add below the contents of cell.
	BottomPadding float64 `json:"BottomPadding,omitempty"`

	// Gets or sets if true, fits text in the cell, compressing each paragraph to the width of the cell.
	FitText bool `json:"FitText,omitempty"`

	// Gets or sets specifies how the cell is merged horizontally with other cells in the row.
	HorizontalMerge string `json:"HorizontalMerge,omitempty"`

	// Gets or sets returns or sets the amount of space (in points) to add to the left of the contents of cell.
	LeftPadding float64 `json:"LeftPadding,omitempty"`

	// Gets or sets returns or sets the orientation of text in a table cell.
	Orientation string `json:"Orientation,omitempty"`

	// Gets or sets returns or sets the preferred width of the cell.
	PreferredWidth *PreferredWidth `json:"PreferredWidth,omitempty"`

	// Gets or sets returns or sets the amount of space (in points) to add to the right of the contents of cell.
	RightPadding float64 `json:"RightPadding,omitempty"`

	// Gets or sets returns or sets the amount of space (in points) to add above the contents of cell.
	TopPadding float64 `json:"TopPadding,omitempty"`

	// Gets or sets returns or sets the vertical alignment of text in the cell.
	VerticalAlignment string `json:"VerticalAlignment,omitempty"`

	// Gets or sets specifies how the cell is merged with other cells vertically.
	VerticalMerge string `json:"VerticalMerge,omitempty"`

	// Gets or sets the width of the cell in points.
	Width float64 `json:"Width,omitempty"`

	// Gets or sets if true, wrap text for the cell.
	WrapText bool `json:"WrapText,omitempty"`
}
type ITableCellFormat interface {
	IsTableCellFormat() bool
}
func (TableCellFormat) IsTableCellFormat() bool {
	return true;
}
func (TableCellFormat) IsLinkElement() bool {
	return true;
}

