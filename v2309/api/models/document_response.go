/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="document_response.go">
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

// The REST response with a document description.

type IDocumentResponse interface {
    IsDocumentResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetRequestId() *string
    SetRequestId(value *string)
    GetDocument() IDocument
    SetDocument(value IDocument)
}

type DocumentResponse struct {
    // The REST response with a document description.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a document description.
    Document IDocument `json:"Document,omitempty"`
}

func (DocumentResponse) IsDocumentResponse() bool {
    return true
}

func (DocumentResponse) IsWordsResponse() bool {
    return true
}

func (obj *DocumentResponse) Initialize() {
    if (obj.Document != nil) {
        obj.Document.Initialize()
    }


}

func (obj *DocumentResponse) Deserialize(json map[string]interface{}) {
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

func (obj *DocumentResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *DocumentResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *DocumentResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *DocumentResponse) GetDocument() IDocument {
    return obj.Document
}

func (obj *DocumentResponse) SetDocument(value IDocument) {
    obj.Document = value
}

