/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="paragraph_format_update.go">
 *   Copyright (c) 2023 Aspose.Words for Cloud
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

// Paragraph format element update DTO.

type IParagraphFormatUpdate interface {
    IsParagraphFormatUpdate() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetAddSpaceBetweenFarEastAndAlpha() *bool
    SetAddSpaceBetweenFarEastAndAlpha(value *bool)
    GetAddSpaceBetweenFarEastAndDigit() *bool
    SetAddSpaceBetweenFarEastAndDigit(value *bool)
    GetAlignment() *string
    SetAlignment(value *string)
    GetBidi() *bool
    SetBidi(value *bool)
    GetDropCapPosition() *string
    SetDropCapPosition(value *string)
    GetFirstLineIndent() *float64
    SetFirstLineIndent(value *float64)
    GetKeepTogether() *bool
    SetKeepTogether(value *bool)
    GetKeepWithNext() *bool
    SetKeepWithNext(value *bool)
    GetLeftIndent() *float64
    SetLeftIndent(value *float64)
    GetLineSpacing() *float64
    SetLineSpacing(value *float64)
    GetLineSpacingRule() *string
    SetLineSpacingRule(value *string)
    GetLinesToDrop() *int32
    SetLinesToDrop(value *int32)
    GetNoSpaceBetweenParagraphsOfSameStyle() *bool
    SetNoSpaceBetweenParagraphsOfSameStyle(value *bool)
    GetOutlineLevel() *string
    SetOutlineLevel(value *string)
    GetPageBreakBefore() *bool
    SetPageBreakBefore(value *bool)
    GetRightIndent() *float64
    SetRightIndent(value *float64)
    GetSpaceAfter() *float64
    SetSpaceAfter(value *float64)
    GetSpaceAfterAuto() *bool
    SetSpaceAfterAuto(value *bool)
    GetSpaceBefore() *float64
    SetSpaceBefore(value *float64)
    GetSpaceBeforeAuto() *bool
    SetSpaceBeforeAuto(value *bool)
    GetStyleIdentifier() *string
    SetStyleIdentifier(value *string)
    GetStyleName() *string
    SetStyleName(value *string)
    GetSuppressAutoHyphens() *bool
    SetSuppressAutoHyphens(value *bool)
    GetSuppressLineNumbers() *bool
    SetSuppressLineNumbers(value *bool)
    GetWidowControl() *bool
    SetWidowControl(value *bool)
    GetShading() IShading
    SetShading(value IShading)
}

type ParagraphFormatUpdate struct {
    // Paragraph format element update DTO.
    Link IWordsApiLink `json:"Link,omitempty"`

    // Paragraph format element update DTO.
    AddSpaceBetweenFarEastAndAlpha *bool `json:"AddSpaceBetweenFarEastAndAlpha,omitempty"`

    // Paragraph format element update DTO.
    AddSpaceBetweenFarEastAndDigit *bool `json:"AddSpaceBetweenFarEastAndDigit,omitempty"`

    // Paragraph format element update DTO.
    Alignment *string `json:"Alignment,omitempty"`

    // Paragraph format element update DTO.
    Bidi *bool `json:"Bidi,omitempty"`

    // Paragraph format element update DTO.
    DropCapPosition *string `json:"DropCapPosition,omitempty"`

    // Paragraph format element update DTO.
    FirstLineIndent *float64 `json:"FirstLineIndent,omitempty"`

    // Paragraph format element update DTO.
    KeepTogether *bool `json:"KeepTogether,omitempty"`

    // Paragraph format element update DTO.
    KeepWithNext *bool `json:"KeepWithNext,omitempty"`

    // Paragraph format element update DTO.
    LeftIndent *float64 `json:"LeftIndent,omitempty"`

    // Paragraph format element update DTO.
    LineSpacing *float64 `json:"LineSpacing,omitempty"`

    // Paragraph format element update DTO.
    LineSpacingRule *string `json:"LineSpacingRule,omitempty"`

    // Paragraph format element update DTO.
    LinesToDrop *int32 `json:"LinesToDrop,omitempty"`

    // Paragraph format element update DTO.
    NoSpaceBetweenParagraphsOfSameStyle *bool `json:"NoSpaceBetweenParagraphsOfSameStyle,omitempty"`

    // Paragraph format element update DTO.
    OutlineLevel *string `json:"OutlineLevel,omitempty"`

    // Paragraph format element update DTO.
    PageBreakBefore *bool `json:"PageBreakBefore,omitempty"`

    // Paragraph format element update DTO.
    RightIndent *float64 `json:"RightIndent,omitempty"`

    // Paragraph format element update DTO.
    SpaceAfter *float64 `json:"SpaceAfter,omitempty"`

    // Paragraph format element update DTO.
    SpaceAfterAuto *bool `json:"SpaceAfterAuto,omitempty"`

    // Paragraph format element update DTO.
    SpaceBefore *float64 `json:"SpaceBefore,omitempty"`

    // Paragraph format element update DTO.
    SpaceBeforeAuto *bool `json:"SpaceBeforeAuto,omitempty"`

    // Paragraph format element update DTO.
    StyleIdentifier *string `json:"StyleIdentifier,omitempty"`

    // Paragraph format element update DTO.
    StyleName *string `json:"StyleName,omitempty"`

    // Paragraph format element update DTO.
    SuppressAutoHyphens *bool `json:"SuppressAutoHyphens,omitempty"`

    // Paragraph format element update DTO.
    SuppressLineNumbers *bool `json:"SuppressLineNumbers,omitempty"`

    // Paragraph format element update DTO.
    WidowControl *bool `json:"WidowControl,omitempty"`

    // Paragraph format element update DTO.
    Shading IShading `json:"Shading,omitempty"`
}

