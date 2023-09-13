/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="save_as_tiff_online_response.go">
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


// SaveAsTiffOnlineResponse struct
// Converts a document to TIFF format using detailed conversion settings.
type SaveAsTiffOnlineResponse struct {
    // The response model.
    Model ISaveResponse


    // The document after modification.
    Document map[string]io.Reader
}

func (obj *SaveAsTiffOnlineResponse) GetModel() ISaveResponse {
    return obj.Model
}

func (obj *SaveAsTiffOnlineResponse) SetModel(value ISaveResponse) {
    obj.Model = value
}

func (obj *SaveAsTiffOnlineResponse) GetDocument() map[string]io.Reader {
    return obj.Document
}

func (obj *SaveAsTiffOnlineResponse) SetDocument(value map[string]io.Reader) {
    obj.Document = value
}