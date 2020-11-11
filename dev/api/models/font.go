/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="font.go">
 *   Copyright (c) 2020 Aspose.Words for Cloud
 * </copyright>
 * <summary>
 *   Permission is hereby granted, free of charge, to any person obtaining a copy
 *  of this software and associated documentation files (the "Software"), to deal
 *  in the Software without restriction, including without limitation the rights
 *  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *  copies of the Software, and to permit persons to whom the Software is
 *  furnished to do so, subject to the following conditions:
 * 
 *  The above copyright notice and this permission notice shall be included in all
 *  copies or substantial portions of the Software.
 * 
 *  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 *  SOFTWARE.
 * </summary>
 * --------------------------------------------------------------------------------
 */

package models

// Font element.
type Font struct {
    // Font element.
    Link *WordsApiLink `json:"Link,omitempty"`

    // Font element.
    AllCaps *bool `json:"AllCaps,omitempty"`

    // Font element.
    Bidi *bool `json:"Bidi,omitempty"`

    // Font element.
    Bold *bool `json:"Bold,omitempty"`

    // Font element.
    BoldBi *bool `json:"BoldBi,omitempty"`

    // Font element.
    Border *Border `json:"Border,omitempty"`

    // Font element.
    Color *XmlColor `json:"Color,omitempty"`

    // Font element.
    ComplexScript *bool `json:"ComplexScript,omitempty"`

    // Font element.
    DoubleStrikeThrough *bool `json:"DoubleStrikeThrough,omitempty"`

    // Font element.
    Emboss *bool `json:"Emboss,omitempty"`

    // Font element.
    Engrave *bool `json:"Engrave,omitempty"`

    // Font element.
    Hidden *bool `json:"Hidden,omitempty"`

    // Font element.
    HighlightColor *XmlColor `json:"HighlightColor,omitempty"`

    // Font element.
    Italic *bool `json:"Italic,omitempty"`

    // Font element.
    ItalicBi *bool `json:"ItalicBi,omitempty"`

    // Font element.
    Kerning *float64 `json:"Kerning,omitempty"`

    // Font element.
    LocaleId *int32 `json:"LocaleId,omitempty"`

    // Font element.
    LocaleIdBi *int32 `json:"LocaleIdBi,omitempty"`

    // Font element.
    LocaleIdFarEast *int32 `json:"LocaleIdFarEast,omitempty"`

    // Font element.
    Name *string `json:"Name,omitempty"`

    // Font element.
    NameAscii *string `json:"NameAscii,omitempty"`

    // Font element.
    NameBi *string `json:"NameBi,omitempty"`

    // Font element.
    NameFarEast *string `json:"NameFarEast,omitempty"`

    // Font element.
    NameOther *string `json:"NameOther,omitempty"`

    // Font element.
    NoProofing *bool `json:"NoProofing,omitempty"`

    // Font element.
    Outline *bool `json:"Outline,omitempty"`

    // Font element.
    Position *float64 `json:"Position,omitempty"`

    // Font element.
    Scaling *int32 `json:"Scaling,omitempty"`

    // Font element.
    Shadow *bool `json:"Shadow,omitempty"`

    // Font element.
    Size *float64 `json:"Size,omitempty"`

    // Font element.
    SizeBi *float64 `json:"SizeBi,omitempty"`

    // Font element.
    SmallCaps *bool `json:"SmallCaps,omitempty"`

    // Font element.
    Spacing *float64 `json:"Spacing,omitempty"`

    // Font element.
    StrikeThrough *bool `json:"StrikeThrough,omitempty"`

    // Font element.
    StyleIdentifier string `json:"StyleIdentifier,omitempty"`

    // Font element.
    StyleName *string `json:"StyleName,omitempty"`

    // Font element.
    Subscript *bool `json:"Subscript,omitempty"`

    // Font element.
    Superscript *bool `json:"Superscript,omitempty"`

    // Font element.
    TextEffect string `json:"TextEffect,omitempty"`

    // Font element.
    Underline string `json:"Underline,omitempty"`

    // Font element.
    UnderlineColor *XmlColor `json:"UnderlineColor,omitempty"`
}

type IFont interface {
    IsFont() bool
}
func (Font) IsFont() bool {
    return true
}

func (Font) IsLinkElement() bool {
    return true
}
