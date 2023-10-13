/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="comment_response.go">
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

import (
    "errors"
)

// The REST response with a comment.
// This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/comments/0" REST API requests.

type ICommentResponse interface {
    IsCommentResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetComment() IComment
    SetComment(value IComment)
}

type CommentResponse struct {
    // The REST response with a comment.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/comments/0" REST API requests.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a comment.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/comments/0" REST API requests.
    Comment IComment `json:"Comment,omitempty"`
}

func (CommentResponse) IsCommentResponse() bool {
    return true
}

func (CommentResponse) IsWordsResponse() bool {
    return true
}

func (obj *CommentResponse) Initialize() {
    if (obj.Comment != nil) {
        obj.Comment.Initialize()
    }


}

func (obj *CommentResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["Comment"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IComment = new(Comment)
            modelInstance.Deserialize(parsedValue)
            obj.Comment = modelInstance
        }

    } else if jsonValue, exists := json["comment"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IComment = new(Comment)
            modelInstance.Deserialize(parsedValue)
            obj.Comment = modelInstance
        }

    }
}

func (obj *CommentResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *CommentResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    return nil;
}

func (obj *CommentResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *CommentResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *CommentResponse) GetComment() IComment {
    return obj.Comment
}

func (obj *CommentResponse) SetComment(value IComment) {
    obj.Comment = value
}

