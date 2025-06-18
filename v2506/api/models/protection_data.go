/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="protection_data.go">
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

// Container for the data about protection of the document.

type IProtectionData interface {
    IsProtectionData() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetProtectionType() *string
    SetProtectionType(value *string)
}

type ProtectionData struct {
    // Container for the data about protection of the document.
    ProtectionType *string `json:"ProtectionType,omitempty"`
}

func (ProtectionData) IsProtectionData() bool {
    return true
}


func (obj *ProtectionData) Initialize() {
}

func (obj *ProtectionData) Deserialize(json map[string]interface{}) {
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

func (obj *ProtectionData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *ProtectionData) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.ProtectionType == nil {
        return errors.New("Property ProtectionType in ProtectionData is required.")
    }
    return nil;
}

func (obj *ProtectionData) GetProtectionType() *string {
    return obj.ProtectionType
}

func (obj *ProtectionData) SetProtectionType(value *string) {
    obj.ProtectionType = value
}

