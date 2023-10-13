/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="compress_response.go">
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

// The REST response of compressed document.

type ICompressResponse interface {
    IsCompressResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetDocument() IDocument
    SetDocument(value IDocument)
}

type CompressResponse struct {
    // The REST response of compressed document.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response of compressed document.
    Document IDocument `json:"Document,omitempty"`
}

func (CompressResponse) IsCompressResponse() bool {
    return true
}

func (CompressResponse) IsWordsResponse() bool {
    return true
}

func (obj *CompressResponse) Initialize() {
    if (obj.Document != nil) {
        obj.Document.Initialize()
    }


}

func (obj *CompressResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["Document"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IDocument = new(Document)
            modelInstance.Deserialize(parsedValue)
            obj.Document = modelInstance
        }

    } else if jsonValue, exists := json["document"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IDocument = new(Document)
            modelInstance.Deserialize(parsedValue)
            obj.Document = modelInstance
        }

    }
}

func (obj *CompressResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *CompressResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    return nil;
}

func (obj *CompressResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *CompressResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *CompressResponse) GetDocument() IDocument {
    return obj.Document
}

func (obj *CompressResponse) SetDocument(value IDocument) {
    obj.Document = value
}

