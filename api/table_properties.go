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

// Represents the table properties.             
type TableProperties struct {

	// Gets or sets link to the document.
	Link *WordsApiLink `json:"link,omitempty"`

	// Gets or sets specifies how an inline table is aligned in the document.
	Alignment string `json:"Alignment,omitempty"`

	// Gets or sets allows Microsoft Word and Aspose.Words to automatically resize cells in a table to fit their contents.
	AllowAutoFit bool `json:"AllowAutoFit,omitempty"`

	// Gets or sets whether this is a right-to-left table.
	Bidi bool `json:"Bidi,omitempty"`

	// Gets or sets the amount of space (in points) to add below the contents of cells.
	BottomPadding float64 `json:"BottomPadding,omitempty"`

	// Gets or sets the amount of space (in points) between the cells.
	CellSpacing float64 `json:"CellSpacing,omitempty"`

	// Gets or sets the value that represents the left indent of the table.
	LeftIndent float64 `json:"LeftIndent,omitempty"`

	// Gets or sets the amount of space (in points) to add to the left of the contents of cells.
	LeftPadding float64 `json:"LeftPadding,omitempty"`

	// Gets or sets the table preferred width. Preferred width can be specified as a percentage, number of points or a special \"auto\" value.
	PreferredWidth *PreferredWidth `json:"PreferredWidth,omitempty"`

	// Gets or sets the amount of space (in points) to add to the right of the contents of cells.
	RightPadding float64 `json:"RightPadding,omitempty"`

	// Gets or sets the locale independent style identifier of the table style applied to this table.
	StyleIdentifier string `json:"StyleIdentifier,omitempty"`

	// Gets or sets the name of the table style applied to this table.
	StyleName string `json:"StyleName,omitempty"`

	// Gets or sets bit flags that specify how a table style is applied to this table.
	StyleOptions string `json:"StyleOptions,omitempty"`

	// Gets or sets get or sets TextWrapping  for table.
	TextWrapping string `json:"TextWrapping,omitempty"`

	// Gets or sets the amount of space (in points) to add above the contents of cells.
	TopPadding float64 `json:"TopPadding,omitempty"`
}
type ITableProperties interface {
	IsTableProperties() bool
}
func (TableProperties) IsTableProperties() bool {
	return true;
}
func (TableProperties) IsLinkElement() bool {
	return true;
}

