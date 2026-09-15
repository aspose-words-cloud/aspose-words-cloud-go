/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="font.go">
 *   Copyright (c) 2026 Aspose.Words for Cloud
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

import (
    "errors"
)

// DTO container with a font element.

type IFont interface {
    IsFont() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetAllCaps() *bool
    SetAllCaps(value *bool)
    GetBidi() *bool
    SetBidi(value *bool)
    GetBold() *bool
    SetBold(value *bool)
    GetBoldBi() *bool
    SetBoldBi(value *bool)
    GetBorder() IBorder
    SetBorder(value IBorder)
    GetColor() IXmlColor
    SetColor(value IXmlColor)
    GetComplexScript() *bool
    SetComplexScript(value *bool)
    GetDoubleStrikeThrough() *bool
    SetDoubleStrikeThrough(value *bool)
    GetEmboss() *bool
    SetEmboss(value *bool)
    GetEngrave() *bool
    SetEngrave(value *bool)
    GetHidden() *bool
    SetHidden(value *bool)
    GetHighlightColor() IXmlColor
    SetHighlightColor(value IXmlColor)
    GetItalic() *bool
    SetItalic(value *bool)
    GetItalicBi() *bool
    SetItalicBi(value *bool)
    GetKerning() *float64
    SetKerning(value *float64)
    GetLocaleId() *int32
    SetLocaleId(value *int32)
    GetLocaleIdBi() *int32
    SetLocaleIdBi(value *int32)
    GetLocaleIdFarEast() *int32
    SetLocaleIdFarEast(value *int32)
    GetName() *string
    SetName(value *string)
    GetNameAscii() *string
    SetNameAscii(value *string)
    GetNameBi() *string
    SetNameBi(value *string)
    GetNameFarEast() *string
    SetNameFarEast(value *string)
    GetNameOther() *string
    SetNameOther(value *string)
    GetNoProofing() *bool
    SetNoProofing(value *bool)
    GetOutline() *bool
    SetOutline(value *bool)
    GetPosition() *float64
    SetPosition(value *float64)
    GetScaling() *int32
    SetScaling(value *int32)
    GetShadow() *bool
    SetShadow(value *bool)
    GetSize() *float64
    SetSize(value *float64)
    GetSizeBi() *float64
    SetSizeBi(value *float64)
    GetSmallCaps() *bool
    SetSmallCaps(value *bool)
    GetSpacing() *float64
    SetSpacing(value *float64)
    GetStrikeThrough() *bool
    SetStrikeThrough(value *bool)
    GetStyleIdentifier() *string
    SetStyleIdentifier(value *string)
    GetStyleName() *string
    SetStyleName(value *string)
    GetSubscript() *bool
    SetSubscript(value *bool)
    GetSuperscript() *bool
    SetSuperscript(value *bool)
    GetTextEffect() *string
    SetTextEffect(value *string)
    GetUnderline() *string
    SetUnderline(value *string)
    GetUnderlineColor() IXmlColor
    SetUnderlineColor(value IXmlColor)
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

func (obj *Font) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["Link"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IWordsApiLink = new(WordsApiLink)
            modelInstance.Deserialize(parsedValue)
            obj.Link = modelInstance
        }

    } else if jsonValue, exists := json["link"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IWordsApiLink = new(WordsApiLink)
            modelInstance.Deserialize(parsedValue)
            obj.Link = modelInstance
        }

    }

    if jsonValue, exists := json["AllCaps"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.AllCaps = &parsedValue
        }

    } else if jsonValue, exists := json["allCaps"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.AllCaps = &parsedValue
        }

    }

    if jsonValue, exists := json["Bidi"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.Bidi = &parsedValue
        }

    } else if jsonValue, exists := json["bidi"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.Bidi = &parsedValue
        }

    }

    if jsonValue, exists := json["Bold"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.Bold = &parsedValue
        }

    } else if jsonValue, exists := json["bold"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.Bold = &parsedValue
        }

    }

    if jsonValue, exists := json["BoldBi"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.BoldBi = &parsedValue
        }

    } else if jsonValue, exists := json["boldBi"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.BoldBi = &parsedValue
        }

    }

    if jsonValue, exists := json["Border"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IBorder = new(Border)
            modelInstance.Deserialize(parsedValue)
            obj.Border = modelInstance
        }

    } else if jsonValue, exists := json["border"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IBorder = new(Border)
            modelInstance.Deserialize(parsedValue)
            obj.Border = modelInstance
        }

    }

    if jsonValue, exists := json["Color"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IXmlColor = new(XmlColor)
            modelInstance.Deserialize(parsedValue)
            obj.Color = modelInstance
        }

    } else if jsonValue, exists := json["color"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IXmlColor = new(XmlColor)
            modelInstance.Deserialize(parsedValue)
            obj.Color = modelInstance
        }

    }

    if jsonValue, exists := json["ComplexScript"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ComplexScript = &parsedValue
        }

    } else if jsonValue, exists := json["complexScript"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ComplexScript = &parsedValue
        }

    }

    if jsonValue, exists := json["DoubleStrikeThrough"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.DoubleStrikeThrough = &parsedValue
        }

    } else if jsonValue, exists := json["doubleStrikeThrough"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.DoubleStrikeThrough = &parsedValue
        }

    }

    if jsonValue, exists := json["Emboss"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.Emboss = &parsedValue
        }

    } else if jsonValue, exists := json["emboss"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.Emboss = &parsedValue
        }

    }

    if jsonValue, exists := json["Engrave"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.Engrave = &parsedValue
        }

    } else if jsonValue, exists := json["engrave"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.Engrave = &parsedValue
        }

    }

    if jsonValue, exists := json["Hidden"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.Hidden = &parsedValue
        }

    } else if jsonValue, exists := json["hidden"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.Hidden = &parsedValue
        }

    }

    if jsonValue, exists := json["HighlightColor"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IXmlColor = new(XmlColor)
            modelInstance.Deserialize(parsedValue)
            obj.HighlightColor = modelInstance
        }

    } else if jsonValue, exists := json["highlightColor"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IXmlColor = new(XmlColor)
            modelInstance.Deserialize(parsedValue)
            obj.HighlightColor = modelInstance
        }

    }

    if jsonValue, exists := json["Italic"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.Italic = &parsedValue
        }

    } else if jsonValue, exists := json["italic"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.Italic = &parsedValue
        }

    }

    if jsonValue, exists := json["ItalicBi"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ItalicBi = &parsedValue
        }

    } else if jsonValue, exists := json["italicBi"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ItalicBi = &parsedValue
        }

    }

    if jsonValue, exists := json["Kerning"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Kerning = &parsedValue
        }

    } else if jsonValue, exists := json["kerning"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Kerning = &parsedValue
        }

    }

    if jsonValue, exists := json["LocaleId"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.LocaleId = new(int32)
            *obj.LocaleId = int32(parsedValue)
        }

    } else if jsonValue, exists := json["localeId"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.LocaleId = new(int32)
            *obj.LocaleId = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["LocaleIdBi"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.LocaleIdBi = new(int32)
            *obj.LocaleIdBi = int32(parsedValue)
        }

    } else if jsonValue, exists := json["localeIdBi"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.LocaleIdBi = new(int32)
            *obj.LocaleIdBi = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["LocaleIdFarEast"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.LocaleIdFarEast = new(int32)
            *obj.LocaleIdFarEast = int32(parsedValue)
        }

    } else if jsonValue, exists := json["localeIdFarEast"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.LocaleIdFarEast = new(int32)
            *obj.LocaleIdFarEast = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["Name"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Name = &parsedValue
        }

    } else if jsonValue, exists := json["name"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Name = &parsedValue
        }

    }

    if jsonValue, exists := json["NameAscii"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NameAscii = &parsedValue
        }

    } else if jsonValue, exists := json["nameAscii"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NameAscii = &parsedValue
        }

    }

    if jsonValue, exists := json["NameBi"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NameBi = &parsedValue
        }

    } else if jsonValue, exists := json["nameBi"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NameBi = &parsedValue
        }

    }

    if jsonValue, exists := json["NameFarEast"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NameFarEast = &parsedValue
        }

    } else if jsonValue, exists := json["nameFarEast"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NameFarEast = &parsedValue
        }

    }

    if jsonValue, exists := json["NameOther"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NameOther = &parsedValue
        }

    } else if jsonValue, exists := json["nameOther"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NameOther = &parsedValue
        }

    }

    if jsonValue, exists := json["NoProofing"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.NoProofing = &parsedValue
        }

    } else if jsonValue, exists := json["noProofing"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.NoProofing = &parsedValue
        }

    }

    if jsonValue, exists := json["Outline"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.Outline = &parsedValue
        }

    } else if jsonValue, exists := json["outline"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.Outline = &parsedValue
        }

    }

    if jsonValue, exists := json["Position"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Position = &parsedValue
        }

    } else if jsonValue, exists := json["position"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Position = &parsedValue
        }

    }

    if jsonValue, exists := json["Scaling"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Scaling = new(int32)
            *obj.Scaling = int32(parsedValue)
        }

    } else if jsonValue, exists := json["scaling"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Scaling = new(int32)
            *obj.Scaling = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["Shadow"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.Shadow = &parsedValue
        }

    } else if jsonValue, exists := json["shadow"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.Shadow = &parsedValue
        }

    }

    if jsonValue, exists := json["Size"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Size = &parsedValue
        }

    } else if jsonValue, exists := json["size"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Size = &parsedValue
        }

    }

    if jsonValue, exists := json["SizeBi"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.SizeBi = &parsedValue
        }

    } else if jsonValue, exists := json["sizeBi"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.SizeBi = &parsedValue
        }

    }

    if jsonValue, exists := json["SmallCaps"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.SmallCaps = &parsedValue
        }

    } else if jsonValue, exists := json["smallCaps"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.SmallCaps = &parsedValue
        }

    }

    if jsonValue, exists := json["Spacing"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Spacing = &parsedValue
        }

    } else if jsonValue, exists := json["spacing"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Spacing = &parsedValue
        }

    }

    if jsonValue, exists := json["StrikeThrough"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.StrikeThrough = &parsedValue
        }

    } else if jsonValue, exists := json["strikeThrough"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.StrikeThrough = &parsedValue
        }

    }

    if jsonValue, exists := json["StyleIdentifier"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.StyleIdentifier = &parsedValue
        }

    } else if jsonValue, exists := json["styleIdentifier"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.StyleIdentifier = &parsedValue
        }

    }

    if jsonValue, exists := json["StyleName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.StyleName = &parsedValue
        }

    } else if jsonValue, exists := json["styleName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.StyleName = &parsedValue
        }

    }

    if jsonValue, exists := json["Subscript"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.Subscript = &parsedValue
        }

    } else if jsonValue, exists := json["subscript"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.Subscript = &parsedValue
        }

    }

    if jsonValue, exists := json["Superscript"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.Superscript = &parsedValue
        }

    } else if jsonValue, exists := json["superscript"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.Superscript = &parsedValue
        }

    }

    if jsonValue, exists := json["TextEffect"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.TextEffect = &parsedValue
        }

    } else if jsonValue, exists := json["textEffect"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.TextEffect = &parsedValue
        }

    }

    if jsonValue, exists := json["Underline"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Underline = &parsedValue
        }

    } else if jsonValue, exists := json["underline"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Underline = &parsedValue
        }

    }

    if jsonValue, exists := json["UnderlineColor"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IXmlColor = new(XmlColor)
            modelInstance.Deserialize(parsedValue)
            obj.UnderlineColor = modelInstance
        }

    } else if jsonValue, exists := json["underlineColor"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IXmlColor = new(XmlColor)
            modelInstance.Deserialize(parsedValue)
            obj.UnderlineColor = modelInstance
        }

    }
}

