/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="office_math_link.go">
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

// OfficeMath object link element.

type IOfficeMathLink interface {
    IsOfficeMathLink() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetNodeId() *string
    SetNodeId(value *string)
}

type OfficeMathLink struct {
    // OfficeMath object link element.
    Link IWordsApiLink `json:"Link,omitempty"`

    // OfficeMath object link element.
    NodeId *string `json:"NodeId,omitempty"`
}

func (OfficeMathLink) IsOfficeMathLink() bool {
    return true
}

func (OfficeMathLink) IsNodeLink() bool {
    return true
}

func (OfficeMathLink) IsLinkElement() bool {
    return true
}

func (obj *OfficeMathLink) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }


}

func (obj *OfficeMathLink) Deserialize(json map[string]interface{}) {
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
}

func (obj *OfficeMathLink) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *OfficeMathLink) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Link == nil {
        return errors.New("Property Link in OfficeMathLink is required.")
    }

    if obj.Link != nil {
        if err := obj.Link.Validate(); err != nil {
            return err
        }
    }

    if obj.NodeId == nil {
        return errors.New("Property NodeId in OfficeMathLink is required.")
    }

    return nil;
}

func (obj *OfficeMathLink) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *OfficeMathLink) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *OfficeMathLink) GetNodeId() *string {
    return obj.NodeId
}

func (obj *OfficeMathLink) SetNodeId(value *string) {
    obj.NodeId = value
}

