/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="field_response.go">
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

import (
    "errors"
)

// The REST response with a field.
// This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/paragraphs/{0}/fields/{1}" REST API requests.

type IFieldResponse interface {
    IsFieldResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetField() IField
    SetField(value IField)
}

type FieldResponse struct {
    // The REST response with a field.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/paragraphs/{0}/fields/{1}" REST API requests.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a field.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/paragraphs/{0}/fields/{1}" REST API requests.
    Field IField `json:"Field,omitempty"`
}

func (FieldResponse) IsFieldResponse() bool {
    return true
}

func (FieldResponse) IsWordsResponse() bool {
    return true
}

func (obj *FieldResponse) Initialize() {
    if (obj.Field != nil) {
        obj.Field.Initialize()
    }


}

func (obj *FieldResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["Field"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IField = new(Field)
            modelInstance.Deserialize(parsedValue)
            obj.Field = modelInstance
        }

    } else if jsonValue, exists := json["field"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IField = new(Field)
            modelInstance.Deserialize(parsedValue)
            obj.Field = modelInstance
        }

    }
}

func (obj *FieldResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *FieldResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Field != nil {
        if err := obj.Field.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *FieldResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *FieldResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *FieldResponse) GetField() IField {
    return obj.Field
}

func (obj *FieldResponse) SetField(value IField) {
    obj.Field = value
}

