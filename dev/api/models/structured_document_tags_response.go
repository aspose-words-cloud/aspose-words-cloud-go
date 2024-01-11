/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="structured_document_tags_response.go">
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

import (
    "errors"
)

// The REST response with a collection of StructuredDocumentTags.
// This response should be returned by the service when handling: GET /structuredDocumentTags.

type IStructuredDocumentTagsResponse interface {
    IsStructuredDocumentTagsResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetStructuredDocumentTags() IStructuredDocumentTagCollection
    SetStructuredDocumentTags(value IStructuredDocumentTagCollection)
}

type StructuredDocumentTagsResponse struct {
    // The REST response with a collection of StructuredDocumentTags.
    // This response should be returned by the service when handling: GET /structuredDocumentTags.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a collection of StructuredDocumentTags.
    // This response should be returned by the service when handling: GET /structuredDocumentTags.
    StructuredDocumentTags IStructuredDocumentTagCollection `json:"StructuredDocumentTags,omitempty"`
}

func (StructuredDocumentTagsResponse) IsStructuredDocumentTagsResponse() bool {
    return true
}

func (StructuredDocumentTagsResponse) IsWordsResponse() bool {
    return true
}

func (obj *StructuredDocumentTagsResponse) Initialize() {
    if (obj.StructuredDocumentTags != nil) {
        obj.StructuredDocumentTags.Initialize()
    }


}

func (obj *StructuredDocumentTagsResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["StructuredDocumentTags"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IStructuredDocumentTagCollection = new(StructuredDocumentTagCollection)
            modelInstance.Deserialize(parsedValue)
            obj.StructuredDocumentTags = modelInstance
        }

    } else if jsonValue, exists := json["structuredDocumentTags"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IStructuredDocumentTagCollection = new(StructuredDocumentTagCollection)
            modelInstance.Deserialize(parsedValue)
            obj.StructuredDocumentTags = modelInstance
        }

    }
}

func (obj *StructuredDocumentTagsResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *StructuredDocumentTagsResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.StructuredDocumentTags != nil {
        if err := obj.StructuredDocumentTags.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *StructuredDocumentTagsResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *StructuredDocumentTagsResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *StructuredDocumentTagsResponse) GetStructuredDocumentTags() IStructuredDocumentTagCollection {
    return obj.StructuredDocumentTags
}

func (obj *StructuredDocumentTagsResponse) SetStructuredDocumentTags(value IStructuredDocumentTagCollection) {
    obj.StructuredDocumentTags = value
}

