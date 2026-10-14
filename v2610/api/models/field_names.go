/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="field_names.go">
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

// Represents a collection of merge fields within a document.

type IFieldNames interface {
    IsFieldNames() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetNames() []string
    SetNames(value []string)
}

type FieldNames struct {
    // Represents a collection of merge fields within a document.
    Link IWordsApiLink `json:"Link,omitempty"`

    // Represents a collection of merge fields within a document.
    Names []string `json:"Names,omitempty"`
}

func (FieldNames) IsFieldNames() bool {
    return true
}

func (FieldNames) IsLinkElement() bool {
    return true
}

func (obj *FieldNames) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }


}

func (obj *FieldNames) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["Names"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.Names = make([]string, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(string); valid {
                    obj.Names = append(obj.Names, elementValue)
                }

            }
        }

    } else if jsonValue, exists := json["names"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.Names = make([]string, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(string); valid {
                    obj.Names = append(obj.Names, elementValue)
                }

            }
        }

    }
}

func (obj *FieldNames) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *FieldNames) Validate() error {
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

func (obj *FieldNames) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *FieldNames) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *FieldNames) GetNames() []string {
    return obj.Names
}

func (obj *FieldNames) SetNames(value []string) {
    obj.Names = value
}

