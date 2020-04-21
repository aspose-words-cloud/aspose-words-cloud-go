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



// Represents a single document style properties to update.
type StyleUpdate struct {

	// Gets or sets /sets the name of the style to be applied automatically to a new paragraph inserted after a paragraph formatted with the specified style.
	NextParagraphStyleName string `json:"NextParagraphStyleName,omitempty"`

	// Gets or sets /sets the name of the style this style is based on.
	BaseStyleName string `json:"BaseStyleName,omitempty"`

	// Gets or sets a value indicating whether specifies whether this style is shown in the Quick Style gallery inside MS Word UI.
	IsQuickStyle bool `json:"IsQuickStyle,omitempty"`

	// Gets or sets the name of the style.
	Name string `json:"Name,omitempty"`
}

type IStyleUpdate interface {
	IsStyleUpdate() bool
}
func (StyleUpdate) IsStyleUpdate() bool {
	return true;
}
