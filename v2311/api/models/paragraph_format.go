/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="paragraph_format.go">
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

// Paragraph format element.

type IParagraphFormat interface {
    IsParagraphFormat() bool
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
    GetIsListItem() *bool
    SetIsListItem(value *bool)
    GetIsHeading() *bool
    SetIsHeading(value *bool)
}

type ParagraphFormat struct {
    // Paragraph format element.
    Link IWordsApiLink `json:"Link,omitempty"`

    // Paragraph format element.
    AddSpaceBetweenFarEastAndAlpha *bool `json:"AddSpaceBetweenFarEastAndAlpha,omitempty"`

    // Paragraph format element.
    AddSpaceBetweenFarEastAndDigit *bool `json:"AddSpaceBetweenFarEastAndDigit,omitempty"`

    // Paragraph format element.
    Alignment *string `json:"Alignment,omitempty"`

    // Paragraph format element.
    Bidi *bool `json:"Bidi,omitempty"`

    // Paragraph format element.
    DropCapPosition *string `json:"DropCapPosition,omitempty"`

    // Paragraph format element.
    FirstLineIndent *float64 `json:"FirstLineIndent,omitempty"`

    // Paragraph format element.
    KeepTogether *bool `json:"KeepTogether,omitempty"`

    // Paragraph format element.
    KeepWithNext *bool `json:"KeepWithNext,omitempty"`

    // Paragraph format element.
    LeftIndent *float64 `json:"LeftIndent,omitempty"`

    // Paragraph format element.
    LineSpacing *float64 `json:"LineSpacing,omitempty"`

    // Paragraph format element.
    LineSpacingRule *string `json:"LineSpacingRule,omitempty"`

    // Paragraph format element.
    LinesToDrop *int32 `json:"LinesToDrop,omitempty"`

    // Paragraph format element.
    NoSpaceBetweenParagraphsOfSameStyle *bool `json:"NoSpaceBetweenParagraphsOfSameStyle,omitempty"`

    // Paragraph format element.
    OutlineLevel *string `json:"OutlineLevel,omitempty"`

    // Paragraph format element.
    PageBreakBefore *bool `json:"PageBreakBefore,omitempty"`

    // Paragraph format element.
    RightIndent *float64 `json:"RightIndent,omitempty"`

    // Paragraph format element.
    SpaceAfter *float64 `json:"SpaceAfter,omitempty"`

    // Paragraph format element.
    SpaceAfterAuto *bool `json:"SpaceAfterAuto,omitempty"`

    // Paragraph format element.
    SpaceBefore *float64 `json:"SpaceBefore,omitempty"`

    // Paragraph format element.
    SpaceBeforeAuto *bool `json:"SpaceBeforeAuto,omitempty"`

    // Paragraph format element.
    StyleIdentifier *string `json:"StyleIdentifier,omitempty"`

    // Paragraph format element.
    StyleName *string `json:"StyleName,omitempty"`

    // Paragraph format element.
    SuppressAutoHyphens *bool `json:"SuppressAutoHyphens,omitempty"`

    // Paragraph format element.
    SuppressLineNumbers *bool `json:"SuppressLineNumbers,omitempty"`

    // Paragraph format element.
    WidowControl *bool `json:"WidowControl,omitempty"`

    // Paragraph format element.
    Shading IShading `json:"Shading,omitempty"`

    // Paragraph format element.
    IsListItem *bool `json:"IsListItem,omitempty"`

    // Paragraph format element.
    IsHeading *bool `json:"IsHeading,omitempty"`
}

func (ParagraphFormat) IsParagraphFormat() bool {
    return true
}

func (ParagraphFormat) IsParagraphFormatBase() bool {
    return true
}

func (ParagraphFormat) IsLinkElement() bool {
    return true
}

func (obj *ParagraphFormat) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.Shading != nil) {
        obj.Shading.Initialize()
    }


}

func (obj *ParagraphFormat) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["IsListItem"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsListItem = &parsedValue
        }

    } else if jsonValue, exists := json["isListItem"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsListItem = &parsedValue
        }

    }

    if jsonValue, exists := json["IsHeading"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsHeading = &parsedValue
        }

    } else if jsonValue, exists := json["isHeading"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsHeading = &parsedValue
        }

    }
}

func (obj *ParagraphFormat) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *ParagraphFormat) Validate() error {
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

func (obj *ParagraphFormat) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *ParagraphFormat) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *ParagraphFormat) GetAddSpaceBetweenFarEastAndAlpha() *bool {
    return obj.AddSpaceBetweenFarEastAndAlpha
}

func (obj *ParagraphFormat) SetAddSpaceBetweenFarEastAndAlpha(value *bool) {
    obj.AddSpaceBetweenFarEastAndAlpha = value
}

func (obj *ParagraphFormat) GetAddSpaceBetweenFarEastAndDigit() *bool {
    return obj.AddSpaceBetweenFarEastAndDigit
}

func (obj *ParagraphFormat) SetAddSpaceBetweenFarEastAndDigit(value *bool) {
    obj.AddSpaceBetweenFarEastAndDigit = value
}

func (obj *ParagraphFormat) GetAlignment() *string {
    return obj.Alignment
}

func (obj *ParagraphFormat) SetAlignment(value *string) {
    obj.Alignment = value
}

func (obj *ParagraphFormat) GetBidi() *bool {
    return obj.Bidi
}

func (obj *ParagraphFormat) SetBidi(value *bool) {
    obj.Bidi = value
}

func (obj *ParagraphFormat) GetDropCapPosition() *string {
    return obj.DropCapPosition
}

