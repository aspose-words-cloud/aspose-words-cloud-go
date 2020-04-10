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



// Represents a document list levels.
type ListLevel struct {

	// Gets or sets returns or sets the starting number for this list level.
	StartAt int32 `json:"StartAt,omitempty"`

	// Gets or sets returns or sets the number style for this list level.
	NumberStyle string `json:"NumberStyle,omitempty"`

	// Gets or sets returns or sets the number format for the list level.
	NumberFormat string `json:"NumberFormat,omitempty"`

	// Gets or sets the justification of the actual number of the list item.
	Alignment string `json:"Alignment,omitempty"`

	// Gets or sets a value indicating whether true if the level turns all inherited numbers to Arabic, false if it preserves their number style.
	IsLegal bool `json:"IsLegal,omitempty"`

	// Gets or sets or returns the list level that must appear before the specified list level restarts numbering.
	RestartAfterLevel int32 `json:"RestartAfterLevel,omitempty"`

	// Gets or sets returns or sets the character inserted after the number for the list level.
	TrailingCharacter string `json:"TrailingCharacter,omitempty"`

	Font *Font `json:"Font,omitempty"`

	// Gets or sets returns or sets the tab position (in points) for the list level.
	TabPosition float64 `json:"TabPosition,omitempty"`

	// Gets or sets returns or sets the position (in points) of the number or bullet for the list level.
	NumberPosition float64 `json:"NumberPosition,omitempty"`

	// Gets or sets returns or sets the position (in points) for the second line of wrapping text for the list level.
	TextPosition float64 `json:"TextPosition,omitempty"`

	LinkedStyle *Style `json:"LinkedStyle,omitempty"`
}

type IListLevel interface {
	IsListLevel() bool
}
func (ListLevel) IsListLevel() bool {
	return true;
}
