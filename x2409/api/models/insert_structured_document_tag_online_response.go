/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="insert_structured_document_tag_online_response.go">
 *   Copyright (c) 2024 Aspose.Words for Cloud
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


// InsertStructuredDocumentTagOnlineResponse struct
// Inserts a new StructuredDocumentTag (SDT) to the document node.
type InsertStructuredDocumentTagOnlineResponse struct {
    // The REST response with a StructuredDocumentTag.
    Model IStructuredDocumentTagResponse


    // The document after modification.
    Document map[string]io.Reader
}

func (obj *InsertStructuredDocumentTagOnlineResponse) GetModel() IStructuredDocumentTagResponse {
    return obj.Model
}

func (obj *InsertStructuredDocumentTagOnlineResponse) SetModel(value IStructuredDocumentTagResponse) {
    obj.Model = value
}

func (obj *InsertStructuredDocumentTagOnlineResponse) GetDocument() map[string]io.Reader {
    return obj.Document
}

func (obj *InsertStructuredDocumentTagOnlineResponse) SetDocument(value map[string]io.Reader) {
    obj.Document = value
}