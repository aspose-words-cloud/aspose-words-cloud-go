/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="paragraph_link_collection.go">
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

// The collection of paragraph's links.

type IParagraphLinkCollection interface {
    IsParagraphLinkCollection() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetParagraphLinkList() []IParagraphLink
    SetParagraphLinkList(value []IParagraphLink)
}

type ParagraphLinkCollection struct {
    // The collection of paragraph's links.
    Link IWordsApiLink `json:"Link,omitempty"`

    // The collection of paragraph's links.
    ParagraphLinkList []IParagraphLink `json:"ParagraphLinkList,omitempty"`
}

func (ParagraphLinkCollection) IsParagraphLinkCollection() bool {
    return true
}

func (ParagraphLinkCollection) IsLinkElement() bool {
    return true
}

func (obj *ParagraphLinkCollection) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.ParagraphLinkList != nil) {
        for _, objElementParagraphLinkList := range obj.ParagraphLinkList {
            objElementParagraphLinkList.Initialize()
        }
    }

}

func (obj *ParagraphLinkCollection) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["ParagraphLinkList"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.ParagraphLinkList = make([]IParagraphLink, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IParagraphLink = new(ParagraphLink)
                    modelElementInstance.Deserialize(elementValue)
                    obj.ParagraphLinkList = append(obj.ParagraphLinkList, modelElementInstance)
                }

            }
        }

    } else if jsonValue, exists := json["paragraphLinkList"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.ParagraphLinkList = make([]IParagraphLink, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IParagraphLink = new(ParagraphLink)
                    modelElementInstance.Deserialize(elementValue)
                    obj.ParagraphLinkList = append(obj.ParagraphLinkList, modelElementInstance)
                }

            }
        }

    }
}

func (obj *ParagraphLinkCollection) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *ParagraphLinkCollection) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Link == nil {
        return errors.New("Property Link in ParagraphLinkCollection is required.")
    }

    if obj.Link != nil {
        if err := obj.Link.Validate(); err != nil {
            return err
        }
    }

    if obj.ParagraphLinkList == nil {
        return errors.New("Property ParagraphLinkList in ParagraphLinkCollection is required.")
    }

    if obj.ParagraphLinkList != nil {
        for _, elementParagraphLinkList := range obj.ParagraphLinkList {
            if elementParagraphLinkList != nil {
                if err := elementParagraphLinkList.Validate(); err != nil {
                    return err
                }
            }
        }
    }

    return nil;
}

func (obj *ParagraphLinkCollection) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *ParagraphLinkCollection) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *ParagraphLinkCollection) GetParagraphLinkList() []IParagraphLink {
    return obj.ParagraphLinkList
}

func (obj *ParagraphLinkCollection) SetParagraphLinkList(value []IParagraphLink) {
    obj.ParagraphLinkList = value
}

