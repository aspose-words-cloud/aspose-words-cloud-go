/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="document_property_response.go">
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

// The REST response with a document property.

type IDocumentPropertyResponse interface {
    IsDocumentPropertyResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetRequestId() *string
    SetRequestId(value *string)
    GetDocumentProperty() IDocumentProperty
    SetDocumentProperty(value IDocumentProperty)
}

type DocumentPropertyResponse struct {
    // The REST response with a document property.
    RequestId *string

    // The REST response with a document property.
    DocumentProperty IDocumentProperty
}

func (DocumentPropertyResponse) IsDocumentPropertyResponse() bool {
    return true
}

func (DocumentPropertyResponse) IsWordsResponse() bool {
    return true
}

func (obj *DocumentPropertyResponse) Initialize() {
    if (obj.DocumentProperty != nil) {
        obj.DocumentProperty.Initialize()
    }


}

func (obj *DocumentPropertyResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["DocumentProperty"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IDocumentProperty = new(DocumentProperty)
            modelInstance.Deserialize(parsedValue)
            obj.DocumentProperty = modelInstance
        }

    } else if jsonValue, exists := json["documentProperty"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IDocumentProperty = new(DocumentProperty)
            modelInstance.Deserialize(parsedValue)
            obj.DocumentProperty = modelInstance
        }

    }
}

func (obj *DocumentPropertyResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *DocumentPropertyResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *DocumentPropertyResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *DocumentPropertyResponse) GetDocumentProperty() IDocumentProperty {
    return obj.DocumentProperty
}

func (obj *DocumentPropertyResponse) SetDocumentProperty(value IDocumentProperty) {
    obj.DocumentProperty = value
}

