/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="tab_stop.go">
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

// DTO container with paragraph format tab stop.

type ITabStop interface {
    IsTabStop() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetAlignment() *string
    SetAlignment(value *string)
    GetLeader() *string
    SetLeader(value *string)
    GetPosition() *float64
    SetPosition(value *float64)
    GetIsClear() *bool
    SetIsClear(value *bool)
}

type TabStop struct {
    // DTO container with paragraph format tab stop.
    Alignment *string `json:"Alignment,omitempty"`

    // DTO container with paragraph format tab stop.
    Leader *string `json:"Leader,omitempty"`

    // DTO container with paragraph format tab stop.
    Position *float64 `json:"Position,omitempty"`

    // DTO container with paragraph format tab stop.
    IsClear *bool `json:"IsClear,omitempty"`
}

func (TabStop) IsTabStop() bool {
    return true
}

func (TabStop) IsTabStopBase() bool {
    return true
}

func (obj *TabStop) Initialize() {
}

func (obj *TabStop) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["Alignment"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Alignment = &parsedValue
        }

    } else if jsonValue, exists := json["alignment"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Alignment = &parsedValue
        }

    }

    if jsonValue, exists := json["Leader"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Leader = &parsedValue
        }

    } else if jsonValue, exists := json["leader"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Leader = &parsedValue
        }

    }

    if jsonValue, exists := json["Position"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Position = &parsedValue
        }

    } else if jsonValue, exists := json["position"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Position = &parsedValue
        }

    }

    if jsonValue, exists := json["IsClear"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsClear = &parsedValue
        }

    } else if jsonValue, exists := json["isClear"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsClear = &parsedValue
        }

    }
}

func (obj *TabStop) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *TabStop) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Alignment == nil {
        return errors.New("Property Alignment in TabStop is required.")
    }
    if obj.Leader == nil {
        return errors.New("Property Leader in TabStop is required.")
    }
    if obj.Position == nil {
        return errors.New("Property Position in TabStop is required.")
    }
    if obj.IsClear == nil {
        return errors.New("Property IsClear in TabStop is required.")
    }
    return nil;
}

func (obj *TabStop) GetAlignment() *string {
    return obj.Alignment
}

func (obj *TabStop) SetAlignment(value *string) {
    obj.Alignment = value
}

func (obj *TabStop) GetLeader() *string {
    return obj.Leader
}

func (obj *TabStop) SetLeader(value *string) {
    obj.Leader = value
}

func (obj *TabStop) GetPosition() *float64 {
    return obj.Position
}

func (obj *TabStop) SetPosition(value *float64) {
    obj.Position = value
}

func (obj *TabStop) GetIsClear() *bool {
    return obj.IsClear
}

func (obj *TabStop) SetIsClear(value *bool) {
    obj.IsClear = value
}

