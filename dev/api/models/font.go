/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="font.go">
 *   Copyright (c) 2022 Aspose.Words for Cloud
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

// DTO container with a font element.
type FontResult struct {
    // DTO container with a font element.
    Link WordsApiLinkResult `json:"Link,omitempty"`

    // DTO container with a font element.
    AllCaps bool `json:"AllCaps,omitempty"`

    // DTO container with a font element.
    Bidi bool `json:"Bidi,omitempty"`

    // DTO container with a font element.
    Bold bool `json:"Bold,omitempty"`

    // DTO container with a font element.
    BoldBi bool `json:"BoldBi,omitempty"`

    // DTO container with a font element.
    Border BorderResult `json:"Border,omitempty"`

    // DTO container with a font element.
    Color XmlColorResult `json:"Color,omitempty"`

    // DTO container with a font element.
    ComplexScript bool `json:"ComplexScript,omitempty"`

    // DTO container with a font element.
    DoubleStrikeThrough bool `json:"DoubleStrikeThrough,omitempty"`

    // DTO container with a font element.
    Emboss bool `json:"Emboss,omitempty"`

    // DTO container with a font element.
    Engrave bool `json:"Engrave,omitempty"`

    // DTO container with a font element.
    Hidden bool `json:"Hidden,omitempty"`

    // DTO container with a font element.
    HighlightColor XmlColorResult `json:"HighlightColor,omitempty"`

    // DTO container with a font element.
    Italic bool `json:"Italic,omitempty"`

    // DTO container with a font element.
    ItalicBi bool `json:"ItalicBi,omitempty"`

    // DTO container with a font element.
    Kerning float64 `json:"Kerning,omitempty"`

    // DTO container with a font element.
    LocaleId int32 `json:"LocaleId,omitempty"`

    // DTO container with a font element.
    LocaleIdBi int32 `json:"LocaleIdBi,omitempty"`

    // DTO container with a font element.
    LocaleIdFarEast int32 `json:"LocaleIdFarEast,omitempty"`

    // DTO container with a font element.
    Name string `json:"Name,omitempty"`

    // DTO container with a font element.
    NameAscii string `json:"NameAscii,omitempty"`

    // DTO container with a font element.
    NameBi string `json:"NameBi,omitempty"`

    // DTO container with a font element.
    NameFarEast string `json:"NameFarEast,omitempty"`

    // DTO container with a font element.
    NameOther string `json:"NameOther,omitempty"`

    // DTO container with a font element.
    NoProofing bool `json:"NoProofing,omitempty"`

    // DTO container with a font element.
    Outline bool `json:"Outline,omitempty"`

    // DTO container with a font element.
    Position float64 `json:"Position,omitempty"`

    // DTO container with a font element.
    Scaling int32 `json:"Scaling,omitempty"`

    // DTO container with a font element.
    Shadow bool `json:"Shadow,omitempty"`

    // DTO container with a font element.
    Size float64 `json:"Size,omitempty"`

    // DTO container with a font element.
    SizeBi float64 `json:"SizeBi,omitempty"`

    // DTO container with a font element.
    SmallCaps bool `json:"SmallCaps,omitempty"`

    // DTO container with a font element.
    Spacing float64 `json:"Spacing,omitempty"`

    // DTO container with a font element.
    StrikeThrough bool `json:"StrikeThrough,omitempty"`

    // DTO container with a font element.
    StyleIdentifier string `json:"StyleIdentifier,omitempty"`

    // DTO container with a font element.
    StyleName string `json:"StyleName,omitempty"`

    // DTO container with a font element.
    Subscript bool `json:"Subscript,omitempty"`

    // DTO container with a font element.
    Superscript bool `json:"Superscript,omitempty"`

    // DTO container with a font element.
    TextEffect string `json:"TextEffect,omitempty"`

    // DTO container with a font element.
    Underline string `json:"Underline,omitempty"`

    // DTO container with a font element.
    UnderlineColor XmlColorResult `json:"UnderlineColor,omitempty"`
}

