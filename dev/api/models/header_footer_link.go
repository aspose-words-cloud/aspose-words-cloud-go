/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="header_footer_link.go">
 *   Copyright (c) 2025 Aspose.Words for Cloud
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

// HeaderFooter link element.

type IHeaderFooterLink interface {
    IsHeaderFooterLink() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetType() *string
    SetType(value *string)
}

type HeaderFooterLink struct {
    // HeaderFooter link element.
    Link IWordsApiLink `json:"Link,omitempty"`

    // HeaderFooter link element.
    Type *string `json:"Type,omitempty"`
}

func (HeaderFooterLink) IsHeaderFooterLink() bool {
    return true
}

func (HeaderFooterLink) IsLinkElement() bool {
    return true
}

func (obj *HeaderFooterLink) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }


}

func (obj *HeaderFooterLink) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["Type"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Type = &parsedValue
        }

    } else if jsonValue, exists := json["type"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Type = &parsedValue
        }

    }
}

func (obj *HeaderFooterLink) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *HeaderFooterLink) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Type == nil {
        return errors.New("Property Type in HeaderFooterLink is required.")
    }
    if obj.Link != nil {
        if err := obj.Link.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *HeaderFooterLink) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *HeaderFooterLink) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *HeaderFooterLink) GetType() *string {
    return obj.Type
}

func (obj *HeaderFooterLink) SetType(value *string) {
    obj.Type = value
}

