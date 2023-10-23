/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="document_properties_response.go">
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

// The REST response with a collection of document properties.
// This response should be returned by the service when handling: GET /documentProperties.

type IDocumentPropertiesResponse interface {
    IsDocumentPropertiesResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetDocumentProperties() IDocumentProperties
    SetDocumentProperties(value IDocumentProperties)
}

type DocumentPropertiesResponse struct {
    // The REST response with a collection of document properties.
    // This response should be returned by the service when handling: GET /documentProperties.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a collection of document properties.
    // This response should be returned by the service when handling: GET /documentProperties.
    DocumentProperties IDocumentProperties `json:"DocumentProperties,omitempty"`
}

func (DocumentPropertiesResponse) IsDocumentPropertiesResponse() bool {
    return true
}

func (DocumentPropertiesResponse) IsWordsResponse() bool {
    return true
}

func (obj *DocumentPropertiesResponse) Initialize() {
    if (obj.DocumentProperties != nil) {
        obj.DocumentProperties.Initialize()
    }


}

func (obj *DocumentPropertiesResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["DocumentProperties"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IDocumentProperties = new(DocumentProperties)
            modelInstance.Deserialize(parsedValue)
            obj.DocumentProperties = modelInstance
        }

    } else if jsonValue, exists := json["documentProperties"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IDocumentProperties = new(DocumentProperties)
            modelInstance.Deserialize(parsedValue)
            obj.DocumentProperties = modelInstance
        }

    }
}

func (obj *DocumentPropertiesResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *DocumentPropertiesResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.RequestId == nil {
        return errors.New("Property RequestId in DocumentPropertiesResponse is required.")
    }

    if obj.DocumentProperties == nil {
        return errors.New("Property DocumentProperties in DocumentPropertiesResponse is required.")
    }

    if obj.DocumentProperties != nil {
        if err := obj.DocumentProperties.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *DocumentPropertiesResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *DocumentPropertiesResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *DocumentPropertiesResponse) GetDocumentProperties() IDocumentProperties {
    return obj.DocumentProperties
}

func (obj *DocumentPropertiesResponse) SetDocumentProperties(value IDocumentProperties) {
    obj.DocumentProperties = value
}

