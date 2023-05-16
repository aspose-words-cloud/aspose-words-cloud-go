/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="footnote_base.go">
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

// Footnote base class.
type FootnoteBaseResult struct {
    // Footnote base class.
    FootnoteType string `json:"FootnoteType,omitempty"`

    // Footnote base class.
    Position NewDocumentPositionResult `json:"Position,omitempty"`

    // Footnote base class.
    ReferenceMark string `json:"ReferenceMark,omitempty"`

    // Footnote base class.
    Text string `json:"Text,omitempty"`
}

type FootnoteBase struct {
    // Footnote base class.
    FootnoteType *string `json:"FootnoteType,omitempty"`

    // Footnote base class.
    Position INewDocumentPosition `json:"Position,omitempty"`

    // Footnote base class.
    ReferenceMark *string `json:"ReferenceMark,omitempty"`

    // Footnote base class.
    Text *string `json:"Text,omitempty"`
}

type IFootnoteBase interface {
    IsFootnoteBase() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
}

func (FootnoteBase) IsFootnoteBase() bool {
    return true
}


func (obj *FootnoteBase) Initialize() {
    if (obj.Position != nil) {
        obj.Position.Initialize()
    }


}

func (obj *FootnoteBase) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}


