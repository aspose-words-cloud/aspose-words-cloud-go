/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="protection_request.go">
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

type IProtectionRequest interface {
    IsProtectionRequest() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetNewPassword() *string
    SetNewPassword(value *string)
    GetPassword() *string
    SetPassword(value *string)
    GetProtectionType() *string
    SetProtectionType(value *string)
}

type ProtectionRequest struct {
    // Request on changing of protection.
    NewPassword *string `json:"NewPassword,omitempty"`

    // Request on changing of protection.
    Password *string `json:"Password,omitempty"`

    // Request on changing of protection.
    ProtectionType *string `json:"ProtectionType,omitempty"`
}

func (ProtectionRequest) IsProtectionRequest() bool {
    return true
}


func (obj *ProtectionRequest) Initialize() {
}

func (obj *ProtectionRequest) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["NewPassword"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NewPassword = &parsedValue
        }

    } else if jsonValue, exists := json["newPassword"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NewPassword = &parsedValue
        }

    }

    if jsonValue, exists := json["Password"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Password = &parsedValue
        }

    } else if jsonValue, exists := json["password"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Password = &parsedValue
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

func (obj *ProtectionRequest) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *ProtectionRequest) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Password == nil {
        return errors.New("Property Password in ProtectionRequest is required.")
    }

    return nil;
}

func (obj *ProtectionRequest) GetNewPassword() *string {
    return obj.NewPassword
}

func (obj *ProtectionRequest) SetNewPassword(value *string) {
    obj.NewPassword = value
}

func (obj *ProtectionRequest) GetPassword() *string {
    return obj.Password
}

func (obj *ProtectionRequest) SetPassword(value *string) {
    obj.Password = value
}

func (obj *ProtectionRequest) GetProtectionType() *string {
    return obj.ProtectionType
}

func (obj *ProtectionRequest) SetProtectionType(value *string) {
    obj.ProtectionType = value
}

