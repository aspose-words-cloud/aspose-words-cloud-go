/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="paragraph_format_base.go">
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

// Paragraph format element base class.

type IParagraphFormatBase interface {
    IsParagraphFormatBase() bool
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

