/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="field.go">
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

// DTO container with a field.

type IField interface {
    IsField() bool
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
    GetLocaleId() *string
    SetLocaleId(value *string)
    GetResult() *string
    SetResult(value *string)
}

type Field struct {
    // DTO container with a field.
    Link IWordsApiLink `json:"Link,omitempty"`

    // DTO container with a field.
    NodeId *string `json:"NodeId,omitempty"`

    // DTO container with a field.
    FieldCode *string `json:"FieldCode,omitempty"`

    // DTO container with a field.
    LocaleId *string `json:"LocaleId,omitempty"`

    // DTO container with a field.
    Result *string `json:"Result,omitempty"`
}

func (Field) IsField() bool {
    return true
}

func (Field) IsFieldLink() bool {
    return true
}

func (Field) IsNodeLink() bool {
    return true
}

func (Field) IsLinkElement() bool {
    return true
}

func (obj *Field) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }


}

func (obj *Field) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["LocaleId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.LocaleId = &parsedValue
        }

    } else if jsonValue, exists := json["localeId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.LocaleId = &parsedValue
        }

    }

    if jsonValue, exists := json["Result"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Result = &parsedValue
        }

    } else if jsonValue, exists := json["result"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Result = &parsedValue
        }

    }
}

func (obj *Field) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *Field) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    return nil;
}

func (obj *Field) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *Field) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *Field) GetNodeId() *string {
    return obj.NodeId
}

func (obj *Field) SetNodeId(value *string) {
    obj.NodeId = value
}

func (obj *Field) GetFieldCode() *string {
    return obj.FieldCode
}

func (obj *Field) SetFieldCode(value *string) {
    obj.FieldCode = value
}

func (obj *Field) GetLocaleId() *string {
    return obj.LocaleId
}

func (obj *Field) SetLocaleId(value *string) {
    obj.LocaleId = value
}

func (obj *Field) GetResult() *string {
    return obj.Result
}

func (obj *Field) SetResult(value *string) {
    obj.Result = value
}

