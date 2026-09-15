/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="list_info.go">
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

// DTO container with a single document list.

type IListInfo interface {
    IsListInfo() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetListId() *int32
    SetListId(value *int32)
    GetIsMultiLevel() *bool
    SetIsMultiLevel(value *bool)
    GetIsRestartAtEachSection() *bool
    SetIsRestartAtEachSection(value *bool)
    GetIsListStyleDefinition() *bool
    SetIsListStyleDefinition(value *bool)
    GetIsListStyleReference() *bool
    SetIsListStyleReference(value *bool)
    GetStyle() IStyle
    SetStyle(value IStyle)
    GetListLevels() IListLevels
    SetListLevels(value IListLevels)
}

type ListInfo struct {
    // DTO container with a single document list.
    Link IWordsApiLink `json:"Link,omitempty"`

    // DTO container with a single document list.
    ListId *int32 `json:"ListId,omitempty"`

    // DTO container with a single document list.
    IsMultiLevel *bool `json:"IsMultiLevel,omitempty"`

    // DTO container with a single document list.
    IsRestartAtEachSection *bool `json:"IsRestartAtEachSection,omitempty"`

    // DTO container with a single document list.
    IsListStyleDefinition *bool `json:"IsListStyleDefinition,omitempty"`

    // DTO container with a single document list.
    IsListStyleReference *bool `json:"IsListStyleReference,omitempty"`

    // DTO container with a single document list.
    Style IStyle `json:"Style,omitempty"`

    // DTO container with a single document list.
    ListLevels IListLevels `json:"ListLevels,omitempty"`
}

func (ListInfo) IsListInfo() bool {
    return true
}

func (ListInfo) IsLinkElement() bool {
    return true
}

func (obj *ListInfo) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.Style != nil) {
        obj.Style.Initialize()
    }

    if (obj.ListLevels != nil) {
        obj.ListLevels.Initialize()
    }


}

func (obj *ListInfo) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["IsMultiLevel"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsMultiLevel = &parsedValue
        }

    } else if jsonValue, exists := json["isMultiLevel"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsMultiLevel = &parsedValue
        }

    }

    if jsonValue, exists := json["IsRestartAtEachSection"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsRestartAtEachSection = &parsedValue
        }

    } else if jsonValue, exists := json["isRestartAtEachSection"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsRestartAtEachSection = &parsedValue
        }

    }

    if jsonValue, exists := json["IsListStyleDefinition"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsListStyleDefinition = &parsedValue
        }

    } else if jsonValue, exists := json["isListStyleDefinition"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsListStyleDefinition = &parsedValue
        }

    }

    if jsonValue, exists := json["IsListStyleReference"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsListStyleReference = &parsedValue
        }

    } else if jsonValue, exists := json["isListStyleReference"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsListStyleReference = &parsedValue
        }

    }

    if jsonValue, exists := json["Style"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IStyle = new(Style)
            modelInstance.Deserialize(parsedValue)
            obj.Style = modelInstance
        }

    } else if jsonValue, exists := json["style"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IStyle = new(Style)
            modelInstance.Deserialize(parsedValue)
            obj.Style = modelInstance
        }

    }

    if jsonValue, exists := json["ListLevels"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IListLevels = new(ListLevels)
            modelInstance.Deserialize(parsedValue)
            obj.ListLevels = modelInstance
        }

    } else if jsonValue, exists := json["listLevels"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IListLevels = new(ListLevels)
            modelInstance.Deserialize(parsedValue)
            obj.ListLevels = modelInstance
        }

    }
}

func (obj *ListInfo) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *ListInfo) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.ListId == nil {
        return errors.New("Property ListId in ListInfo is required.")
    }
    if obj.IsMultiLevel == nil {
        return errors.New("Property IsMultiLevel in ListInfo is required.")
    }
    if obj.IsRestartAtEachSection == nil {
        return errors.New("Property IsRestartAtEachSection in ListInfo is required.")
    }
    if obj.IsListStyleDefinition == nil {
        return errors.New("Property IsListStyleDefinition in ListInfo is required.")
    }
    if obj.IsListStyleReference == nil {
        return errors.New("Property IsListStyleReference in ListInfo is required.")
    }
    if obj.Link != nil {
        if err := obj.Link.Validate(); err != nil {
            return err
        }
    }
    if obj.Style != nil {
        if err := obj.Style.Validate(); err != nil {
            return err
        }
    }
    if obj.ListLevels != nil {
        if err := obj.ListLevels.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *ListInfo) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *ListInfo) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *ListInfo) GetListId() *int32 {
    return obj.ListId
}

func (obj *ListInfo) SetListId(value *int32) {
    obj.ListId = value
}

func (obj *ListInfo) GetIsMultiLevel() *bool {
    return obj.IsMultiLevel
}

func (obj *ListInfo) SetIsMultiLevel(value *bool) {
    obj.IsMultiLevel = value
}

func (obj *ListInfo) GetIsRestartAtEachSection() *bool {
    return obj.IsRestartAtEachSection
}

func (obj *ListInfo) SetIsRestartAtEachSection(value *bool) {
    obj.IsRestartAtEachSection = value
}

func (obj *ListInfo) GetIsListStyleDefinition() *bool {
    return obj.IsListStyleDefinition
}

func (obj *ListInfo) SetIsListStyleDefinition(value *bool) {
    obj.IsListStyleDefinition = value
}

func (obj *ListInfo) GetIsListStyleReference() *bool {
    return obj.IsListStyleReference
}

func (obj *ListInfo) SetIsListStyleReference(value *bool) {
    obj.IsListStyleReference = value
}

func (obj *ListInfo) GetStyle() IStyle {
    return obj.Style
}

func (obj *ListInfo) SetStyle(value IStyle) {
    obj.Style = value
}

func (obj *ListInfo) GetListLevels() IListLevels {
    return obj.ListLevels
}

func (obj *ListInfo) SetListLevels(value IListLevels) {
    obj.ListLevels = value
}