func (ParagraphFormatUpdate) IsParagraphFormatUpdate() bool {
    return true
}

func (ParagraphFormatUpdate) IsParagraphFormatBase() bool {
    return true
}

func (ParagraphFormatUpdate) IsLinkElement() bool {
    return true
}

func (obj *ParagraphFormatUpdate) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.Shading != nil) {
        obj.Shading.Initialize()
    }


}

func (obj *ParagraphFormatUpdate) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["AddSpaceBetweenFarEastAndAlpha"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.AddSpaceBetweenFarEastAndAlpha = &parsedValue
        }

    } else if jsonValue, exists := json["addSpaceBetweenFarEastAndAlpha"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.AddSpaceBetweenFarEastAndAlpha = &parsedValue
        }

    }

    if jsonValue, exists := json["AddSpaceBetweenFarEastAndDigit"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.AddSpaceBetweenFarEastAndDigit = &parsedValue
        }

    } else if jsonValue, exists := json["addSpaceBetweenFarEastAndDigit"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.AddSpaceBetweenFarEastAndDigit = &parsedValue
        }

    }

    if jsonValue, exists := json["Alignment"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Alignment = &parsedValue
        }

    } else if jsonValue, exists := json["alignment"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Alignment = &parsedValue
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

    if jsonValue, exists := json["DropCapPosition"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DropCapPosition = &parsedValue
        }

    } else if jsonValue, exists := json["dropCapPosition"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DropCapPosition = &parsedValue
        }

    }

    if jsonValue, exists := json["FirstLineIndent"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.FirstLineIndent = &parsedValue
        }

    } else if jsonValue, exists := json["firstLineIndent"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.FirstLineIndent = &parsedValue
        }

    }

    if jsonValue, exists := json["KeepTogether"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.KeepTogether = &parsedValue
        }

    } else if jsonValue, exists := json["keepTogether"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.KeepTogether = &parsedValue
        }

    }

    if jsonValue, exists := json["KeepWithNext"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.KeepWithNext = &parsedValue
        }

    } else if jsonValue, exists := json["keepWithNext"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.KeepWithNext = &parsedValue
        }

    }

    if jsonValue, exists := json["LeftIndent"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.LeftIndent = &parsedValue
        }

    } else if jsonValue, exists := json["leftIndent"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.LeftIndent = &parsedValue
        }

    }

    if jsonValue, exists := json["LineSpacing"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.LineSpacing = &parsedValue
        }

    } else if jsonValue, exists := json["lineSpacing"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.LineSpacing = &parsedValue
        }

    }

    if jsonValue, exists := json["LineSpacingRule"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.LineSpacingRule = &parsedValue
        }

    } else if jsonValue, exists := json["lineSpacingRule"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.LineSpacingRule = &parsedValue
        }

    }

    if jsonValue, exists := json["LinesToDrop"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.LinesToDrop = new(int32)
            *obj.LinesToDrop = int32(parsedValue)
        }

    } else if jsonValue, exists := json["linesToDrop"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.LinesToDrop = new(int32)
            *obj.LinesToDrop = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["NoSpaceBetweenParagraphsOfSameStyle"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.NoSpaceBetweenParagraphsOfSameStyle = &parsedValue
        }

    } else if jsonValue, exists := json["noSpaceBetweenParagraphsOfSameStyle"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.NoSpaceBetweenParagraphsOfSameStyle = &parsedValue
        }

    }

    if jsonValue, exists := json["OutlineLevel"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.OutlineLevel = &parsedValue
        }

    } else if jsonValue, exists := json["outlineLevel"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.OutlineLevel = &parsedValue
        }

    }

    if jsonValue, exists := json["PageBreakBefore"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.PageBreakBefore = &parsedValue
        }

    } else if jsonValue, exists := json["pageBreakBefore"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.PageBreakBefore = &parsedValue
        }

    }

    if jsonValue, exists := json["RightIndent"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.RightIndent = &parsedValue
        }

    } else if jsonValue, exists := json["rightIndent"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.RightIndent = &parsedValue
        }

    }

    if jsonValue, exists := json["SpaceAfter"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.SpaceAfter = &parsedValue
        }

    } else if jsonValue, exists := json["spaceAfter"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.SpaceAfter = &parsedValue
        }

    }

    if jsonValue, exists := json["SpaceAfterAuto"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.SpaceAfterAuto = &parsedValue
        }

    } else if jsonValue, exists := json["spaceAfterAuto"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.SpaceAfterAuto = &parsedValue
        }

    }

    if jsonValue, exists := json["SpaceBefore"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.SpaceBefore = &parsedValue
        }

    } else if jsonValue, exists := json["spaceBefore"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.SpaceBefore = &parsedValue
        }

    }

    if jsonValue, exists := json["SpaceBeforeAuto"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.SpaceBeforeAuto = &parsedValue
        }

    } else if jsonValue, exists := json["spaceBeforeAuto"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.SpaceBeforeAuto = &parsedValue
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

    if jsonValue, exists := json["SuppressAutoHyphens"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.SuppressAutoHyphens = &parsedValue
        }

    } else if jsonValue, exists := json["suppressAutoHyphens"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.SuppressAutoHyphens = &parsedValue
        }

    }

    if jsonValue, exists := json["SuppressLineNumbers"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.SuppressLineNumbers = &parsedValue
        }

    } else if jsonValue, exists := json["suppressLineNumbers"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.SuppressLineNumbers = &parsedValue
        }

    }

    if jsonValue, exists := json["WidowControl"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.WidowControl = &parsedValue
        }

    } else if jsonValue, exists := json["widowControl"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.WidowControl = &parsedValue
        }

    }

    if jsonValue, exists := json["Shading"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IShading = new(Shading)
            modelInstance.Deserialize(parsedValue)
            obj.Shading = modelInstance
        }

    } else if jsonValue, exists := json["shading"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IShading = new(Shading)
            modelInstance.Deserialize(parsedValue)
            obj.Shading = modelInstance
        }

    }
}

