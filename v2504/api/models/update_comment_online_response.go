/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="update_comment_online_response.go">
 *   Copyright (c) 2025 Aspose.Words for Cloud
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

import "io"


// UpdateCommentOnlineResponse struct
// Updates a comment in the document.
type UpdateCommentOnlineResponse struct {
    // The REST response with a comment.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/comments/0" REST API requests.
    Model ICommentResponse


    // The document after modification.
    Document map[string]io.Reader
}

func (obj *UpdateCommentOnlineResponse) GetModel() ICommentResponse {
    return obj.Model
}

func (obj *UpdateCommentOnlineResponse) SetModel(value ICommentResponse) {
    obj.Model = value
}

func (obj *UpdateCommentOnlineResponse) GetDocument() map[string]io.Reader {
    return obj.Document
}

func (obj *UpdateCommentOnlineResponse) SetDocument(value map[string]io.Reader) {
    obj.Document = value
}