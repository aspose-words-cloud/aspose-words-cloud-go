/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="advanced_compare_options.go">
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

// Allows to set advanced compare options.

type IAdvancedCompareOptions interface {
    IsAdvancedCompareOptions() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetCompareListDefinitions() *bool
    SetCompareListDefinitions(value *bool)
    GetIgnoreDmlUniqueId() *bool
    SetIgnoreDmlUniqueId(value *bool)
    GetIgnoreStoreItemId() *bool
    SetIgnoreStoreItemId(value *bool)
}

type AdvancedCompareOptions struct {
    // Allows to set advanced compare options.
    CompareListDefinitions *bool `json:"CompareListDefinitions,omitempty"`

    // Allows to set advanced compare options.
    IgnoreDmlUniqueId *bool `json:"IgnoreDmlUniqueId,omitempty"`

    // Allows to set advanced compare options.
    IgnoreStoreItemId *bool `json:"IgnoreStoreItemId,omitempty"`
}

func (AdvancedCompareOptions) IsAdvancedCompareOptions() bool {
    return true
}


func (obj *AdvancedCompareOptions) Initialize() {
}

func (obj *AdvancedCompareOptions) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["CompareListDefinitions"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.CompareListDefinitions = &parsedValue
        }

    } else if jsonValue, exists := json["compareListDefinitions"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.CompareListDefinitions = &parsedValue
        }

    }

    if jsonValue, exists := json["IgnoreDmlUniqueId"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IgnoreDmlUniqueId = &parsedValue
        }

    } else if jsonValue, exists := json["ignoreDmlUniqueId"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IgnoreDmlUniqueId = &parsedValue
        }

    }

    if jsonValue, exists := json["IgnoreStoreItemId"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IgnoreStoreItemId = &parsedValue
        }

    } else if jsonValue, exists := json["ignoreStoreItemId"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IgnoreStoreItemId = &parsedValue
        }

    }
}

func (obj *AdvancedCompareOptions) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *AdvancedCompareOptions) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    return nil;
}

func (obj *AdvancedCompareOptions) GetCompareListDefinitions() *bool {
    return obj.CompareListDefinitions
}

func (obj *AdvancedCompareOptions) SetCompareListDefinitions(value *bool) {
    obj.CompareListDefinitions = value
}

func (obj *AdvancedCompareOptions) GetIgnoreDmlUniqueId() *bool {
    return obj.IgnoreDmlUniqueId
}

func (obj *AdvancedCompareOptions) SetIgnoreDmlUniqueId(value *bool) {
    obj.IgnoreDmlUniqueId = value
}

func (obj *AdvancedCompareOptions) GetIgnoreStoreItemId() *bool {
    return obj.IgnoreStoreItemId
}

func (obj *AdvancedCompareOptions) SetIgnoreStoreItemId(value *bool) {
    obj.IgnoreStoreItemId = value
}