func (obj *Font) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *Font) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Link != nil {
        if err := obj.Link.Validate(); err != nil {
            return err
        }
    }
    if obj.Border != nil {
        if err := obj.Border.Validate(); err != nil {
            return err
        }
    }
    if obj.Color != nil {
        if err := obj.Color.Validate(); err != nil {
            return err
        }
    }
    if obj.HighlightColor != nil {
        if err := obj.HighlightColor.Validate(); err != nil {
            return err
        }
    }
    if obj.UnderlineColor != nil {
        if err := obj.UnderlineColor.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *Font) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *Font) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *Font) GetAllCaps() *bool {
    return obj.AllCaps
}

func (obj *Font) SetAllCaps(value *bool) {
    obj.AllCaps = value
}

func (obj *Font) GetBidi() *bool {
    return obj.Bidi
}

func (obj *Font) SetBidi(value *bool) {
    obj.Bidi = value
}

func (obj *Font) GetBold() *bool {
    return obj.Bold
}

func (obj *Font) SetBold(value *bool) {
    obj.Bold = value
}

func (obj *Font) GetBoldBi() *bool {
    return obj.BoldBi
}

func (obj *Font) SetBoldBi(value *bool) {
    obj.BoldBi = value
}

func (obj *Font) GetBorder() IBorder {
    return obj.Border
}

