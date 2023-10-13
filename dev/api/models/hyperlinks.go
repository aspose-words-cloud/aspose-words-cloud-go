/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="hyperlinks.go">
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

// Collection of Hyperlink.

type IHyperlinks interface {
    IsHyperlinks() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetHyperlinkList() []IHyperlink
    SetHyperlinkList(value []IHyperlink)
}

type Hyperlinks struct {
    // Collection of Hyperlink.
    Link IWordsApiLink `json:"Link,omitempty"`

    // Collection of Hyperlink.
    HyperlinkList []IHyperlink `json:"HyperlinkList,omitempty"`
}

func (Hyperlinks) IsHyperlinks() bool {
    return true
}

func (Hyperlinks) IsLinkElement() bool {
    return true
}

func (obj *Hyperlinks) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.HyperlinkList != nil) {
        for _, objElementHyperlinkList := range obj.HyperlinkList {
            objElementHyperlinkList.Initialize()
        }
    }

}

func (obj *Hyperlinks) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["HyperlinkList"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.HyperlinkList = make([]IHyperlink, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IHyperlink = new(Hyperlink)
                    modelElementInstance.Deserialize(elementValue)
                    obj.HyperlinkList = append(obj.HyperlinkList, modelElementInstance)
                }

            }
        }

    } else if jsonValue, exists := json["hyperlinkList"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.HyperlinkList = make([]IHyperlink, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IHyperlink = new(Hyperlink)
                    modelElementInstance.Deserialize(elementValue)
                    obj.HyperlinkList = append(obj.HyperlinkList, modelElementInstance)
                }

            }
        }

    }
}

func (obj *Hyperlinks) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *Hyperlinks) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    return nil;
}

func (obj *Hyperlinks) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *Hyperlinks) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *Hyperlinks) GetHyperlinkList() []IHyperlink {
    return obj.HyperlinkList
}

func (obj *Hyperlinks) SetHyperlinkList(value []IHyperlink) {
    obj.HyperlinkList = value
}

