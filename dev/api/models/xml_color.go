/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="xml_color.go">
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

// Utility class for Color serialization.

type IXmlColor interface {
    IsXmlColor() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetAlpha() *int32
    SetAlpha(value *int32)
    GetWeb() *string
    SetWeb(value *string)
    GetXmlAlpha() *int32
    SetXmlAlpha(value *int32)
    GetXmlAlphaSpecified() *bool
    SetXmlAlphaSpecified(value *bool)
}

type XmlColor struct {
    // Utility class for Color serialization.
    Alpha *int32 `json:"Alpha,omitempty"`

    // Utility class for Color serialization.
    Web *string `json:"Web,omitempty"`

    // Utility class for Color serialization.
    XmlAlpha *int32 `json:"XmlAlpha,omitempty"`

    // Utility class for Color serialization.
    XmlAlphaSpecified *bool `json:"XmlAlphaSpecified,omitempty"`
}

func (XmlColor) IsXmlColor() bool {
    return true
}


func (obj *XmlColor) Initialize() {
}

func (obj *XmlColor) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["Alpha"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Alpha = new(int32)
            *obj.Alpha = int32(parsedValue)
        }

    } else if jsonValue, exists := json["alpha"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Alpha = new(int32)
            *obj.Alpha = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["Web"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Web = &parsedValue
        }

    } else if jsonValue, exists := json["web"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Web = &parsedValue
        }

    }

    if jsonValue, exists := json["XmlAlpha"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.XmlAlpha = new(int32)
            *obj.XmlAlpha = int32(parsedValue)
        }

    } else if jsonValue, exists := json["xmlAlpha"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.XmlAlpha = new(int32)
            *obj.XmlAlpha = int32(parsedValue)
        }

    }
}

func (obj *XmlColor) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *XmlColor) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.XmlAlpha == nil {
        return errors.New("Property XmlAlpha in XmlColor is required.")
    }

    if obj.XmlAlphaSpecified == nil {
        return errors.New("Property XmlAlphaSpecified in XmlColor is required.")
    }

    return nil;
}

func (obj *XmlColor) GetAlpha() *int32 {
    return obj.Alpha
}

func (obj *XmlColor) SetAlpha(value *int32) {
    obj.Alpha = value
}

func (obj *XmlColor) GetWeb() *string {
    return obj.Web
}

func (obj *XmlColor) SetWeb(value *string) {
    obj.Web = value
}

func (obj *XmlColor) GetXmlAlpha() *int32 {
    return obj.XmlAlpha
}

func (obj *XmlColor) SetXmlAlpha(value *int32) {
    obj.XmlAlpha = value
}

func (obj *XmlColor) GetXmlAlphaSpecified() *bool {
    return obj.XmlAlphaSpecified
}

func (obj *XmlColor) SetXmlAlphaSpecified(value *bool) {
    obj.XmlAlphaSpecified = value
}