func (obj *Font) SetBorder(value IBorder) {
    obj.Border = value
}

func (obj *Font) GetColor() IXmlColor {
    return obj.Color
}

func (obj *Font) SetColor(value IXmlColor) {
    obj.Color = value
}

func (obj *Font) GetComplexScript() *bool {
    return obj.ComplexScript
}

func (obj *Font) SetComplexScript(value *bool) {
    obj.ComplexScript = value
}

func (obj *Font) GetDoubleStrikeThrough() *bool {
    return obj.DoubleStrikeThrough
}

func (obj *Font) SetDoubleStrikeThrough(value *bool) {
    obj.DoubleStrikeThrough = value
}

func (obj *Font) GetEmboss() *bool {
    return obj.Emboss
}

func (obj *Font) SetEmboss(value *bool) {
    obj.Emboss = value
}

func (obj *Font) GetEngrave() *bool {
    return obj.Engrave
}

func (obj *Font) SetEngrave(value *bool) {
    obj.Engrave = value
}

func (obj *Font) GetHidden() *bool {
    return obj.Hidden
}

func (obj *Font) SetHidden(value *bool) {
    obj.Hidden = value
}

func (obj *Font) GetHighlightColor() IXmlColor {
    return obj.HighlightColor
}

func (obj *Font) SetHighlightColor(value IXmlColor) {
    obj.HighlightColor = value
}

