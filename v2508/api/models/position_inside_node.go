/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="position_inside_node.go">
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

// DTO container with a new position in the document tree.

type IPositionInsideNode interface {
    IsPositionInsideNode() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetNodeId() *string
    SetNodeId(value *string)
    GetType() *string
    SetType(value *string)
    GetOffset() *int32
    SetOffset(value *int32)
}

type PositionInsideNode struct {
    // DTO container with a new position in the document tree.
    NodeId *string `json:"NodeId,omitempty"`

    // DTO container with a new position in the document tree.
    Type *string `json:"Type,omitempty"`

    // DTO container with a new position in the document tree.
    Offset *int32 `json:"Offset,omitempty"`
}

func (PositionInsideNode) IsPositionInsideNode() bool {
    return true
}

func (PositionInsideNode) IsPosition() bool {
    return true
}

func (obj *PositionInsideNode) Initialize() {
    var _Type = "Inside"
    obj.Type = &_Type


}

func (obj *PositionInsideNode) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["NodeId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NodeId = &parsedValue
        }

    } else if jsonValue, exists := json["nodeId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NodeId = &parsedValue
        }

    }

    if jsonValue, exists := json["Offset"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Offset = new(int32)
            *obj.Offset = int32(parsedValue)
        }

    } else if jsonValue, exists := json["offset"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Offset = new(int32)
            *obj.Offset = int32(parsedValue)
        }

    }
}

func (obj *PositionInsideNode) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *PositionInsideNode) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.NodeId == nil {
        return errors.New("Property NodeId in PositionInsideNode is required.")
    }
    return nil;
}

func (obj *PositionInsideNode) GetNodeId() *string {
    return obj.NodeId
}

func (obj *PositionInsideNode) SetNodeId(value *string) {
    obj.NodeId = value
}

func (obj *PositionInsideNode) GetType() *string {
    return obj.Type
}

func (obj *PositionInsideNode) SetType(value *string) {
    obj.Type = value
}

func (obj *PositionInsideNode) GetOffset() *int32 {
    return obj.Offset
}

func (obj *PositionInsideNode) SetOffset(value *int32) {
    obj.Offset = value
}

