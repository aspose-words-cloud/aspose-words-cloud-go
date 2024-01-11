/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="structured_document_tag_collection.go">
 *   Copyright (c) 2024 Aspose.Words for Cloud
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

// DTO container with a collection of StructuredDocumentTags links.

type IStructuredDocumentTagCollection interface {
    IsStructuredDocumentTagCollection() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetList() []IStructuredDocumentTag
    SetList(value []IStructuredDocumentTag)
}

type StructuredDocumentTagCollection struct {
    // DTO container with a collection of StructuredDocumentTags links.
    Link IWordsApiLink `json:"Link,omitempty"`

    // DTO container with a collection of StructuredDocumentTags links.
    List []IStructuredDocumentTag `json:"List,omitempty"`
}

func (StructuredDocumentTagCollection) IsStructuredDocumentTagCollection() bool {
    return true
}

func (StructuredDocumentTagCollection) IsLinkElement() bool {
    return true
}

func (obj *StructuredDocumentTagCollection) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.List != nil) {
        for _, objElementList := range obj.List {
            objElementList.Initialize()
        }
    }

}

func (obj *StructuredDocumentTagCollection) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["List"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.List = make([]IStructuredDocumentTag, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IStructuredDocumentTag = nil
                    if jsonElementType, found := elementValue["$type"]; found {
                        jsonTypeStr := jsonElementType.(string)
                        if jsonTypeStr == "StructuredDocumentTagUpdate, _" { modelElementInstance = new(StructuredDocumentTagUpdate) }
                    }

                    if modelElementInstance == nil { modelElementInstance = new(StructuredDocumentTag) }
                    modelElementInstance.Deserialize(elementValue)
                    obj.List = append(obj.List, modelElementInstance)
                }

            }
        }

    } else if jsonValue, exists := json["list"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.List = make([]IStructuredDocumentTag, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IStructuredDocumentTag = nil
                    if jsonElementType, found := elementValue["$type"]; found {
                        jsonTypeStr := jsonElementType.(string)
                        if jsonTypeStr == "StructuredDocumentTagUpdate, _" { modelElementInstance = new(StructuredDocumentTagUpdate) }
                    }

                    if modelElementInstance == nil { modelElementInstance = new(StructuredDocumentTag) }
                    modelElementInstance.Deserialize(elementValue)
                    obj.List = append(obj.List, modelElementInstance)
                }

            }
        }

    }
}

func (obj *StructuredDocumentTagCollection) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *StructuredDocumentTagCollection) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Link != nil {
        if err := obj.Link.Validate(); err != nil {
            return err
        }
    }
    if obj.List != nil {
        for _, elementList := range obj.List {
            if elementList != nil {
                if err := elementList.Validate(); err != nil {
                    return err
                }
            }
        }
    }

    return nil;
}

func (obj *StructuredDocumentTagCollection) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *StructuredDocumentTagCollection) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *StructuredDocumentTagCollection) GetList() []IStructuredDocumentTag {
    return obj.List
}

func (obj *StructuredDocumentTagCollection) SetList(value []IStructuredDocumentTag) {
    obj.List = value
}

