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



// Represents a single document list.
type Style struct {

	Font *Font `json:"Font,omitempty"`

	// Gets or sets a value indicating whether true if this style is one of the built-in styles in MS Word.
	BuiltIn bool `json:"BuiltIn,omitempty"`

	// Gets or sets /sets the name of the style to be applied automatically to a new paragraph inserted after a paragraph formatted with the specified style.
	NextParagraphStyleName string `json:"NextParagraphStyleName,omitempty"`

	// Gets or sets /sets the name of the style this style is based on.
	BaseStyleName string `json:"BaseStyleName,omitempty"`

	// Gets or sets a value indicating whether specifies whether this style is shown in the Quick Style gallery inside MS Word UI.
	IsQuickStyle bool `json:"IsQuickStyle,omitempty"`

	// Gets or sets the name of the Style linked to this one. Returns Empty string if no styles are linked.
	LinkedStyleName string `json:"LinkedStyleName,omitempty"`

	// Gets or sets the style type (paragraph or character).
	Type_ string `json:"Type,omitempty"`

	// Gets or sets a value indicating whether true when the style is one of the built-in Heading styles.
	IsHeading bool `json:"IsHeading,omitempty"`

	// Gets or sets all aliases of this style. If style has no aliases then empty array of string is returned.
	Aliases []string `json:"Aliases,omitempty"`

	// Gets or sets the locale independent style identifier for a built-in style.
	StyleIdentifier string `json:"StyleIdentifier,omitempty"`

	// Gets or sets the name of the style.
	Name string `json:"Name,omitempty"`
}

type IStyle interface {
	IsStyle() bool
}
func (Style) IsStyle() bool {
	return true;
}
