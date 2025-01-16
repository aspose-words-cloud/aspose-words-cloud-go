/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="position_before_node.go">
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

// Describes the location of the node before specified node.

type IPositionBeforeNode interface {
    IsPositionBeforeNode() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetNodeId() *string
    SetNodeId(value *string)
    GetType() *string
    SetType(value *string)
}

type PositionBeforeNode struct {
    // Describes the location of the node before specified node.
    NodeId *string `json:"NodeId,omitempty"`

    // Describes the location of the node before specified node.
    Type *string `json:"Type,omitempty"`
}

func (PositionBeforeNode) IsPositionBeforeNode() bool {
    return true
}

func (PositionBeforeNode) IsPosition() bool {
    return true
}

func (obj *PositionBeforeNode) Initialize() {
    var _Type = "Before"
    obj.Type = &_Type


}

func (obj *PositionBeforeNode) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["NodeId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NodeId = &parsedValue
        }

    } else if jsonValue, exists := json["nodeId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NodeId = &parsedValue
        }

    }
}

func (obj *PositionBeforeNode) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *PositionBeforeNode) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.NodeId == nil {
        return errors.New("Property NodeId in PositionBeforeNode is required.")
    }
    return nil;
}

func (obj *PositionBeforeNode) GetNodeId() *string {
    return obj.NodeId
}

func (obj *PositionBeforeNode) SetNodeId(value *string) {
    obj.NodeId = value
}

func (obj *PositionBeforeNode) GetType() *string {
    return obj.Type
}

func (obj *PositionBeforeNode) SetType(value *string) {
    obj.Type = value
}

