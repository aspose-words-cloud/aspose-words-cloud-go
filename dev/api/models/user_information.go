/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="user_information.go">
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

// DTO for user information.

type IUserInformation interface {
    IsUserInformation() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetAddress() *string
    SetAddress(value *string)
    GetInitials() *string
    SetInitials(value *string)
    GetName() *string
    SetName(value *string)
}

type UserInformation struct {
    // DTO for user information.
    Address *string `json:"Address,omitempty"`

    // DTO for user information.
    Initials *string `json:"Initials,omitempty"`

    // DTO for user information.
    Name *string `json:"Name,omitempty"`
}

func (UserInformation) IsUserInformation() bool {
    return true
}


func (obj *UserInformation) Initialize() {
}

func (obj *UserInformation) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["Address"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Address = &parsedValue
        }

    } else if jsonValue, exists := json["address"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Address = &parsedValue
        }

    }

    if jsonValue, exists := json["Initials"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Initials = &parsedValue
        }

    } else if jsonValue, exists := json["initials"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Initials = &parsedValue
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
}

func (obj *UserInformation) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *UserInformation) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    return nil;
}

func (obj *UserInformation) GetAddress() *string {
    return obj.Address
}

func (obj *UserInformation) SetAddress(value *string) {
    obj.Address = value
}

func (obj *UserInformation) GetInitials() *string {
    return obj.Initials
}

func (obj *UserInformation) SetInitials(value *string) {
    obj.Initials = value
}

func (obj *UserInformation) GetName() *string {
    return obj.Name
}

func (obj *UserInformation) SetName(value *string) {
    obj.Name = value
}

