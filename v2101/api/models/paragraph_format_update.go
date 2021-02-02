/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="paragraph_format_update.go">
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

// Paragraph format element update DTO.
type ParagraphFormatUpdateResult struct {
    // Paragraph format element update DTO.
    Link WordsApiLinkResult `json:"Link,omitempty"`

    // Paragraph format element update DTO.
    AddSpaceBetweenFarEastAndAlpha bool `json:"AddSpaceBetweenFarEastAndAlpha,omitempty"`

    // Paragraph format element update DTO.
    AddSpaceBetweenFarEastAndDigit bool `json:"AddSpaceBetweenFarEastAndDigit,omitempty"`

    // Paragraph format element update DTO.
    Alignment string `json:"Alignment,omitempty"`

    // Paragraph format element update DTO.
    Bidi bool `json:"Bidi,omitempty"`

    // Paragraph format element update DTO.
    DropCapPosition string `json:"DropCapPosition,omitempty"`

    // Paragraph format element update DTO.
    FirstLineIndent float64 `json:"FirstLineIndent,omitempty"`

    // Paragraph format element update DTO.
    KeepTogether bool `json:"KeepTogether,omitempty"`

    // Paragraph format element update DTO.
    KeepWithNext bool `json:"KeepWithNext,omitempty"`

    // Paragraph format element update DTO.
    LeftIndent float64 `json:"LeftIndent,omitempty"`

    // Paragraph format element update DTO.
    LineSpacing float64 `json:"LineSpacing,omitempty"`

    // Paragraph format element update DTO.
    LineSpacingRule string `json:"LineSpacingRule,omitempty"`

    // Paragraph format element update DTO.
    LinesToDrop int32 `json:"LinesToDrop,omitempty"`

    // Paragraph format element update DTO.
    NoSpaceBetweenParagraphsOfSameStyle bool `json:"NoSpaceBetweenParagraphsOfSameStyle,omitempty"`

    // Paragraph format element update DTO.
    OutlineLevel string `json:"OutlineLevel,omitempty"`

    // Paragraph format element update DTO.
    PageBreakBefore bool `json:"PageBreakBefore,omitempty"`

    // Paragraph format element update DTO.
    RightIndent float64 `json:"RightIndent,omitempty"`

    // Paragraph format element update DTO.
    Shading ShadingResult `json:"Shading,omitempty"`

    // Paragraph format element update DTO.
    SpaceAfter float64 `json:"SpaceAfter,omitempty"`

    // Paragraph format element update DTO.
    SpaceAfterAuto bool `json:"SpaceAfterAuto,omitempty"`

    // Paragraph format element update DTO.
    SpaceBefore float64 `json:"SpaceBefore,omitempty"`

    // Paragraph format element update DTO.
    SpaceBeforeAuto bool `json:"SpaceBeforeAuto,omitempty"`

    // Paragraph format element update DTO.
    StyleIdentifier string `json:"StyleIdentifier,omitempty"`

    // Paragraph format element update DTO.
    StyleName string `json:"StyleName,omitempty"`

    // Paragraph format element update DTO.
    SuppressAutoHyphens bool `json:"SuppressAutoHyphens,omitempty"`

    // Paragraph format element update DTO.
    SuppressLineNumbers bool `json:"SuppressLineNumbers,omitempty"`

    // Paragraph format element update DTO.
    WidowControl bool `json:"WidowControl,omitempty"`
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
    Shading IShading `json:"Shading,omitempty"`

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
}

type IParagraphFormatUpdate interface {
    IsParagraphFormatUpdate() bool
}
func (ParagraphFormatUpdate) IsParagraphFormatUpdate() bool {
    return true
}

func (ParagraphFormatUpdate) IsParagraphFormatBase() bool {
    return true
}


