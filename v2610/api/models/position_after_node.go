/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="position_after_node.go">
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

// Describes the location of the node after specified node.

type IPositionAfterNode interface {
    IsPositionAfterNode() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetNodeId() *string
    SetNodeId(value *string)
    GetType() *string
    SetType(value *string)
}

type PositionAfterNode struct {
    // Describes the location of the node after specified node.
    NodeId *string `json:"NodeId,omitempty"`

    // Describes the location of the node after specified node.
    Type *string `json:"Type,omitempty"`
}

func (PositionAfterNode) IsPositionAfterNode() bool {
    return true
}

func (PositionAfterNode) IsPosition() bool {
    return true
}

func (obj *PositionAfterNode) Initialize() {
    var _Type = "After"
    obj.Type = &_Type


}

func (obj *PositionAfterNode) Deserialize(json map[string]interface{}) {
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

func (obj *PositionAfterNode) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *PositionAfterNode) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.NodeId == nil {
        return errors.New("Property NodeId in PositionAfterNode is required.")
    }
    return nil;
}

func (obj *PositionAfterNode) GetNodeId() *string {
    return obj.NodeId
}

func (obj *PositionAfterNode) SetNodeId(value *string) {
    obj.NodeId = value
}

func (obj *PositionAfterNode) GetType() *string {
    return obj.Type
}

func (obj *PositionAfterNode) SetType(value *string) {
    obj.Type = value
}

