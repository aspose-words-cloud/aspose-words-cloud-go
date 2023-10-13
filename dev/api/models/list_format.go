/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="list_format.go">
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

// DTO container with a paragraph list format element.

type IListFormat interface {
    IsListFormat() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetListLevelNumber() *int32
    SetListLevelNumber(value *int32)
    GetListId() *int32
    SetListId(value *int32)
    GetIsListItem() *bool
    SetIsListItem(value *bool)
}

type ListFormat struct {
    // DTO container with a paragraph list format element.
    Link IWordsApiLink `json:"Link,omitempty"`

    // DTO container with a paragraph list format element.
    ListLevelNumber *int32 `json:"ListLevelNumber,omitempty"`

    // DTO container with a paragraph list format element.
    ListId *int32 `json:"ListId,omitempty"`

    // DTO container with a paragraph list format element.
    IsListItem *bool `json:"IsListItem,omitempty"`
}

func (ListFormat) IsListFormat() bool {
    return true
}

func (ListFormat) IsLinkElement() bool {
    return true
}

func (obj *ListFormat) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }


}

func (obj *ListFormat) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["ListLevelNumber"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ListLevelNumber = new(int32)
            *obj.ListLevelNumber = int32(parsedValue)
        }

    } else if jsonValue, exists := json["listLevelNumber"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ListLevelNumber = new(int32)
            *obj.ListLevelNumber = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["ListId"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ListId = new(int32)
            *obj.ListId = int32(parsedValue)
        }

    } else if jsonValue, exists := json["listId"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ListId = new(int32)
            *obj.ListId = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["IsListItem"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsListItem = &parsedValue
        }

    } else if jsonValue, exists := json["isListItem"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsListItem = &parsedValue
        }

    }
}

func (obj *ListFormat) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *ListFormat) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.ListLevelNumber == nil {
        return errors.New("Property ListLevelNumber in ListFormat is required.")
    }

    if obj.IsListItem == nil {
        return errors.New("Property IsListItem in ListFormat is required.")
    }

    return nil;
}

func (obj *ListFormat) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *ListFormat) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *ListFormat) GetListLevelNumber() *int32 {
    return obj.ListLevelNumber
}

func (obj *ListFormat) SetListLevelNumber(value *int32) {
    obj.ListLevelNumber = value
}

func (obj *ListFormat) GetListId() *int32 {
    return obj.ListId
}

func (obj *ListFormat) SetListId(value *int32) {
    obj.ListId = value
}

func (obj *ListFormat) GetIsListItem() *bool {
    return obj.IsListItem
}

func (obj *ListFormat) SetIsListItem(value *bool) {
    obj.IsListItem = value
}

