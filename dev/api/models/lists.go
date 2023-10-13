/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="lists.go">
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

// DTO container with an array of document lists.

type ILists interface {
    IsLists() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetListInfo() []IListInfo
    SetListInfo(value []IListInfo)
}

type Lists struct {
    // DTO container with an array of document lists.
    Link IWordsApiLink `json:"Link,omitempty"`

    // DTO container with an array of document lists.
    ListInfo []IListInfo `json:"ListInfo,omitempty"`
}

func (Lists) IsLists() bool {
    return true
}

func (Lists) IsLinkElement() bool {
    return true
}

func (obj *Lists) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.ListInfo != nil) {
        for _, objElementListInfo := range obj.ListInfo {
            objElementListInfo.Initialize()
        }
    }

}

func (obj *Lists) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["ListInfo"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.ListInfo = make([]IListInfo, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IListInfo = new(ListInfo)
                    modelElementInstance.Deserialize(elementValue)
                    obj.ListInfo = append(obj.ListInfo, modelElementInstance)
                }

            }
        }

    } else if jsonValue, exists := json["listInfo"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.ListInfo = make([]IListInfo, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IListInfo = new(ListInfo)
                    modelElementInstance.Deserialize(elementValue)
                    obj.ListInfo = append(obj.ListInfo, modelElementInstance)
                }

            }
        }

    }
}

func (obj *Lists) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *Lists) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    return nil;
}

func (obj *Lists) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *Lists) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *Lists) GetListInfo() []IListInfo {
    return obj.ListInfo
}

func (obj *Lists) SetListInfo(value []IListInfo) {
    obj.ListInfo = value
}