func (obj *Font) GetItalic() *bool {
    return obj.Italic
}

func (obj *Font) SetItalic(value *bool) {
    obj.Italic = value
}

func (obj *Font) GetItalicBi() *bool {
    return obj.ItalicBi
}

func (obj *Font) SetItalicBi(value *bool) {
    obj.ItalicBi = value
}

func (obj *Font) GetKerning() *float64 {
    return obj.Kerning
}

func (obj *Font) SetKerning(value *float64) {
    obj.Kerning = value
}

func (obj *Font) GetLocaleId() *int32 {
    return obj.LocaleId
}

func (obj *Font) SetLocaleId(value *int32) {
    obj.LocaleId = value
}

func (obj *Font) GetLocaleIdBi() *int32 {
    return obj.LocaleIdBi
}

func (obj *Font) SetLocaleIdBi(value *int32) {
    obj.LocaleIdBi = value
}

func (obj *Font) GetLocaleIdFarEast() *int32 {
    return obj.LocaleIdFarEast
}

func (obj *Font) SetLocaleIdFarEast(value *int32) {
    obj.LocaleIdFarEast = value
}

func (obj *Font) GetName() *string {
    return obj.Name
}

func (obj *Font) SetName(value *string) {
    obj.Name = value
}

func (obj *Font) GetNameAscii() *string {
    return obj.NameAscii
}

