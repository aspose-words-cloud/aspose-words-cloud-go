/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="footnote_insert.go">
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

// Footnote for insert.
type FootnoteInsertResult struct {
    // Footnote for insert.
    FootnoteType string `json:"FootnoteType,omitempty"`

    // Footnote for insert.
    Position DocumentPositionResult `json:"Position,omitempty"`

    // Footnote for insert.
    ReferenceMark string `json:"ReferenceMark,omitempty"`

    // Footnote for insert.
    Text string `json:"Text,omitempty"`
}

type FootnoteInsert struct {
    // Footnote for insert.
    FootnoteType *string `json:"FootnoteType,omitempty"`

    // Footnote for insert.
    Position IDocumentPosition `json:"Position,omitempty"`

    // Footnote for insert.
    ReferenceMark *string `json:"ReferenceMark,omitempty"`

    // Footnote for insert.
    Text *string `json:"Text,omitempty"`
}

type IFootnoteInsert interface {
    IsFootnoteInsert() bool
    Initialize()
}

func (FootnoteInsert) IsFootnoteInsert() bool {
    return true
}

func (FootnoteInsert) IsFootnoteBase() bool {
    return true
}

func (obj *FootnoteInsert) Initialize() {
}


