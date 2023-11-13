/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="custom_xml_part.go">
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

// DTO container with a CustomXmlPart.

type ICustomXmlPart interface {
    IsCustomXmlPart() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetId() *string
    SetId(value *string)
    GetData() *string
    SetData(value *string)
}

type CustomXmlPart struct {
    // DTO container with a CustomXmlPart.
    Link IWordsApiLink `json:"Link,omitempty"`

    // DTO container with a CustomXmlPart.
    Id *string `json:"Id,omitempty"`

    // DTO container with a CustomXmlPart.
    Data *string `json:"Data,omitempty"`
}

func (CustomXmlPart) IsCustomXmlPart() bool {
    return true
}

func (CustomXmlPart) IsCustomXmlPartLink() bool {
    return true
}

func (CustomXmlPart) IsLinkElement() bool {
    return true
}

func (obj *CustomXmlPart) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }


}

func (obj *CustomXmlPart) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["Id"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Id = &parsedValue
        }

    } else if jsonValue, exists := json["id"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Id = &parsedValue
        }

    }

    if jsonValue, exists := json["Data"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Data = &parsedValue
        }

    } else if jsonValue, exists := json["data"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Data = &parsedValue
        }

    }
}

func (obj *CustomXmlPart) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *CustomXmlPart) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Data == nil {
        return errors.New("Property Data in CustomXmlPart is required.")
    }
    if obj.Link != nil {
        if err := obj.Link.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *CustomXmlPart) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *CustomXmlPart) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *CustomXmlPart) GetId() *string {
    return obj.Id
}

func (obj *CustomXmlPart) SetId(value *string) {
    obj.Id = value
}

func (obj *CustomXmlPart) GetData() *string {
    return obj.Data
}

func (obj *CustomXmlPart) SetData(value *string) {
    obj.Data = value
}

