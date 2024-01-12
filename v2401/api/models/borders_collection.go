/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="borders_collection.go">
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

// The collection of borders.

type IBordersCollection interface {
    IsBordersCollection() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetList() []IBorder
    SetList(value []IBorder)
}

type BordersCollection struct {
    // The collection of borders.
    Link IWordsApiLink `json:"Link,omitempty"`

    // The collection of borders.
    List []IBorder `json:"List,omitempty"`
}

func (BordersCollection) IsBordersCollection() bool {
    return true
}

func (BordersCollection) IsLinkElement() bool {
    return true
}

func (obj *BordersCollection) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.List != nil) {
        for _, objElementList := range obj.List {
            objElementList.Initialize()
        }
    }

}

func (obj *BordersCollection) Deserialize(json map[string]interface{}) {
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
            obj.List = make([]IBorder, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IBorder = new(Border)
                    modelElementInstance.Deserialize(elementValue)
                    obj.List = append(obj.List, modelElementInstance)
                }

            }
        }

    } else if jsonValue, exists := json["list"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.List = make([]IBorder, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IBorder = new(Border)
                    modelElementInstance.Deserialize(elementValue)
                    obj.List = append(obj.List, modelElementInstance)
                }

            }
        }

    }
}

func (obj *BordersCollection) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *BordersCollection) Validate() error {
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

func (obj *BordersCollection) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *BordersCollection) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *BordersCollection) GetList() []IBorder {
    return obj.List
}

func (obj *BordersCollection) SetList(value []IBorder) {
    obj.List = value
}

