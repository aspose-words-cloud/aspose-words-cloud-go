/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="comment.go">
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

// DTO container with a comment.
type Comment struct {
    // DTO container with a comment.
    Link *WordsApiLink `json:"Link,omitempty"`

    // DTO container with a comment.
    Author *string `json:"Author,omitempty"`

    // DTO container with a comment.
    Content *StoryChildNodes `json:"Content,omitempty"`

    // DTO container with a comment.
    DateTime time.Time `json:"DateTime,omitempty"`

    // DTO container with a comment.
    Initial *string `json:"Initial,omitempty"`

    // DTO container with a comment.
    RangeEnd *DocumentPosition `json:"RangeEnd,omitempty"`

    // DTO container with a comment.
    RangeStart *DocumentPosition `json:"RangeStart,omitempty"`

    // DTO container with a comment.
    Text *string `json:"Text,omitempty"`
}

type IComment interface {
    IsComment() bool
}
func (Comment) IsComment() bool {
    return true
}

func (Comment) IsCommentLink() bool {
    return true
}
