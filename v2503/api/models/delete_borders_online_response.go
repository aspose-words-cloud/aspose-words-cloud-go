/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="delete_borders_online_response.go">
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


// DeleteBordersOnlineResponse struct
// Removes borders from the document node.
// The 'nodePath' parameter should refer to a paragraph, a cell or a row.
type DeleteBordersOnlineResponse struct {
    // The REST response with a collection of borders.
    // This response is returned by the Service when handling "GET {nodeWithBorders}/borders" REST API requests.
    Model IBordersResponse


    // The document after modification.
    Document map[string]io.Reader
}

func (obj *DeleteBordersOnlineResponse) GetModel() IBordersResponse {
    return obj.Model
}

func (obj *DeleteBordersOnlineResponse) SetModel(value IBordersResponse) {
    obj.Model = value
}

func (obj *DeleteBordersOnlineResponse) GetDocument() map[string]io.Reader {
    return obj.Document
}

func (obj *DeleteBordersOnlineResponse) SetDocument(value map[string]io.Reader) {
    obj.Document = value
}