func (obj *ParagraphFormatUpdate) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *ParagraphFormatUpdate) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Link != nil {
        if err := obj.Link.Validate(); err != nil {
            return err
        }
    }
    if obj.Shading != nil {
        if err := obj.Shading.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *ParagraphFormatUpdate) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *ParagraphFormatUpdate) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *ParagraphFormatUpdate) GetAddSpaceBetweenFarEastAndAlpha() *bool {
    return obj.AddSpaceBetweenFarEastAndAlpha
}

func (obj *ParagraphFormatUpdate) SetAddSpaceBetweenFarEastAndAlpha(value *bool) {
    obj.AddSpaceBetweenFarEastAndAlpha = value
}

func (obj *ParagraphFormatUpdate) GetAddSpaceBetweenFarEastAndDigit() *bool {
    return obj.AddSpaceBetweenFarEastAndDigit
}

func (obj *ParagraphFormatUpdate) SetAddSpaceBetweenFarEastAndDigit(value *bool) {
    obj.AddSpaceBetweenFarEastAndDigit = value
}

func (obj *ParagraphFormatUpdate) GetAlignment() *string {
    return obj.Alignment
}

func (obj *ParagraphFormatUpdate) SetAlignment(value *string) {
    obj.Alignment = value
}

func (obj *ParagraphFormatUpdate) GetBidi() *bool {
    return obj.Bidi
}

func (obj *ParagraphFormatUpdate) SetBidi(value *bool) {
    obj.Bidi = value
}

func (obj *ParagraphFormatUpdate) GetDropCapPosition() *string {
    return obj.DropCapPosition
}

func (obj *ParagraphFormatUpdate) SetDropCapPosition(value *string) {
    obj.DropCapPosition = value
}

func (obj *ParagraphFormatUpdate) GetFirstLineIndent() *float64 {
    return obj.FirstLineIndent
}

func (obj *ParagraphFormatUpdate) SetFirstLineIndent(value *float64) {
    obj.FirstLineIndent = value
}

func (obj *ParagraphFormatUpdate) GetKeepTogether() *bool {
    return obj.KeepTogether
}

func (obj *ParagraphFormatUpdate) SetKeepTogether(value *bool) {
    obj.KeepTogether = value
}

func (obj *ParagraphFormatUpdate) GetKeepWithNext() *bool {
    return obj.KeepWithNext
}

