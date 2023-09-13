/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="update_paragraph_format_online_response.go">
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

import "io"


// UpdateParagraphFormatOnlineResponse struct
// Updates the formatting properties of a paragraph in the document node.
type UpdateParagraphFormatOnlineResponse struct {
    // The response model.
    Model IParagraphFormatResponse


    // The document after modification.
    Document map[string]io.Reader
}

func (obj *UpdateParagraphFormatOnlineResponse) GetModel() IParagraphFormatResponse {
    return obj.Model
}

func (obj *UpdateParagraphFormatOnlineResponse) SetModel(value IParagraphFormatResponse) {
    obj.Model = value
}

func (obj *UpdateParagraphFormatOnlineResponse) GetDocument() map[string]io.Reader {
    return obj.Document
}

func (obj *UpdateParagraphFormatOnlineResponse) SetDocument(value map[string]io.Reader) {
    obj.Document = value
}