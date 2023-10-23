/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="field_names_response.go">
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

// The REST response with a collection of mail merge fields.
// This response should be returned by the service when handling: GET /{name}/mailMergeFieldNames.

type IFieldNamesResponse interface {
    IsFieldNamesResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetFieldNames() IFieldNames
    SetFieldNames(value IFieldNames)
}

type FieldNamesResponse struct {
    // The REST response with a collection of mail merge fields.
    // This response should be returned by the service when handling: GET /{name}/mailMergeFieldNames.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a collection of mail merge fields.
    // This response should be returned by the service when handling: GET /{name}/mailMergeFieldNames.
    FieldNames IFieldNames `json:"FieldNames,omitempty"`
}

func (FieldNamesResponse) IsFieldNamesResponse() bool {
    return true
}

func (FieldNamesResponse) IsWordsResponse() bool {
    return true
}

func (obj *FieldNamesResponse) Initialize() {
    if (obj.FieldNames != nil) {
        obj.FieldNames.Initialize()
    }


}

func (obj *FieldNamesResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["FieldNames"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFieldNames = new(FieldNames)
            modelInstance.Deserialize(parsedValue)
            obj.FieldNames = modelInstance
        }

    } else if jsonValue, exists := json["fieldNames"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFieldNames = new(FieldNames)
            modelInstance.Deserialize(parsedValue)
            obj.FieldNames = modelInstance
        }

    }
}

func (obj *FieldNamesResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *FieldNamesResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.RequestId == nil {
        return errors.New("Property RequestId in FieldNamesResponse is required.")
    }

    if obj.FieldNames == nil {
        return errors.New("Property FieldNames in FieldNamesResponse is required.")
    }

    if obj.FieldNames != nil {
        if err := obj.FieldNames.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *FieldNamesResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *FieldNamesResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *FieldNamesResponse) GetFieldNames() IFieldNames {
    return obj.FieldNames
}

func (obj *FieldNamesResponse) SetFieldNames(value IFieldNames) {
    obj.FieldNames = value
}

