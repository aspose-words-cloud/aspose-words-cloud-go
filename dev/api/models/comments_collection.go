/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="comments_collection.go">
 *   Copyright (c) 2021 Aspose.Words for Cloud
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

// The collection of comments.
type CommentsCollectionResult struct {
    // The collection of comments.
    Link WordsApiLinkResult `json:"Link,omitempty"`

    // The collection of comments.
    CommentList []CommentResult `json:"CommentList,omitempty"`
}

type CommentsCollection struct {
    // The collection of comments.
    Link IWordsApiLink `json:"Link,omitempty"`

    // The collection of comments.
    CommentList []Comment `json:"CommentList,omitempty"`
}

type ICommentsCollection interface {
    IsCommentsCollection() bool
}
func (CommentsCollection) IsCommentsCollection() bool {
    return true
}

func (CommentsCollection) IsLinkElement() bool {
    return true
}


