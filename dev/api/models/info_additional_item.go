/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="info_additional_item.go">
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

// Info additional item.

type IInfoAdditionalItem interface {
    IsInfoAdditionalItem() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetKey() *string
    SetKey(value *string)
    GetValue() *string
    SetValue(value *string)
}

type InfoAdditionalItem struct {
    // Info additional item.
    Key *string `json:"Key,omitempty"`

    // Info additional item.
    Value *string `json:"Value,omitempty"`
}

func (InfoAdditionalItem) IsInfoAdditionalItem() bool {
    return true
}


func (obj *InfoAdditionalItem) Initialize() {
}

func (obj *InfoAdditionalItem) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["Key"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Key = &parsedValue
        }

    } else if jsonValue, exists := json["key"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Key = &parsedValue
        }

    }

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

func (obj *InfoAdditionalItem) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *InfoAdditionalItem) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    return nil;
}

func (obj *InfoAdditionalItem) GetKey() *string {
    return obj.Key
}

func (obj *InfoAdditionalItem) SetKey(value *string) {
    obj.Key = value
}

func (obj *InfoAdditionalItem) GetValue() *string {
    return obj.Value
}

func (obj *InfoAdditionalItem) SetValue(value *string) {
    obj.Value = value
}

