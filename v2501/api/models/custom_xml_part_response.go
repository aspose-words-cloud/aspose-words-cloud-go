/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="custom_xml_part_response.go">
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

// The REST response with a custom xml part.
// This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/customXmlParts/0" REST API requests.

type ICustomXmlPartResponse interface {
    IsCustomXmlPartResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetCustomXmlPart() ICustomXmlPart
    SetCustomXmlPart(value ICustomXmlPart)
}

type CustomXmlPartResponse struct {
    // The REST response with a custom xml part.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/customXmlParts/0" REST API requests.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a custom xml part.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/customXmlParts/0" REST API requests.
    CustomXmlPart ICustomXmlPart `json:"CustomXmlPart,omitempty"`
}

func (CustomXmlPartResponse) IsCustomXmlPartResponse() bool {
    return true
}

func (CustomXmlPartResponse) IsWordsResponse() bool {
    return true
}

func (obj *CustomXmlPartResponse) Initialize() {
    if (obj.CustomXmlPart != nil) {
        obj.CustomXmlPart.Initialize()
    }


}

func (obj *CustomXmlPartResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["CustomXmlPart"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ICustomXmlPart = nil
            if jsonType, found := parsedValue["$type"]; found {
                jsonTypeStr := jsonType.(string)
                if jsonTypeStr == "CustomXmlPartInsert, _" { modelInstance = new(CustomXmlPartInsert) }
                if jsonTypeStr == "CustomXmlPartUpdate, _" { modelInstance = new(CustomXmlPartUpdate) }
            }

            if modelInstance == nil { modelInstance = new(CustomXmlPart) }
            modelInstance.Deserialize(parsedValue)
            obj.CustomXmlPart = modelInstance
        }

    } else if jsonValue, exists := json["customXmlPart"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ICustomXmlPart = nil
            if jsonType, found := parsedValue["$type"]; found {
                jsonTypeStr := jsonType.(string)
                if jsonTypeStr == "CustomXmlPartInsert, _" { modelInstance = new(CustomXmlPartInsert) }
                if jsonTypeStr == "CustomXmlPartUpdate, _" { modelInstance = new(CustomXmlPartUpdate) }
            }

            if modelInstance == nil { modelInstance = new(CustomXmlPart) }
            modelInstance.Deserialize(parsedValue)
            obj.CustomXmlPart = modelInstance
        }

    }
}

func (obj *CustomXmlPartResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *CustomXmlPartResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.CustomXmlPart != nil {
        if err := obj.CustomXmlPart.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *CustomXmlPartResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *CustomXmlPartResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *CustomXmlPartResponse) GetCustomXmlPart() ICustomXmlPart {
    return obj.CustomXmlPart
}

func (obj *CustomXmlPartResponse) SetCustomXmlPart(value ICustomXmlPart) {
    obj.CustomXmlPart = value
}

