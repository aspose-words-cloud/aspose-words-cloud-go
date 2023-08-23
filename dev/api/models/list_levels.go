/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="list_levels.go">
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

// DTO container with a single document list.

type IListLevels interface {
    IsListLevels() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetListLevel() []IListLevel
    SetListLevel(value []IListLevel)
}

type ListLevels struct {
    // DTO container with a single document list.
    Link IWordsApiLink

    // DTO container with a single document list.
    ListLevel []IListLevel
}

func (ListLevels) IsListLevels() bool {
    return true
}

func (ListLevels) IsLinkElement() bool {
    return true
}

func (obj *ListLevels) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.ListLevel != nil) {
        for _, objElementListLevel := range obj.ListLevel {
            objElementListLevel.Initialize()
        }
    }

}

func (obj *ListLevels) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["ListLevel"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IListLevel = new(ListLevel)
                    modelElementInstance.Deserialize(elementValue)
                    obj.ListLevel = append(obj.ListLevel, modelElementInstance)
                }

            }
        }

    } else if jsonValue, exists := json["listLevel"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IListLevel = new(ListLevel)
                    modelElementInstance.Deserialize(elementValue)
                    obj.ListLevel = append(obj.ListLevel, modelElementInstance)
                }

            }
        }

    }
}

func (obj *ListLevels) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *ListLevels) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *ListLevels) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *ListLevels) GetListLevel() []IListLevel {
    return obj.ListLevel
}

func (obj *ListLevels) SetListLevel(value []IListLevel) {
    obj.ListLevel = value
}

