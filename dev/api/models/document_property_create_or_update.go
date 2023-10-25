/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="document_property_create_or_update.go">
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

// Words document property DTO for create or update.

type IDocumentPropertyCreateOrUpdate interface {
    IsDocumentPropertyCreateOrUpdate() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetValue() *string
    SetValue(value *string)
}

type DocumentPropertyCreateOrUpdate struct {
    // Words document property DTO for create or update.
    Value *string `json:"Value,omitempty"`
}

func (DocumentPropertyCreateOrUpdate) IsDocumentPropertyCreateOrUpdate() bool {
    return true
}

func (DocumentPropertyCreateOrUpdate) IsDocumentPropertyBase() bool {
    return true
}

func (obj *DocumentPropertyCreateOrUpdate) Initialize() {
}

func (obj *DocumentPropertyCreateOrUpdate) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["Value"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Value = &parsedValue
        }

    } else if jsonValue, exists := json["value"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Value = &parsedValue
        }

    }
}

func (obj *DocumentPropertyCreateOrUpdate) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *DocumentPropertyCreateOrUpdate) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    return nil;
}

func (obj *DocumentPropertyCreateOrUpdate) GetValue() *string {
    return obj.Value
}

func (obj *DocumentPropertyCreateOrUpdate) SetValue(value *string) {
    obj.Value = value
}

