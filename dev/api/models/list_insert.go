/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="list_insert.go">
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

// Insert document to document list.

type IListInsert interface {
    IsListInsert() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetTemplate() *string
    SetTemplate(value *string)
}

type ListInsert struct {
    // Insert document to document list.
    Template *string `json:"Template,omitempty"`
}

func (ListInsert) IsListInsert() bool {
    return true
}


func (obj *ListInsert) Initialize() {
}

func (obj *ListInsert) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["Template"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Template = &parsedValue
        }

    } else if jsonValue, exists := json["template"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Template = &parsedValue
        }

    }
}

func (obj *ListInsert) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *ListInsert) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Template == nil {
        return errors.New("Property Template in ListInsert is required.")
    }
    return nil;
}

func (obj *ListInsert) GetTemplate() *string {
    return obj.Template
}

func (obj *ListInsert) SetTemplate(value *string) {
    obj.Template = value
}

