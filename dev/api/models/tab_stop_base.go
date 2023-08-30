/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="tab_stop_base.go">
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

// Base class for paragraph format tab stop DTO.

type ITabStopBase interface {
    IsTabStopBase() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetAlignment() *string
    SetAlignment(value *string)
    GetLeader() *string
    SetLeader(value *string)
    GetPosition() *float64
    SetPosition(value *float64)
}

type TabStopBase struct {
    // Base class for paragraph format tab stop DTO.
    Alignment *string `json:"Alignment,omitempty"`

    // Base class for paragraph format tab stop DTO.
    Leader *string `json:"Leader,omitempty"`

    // Base class for paragraph format tab stop DTO.
    Position *float64 `json:"Position,omitempty"`
}

func (TabStopBase) IsTabStopBase() bool {
    return true
}


func (obj *TabStopBase) Initialize() {
}

func (obj *TabStopBase) Deserialize(json map[string]interface{}) {
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

func (obj *TabStopBase) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *TabStopBase) GetAlignment() *string {
    return obj.Alignment
}

func (obj *TabStopBase) SetAlignment(value *string) {
    obj.Alignment = value
}

func (obj *TabStopBase) GetLeader() *string {
    return obj.Leader
}

func (obj *TabStopBase) SetLeader(value *string) {
    obj.Leader = value
}

func (obj *TabStopBase) GetPosition() *float64 {
    return obj.Position
}

func (obj *TabStopBase) SetPosition(value *float64) {
    obj.Position = value
}

