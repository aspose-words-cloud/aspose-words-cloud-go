/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="field_link.go">
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

// Field link.

type IFieldLink interface {
    IsFieldLink() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetNodeId() *string
    SetNodeId(value *string)
    GetFieldCode() *string
    SetFieldCode(value *string)
}

type FieldLink struct {
    // Field link.
    Link IWordsApiLink `json:"Link,omitempty"`

    // Field link.
    NodeId *string `json:"NodeId,omitempty"`

    // Field link.
    FieldCode *string `json:"FieldCode,omitempty"`
}

func (FieldLink) IsFieldLink() bool {
    return true
}

func (FieldLink) IsNodeLink() bool {
    return true
}

func (FieldLink) IsLinkElement() bool {
    return true
}

func (obj *FieldLink) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }


}

func (obj *FieldLink) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["Link"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IWordsApiLink = new(WordsApiLink)
            modelInstance.Deserialize(parsedValue)
            obj.Link = modelInstance
        }

    } else if jsonValue, exists := json["link"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IWordsApiLink = new(WordsApiLink)
            modelInstance.Deserialize(parsedValue)
            obj.Link = modelInstance
        }

    }

    if jsonValue, exists := json["NodeId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NodeId = &parsedValue
        }

    } else if jsonValue, exists := json["nodeId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NodeId = &parsedValue
        }

    }

    if jsonValue, exists := json["FieldCode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FieldCode = &parsedValue
        }

    } else if jsonValue, exists := json["fieldCode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FieldCode = &parsedValue
        }

    }
}

func (obj *FieldLink) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *FieldLink) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Link != nil {
        if err := obj.Link.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *FieldLink) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *FieldLink) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *FieldLink) GetNodeId() *string {
    return obj.NodeId
}

func (obj *FieldLink) SetNodeId(value *string) {
    obj.NodeId = value
}

func (obj *FieldLink) GetFieldCode() *string {
    return obj.FieldCode
}

func (obj *FieldLink) SetFieldCode(value *string) {
    obj.FieldCode = value
}

