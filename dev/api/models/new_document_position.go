/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="new_document_position.go">
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

// DTO container with a new position in the document tree.

type INewDocumentPosition interface {
    IsNewDocumentPosition() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetNodeId() *string
    SetNodeId(value *string)
    GetOffset() *int32
    SetOffset(value *int32)
}

type NewDocumentPosition struct {
    // DTO container with a new position in the document tree.
    NodeId *string `json:"NodeId,omitempty"`

    // DTO container with a new position in the document tree.
    Offset *int32 `json:"Offset,omitempty"`
}

func (NewDocumentPosition) IsNewDocumentPosition() bool {
    return true
}


func (obj *NewDocumentPosition) Initialize() {
}

func (obj *NewDocumentPosition) Deserialize(json map[string]interface{}) {
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

func (obj *NewDocumentPosition) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *NewDocumentPosition) GetNodeId() *string {
    return obj.NodeId
}

func (obj *NewDocumentPosition) SetNodeId(value *string) {
    obj.NodeId = value
}

func (obj *NewDocumentPosition) GetOffset() *int32 {
    return obj.Offset
}

func (obj *NewDocumentPosition) SetOffset(value *int32) {
    obj.Offset = value
}

