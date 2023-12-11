/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="protection_request_v2.go">
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

// Request on changing of protection.

type IProtectionRequestV2 interface {
    IsProtectionRequestV2() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetProtectionPassword() *string
    SetProtectionPassword(value *string)
    GetProtectionType() *string
    SetProtectionType(value *string)
}

type ProtectionRequestV2 struct {
    // Request on changing of protection.
    ProtectionPassword *string `json:"ProtectionPassword,omitempty"`

    // Request on changing of protection.
    ProtectionType *string `json:"ProtectionType,omitempty"`
}

func (ProtectionRequestV2) IsProtectionRequestV2() bool {
    return true
}

func (ProtectionRequestV2) IsProtectionRequestBase() bool {
    return true
}

func (obj *ProtectionRequestV2) Initialize() {
}

func (obj *ProtectionRequestV2) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["ProtectionPassword"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ProtectionPassword = &parsedValue
        }

    } else if jsonValue, exists := json["protectionPassword"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ProtectionPassword = &parsedValue
        }

    }

    if jsonValue, exists := json["ProtectionType"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ProtectionType = &parsedValue
        }

    } else if jsonValue, exists := json["protectionType"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ProtectionType = &parsedValue
        }

    }
}

func (obj *ProtectionRequestV2) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *ProtectionRequestV2) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.ProtectionPassword == nil {
        return errors.New("Property ProtectionPassword in ProtectionRequestV2 is required.")
    }
    if obj.ProtectionType == nil {
        return errors.New("Property ProtectionType in ProtectionRequestV2 is required.")
    }
    return nil;
}

func (obj *ProtectionRequestV2) GetProtectionPassword() *string {
    return obj.ProtectionPassword
}

func (obj *ProtectionRequestV2) SetProtectionPassword(value *string) {
    obj.ProtectionPassword = value
}

func (obj *ProtectionRequestV2) GetProtectionType() *string {
    return obj.ProtectionType
}

func (obj *ProtectionRequestV2) SetProtectionType(value *string) {
    obj.ProtectionType = value
}

