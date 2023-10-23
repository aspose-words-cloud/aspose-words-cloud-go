/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="run_insert.go">
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

// Run element for insert.

type IRunInsert interface {
    IsRunInsert() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetText() *string
    SetText(value *string)
}

type RunInsert struct {
    // Run element for insert.
    Text *string `json:"Text,omitempty"`
}

func (RunInsert) IsRunInsert() bool {
    return true
}

func (RunInsert) IsRunBase() bool {
    return true
}

func (obj *RunInsert) Initialize() {
}

func (obj *RunInsert) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["Text"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Text = &parsedValue
        }

    } else if jsonValue, exists := json["text"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Text = &parsedValue
        }

    }
}

func (obj *RunInsert) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *RunInsert) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Text == nil {
        return errors.New("Property Text in RunInsert is required.")
    }
    if obj.Text == nil {
        return errors.New("Property Text in RunInsert is required.")
    }

    return nil;
}

func (obj *RunInsert) GetText() *string {
    return obj.Text
}

func (obj *RunInsert) SetText(value *string) {
    obj.Text = value
}