func (obj *ParagraphFormat) SetDropCapPosition(value *string) {
    obj.DropCapPosition = value
}

func (obj *ParagraphFormat) GetFirstLineIndent() *float64 {
    return obj.FirstLineIndent
}

func (obj *ParagraphFormat) SetFirstLineIndent(value *float64) {
    obj.FirstLineIndent = value
}

func (obj *ParagraphFormat) GetKeepTogether() *bool {
    return obj.KeepTogether
}

func (obj *ParagraphFormat) SetKeepTogether(value *bool) {
    obj.KeepTogether = value
}

func (obj *ParagraphFormat) GetKeepWithNext() *bool {
    return obj.KeepWithNext
}

func (obj *ParagraphFormat) SetKeepWithNext(value *bool) {
    obj.KeepWithNext = value
}

func (obj *ParagraphFormat) GetLeftIndent() *float64 {
    return obj.LeftIndent
}

func (obj *ParagraphFormat) SetLeftIndent(value *float64) {
    obj.LeftIndent = value
}

func (obj *ParagraphFormat) GetLineSpacing() *float64 {
    return obj.LineSpacing
}

func (obj *ParagraphFormat) SetLineSpacing(value *float64) {
    obj.LineSpacing = value
}

func (obj *ParagraphFormat) GetLineSpacingRule() *string {
    return obj.LineSpacingRule
}

func (obj *ParagraphFormat) SetLineSpacingRule(value *string) {
    obj.LineSpacingRule = value
}

func (obj *ParagraphFormat) GetLinesToDrop() *int32 {
    return obj.LinesToDrop
}

func (obj *ParagraphFormat) SetLinesToDrop(value *int32) {
    obj.LinesToDrop = value
}

func (obj *ParagraphFormat) GetNoSpaceBetweenParagraphsOfSameStyle() *bool {
    return obj.NoSpaceBetweenParagraphsOfSameStyle
}

func (obj *ParagraphFormat) SetNoSpaceBetweenParagraphsOfSameStyle(value *bool) {
    obj.NoSpaceBetweenParagraphsOfSameStyle = value
}

func (obj *ParagraphFormat) GetOutlineLevel() *string {
    return obj.OutlineLevel
}

func (obj *ParagraphFormat) SetOutlineLevel(value *string) {
    obj.OutlineLevel = value
}

func (obj *ParagraphFormat) GetPageBreakBefore() *bool {
    return obj.PageBreakBefore
}

func (obj *ParagraphFormat) SetPageBreakBefore(value *bool) {
    obj.PageBreakBefore = value
}

func (obj *ParagraphFormat) GetRightIndent() *float64 {
    return obj.RightIndent
}

func (obj *ParagraphFormat) SetRightIndent(value *float64) {
    obj.RightIndent = value
}

func (obj *ParagraphFormat) GetSpaceAfter() *float64 {
    return obj.SpaceAfter
}

func (obj *ParagraphFormat) SetSpaceAfter(value *float64) {
    obj.SpaceAfter = value
}

func (obj *ParagraphFormat) GetSpaceAfterAuto() *bool {
    return obj.SpaceAfterAuto
}

func (obj *ParagraphFormat) SetSpaceAfterAuto(value *bool) {
    obj.SpaceAfterAuto = value
}

func (obj *ParagraphFormat) GetSpaceBefore() *float64 {
    return obj.SpaceBefore
}

func (obj *ParagraphFormat) SetSpaceBefore(value *float64) {
    obj.SpaceBefore = value
}

func (obj *ParagraphFormat) GetSpaceBeforeAuto() *bool {
    return obj.SpaceBeforeAuto
}

func (obj *ParagraphFormat) SetSpaceBeforeAuto(value *bool) {
    obj.SpaceBeforeAuto = value
}

func (obj *ParagraphFormat) GetStyleIdentifier() *string {
    return obj.StyleIdentifier
}

func (obj *ParagraphFormat) SetStyleIdentifier(value *string) {
    obj.StyleIdentifier = value
}

func (obj *ParagraphFormat) GetStyleName() *string {
    return obj.StyleName
}

func (obj *ParagraphFormat) SetStyleName(value *string) {
    obj.StyleName = value
}

func (obj *ParagraphFormat) GetSuppressAutoHyphens() *bool {
    return obj.SuppressAutoHyphens
}

func (obj *ParagraphFormat) SetSuppressAutoHyphens(value *bool) {
    obj.SuppressAutoHyphens = value
}

func (obj *ParagraphFormat) GetSuppressLineNumbers() *bool {
    return obj.SuppressLineNumbers
}

func (obj *ParagraphFormat) SetSuppressLineNumbers(value *bool) {
    obj.SuppressLineNumbers = value
}

func (obj *ParagraphFormat) GetWidowControl() *bool {
    return obj.WidowControl
}

func (obj *ParagraphFormat) SetWidowControl(value *bool) {
    obj.WidowControl = value
}

func (obj *ParagraphFormat) GetShading() IShading {
    return obj.Shading
}

func (obj *ParagraphFormat) SetShading(value IShading) {
    obj.Shading = value
}

func (obj *ParagraphFormat) GetIsListItem() *bool {
    return obj.IsListItem
}

func (obj *ParagraphFormat) SetIsListItem(value *bool) {
    obj.IsListItem = value
}

func (obj *ParagraphFormat) GetIsHeading() *bool {
    return obj.IsHeading
}

func (obj *ParagraphFormat) SetIsHeading(value *bool) {
    obj.IsHeading = value
}

