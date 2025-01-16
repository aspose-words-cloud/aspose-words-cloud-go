/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="words_api_link.go">
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

// Provides information for the words API resource link.

type IWordsApiLink interface {
    IsWordsApiLink() bool
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

type WordsApiLink struct {
    // Provides information for the words API resource link.
    Href *string `json:"Href,omitempty"`

    // Provides information for the words API resource link.
    Rel *string `json:"Rel,omitempty"`

    // Provides information for the words API resource link.
    Title *string `json:"Title,omitempty"`

    // Provides information for the words API resource link.
    Type *string `json:"Type,omitempty"`
}

func (WordsApiLink) IsWordsApiLink() bool {
    return true
}

func (WordsApiLink) IsLink() bool {
    return true
}

func (obj *WordsApiLink) Initialize() {
}

func (obj *WordsApiLink) Deserialize(json map[string]interface{}) {
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

func (obj *WordsApiLink) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *WordsApiLink) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    return nil;
}

func (obj *WordsApiLink) GetHref() *string {
    return obj.Href
}

func (obj *WordsApiLink) SetHref(value *string) {
    obj.Href = value
}

func (obj *WordsApiLink) GetRel() *string {
    return obj.Rel
}

func (obj *WordsApiLink) SetRel(value *string) {
    obj.Rel = value
}

func (obj *WordsApiLink) GetTitle() *string {
    return obj.Title
}

func (obj *WordsApiLink) SetTitle(value *string) {
    obj.Title = value
}

func (obj *WordsApiLink) GetType() *string {
    return obj.Type
}

func (obj *WordsApiLink) SetType(value *string) {
    obj.Type = value
}

