/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="section_link_collection.go">
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

// The collection of section's links.

type ISectionLinkCollection interface {
    IsSectionLinkCollection() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetSectionLinkList() []ISectionLink
    SetSectionLinkList(value []ISectionLink)
}

type SectionLinkCollection struct {
    // The collection of section's links.
    Link IWordsApiLink `json:"Link,omitempty"`

    // The collection of section's links.
    SectionLinkList []ISectionLink `json:"SectionLinkList,omitempty"`
}

func (SectionLinkCollection) IsSectionLinkCollection() bool {
    return true
}

func (SectionLinkCollection) IsLinkElement() bool {
    return true
}

func (obj *SectionLinkCollection) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.SectionLinkList != nil) {
        for _, objElementSectionLinkList := range obj.SectionLinkList {
            objElementSectionLinkList.Initialize()
        }
    }

}

func (obj *SectionLinkCollection) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["SectionLinkList"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.SectionLinkList = make([]ISectionLink, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance ISectionLink = new(SectionLink)
                    modelElementInstance.Deserialize(elementValue)
                    obj.SectionLinkList = append(obj.SectionLinkList, modelElementInstance)
                }

            }
        }

    } else if jsonValue, exists := json["sectionLinkList"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.SectionLinkList = make([]ISectionLink, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance ISectionLink = new(SectionLink)
                    modelElementInstance.Deserialize(elementValue)
                    obj.SectionLinkList = append(obj.SectionLinkList, modelElementInstance)
                }

            }
        }

    }
}

func (obj *SectionLinkCollection) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *SectionLinkCollection) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Link != nil {
        if err := obj.Link.Validate(); err != nil {
            return err
        }
    }
    if obj.SectionLinkList != nil {
        for _, elementSectionLinkList := range obj.SectionLinkList {
            if elementSectionLinkList != nil {
                if err := elementSectionLinkList.Validate(); err != nil {
                    return err
                }
            }
        }
    }

    return nil;
}

func (obj *SectionLinkCollection) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *SectionLinkCollection) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *SectionLinkCollection) GetSectionLinkList() []ISectionLink {
    return obj.SectionLinkList
}

func (obj *SectionLinkCollection) SetSectionLinkList(value []ISectionLink) {
    obj.SectionLinkList = value
}

