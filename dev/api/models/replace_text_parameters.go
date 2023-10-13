/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="replace_text_parameters.go">
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

// Class for document replace text request building.

type IReplaceTextParameters interface {
    IsReplaceTextParameters() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetIsMatchCase() *bool
    SetIsMatchCase(value *bool)
    GetIsMatchWholeWord() *bool
    SetIsMatchWholeWord(value *bool)
    GetIsOldValueRegex() *bool
    SetIsOldValueRegex(value *bool)
    GetNewValue() *string
    SetNewValue(value *string)
    GetOldValue() *string
    SetOldValue(value *string)
}

type ReplaceTextParameters struct {
    // Class for document replace text request building.
    IsMatchCase *bool `json:"IsMatchCase,omitempty"`

    // Class for document replace text request building.
    IsMatchWholeWord *bool `json:"IsMatchWholeWord,omitempty"`

    // Class for document replace text request building.
    IsOldValueRegex *bool `json:"IsOldValueRegex,omitempty"`

    // Class for document replace text request building.
    NewValue *string `json:"NewValue,omitempty"`

    // Class for document replace text request building.
    OldValue *string `json:"OldValue,omitempty"`
}

func (ReplaceTextParameters) IsReplaceTextParameters() bool {
    return true
}


func (obj *ReplaceTextParameters) Initialize() {
}

func (obj *ReplaceTextParameters) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["IsMatchCase"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsMatchCase = &parsedValue
        }

    } else if jsonValue, exists := json["isMatchCase"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsMatchCase = &parsedValue
        }

    }

    if jsonValue, exists := json["IsMatchWholeWord"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsMatchWholeWord = &parsedValue
        }

    } else if jsonValue, exists := json["isMatchWholeWord"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsMatchWholeWord = &parsedValue
        }

    }

    if jsonValue, exists := json["IsOldValueRegex"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsOldValueRegex = &parsedValue
        }

    } else if jsonValue, exists := json["isOldValueRegex"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsOldValueRegex = &parsedValue
        }

    }

    if jsonValue, exists := json["NewValue"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NewValue = &parsedValue
        }

    } else if jsonValue, exists := json["newValue"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NewValue = &parsedValue
        }

    }

    if jsonValue, exists := json["OldValue"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.OldValue = &parsedValue
        }

    } else if jsonValue, exists := json["oldValue"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.OldValue = &parsedValue
        }

    }
}

func (obj *ReplaceTextParameters) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *ReplaceTextParameters) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.IsMatchCase == nil {
        return errors.New("Property IsMatchCase in ReplaceTextParameters is required.")
    }

    if obj.IsMatchWholeWord == nil {
        return errors.New("Property IsMatchWholeWord in ReplaceTextParameters is required.")
    }

    if obj.IsOldValueRegex == nil {
        return errors.New("Property IsOldValueRegex in ReplaceTextParameters is required.")
    }

    if obj.OldValue == nil {
        return errors.New("Property OldValue in ReplaceTextParameters is required.")
    }

    return nil;
}

func (obj *ReplaceTextParameters) GetIsMatchCase() *bool {
    return obj.IsMatchCase
}

func (obj *ReplaceTextParameters) SetIsMatchCase(value *bool) {
    obj.IsMatchCase = value
}

func (obj *ReplaceTextParameters) GetIsMatchWholeWord() *bool {
    return obj.IsMatchWholeWord
}

func (obj *ReplaceTextParameters) SetIsMatchWholeWord(value *bool) {
    obj.IsMatchWholeWord = value
}

func (obj *ReplaceTextParameters) GetIsOldValueRegex() *bool {
    return obj.IsOldValueRegex
}

func (obj *ReplaceTextParameters) SetIsOldValueRegex(value *bool) {
    obj.IsOldValueRegex = value
}

func (obj *ReplaceTextParameters) GetNewValue() *string {
    return obj.NewValue
}

func (obj *ReplaceTextParameters) SetNewValue(value *string) {
    obj.NewValue = value
}

func (obj *ReplaceTextParameters) GetOldValue() *string {
    return obj.OldValue
}

func (obj *ReplaceTextParameters) SetOldValue(value *string) {
    obj.OldValue = value
}

