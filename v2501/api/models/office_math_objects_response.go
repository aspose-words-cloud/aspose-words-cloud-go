/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="office_math_objects_response.go">
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

// The REST response with a collection of OfficeMath objects.
// This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/OfficeMathObjects" REST API requests.

type IOfficeMathObjectsResponse interface {
    IsOfficeMathObjectsResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetOfficeMathObjects() IOfficeMathObjectsCollection
    SetOfficeMathObjects(value IOfficeMathObjectsCollection)
}

type OfficeMathObjectsResponse struct {
    // The REST response with a collection of OfficeMath objects.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/OfficeMathObjects" REST API requests.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a collection of OfficeMath objects.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/OfficeMathObjects" REST API requests.
    OfficeMathObjects IOfficeMathObjectsCollection `json:"OfficeMathObjects,omitempty"`
}

func (OfficeMathObjectsResponse) IsOfficeMathObjectsResponse() bool {
    return true
}

func (OfficeMathObjectsResponse) IsWordsResponse() bool {
    return true
}

func (obj *OfficeMathObjectsResponse) Initialize() {
    if (obj.OfficeMathObjects != nil) {
        obj.OfficeMathObjects.Initialize()
    }


}

func (obj *OfficeMathObjectsResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["OfficeMathObjects"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IOfficeMathObjectsCollection = new(OfficeMathObjectsCollection)
            modelInstance.Deserialize(parsedValue)
            obj.OfficeMathObjects = modelInstance
        }

    } else if jsonValue, exists := json["officeMathObjects"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IOfficeMathObjectsCollection = new(OfficeMathObjectsCollection)
            modelInstance.Deserialize(parsedValue)
            obj.OfficeMathObjects = modelInstance
        }

    }
}

func (obj *OfficeMathObjectsResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *OfficeMathObjectsResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.OfficeMathObjects != nil {
        if err := obj.OfficeMathObjects.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *OfficeMathObjectsResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *OfficeMathObjectsResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *OfficeMathObjectsResponse) GetOfficeMathObjects() IOfficeMathObjectsCollection {
    return obj.OfficeMathObjects
}

func (obj *OfficeMathObjectsResponse) SetOfficeMathObjects(value IOfficeMathObjectsCollection) {
    obj.OfficeMathObjects = value
}

