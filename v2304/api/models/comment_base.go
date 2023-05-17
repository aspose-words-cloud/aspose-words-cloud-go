/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="comment_base.go">
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

// Comment.
type CommentBaseResult struct {
    // Comment.
    Author string `json:"Author,omitempty"`

    // Comment.
    DateTime Time `json:"DateTime,omitempty"`

    // Comment.
    Initial string `json:"Initial,omitempty"`

    // Comment.
    RangeEnd NewDocumentPositionResult `json:"RangeEnd,omitempty"`

    // Comment.
    RangeStart NewDocumentPositionResult `json:"RangeStart,omitempty"`

    // Comment.
    Text string `json:"Text,omitempty"`
}

type CommentBase struct {
    // Comment.
    Author *string `json:"Author,omitempty"`

    // Comment.
    DateTime *Time `json:"DateTime,omitempty"`

    // Comment.
    Initial *string `json:"Initial,omitempty"`

    // Comment.
    RangeEnd INewDocumentPosition `json:"RangeEnd,omitempty"`

    // Comment.
    RangeStart INewDocumentPosition `json:"RangeStart,omitempty"`

    // Comment.
    Text *string `json:"Text,omitempty"`
}

type ICommentBase interface {
    IsCommentBase() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
}

func (CommentBase) IsCommentBase() bool {
    return true
}


func (obj *CommentBase) Initialize() {
    if (obj.RangeEnd != nil) {
        obj.RangeEnd.Initialize()
    }

    if (obj.RangeStart != nil) {
        obj.RangeStart.Initialize()
    }


}

func (obj *CommentBase) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}


