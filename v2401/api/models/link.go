/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="link.go">
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

// Provides information for the object link.
// This is supposed to be an atom:link, therefore it should have all attributes specified here http://tools.ietf.org/html/rfc4287#section-4.2.7.

type ILink interface {
    IsLink() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetHref() *string
    SetHref(value *string)
    GetRel() *string
    SetRel(value *string)
    GetTitle() *string
    SetTitle(value *string)
    GetType() *string
    SetType(value *string)
}

type Link struct {
    // Provides information for the object link.
    // This is supposed to be an atom:link, therefore it should have all attributes specified here http://tools.ietf.org/html/rfc4287#section-4.2.7.
    Href *string `json:"Href,omitempty"`

    // Provides information for the object link.
    // This is supposed to be an atom:link, therefore it should have all attributes specified here http://tools.ietf.org/html/rfc4287#section-4.2.7.
    Rel *string `json:"Rel,omitempty"`

    // Provides information for the object link.
    // This is supposed to be an atom:link, therefore it should have all attributes specified here http://tools.ietf.org/html/rfc4287#section-4.2.7.
    Title *string `json:"Title,omitempty"`

    // Provides information for the object link.
    // This is supposed to be an atom:link, therefore it should have all attributes specified here http://tools.ietf.org/html/rfc4287#section-4.2.7.
    Type *string `json:"Type,omitempty"`
}

func (Link) IsLink() bool {
    return true
}


func (obj *Link) Initialize() {
}

func (obj *Link) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["Href"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Href = &parsedValue
        }

    } else if jsonValue, exists := json["href"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Href = &parsedValue
        }

    }

    if jsonValue, exists := json["Rel"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Rel = &parsedValue
        }

    } else if jsonValue, exists := json["rel"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Rel = &parsedValue
        }

    }

    if jsonValue, exists := json["Title"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Title = &parsedValue
        }

    } else if jsonValue, exists := json["title"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Title = &parsedValue
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

func (obj *Link) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *Link) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    return nil;
}

func (obj *Link) GetHref() *string {
    return obj.Href
}

func (obj *Link) SetHref(value *string) {
    obj.Href = value
}

func (obj *Link) GetRel() *string {
    return obj.Rel
}

func (obj *Link) SetRel(value *string) {
    obj.Rel = value
}

func (obj *Link) GetTitle() *string {
    return obj.Title
}

func (obj *Link) SetTitle(value *string) {
    obj.Title = value
}

func (obj *Link) GetType() *string {
    return obj.Type
}

func (obj *Link) SetType(value *string) {
    obj.Type = value
}

