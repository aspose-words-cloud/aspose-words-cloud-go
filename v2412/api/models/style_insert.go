/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="style_insert.go">
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

// Represents a single document style to insert.

type IStyleInsert interface {
    IsStyleInsert() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetStyleName() *string
    SetStyleName(value *string)
    GetStyleType() *string
    SetStyleType(value *string)
}

type StyleInsert struct {
    // Represents a single document style to insert.
    StyleName *string `json:"StyleName,omitempty"`

    // Represents a single document style to insert.
    StyleType *string `json:"StyleType,omitempty"`
}

func (StyleInsert) IsStyleInsert() bool {
    return true
}


func (obj *StyleInsert) Initialize() {
}

func (obj *StyleInsert) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["StyleName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.StyleName = &parsedValue
        }

    } else if jsonValue, exists := json["styleName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.StyleName = &parsedValue
        }

    }

    if jsonValue, exists := json["StyleType"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.StyleType = &parsedValue
        }

    } else if jsonValue, exists := json["styleType"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.StyleType = &parsedValue
        }

    }
}

func (obj *StyleInsert) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *StyleInsert) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.StyleName == nil {
        return errors.New("Property StyleName in StyleInsert is required.")
    }
    if obj.StyleType == nil {
        return errors.New("Property StyleType in StyleInsert is required.")
    }
    return nil;
}

func (obj *StyleInsert) GetStyleName() *string {
    return obj.StyleName
}

func (obj *StyleInsert) SetStyleName(value *string) {
    obj.StyleName = value
}

func (obj *StyleInsert) GetStyleType() *string {
    return obj.StyleType
}

func (obj *StyleInsert) SetStyleType(value *string) {
    obj.StyleType = value
}