func (obj *ParagraphFormatUpdate) SetKeepWithNext(value *bool) {
    obj.KeepWithNext = value
}

func (obj *ParagraphFormatUpdate) GetLeftIndent() *float64 {
    return obj.LeftIndent
}

func (obj *ParagraphFormatUpdate) SetLeftIndent(value *float64) {
    obj.LeftIndent = value
}

func (obj *ParagraphFormatUpdate) GetLineSpacing() *float64 {
    return obj.LineSpacing
}

func (obj *ParagraphFormatUpdate) SetLineSpacing(value *float64) {
    obj.LineSpacing = value
}

func (obj *ParagraphFormatUpdate) GetLineSpacingRule() *string {
    return obj.LineSpacingRule
}

func (obj *ParagraphFormatUpdate) SetLineSpacingRule(value *string) {
    obj.LineSpacingRule = value
}

func (obj *ParagraphFormatUpdate) GetLinesToDrop() *int32 {
    return obj.LinesToDrop
}

func (obj *ParagraphFormatUpdate) SetLinesToDrop(value *int32) {
    obj.LinesToDrop = value
}

func (obj *ParagraphFormatUpdate) GetNoSpaceBetweenParagraphsOfSameStyle() *bool {
    return obj.NoSpaceBetweenParagraphsOfSameStyle
}

func (obj *ParagraphFormatUpdate) SetNoSpaceBetweenParagraphsOfSameStyle(value *bool) {
    obj.NoSpaceBetweenParagraphsOfSameStyle = value
}

func (obj *ParagraphFormatUpdate) GetOutlineLevel() *string {
    return obj.OutlineLevel
}

func (obj *ParagraphFormatUpdate) SetOutlineLevel(value *string) {
    obj.OutlineLevel = value
}

func (obj *ParagraphFormatUpdate) GetPageBreakBefore() *bool {
    return obj.PageBreakBefore
}

func (obj *ParagraphFormatUpdate) SetPageBreakBefore(value *bool) {
    obj.PageBreakBefore = value
}

func (obj *ParagraphFormatUpdate) GetRightIndent() *float64 {
    return obj.RightIndent
}

func (obj *ParagraphFormatUpdate) SetRightIndent(value *float64) {
    obj.RightIndent = value
}

func (obj *ParagraphFormatUpdate) GetSpaceAfter() *float64 {
    return obj.SpaceAfter
}

func (obj *ParagraphFormatUpdate) SetSpaceAfter(value *float64) {
    obj.SpaceAfter = value
}

func (obj *ParagraphFormatUpdate) GetSpaceAfterAuto() *bool {
    return obj.SpaceAfterAuto
}

func (obj *ParagraphFormatUpdate) SetSpaceAfterAuto(value *bool) {
    obj.SpaceAfterAuto = value
}

func (obj *ParagraphFormatUpdate) GetSpaceBefore() *float64 {
    return obj.SpaceBefore
}

func (obj *ParagraphFormatUpdate) SetSpaceBefore(value *float64) {
    obj.SpaceBefore = value
}

func (obj *ParagraphFormatUpdate) GetSpaceBeforeAuto() *bool {
    return obj.SpaceBeforeAuto
}

func (obj *ParagraphFormatUpdate) SetSpaceBeforeAuto(value *bool) {
    obj.SpaceBeforeAuto = value
}

func (obj *ParagraphFormatUpdate) GetStyleIdentifier() *string {
    return obj.StyleIdentifier
}

func (obj *ParagraphFormatUpdate) SetStyleIdentifier(value *string) {
    obj.StyleIdentifier = value
}

func (obj *ParagraphFormatUpdate) GetStyleName() *string {
    return obj.StyleName
}

func (obj *ParagraphFormatUpdate) SetStyleName(value *string) {
    obj.StyleName = value
}

func (obj *ParagraphFormatUpdate) GetSuppressAutoHyphens() *bool {
    return obj.SuppressAutoHyphens
}

func (obj *ParagraphFormatUpdate) SetSuppressAutoHyphens(value *bool) {
    obj.SuppressAutoHyphens = value
}

func (obj *ParagraphFormatUpdate) GetSuppressLineNumbers() *bool {
    return obj.SuppressLineNumbers
}

func (obj *ParagraphFormatUpdate) SetSuppressLineNumbers(value *bool) {
    obj.SuppressLineNumbers = value
}

func (obj *ParagraphFormatUpdate) GetWidowControl() *bool {
    return obj.WidowControl
}

func (obj *ParagraphFormatUpdate) SetWidowControl(value *bool) {
    obj.WidowControl = value
}

func (obj *ParagraphFormatUpdate) GetShading() IShading {
    return obj.Shading
}

func (obj *ParagraphFormatUpdate) SetShading(value IShading) {
    obj.Shading = value
}

