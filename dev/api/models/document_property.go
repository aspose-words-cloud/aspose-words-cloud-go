/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="document_property.go">
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

// Words document property DTO.

type IDocumentProperty interface {
    IsDocumentProperty() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetName() *string
    SetName(value *string)
    GetValue() *string
    SetValue(value *string)
    GetBuiltIn() *bool
    SetBuiltIn(value *bool)
}

type DocumentProperty struct {
    // Words document property DTO.
    Link IWordsApiLink `json:"Link,omitempty"`

    // Words document property DTO.
    Name *string `json:"Name,omitempty"`

    // Words document property DTO.
    Value *string `json:"Value,omitempty"`

    // Words document property DTO.
    BuiltIn *bool `json:"BuiltIn,omitempty"`
}

func (DocumentProperty) IsDocumentProperty() bool {
    return true
}

func (DocumentProperty) IsLinkElement() bool {
    return true
}

func (obj *DocumentProperty) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }


}

func (obj *DocumentProperty) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["Name"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Name = &parsedValue
        }

    } else if jsonValue, exists := json["name"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Name = &parsedValue
        }

    }

    if jsonValue, exists := json["Value"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Value = &parsedValue
        }

    } else if jsonValue, exists := json["value"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Value = &parsedValue
        }

    }

    if jsonValue, exists := json["BuiltIn"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.BuiltIn = &parsedValue
        }

    } else if jsonValue, exists := json["builtIn"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.BuiltIn = &parsedValue
        }

    }
}

func (obj *DocumentProperty) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *DocumentProperty) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.BuiltIn == nil {
        return errors.New("Property BuiltIn in DocumentProperty is required.")
    }
    if obj.Link != nil {
        if err := obj.Link.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *DocumentProperty) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *DocumentProperty) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *DocumentProperty) GetName() *string {
    return obj.Name
}

func (obj *DocumentProperty) SetName(value *string) {
    obj.Name = value
}

func (obj *DocumentProperty) GetValue() *string {
    return obj.Value
}

func (obj *DocumentProperty) SetValue(value *string) {
    obj.Value = value
}

func (obj *DocumentProperty) GetBuiltIn() *bool {
    return obj.BuiltIn
}

func (obj *DocumentProperty) SetBuiltIn(value *bool) {
    obj.BuiltIn = value
}

