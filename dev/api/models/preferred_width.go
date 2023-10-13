/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="preferred_width.go">
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

// DTO container with a preferred width value.

type IPreferredWidth interface {
    IsPreferredWidth() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetType() *string
    SetType(value *string)
    GetValue() *float64
    SetValue(value *float64)
}

type PreferredWidth struct {
    // DTO container with a preferred width value.
    Type *string `json:"Type,omitempty"`

    // DTO container with a preferred width value.
    Value *float64 `json:"Value,omitempty"`
}

func (PreferredWidth) IsPreferredWidth() bool {
    return true
}


func (obj *PreferredWidth) Initialize() {
}

func (obj *PreferredWidth) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["Type"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Type = &parsedValue
        }

    } else if jsonValue, exists := json["type"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Type = &parsedValue
        }

    }

    if jsonValue, exists := json["Value"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Value = &parsedValue
        }

    } else if jsonValue, exists := json["value"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Value = &parsedValue
        }

    }
}

func (obj *PreferredWidth) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *PreferredWidth) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Type == nil {
        return errors.New("Property Type in PreferredWidth is required.")
    }

    return nil;
}

func (obj *PreferredWidth) GetType() *string {
    return obj.Type
}

func (obj *PreferredWidth) SetType(value *string) {
    obj.Type = value
}

func (obj *PreferredWidth) GetValue() *float64 {
    return obj.Value
}

func (obj *PreferredWidth) SetValue(value *float64) {
    obj.Value = value
}

