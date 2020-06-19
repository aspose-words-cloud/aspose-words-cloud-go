/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="paragraph_format.go">
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

// Paragraph format element.
type ParagraphFormat struct {
    // Paragraph format element.
    Link *WordsApiLink `json:"Link,omitempty"`

    // Paragraph format element.
    AddSpaceBetweenFarEastAndAlpha bool `json:"AddSpaceBetweenFarEastAndAlpha,omitempty"`

    // Paragraph format element.
    AddSpaceBetweenFarEastAndDigit bool `json:"AddSpaceBetweenFarEastAndDigit,omitempty"`

    // Paragraph format element.
    Alignment string `json:"Alignment,omitempty"`

    // Paragraph format element.
    Bidi bool `json:"Bidi,omitempty"`

    // Paragraph format element.
    DropCapPosition string `json:"DropCapPosition,omitempty"`

    // Paragraph format element.
    FirstLineIndent float64 `json:"FirstLineIndent,omitempty"`

    // Paragraph format element.
    IsListItem bool `json:"IsListItem,omitempty"`

    // Paragraph format element.
    KeepTogether bool `json:"KeepTogether,omitempty"`

    // Paragraph format element.
    KeepWithNext bool `json:"KeepWithNext,omitempty"`

    // Paragraph format element.
    LeftIndent float64 `json:"LeftIndent,omitempty"`

    // Paragraph format element.
    LineSpacing float64 `json:"LineSpacing,omitempty"`

    // Paragraph format element.
    LineSpacingRule string `json:"LineSpacingRule,omitempty"`

    // Paragraph format element.
    LinesToDrop int32 `json:"LinesToDrop,omitempty"`

    // Paragraph format element.
    NoSpaceBetweenParagraphsOfSameStyle bool `json:"NoSpaceBetweenParagraphsOfSameStyle,omitempty"`

    // Paragraph format element.
    OutlineLevel string `json:"OutlineLevel,omitempty"`

    // Paragraph format element.
    PageBreakBefore bool `json:"PageBreakBefore,omitempty"`

    // Paragraph format element.
    RightIndent float64 `json:"RightIndent,omitempty"`

    // Paragraph format element.
    Shading *Shading `json:"Shading,omitempty"`

    // Paragraph format element.
    SpaceAfter float64 `json:"SpaceAfter,omitempty"`

    // Paragraph format element.
    SpaceAfterAuto bool `json:"SpaceAfterAuto,omitempty"`

    // Paragraph format element.
    SpaceBefore float64 `json:"SpaceBefore,omitempty"`

    // Paragraph format element.
    SpaceBeforeAuto bool `json:"SpaceBeforeAuto,omitempty"`

    // Paragraph format element.
    StyleIdentifier string `json:"StyleIdentifier,omitempty"`

    // Paragraph format element.
    StyleName string `json:"StyleName,omitempty"`

    // Paragraph format element.
    SuppressAutoHyphens bool `json:"SuppressAutoHyphens,omitempty"`

    // Paragraph format element.
    SuppressLineNumbers bool `json:"SuppressLineNumbers,omitempty"`

    // Paragraph format element.
    WidowControl bool `json:"WidowControl,omitempty"`
}

type IParagraphFormat interface {
    IsParagraphFormat() bool
}
func (ParagraphFormat) IsParagraphFormat() bool {
    return true
}

func (ParagraphFormat) IsLinkElement() bool {
    return true
}
