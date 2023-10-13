/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="hyperlink.go">
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

// Hyperlink element.

type IHyperlink interface {
    IsHyperlink() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetDisplayText() *string
    SetDisplayText(value *string)
    GetValue() *string
    SetValue(value *string)
}

type Hyperlink struct {
    // Hyperlink element.
    Link IWordsApiLink `json:"Link,omitempty"`

    // Hyperlink element.
    DisplayText *string `json:"DisplayText,omitempty"`

    // Hyperlink element.
    Value *string `json:"Value,omitempty"`
}

func (Hyperlink) IsHyperlink() bool {
    return true
}

func (Hyperlink) IsLinkElement() bool {
    return true
}

func (obj *Hyperlink) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }


}

func (obj *Hyperlink) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["DisplayText"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DisplayText = &parsedValue
        }

    } else if jsonValue, exists := json["displayText"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DisplayText = &parsedValue
        }

    }

    if jsonValue, exists := json["Value"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Value = &parsedValue
        }

    } else if jsonValue, exists := json["value"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Value = &parsedValue
        }

    }
}

func (obj *Hyperlink) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *Hyperlink) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    return nil;
}

func (obj *Hyperlink) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *Hyperlink) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *Hyperlink) GetDisplayText() *string {
    return obj.DisplayText
}

func (obj *Hyperlink) SetDisplayText(value *string) {
    obj.DisplayText = value
}

func (obj *Hyperlink) GetValue() *string {
    return obj.Value
}

func (obj *Hyperlink) SetValue(value *string) {
    obj.Value = value
}

