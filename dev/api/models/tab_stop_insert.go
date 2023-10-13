/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="tab_stop_insert.go">
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

// A DTO to Insert / replace a tab stop.

type ITabStopInsert interface {
    IsTabStopInsert() bool
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
}

type TabStopInsert struct {
    // A DTO to Insert / replace a tab stop.
    Alignment *string `json:"Alignment,omitempty"`

    // A DTO to Insert / replace a tab stop.
    Leader *string `json:"Leader,omitempty"`

    // A DTO to Insert / replace a tab stop.
    Position *float64 `json:"Position,omitempty"`
}

func (TabStopInsert) IsTabStopInsert() bool {
    return true
}

func (TabStopInsert) IsTabStopBase() bool {
    return true
}

func (obj *TabStopInsert) Initialize() {
}

func (obj *TabStopInsert) Deserialize(json map[string]interface{}) {
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
}

func (obj *TabStopInsert) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *TabStopInsert) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Alignment == nil {
        return errors.New("Property Alignment in TabStopInsert is required.")
    }

    if obj.Leader == nil {
        return errors.New("Property Leader in TabStopInsert is required.")
    }

    if obj.Position == nil {
        return errors.New("Property Position in TabStopInsert is required.")
    }

    return nil;
}

func (obj *TabStopInsert) GetAlignment() *string {
    return obj.Alignment
}

func (obj *TabStopInsert) SetAlignment(value *string) {
    obj.Alignment = value
}

func (obj *TabStopInsert) GetLeader() *string {
    return obj.Leader
}

func (obj *TabStopInsert) SetLeader(value *string) {
    obj.Leader = value
}

func (obj *TabStopInsert) GetPosition() *float64 {
    return obj.Position
}

func (obj *TabStopInsert) SetPosition(value *float64) {
    obj.Position = value
}

