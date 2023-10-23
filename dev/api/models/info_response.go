/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="info_response.go">
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

// Response with API info.

type IInfoResponse interface {
    IsInfoResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetAdditionalInfo() []IInfoAdditionalItem
    SetAdditionalInfo(value []IInfoAdditionalItem)
    GetName() *string
    SetName(value *string)
    GetVersion() *string
    SetVersion(value *string)
}

type InfoResponse struct {
    // Response with API info.
    RequestId *string `json:"RequestId,omitempty"`

    // Response with API info.
    AdditionalInfo []IInfoAdditionalItem `json:"AdditionalInfo,omitempty"`

    // Response with API info.
    Name *string `json:"Name,omitempty"`

    // Response with API info.
    Version *string `json:"Version,omitempty"`
}

func (InfoResponse) IsInfoResponse() bool {
    return true
}

func (InfoResponse) IsWordsResponse() bool {
    return true
}

func (obj *InfoResponse) Initialize() {
    if (obj.AdditionalInfo != nil) {
        for _, objElementAdditionalInfo := range obj.AdditionalInfo {
            objElementAdditionalInfo.Initialize()
        }
    }

}

func (obj *InfoResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["AdditionalInfo"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.AdditionalInfo = make([]IInfoAdditionalItem, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IInfoAdditionalItem = new(InfoAdditionalItem)
                    modelElementInstance.Deserialize(elementValue)
                    obj.AdditionalInfo = append(obj.AdditionalInfo, modelElementInstance)
                }

            }
        }

    } else if jsonValue, exists := json["additionalInfo"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.AdditionalInfo = make([]IInfoAdditionalItem, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IInfoAdditionalItem = new(InfoAdditionalItem)
                    modelElementInstance.Deserialize(elementValue)
                    obj.AdditionalInfo = append(obj.AdditionalInfo, modelElementInstance)
                }

            }
        }

    }

    if jsonValue, exists := json["Name"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Name = &parsedValue
        }

    } else if jsonValue, exists := json["name"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Name = &parsedValue
        }

    }

    if jsonValue, exists := json["Version"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Version = &parsedValue
        }

    } else if jsonValue, exists := json["version"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Version = &parsedValue
        }

    }
}

func (obj *InfoResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *InfoResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.RequestId == nil {
        return errors.New("Property RequestId in InfoResponse is required.")
    }

    if obj.AdditionalInfo == nil {
        return errors.New("Property AdditionalInfo in InfoResponse is required.")
    }

    if obj.AdditionalInfo != nil {
        for _, elementAdditionalInfo := range obj.AdditionalInfo {
            if elementAdditionalInfo != nil {
                if err := elementAdditionalInfo.Validate(); err != nil {
                    return err
                }
            }
        }
    }

    if obj.Name == nil {
        return errors.New("Property Name in InfoResponse is required.")
    }

    if obj.Version == nil {
        return errors.New("Property Version in InfoResponse is required.")
    }

    return nil;
}

func (obj *InfoResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *InfoResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *InfoResponse) GetAdditionalInfo() []IInfoAdditionalItem {
    return obj.AdditionalInfo
}

func (obj *InfoResponse) SetAdditionalInfo(value []IInfoAdditionalItem) {
    obj.AdditionalInfo = value
}

func (obj *InfoResponse) GetName() *string {
    return obj.Name
}

func (obj *InfoResponse) SetName(value *string) {
    obj.Name = value
}

func (obj *InfoResponse) GetVersion() *string {
    return obj.Version
}

func (obj *InfoResponse) SetVersion(value *string) {
    obj.Version = value
}

