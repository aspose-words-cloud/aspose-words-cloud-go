/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="custom_xml_part_update.go">
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

// Custom xml part update.

type ICustomXmlPartUpdate interface {
    IsCustomXmlPartUpdate() bool
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

type CustomXmlPartUpdate struct {
    // Custom xml part update.
    Link IWordsApiLink `json:"Link,omitempty"`

    // Custom xml part update.
    Id *string `json:"Id,omitempty"`

    // Custom xml part update.
    Data *string `json:"Data,omitempty"`
}

func (CustomXmlPartUpdate) IsCustomXmlPartUpdate() bool {
    return true
}

func (CustomXmlPartUpdate) IsCustomXmlPart() bool {
    return true
}

func (CustomXmlPartUpdate) IsCustomXmlPartLink() bool {
    return true
}

func (CustomXmlPartUpdate) IsLinkElement() bool {
    return true
}

func (obj *CustomXmlPartUpdate) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }


}

func (obj *CustomXmlPartUpdate) Deserialize(json map[string]interface{}) {
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

func (obj *CustomXmlPartUpdate) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *CustomXmlPartUpdate) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Data == nil {
        return errors.New("Property Data in CustomXmlPartUpdate is required.")
    }
    if obj.Link != nil {
        if err := obj.Link.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *CustomXmlPartUpdate) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *CustomXmlPartUpdate) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *CustomXmlPartUpdate) GetId() *string {
    return obj.Id
}

func (obj *CustomXmlPartUpdate) SetId(value *string) {
    obj.Id = value
}

func (obj *CustomXmlPartUpdate) GetData() *string {
    return obj.Data
}

func (obj *CustomXmlPartUpdate) SetData(value *string) {
    obj.Data = value
}

