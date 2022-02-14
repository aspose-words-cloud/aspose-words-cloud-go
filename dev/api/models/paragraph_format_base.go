/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="paragraph_format_base.go">
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

// Paragraph format element base class.
type ParagraphFormatBaseResult struct {
    // Paragraph format element base class.
    Link WordsApiLinkResult `json:"Link,omitempty"`

    // Paragraph format element base class.
    AddSpaceBetweenFarEastAndAlpha bool `json:"AddSpaceBetweenFarEastAndAlpha,omitempty"`

    // Paragraph format element base class.
    AddSpaceBetweenFarEastAndDigit bool `json:"AddSpaceBetweenFarEastAndDigit,omitempty"`

    // Paragraph format element base class.
    Alignment string `json:"Alignment,omitempty"`

    // Paragraph format element base class.
    Bidi bool `json:"Bidi,omitempty"`

    // Paragraph format element base class.
    DropCapPosition string `json:"DropCapPosition,omitempty"`

    // Paragraph format element base class.
    FirstLineIndent float64 `json:"FirstLineIndent,omitempty"`

    // Paragraph format element base class.
    KeepTogether bool `json:"KeepTogether,omitempty"`

    // Paragraph format element base class.
    KeepWithNext bool `json:"KeepWithNext,omitempty"`

    // Paragraph format element base class.
    LeftIndent float64 `json:"LeftIndent,omitempty"`

    // Paragraph format element base class.
    LineSpacing float64 `json:"LineSpacing,omitempty"`

    // Paragraph format element base class.
    LineSpacingRule string `json:"LineSpacingRule,omitempty"`

    // Paragraph format element base class.
    LinesToDrop int32 `json:"LinesToDrop,omitempty"`

    // Paragraph format element base class.
    NoSpaceBetweenParagraphsOfSameStyle bool `json:"NoSpaceBetweenParagraphsOfSameStyle,omitempty"`

    // Paragraph format element base class.
    OutlineLevel string `json:"OutlineLevel,omitempty"`

    // Paragraph format element base class.
    PageBreakBefore bool `json:"PageBreakBefore,omitempty"`

    // Paragraph format element base class.
    RightIndent float64 `json:"RightIndent,omitempty"`

    // Paragraph format element base class.
    Shading ShadingResult `json:"Shading,omitempty"`

    // Paragraph format element base class.
    SpaceAfter float64 `json:"SpaceAfter,omitempty"`

    // Paragraph format element base class.
    SpaceAfterAuto bool `json:"SpaceAfterAuto,omitempty"`

    // Paragraph format element base class.
    SpaceBefore float64 `json:"SpaceBefore,omitempty"`

    // Paragraph format element base class.
    SpaceBeforeAuto bool `json:"SpaceBeforeAuto,omitempty"`

    // Paragraph format element base class.
    StyleIdentifier string `json:"StyleIdentifier,omitempty"`

    // Paragraph format element base class.
    StyleName string `json:"StyleName,omitempty"`

    // Paragraph format element base class.
    SuppressAutoHyphens bool `json:"SuppressAutoHyphens,omitempty"`

    // Paragraph format element base class.
    SuppressLineNumbers bool `json:"SuppressLineNumbers,omitempty"`

    // Paragraph format element base class.
    WidowControl bool `json:"WidowControl,omitempty"`
}

type ParagraphFormatBase struct {
    // Paragraph format element base class.
    Link IWordsApiLink `json:"Link,omitempty"`

    // Paragraph format element base class.
    AddSpaceBetweenFarEastAndAlpha *bool `json:"AddSpaceBetweenFarEastAndAlpha,omitempty"`

    // Paragraph format element base class.
    AddSpaceBetweenFarEastAndDigit *bool `json:"AddSpaceBetweenFarEastAndDigit,omitempty"`

    // Paragraph format element base class.
    Alignment *string `json:"Alignment,omitempty"`

    // Paragraph format element base class.
    Bidi *bool `json:"Bidi,omitempty"`

    // Paragraph format element base class.
    DropCapPosition *string `json:"DropCapPosition,omitempty"`

    // Paragraph format element base class.
    FirstLineIndent *float64 `json:"FirstLineIndent,omitempty"`

    // Paragraph format element base class.
    KeepTogether *bool `json:"KeepTogether,omitempty"`

    // Paragraph format element base class.
    KeepWithNext *bool `json:"KeepWithNext,omitempty"`

    // Paragraph format element base class.
    LeftIndent *float64 `json:"LeftIndent,omitempty"`

    // Paragraph format element base class.
    LineSpacing *float64 `json:"LineSpacing,omitempty"`

    // Paragraph format element base class.
    LineSpacingRule *string `json:"LineSpacingRule,omitempty"`

    // Paragraph format element base class.
    LinesToDrop *int32 `json:"LinesToDrop,omitempty"`

    // Paragraph format element base class.
    NoSpaceBetweenParagraphsOfSameStyle *bool `json:"NoSpaceBetweenParagraphsOfSameStyle,omitempty"`

    // Paragraph format element base class.
    OutlineLevel *string `json:"OutlineLevel,omitempty"`

    // Paragraph format element base class.
    PageBreakBefore *bool `json:"PageBreakBefore,omitempty"`

    // Paragraph format element base class.
    RightIndent *float64 `json:"RightIndent,omitempty"`

    // Paragraph format element base class.
    Shading IShading `json:"Shading,omitempty"`

    // Paragraph format element base class.
    SpaceAfter *float64 `json:"SpaceAfter,omitempty"`

    // Paragraph format element base class.
    SpaceAfterAuto *bool `json:"SpaceAfterAuto,omitempty"`

    // Paragraph format element base class.
    SpaceBefore *float64 `json:"SpaceBefore,omitempty"`

    // Paragraph format element base class.
    SpaceBeforeAuto *bool `json:"SpaceBeforeAuto,omitempty"`

    // Paragraph format element base class.
    StyleIdentifier *string `json:"StyleIdentifier,omitempty"`

    // Paragraph format element base class.
    StyleName *string `json:"StyleName,omitempty"`

    // Paragraph format element base class.
    SuppressAutoHyphens *bool `json:"SuppressAutoHyphens,omitempty"`

    // Paragraph format element base class.
    SuppressLineNumbers *bool `json:"SuppressLineNumbers,omitempty"`

    // Paragraph format element base class.
    WidowControl *bool `json:"WidowControl,omitempty"`
}

type IParagraphFormatBase interface {
    IsParagraphFormatBase() bool
    Initialize()
}

func (ParagraphFormatBase) IsParagraphFormatBase() bool {
    return true
}

func (ParagraphFormatBase) IsLinkElement() bool {
    return true
}

func (obj *ParagraphFormatBase) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }



































    if (obj.Shading != nil) {
        obj.Shading.Initialize()
    }



















}


