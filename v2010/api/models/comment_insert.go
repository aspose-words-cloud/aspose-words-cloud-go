/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="comment_insert.go">
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

import (
    "time"
)

// Comment insert.
type CommentInsert struct {
    // Comment insert.
    Author string `json:"Author,omitempty"`

    // Comment insert.
    DateTime time.Time `json:"DateTime,omitempty"`

    // Comment insert.
    Initial string `json:"Initial,omitempty"`

    // Comment insert.
    RangeEnd *DocumentPosition `json:"RangeEnd,omitempty"`

    // Comment insert.
    RangeStart *DocumentPosition `json:"RangeStart,omitempty"`

    // Comment insert.
    Text string `json:"Text,omitempty"`
}

type ICommentInsert interface {
    IsCommentInsert() bool
}
func (CommentInsert) IsCommentInsert() bool {
    return true
}

func (CommentInsert) IsCommentBase() bool {
    return true
}
