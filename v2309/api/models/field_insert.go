/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="field_insert.go">
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

// Field for insert.

type IFieldInsert interface {
    IsFieldInsert() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetFieldCode() *string
    SetFieldCode(value *string)
    GetLocaleId() *string
    SetLocaleId(value *string)
}

type FieldInsert struct {
    // Field for insert.
    FieldCode *string `json:"FieldCode,omitempty"`

    // Field for insert.
    LocaleId *string `json:"LocaleId,omitempty"`
}

func (FieldInsert) IsFieldInsert() bool {
    return true
}

func (FieldInsert) IsFieldBase() bool {
    return true
}

func (obj *FieldInsert) Initialize() {
}

func (obj *FieldInsert) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["FieldCode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FieldCode = &parsedValue
        }

    } else if jsonValue, exists := json["fieldCode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FieldCode = &parsedValue
        }

    }

    if jsonValue, exists := json["LocaleId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.LocaleId = &parsedValue
        }

    } else if jsonValue, exists := json["localeId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.LocaleId = &parsedValue
        }

    }
}

func (obj *FieldInsert) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *FieldInsert) GetFieldCode() *string {
    return obj.FieldCode
}

func (obj *FieldInsert) SetFieldCode(value *string) {
    obj.FieldCode = value
}

func (obj *FieldInsert) GetLocaleId() *string {
    return obj.LocaleId
}

func (obj *FieldInsert) SetLocaleId(value *string) {
    obj.LocaleId = value
}

