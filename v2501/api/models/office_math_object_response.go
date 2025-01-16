/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="office_math_object_response.go">
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

// The REST response with a OfficeMath object.
// This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/officeMathObjects/0" REST API requests.

type IOfficeMathObjectResponse interface {
    IsOfficeMathObjectResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetOfficeMathObject() IOfficeMathObject
    SetOfficeMathObject(value IOfficeMathObject)
}

type OfficeMathObjectResponse struct {
    // The REST response with a OfficeMath object.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/officeMathObjects/0" REST API requests.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a OfficeMath object.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/officeMathObjects/0" REST API requests.
    OfficeMathObject IOfficeMathObject `json:"OfficeMathObject,omitempty"`
}

func (OfficeMathObjectResponse) IsOfficeMathObjectResponse() bool {
    return true
}

func (OfficeMathObjectResponse) IsWordsResponse() bool {
    return true
}

func (obj *OfficeMathObjectResponse) Initialize() {
    if (obj.OfficeMathObject != nil) {
        obj.OfficeMathObject.Initialize()
    }


}

func (obj *OfficeMathObjectResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["OfficeMathObject"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IOfficeMathObject = new(OfficeMathObject)
            modelInstance.Deserialize(parsedValue)
            obj.OfficeMathObject = modelInstance
        }

    } else if jsonValue, exists := json["officeMathObject"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IOfficeMathObject = new(OfficeMathObject)
            modelInstance.Deserialize(parsedValue)
            obj.OfficeMathObject = modelInstance
        }

    }
}

func (obj *OfficeMathObjectResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *OfficeMathObjectResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.OfficeMathObject != nil {
        if err := obj.OfficeMathObject.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *OfficeMathObjectResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *OfficeMathObjectResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *OfficeMathObjectResponse) GetOfficeMathObject() IOfficeMathObject {
    return obj.OfficeMathObject
}

func (obj *OfficeMathObjectResponse) SetOfficeMathObject(value IOfficeMathObject) {
    obj.OfficeMathObject = value
}

