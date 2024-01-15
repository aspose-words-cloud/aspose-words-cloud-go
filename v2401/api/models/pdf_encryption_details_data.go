/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="pdf_encryption_details_data.go">
 *   Copyright (c) 2024 Aspose.Words for Cloud
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

// Container class for details of encryption.

type IPdfEncryptionDetailsData interface {
    IsPdfEncryptionDetailsData() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetOwnerPassword() *string
    SetOwnerPassword(value *string)
    GetPermissions() []string
    SetPermissions(value []string)
    GetUserPassword() *string
    SetUserPassword(value *string)
}

type PdfEncryptionDetailsData struct {
    // Container class for details of encryption.
    OwnerPassword *string `json:"OwnerPassword,omitempty"`

    // Container class for details of encryption.
    Permissions []string `json:"Permissions,omitempty"`

    // Container class for details of encryption.
    UserPassword *string `json:"UserPassword,omitempty"`
}

func (PdfEncryptionDetailsData) IsPdfEncryptionDetailsData() bool {
    return true
}


func (obj *PdfEncryptionDetailsData) Initialize() {
}

func (obj *PdfEncryptionDetailsData) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["OwnerPassword"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.OwnerPassword = &parsedValue
        }

    } else if jsonValue, exists := json["ownerPassword"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.OwnerPassword = &parsedValue
        }

    }

    if jsonValue, exists := json["Permissions"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.Permissions = make([]string, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(string); valid {
                    obj.Permissions = append(obj.Permissions, elementValue)
                }

            }
        }

    } else if jsonValue, exists := json["permissions"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.Permissions = make([]string, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(string); valid {
                    obj.Permissions = append(obj.Permissions, elementValue)
                }

            }
        }

    }

    if jsonValue, exists := json["UserPassword"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.UserPassword = &parsedValue
        }

    } else if jsonValue, exists := json["userPassword"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.UserPassword = &parsedValue
        }

    }
}

func (obj *PdfEncryptionDetailsData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *PdfEncryptionDetailsData) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    return nil;
}

func (obj *PdfEncryptionDetailsData) GetOwnerPassword() *string {
    return obj.OwnerPassword
}

func (obj *PdfEncryptionDetailsData) SetOwnerPassword(value *string) {
    obj.OwnerPassword = value
}

func (obj *PdfEncryptionDetailsData) GetPermissions() []string {
    return obj.Permissions
}

func (obj *PdfEncryptionDetailsData) SetPermissions(value []string) {
    obj.Permissions = value
}

func (obj *PdfEncryptionDetailsData) GetUserPassword() *string {
    return obj.UserPassword
}

func (obj *PdfEncryptionDetailsData) SetUserPassword(value *string) {
    obj.UserPassword = value
}

