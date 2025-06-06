/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="comment_update.go">
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

// Comment update.
type CommentUpdateResult struct {
    // Comment update.
    RangeStart NewDocumentPositionResult `json:"RangeStart,omitempty"`

    // Comment update.
    RangeEnd NewDocumentPositionResult `json:"RangeEnd,omitempty"`

    // Comment update.
    Author string `json:"Author,omitempty"`

    // Comment update.
    DateTime Time `json:"DateTime,omitempty"`

    // Comment update.
    Initial string `json:"Initial,omitempty"`

    // Comment update.
    Text string `json:"Text,omitempty"`
}

type CommentUpdate struct {
    // Comment update.
    RangeStart INewDocumentPosition `json:"RangeStart,omitempty"`

    // Comment update.
    RangeEnd INewDocumentPosition `json:"RangeEnd,omitempty"`

    // Comment update.
    Author *string `json:"Author,omitempty"`

    // Comment update.
    DateTime *Time `json:"DateTime,omitempty"`

    // Comment update.
    Initial *string `json:"Initial,omitempty"`

    // Comment update.
    Text *string `json:"Text,omitempty"`
}

type ICommentUpdate interface {
    IsCommentUpdate() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
}

func (CommentUpdate) IsCommentUpdate() bool {
    return true
}

func (CommentUpdate) IsCommentBase() bool {
    return true
}

func (obj *CommentUpdate) Initialize() {
    if (obj.RangeStart != nil) {
        obj.RangeStart.Initialize()
    }

    if (obj.RangeEnd != nil) {
        obj.RangeEnd.Initialize()
    }


}

func (obj *CommentUpdate) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}


