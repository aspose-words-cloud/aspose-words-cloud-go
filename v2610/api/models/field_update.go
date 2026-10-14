/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="field_update.go">
 *   Copyright (c) 2026 Aspose.Words for Cloud
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

// Field for update.

type IFieldUpdate interface {
    IsFieldUpdate() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetLocaleId() *string
    SetLocaleId(value *string)
    GetFieldCode() *string
    SetFieldCode(value *string)
}

type FieldUpdate struct {
    // Field for update.
    LocaleId *string `json:"LocaleId,omitempty"`

    // Field for update.
    FieldCode *string `json:"FieldCode,omitempty"`
}

func (FieldUpdate) IsFieldUpdate() bool {
    return true
}

func (FieldUpdate) IsFieldBase() bool {
    return true
}

func (obj *FieldUpdate) Initialize() {
}

func (obj *FieldUpdate) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["LocaleId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.LocaleId = &parsedValue
        }

    } else if jsonValue, exists := json["localeId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.LocaleId = &parsedValue
        }

    }

    if jsonValue, exists := json["FieldCode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FieldCode = &parsedValue
        }

    } else if jsonValue, exists := json["fieldCode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FieldCode = &parsedValue
        }

    }
}

func (obj *FieldUpdate) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *FieldUpdate) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.FieldCode == nil {
        return errors.New("Property FieldCode in FieldUpdate is required.")
    }
    return nil;
}

func (obj *FieldUpdate) GetLocaleId() *string {
    return obj.LocaleId
}

func (obj *FieldUpdate) SetLocaleId(value *string) {
    obj.LocaleId = value
}

func (obj *FieldUpdate) GetFieldCode() *string {
    return obj.FieldCode
}

func (obj *FieldUpdate) SetFieldCode(value *string) {
    obj.FieldCode = value
}

