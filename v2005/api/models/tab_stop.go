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



// Paragraph format tab stop.             
type TabStop struct {

	// Gets or sets the alignment of text at this tab stop.
	Alignment string `json:"Alignment,omitempty"`

	// Gets or sets the type of the leader line displayed under the tab character.
	Leader string `json:"Leader,omitempty"`

	// Gets or sets the position of the tab stop in points.
	Position float64 `json:"Position,omitempty"`

	// Gets or sets a value indicating whether this tab stop clears any existing tab stops in this position.
	IsClear bool `json:"IsClear,omitempty"`
}

type ITabStop interface {
	IsTabStop() bool
}
func (TabStop) IsTabStop() bool {
	return true;
}
func (TabStop) IsTabStopBase() bool {
	return true;
}
