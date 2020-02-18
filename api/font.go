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

// Font element.             
type Font struct {

	// Gets or sets link to the document.
	Link *WordsApiLink `json:"link,omitempty"`

	// Gets or sets true if the font is formatted as all capital letters.             
	AllCaps bool `json:"AllCaps,omitempty"`

	// Gets or sets specifies whether the contents of this run shall have right-to-left characteristics.             
	Bidi bool `json:"Bidi,omitempty"`

	// Gets or sets true if the font is formatted as bold.             
	Bold bool `json:"Bold,omitempty"`

	// Gets or sets true if the right-to-left text is formatted as bold.             
	BoldBi bool `json:"BoldBi,omitempty"`

	// Gets or sets border object that specifies border for the font.
	Border *Border `json:"Border,omitempty"`

	// Gets or sets the color of the font.             
	Color *XmlColor `json:"Color,omitempty"`

	// Gets or sets specifies whether the contents of this run shall be treated as complex script text regardless of their Unicode character values when determining the formatting for this run.             
	ComplexScript bool `json:"ComplexScript,omitempty"`

	// Gets or sets true if the font is formatted as double strikethrough text.             
	DoubleStrikeThrough bool `json:"DoubleStrikeThrough,omitempty"`

	// Gets or sets true if the font is formatted as embossed.             
	Emboss bool `json:"Emboss,omitempty"`

	// Gets or sets true if the font is formatted as engraved.             
	Engrave bool `json:"Engrave,omitempty"`

	// Gets or sets true if the font is formatted as hidden text.             
	Hidden bool `json:"Hidden,omitempty"`

	// Gets or sets the highlight (marker) color.             
	HighlightColor *XmlColor `json:"HighlightColor,omitempty"`

	// Gets or sets true if the font is formatted as italic.             
	Italic bool `json:"Italic,omitempty"`

	// Gets or sets true if the right-to-left text is formatted as italic.             
	ItalicBi bool `json:"ItalicBi,omitempty"`

	// Gets or sets the font size at which kerning starts.             
	Kerning float64 `json:"Kerning,omitempty"`

	// Gets or sets the locale identifier (language) of the formatted characters.             
	LocaleId int32 `json:"LocaleId,omitempty"`

	// Gets or sets the locale identifier (language) of the formatted right-to-left characters.             
	LocaleIdBi int32 `json:"LocaleIdBi,omitempty"`

	// Gets or sets the locale identifier (language) of the formatted Asian characters.             
	LocaleIdFarEast int32 `json:"LocaleIdFarEast,omitempty"`

	// Gets or sets the name of the font.             
	Name string `json:"Name,omitempty"`

	// Gets or sets returns or sets the font used for Latin text (characters with character codes from 0 (zero) through 127).             
	NameAscii string `json:"NameAscii,omitempty"`

	// Gets or sets returns or sets the name of the font in a right-to-left language document.             
	NameBi string `json:"NameBi,omitempty"`

	// Gets or sets returns or sets an East Asian font name.             
	NameFarEast string `json:"NameFarEast,omitempty"`

	// Gets or sets returns or sets the font used for characters with character codes from 128 through 255.             
	NameOther string `json:"NameOther,omitempty"`

	// Gets or sets true when the formatted characters are not to be spell checked.
	NoProofing bool `json:"NoProofing,omitempty"`

	// Gets or sets true if the font is formatted as outline.             
	Outline bool `json:"Outline,omitempty"`

	// Gets or sets the position of text (in points) relative to the base line. A positive number raises the text, and a negative number lowers it.             
	Position float64 `json:"Position,omitempty"`

	// Gets or sets character width scaling in percent.             
	Scaling int32 `json:"Scaling,omitempty"`

	// Gets or sets true if the font is formatted as shadowed.             
	Shadow bool `json:"Shadow,omitempty"`

	// Gets or sets the font size in points.             
	Size float64 `json:"Size,omitempty"`

	// Gets or sets the font size in points used in a right-to-left document.             
	SizeBi float64 `json:"SizeBi,omitempty"`

	// Gets or sets true if the font is formatted as small capital letters.             
	SmallCaps bool `json:"SmallCaps,omitempty"`

	// Gets or sets returns or sets the spacing (in points) between characters.             
	Spacing float64 `json:"Spacing,omitempty"`

	// Gets or sets true if the font is formatted as strikethrough text.             
	StrikeThrough bool `json:"StrikeThrough,omitempty"`

	// Gets or sets the locale independent style identifier of the character style applied to this formatting.
	StyleIdentifier string `json:"StyleIdentifier,omitempty"`

	// Gets or sets the name of the character style applied to this formatting.             
	StyleName string `json:"StyleName,omitempty"`

	// Gets or sets true if the font is formatted as subscript.             
	Subscript bool `json:"Subscript,omitempty"`

	// Gets or sets true if the font is formatted as superscript.             
	Superscript bool `json:"Superscript,omitempty"`

	// Gets or sets the font animation effect.
	TextEffect string `json:"TextEffect,omitempty"`

	// Gets or sets the type of underline applied to the font.
	Underline string `json:"Underline,omitempty"`

	// Gets or sets the color of the underline applied to the font.
	UnderlineColor *XmlColor `json:"UnderlineColor,omitempty"`
}
type IFont interface {
	IsFont() bool
}
func (Font) IsFont() bool {
	return true;
}
func (Font) IsLinkElement() bool {
	return true;
}

