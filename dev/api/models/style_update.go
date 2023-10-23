/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="style_update.go">
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

// Represents a single document style properties to update.

type IStyleUpdate interface {
    IsStyleUpdate() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetNextParagraphStyleName() *string
    SetNextParagraphStyleName(value *string)
    GetBaseStyleName() *string
    SetBaseStyleName(value *string)
    GetIsQuickStyle() *bool
    SetIsQuickStyle(value *bool)
    GetName() *string
    SetName(value *string)
}

type StyleUpdate struct {
    // Represents a single document style properties to update.
    NextParagraphStyleName *string `json:"NextParagraphStyleName,omitempty"`

    // Represents a single document style properties to update.
    BaseStyleName *string `json:"BaseStyleName,omitempty"`

    // Represents a single document style properties to update.
    IsQuickStyle *bool `json:"IsQuickStyle,omitempty"`

    // Represents a single document style properties to update.
    Name *string `json:"Name,omitempty"`
}

func (StyleUpdate) IsStyleUpdate() bool {
    return true
}


func (obj *StyleUpdate) Initialize() {
}

func (obj *StyleUpdate) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["NextParagraphStyleName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NextParagraphStyleName = &parsedValue
        }

    } else if jsonValue, exists := json["nextParagraphStyleName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NextParagraphStyleName = &parsedValue
        }

    }

    if jsonValue, exists := json["BaseStyleName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.BaseStyleName = &parsedValue
        }

    } else if jsonValue, exists := json["baseStyleName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.BaseStyleName = &parsedValue
        }

    }

    if jsonValue, exists := json["IsQuickStyle"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsQuickStyle = &parsedValue
        }

    } else if jsonValue, exists := json["isQuickStyle"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsQuickStyle = &parsedValue
        }

    }

    if jsonValue, exists := json["Name"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Name = &parsedValue
        }

    } else if jsonValue, exists := json["name"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Name = &parsedValue
        }

    }
}

func (obj *StyleUpdate) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *StyleUpdate) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.NextParagraphStyleName == nil {
        return errors.New("Property NextParagraphStyleName in StyleUpdate is required.")
    }

    if obj.BaseStyleName == nil {
        return errors.New("Property BaseStyleName in StyleUpdate is required.")
    }

    if obj.IsQuickStyle == nil {
        return errors.New("Property IsQuickStyle in StyleUpdate is required.")
    }

    if obj.Name == nil {
        return errors.New("Property Name in StyleUpdate is required.")
    }

    return nil;
}

func (obj *StyleUpdate) GetNextParagraphStyleName() *string {
    return obj.NextParagraphStyleName
}

func (obj *StyleUpdate) SetNextParagraphStyleName(value *string) {
    obj.NextParagraphStyleName = value
}

func (obj *StyleUpdate) GetBaseStyleName() *string {
    return obj.BaseStyleName
}

func (obj *StyleUpdate) SetBaseStyleName(value *string) {
    obj.BaseStyleName = value
}

func (obj *StyleUpdate) GetIsQuickStyle() *bool {
    return obj.IsQuickStyle
}

func (obj *StyleUpdate) SetIsQuickStyle(value *bool) {
    obj.IsQuickStyle = value
}

func (obj *StyleUpdate) GetName() *string {
    return obj.Name
}

func (obj *StyleUpdate) SetName(value *string) {
    obj.Name = value
}

