/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="time_zone_info_data.go">
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

// Class to specify TimeZoneInfo parameters.

type ITimeZoneInfoData interface {
    IsTimeZoneInfoData() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetBaseUtcOffset() *string
    SetBaseUtcOffset(value *string)
    GetDisplayName() *string
    SetDisplayName(value *string)
    GetId() *string
    SetId(value *string)
    GetStandardDisplayName() *string
    SetStandardDisplayName(value *string)
}

type TimeZoneInfoData struct {
    // Class to specify TimeZoneInfo parameters.
    BaseUtcOffset *string `json:"BaseUtcOffset,omitempty"`

    // Class to specify TimeZoneInfo parameters.
    DisplayName *string `json:"DisplayName,omitempty"`

    // Class to specify TimeZoneInfo parameters.
    Id *string `json:"Id,omitempty"`

    // Class to specify TimeZoneInfo parameters.
    StandardDisplayName *string `json:"StandardDisplayName,omitempty"`
}

func (TimeZoneInfoData) IsTimeZoneInfoData() bool {
    return true
}


func (obj *TimeZoneInfoData) Initialize() {
}

func (obj *TimeZoneInfoData) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["BaseUtcOffset"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.BaseUtcOffset = &parsedValue
        }

    } else if jsonValue, exists := json["baseUtcOffset"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.BaseUtcOffset = &parsedValue
        }

    }

    if jsonValue, exists := json["DisplayName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DisplayName = &parsedValue
        }

    } else if jsonValue, exists := json["displayName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DisplayName = &parsedValue
        }

    }

    if jsonValue, exists := json["Id"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Id = &parsedValue
        }

    } else if jsonValue, exists := json["id"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Id = &parsedValue
        }

    }

    if jsonValue, exists := json["StandardDisplayName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.StandardDisplayName = &parsedValue
        }

    } else if jsonValue, exists := json["standardDisplayName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.StandardDisplayName = &parsedValue
        }

    }
}

func (obj *TimeZoneInfoData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *TimeZoneInfoData) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    return nil;
}

func (obj *TimeZoneInfoData) GetBaseUtcOffset() *string {
    return obj.BaseUtcOffset
}

func (obj *TimeZoneInfoData) SetBaseUtcOffset(value *string) {
    obj.BaseUtcOffset = value
}

func (obj *TimeZoneInfoData) GetDisplayName() *string {
    return obj.DisplayName
}

func (obj *TimeZoneInfoData) SetDisplayName(value *string) {
    obj.DisplayName = value
}

func (obj *TimeZoneInfoData) GetId() *string {
    return obj.Id
}

func (obj *TimeZoneInfoData) SetId(value *string) {
    obj.Id = value
}

func (obj *TimeZoneInfoData) GetStandardDisplayName() *string {
    return obj.StandardDisplayName
}

func (obj *TimeZoneInfoData) SetStandardDisplayName(value *string) {
    obj.StandardDisplayName = value
}