func (obj *Font) SetNameAscii(value *string) {
    obj.NameAscii = value
}

func (obj *Font) GetNameBi() *string {
    return obj.NameBi
}

func (obj *Font) SetNameBi(value *string) {
    obj.NameBi = value
}

func (obj *Font) GetNameFarEast() *string {
    return obj.NameFarEast
}

func (obj *Font) SetNameFarEast(value *string) {
    obj.NameFarEast = value
}

func (obj *Font) GetNameOther() *string {
    return obj.NameOther
}

func (obj *Font) SetNameOther(value *string) {
    obj.NameOther = value
}

func (obj *Font) GetNoProofing() *bool {
    return obj.NoProofing
}

func (obj *Font) SetNoProofing(value *bool) {
    obj.NoProofing = value
}

func (obj *Font) GetOutline() *bool {
    return obj.Outline
}

func (obj *Font) SetOutline(value *bool) {
    obj.Outline = value
}

func (obj *Font) GetPosition() *float64 {
    return obj.Position
}

func (obj *Font) SetPosition(value *float64) {
    obj.Position = value
}

func (obj *Font) GetScaling() *int32 {
    return obj.Scaling
}

func (obj *Font) SetScaling(value *int32) {
    obj.Scaling = value
}

func (obj *Font) GetShadow() *bool {
    return obj.Shadow
}

func (obj *Font) SetShadow(value *bool) {
    obj.Shadow = value
}

func (obj *Font) GetSize() *float64 {
    return obj.Size
}

func (obj *Font) SetSize(value *float64) {
    obj.Size = value
}

func (obj *Font) GetSizeBi() *float64 {
    return obj.SizeBi
}

func (obj *Font) SetSizeBi(value *float64) {
    obj.SizeBi = value
}

func (obj *Font) GetSmallCaps() *bool {
    return obj.SmallCaps
}

func (obj *Font) SetSmallCaps(value *bool) {
    obj.SmallCaps = value
}

func (obj *Font) GetSpacing() *float64 {
    return obj.Spacing
}

func (obj *Font) SetSpacing(value *float64) {
    obj.Spacing = value
}

func (obj *Font) GetStrikeThrough() *bool {
    return obj.StrikeThrough
}

func (obj *Font) SetStrikeThrough(value *bool) {
    obj.StrikeThrough = value
}

func (obj *Font) GetStyleIdentifier() *string {
    return obj.StyleIdentifier
}

func (obj *Font) SetStyleIdentifier(value *string) {
    obj.StyleIdentifier = value
}

func (obj *Font) GetStyleName() *string {
    return obj.StyleName
}

func (obj *Font) SetStyleName(value *string) {
    obj.StyleName = value
}

func (obj *Font) GetSubscript() *bool {
    return obj.Subscript
}

func (obj *Font) SetSubscript(value *bool) {
    obj.Subscript = value
}

func (obj *Font) GetSuperscript() *bool {
    return obj.Superscript
}

func (obj *Font) SetSuperscript(value *bool) {
    obj.Superscript = value
}

func (obj *Font) GetTextEffect() *string {
    return obj.TextEffect
}

func (obj *Font) SetTextEffect(value *string) {
    obj.TextEffect = value
}

func (obj *Font) GetUnderline() *string {
    return obj.Underline
}

func (obj *Font) SetUnderline(value *string) {
    obj.Underline = value
}

func (obj *Font) GetUnderlineColor() IXmlColor {
    return obj.UnderlineColor
}

func (obj *Font) SetUnderlineColor(value IXmlColor) {
    obj.UnderlineColor = value
}