type Font struct {
    // DTO container with a font element.
    Link IWordsApiLink `json:"Link,omitempty"`

    // DTO container with a font element.
    AllCaps *bool `json:"AllCaps,omitempty"`

    // DTO container with a font element.
    Bidi *bool `json:"Bidi,omitempty"`

    // DTO container with a font element.
    Bold *bool `json:"Bold,omitempty"`

    // DTO container with a font element.
    BoldBi *bool `json:"BoldBi,omitempty"`

    // DTO container with a font element.
    Border IBorder `json:"Border,omitempty"`

    // DTO container with a font element.
    Color IXmlColor `json:"Color,omitempty"`

    // DTO container with a font element.
    ComplexScript *bool `json:"ComplexScript,omitempty"`

    // DTO container with a font element.
    DoubleStrikeThrough *bool `json:"DoubleStrikeThrough,omitempty"`

    // DTO container with a font element.
    Emboss *bool `json:"Emboss,omitempty"`

    // DTO container with a font element.
    Engrave *bool `json:"Engrave,omitempty"`

    // DTO container with a font element.
    Hidden *bool `json:"Hidden,omitempty"`

    // DTO container with a font element.
    HighlightColor IXmlColor `json:"HighlightColor,omitempty"`

    // DTO container with a font element.
    Italic *bool `json:"Italic,omitempty"`

    // DTO container with a font element.
    ItalicBi *bool `json:"ItalicBi,omitempty"`

    // DTO container with a font element.
    Kerning *float64 `json:"Kerning,omitempty"`

    // DTO container with a font element.
    LocaleId *int32 `json:"LocaleId,omitempty"`

    // DTO container with a font element.
    LocaleIdBi *int32 `json:"LocaleIdBi,omitempty"`

    // DTO container with a font element.
    LocaleIdFarEast *int32 `json:"LocaleIdFarEast,omitempty"`

    // DTO container with a font element.
    Name *string `json:"Name,omitempty"`

    // DTO container with a font element.
    NameAscii *string `json:"NameAscii,omitempty"`

    // DTO container with a font element.
    NameBi *string `json:"NameBi,omitempty"`

    // DTO container with a font element.
    NameFarEast *string `json:"NameFarEast,omitempty"`

    // DTO container with a font element.
    NameOther *string `json:"NameOther,omitempty"`

    // DTO container with a font element.
    NoProofing *bool `json:"NoProofing,omitempty"`

    // DTO container with a font element.
    Outline *bool `json:"Outline,omitempty"`

    // DTO container with a font element.
    Position *float64 `json:"Position,omitempty"`

    // DTO container with a font element.
    Scaling *int32 `json:"Scaling,omitempty"`

    // DTO container with a font element.
    Shadow *bool `json:"Shadow,omitempty"`

    // DTO container with a font element.
    Size *float64 `json:"Size,omitempty"`

    // DTO container with a font element.
    SizeBi *float64 `json:"SizeBi,omitempty"`

    // DTO container with a font element.
    SmallCaps *bool `json:"SmallCaps,omitempty"`

    // DTO container with a font element.
    Spacing *float64 `json:"Spacing,omitempty"`

    // DTO container with a font element.
    StrikeThrough *bool `json:"StrikeThrough,omitempty"`

    // DTO container with a font element.
    StyleIdentifier *string `json:"StyleIdentifier,omitempty"`

    // DTO container with a font element.
    StyleName *string `json:"StyleName,omitempty"`

    // DTO container with a font element.
    Subscript *bool `json:"Subscript,omitempty"`

    // DTO container with a font element.
    Superscript *bool `json:"Superscript,omitempty"`

    // DTO container with a font element.
    TextEffect *string `json:"TextEffect,omitempty"`

    // DTO container with a font element.
    Underline *string `json:"Underline,omitempty"`

    // DTO container with a font element.
    UnderlineColor IXmlColor `json:"UnderlineColor,omitempty"`
}

type IFont interface {
    IsFont() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileContent) []FileContent
}

func (Font) IsFont() bool {
    return true
}

func (Font) IsLinkElement() bool {
    return true
}

func (obj *Font) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.Border != nil) {
        obj.Border.Initialize()
    }

    if (obj.Color != nil) {
        obj.Color.Initialize()
    }

    if (obj.HighlightColor != nil) {
        obj.HighlightColor.Initialize()
    }

    if (obj.UnderlineColor != nil) {
        obj.UnderlineColor.Initialize()
    }


}

func (obj *Font) CollectFilesContent(resultFilesContent []FileContent) []FileContent {
    if (obj.Border != nil) {
        resultFilesContent = obj.Border.CollectFilesContent(resultFilesContent)
    }

    if (obj.Color != nil) {
        resultFilesContent = obj.Color.CollectFilesContent(resultFilesContent)
    }






    if (obj.HighlightColor != nil) {
        resultFilesContent = obj.HighlightColor.CollectFilesContent(resultFilesContent)
    }




























    if (obj.UnderlineColor != nil) {
        resultFilesContent = obj.UnderlineColor.CollectFilesContent(resultFilesContent)
    }

    return resultFilesContent
}


