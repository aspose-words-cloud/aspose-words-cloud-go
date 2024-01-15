/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="table_row_format.go">
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

// DTO container with formatting for a table row.

type ITableRowFormat interface {
    IsTableRowFormat() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetHeight() *float64
    SetHeight(value *float64)
    GetHeightRule() *string
    SetHeightRule(value *string)
    GetAllowBreakAcrossPages() *bool
    SetAllowBreakAcrossPages(value *bool)
    GetHeadingFormat() *bool
    SetHeadingFormat(value *bool)
}

type TableRowFormat struct {
    // DTO container with formatting for a table row.
    Link IWordsApiLink `json:"Link,omitempty"`

    // DTO container with formatting for a table row.
    Height *float64 `json:"Height,omitempty"`

    // DTO container with formatting for a table row.
    HeightRule *string `json:"HeightRule,omitempty"`

    // DTO container with formatting for a table row.
    AllowBreakAcrossPages *bool `json:"AllowBreakAcrossPages,omitempty"`

    // DTO container with formatting for a table row.
    HeadingFormat *bool `json:"HeadingFormat,omitempty"`
}

func (TableRowFormat) IsTableRowFormat() bool {
    return true
}

func (TableRowFormat) IsLinkElement() bool {
    return true
}

func (obj *TableRowFormat) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }


}

func (obj *TableRowFormat) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["Height"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Height = &parsedValue
        }

    } else if jsonValue, exists := json["height"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Height = &parsedValue
        }

    }

    if jsonValue, exists := json["HeightRule"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.HeightRule = &parsedValue
        }

    } else if jsonValue, exists := json["heightRule"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.HeightRule = &parsedValue
        }

    }

    if jsonValue, exists := json["AllowBreakAcrossPages"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.AllowBreakAcrossPages = &parsedValue
        }

    } else if jsonValue, exists := json["allowBreakAcrossPages"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.AllowBreakAcrossPages = &parsedValue
        }

    }

    if jsonValue, exists := json["HeadingFormat"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.HeadingFormat = &parsedValue
        }

    } else if jsonValue, exists := json["headingFormat"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.HeadingFormat = &parsedValue
        }

    }
}

func (obj *TableRowFormat) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *TableRowFormat) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Link != nil {
        if err := obj.Link.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *TableRowFormat) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *TableRowFormat) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *TableRowFormat) GetHeight() *float64 {
    return obj.Height
}

func (obj *TableRowFormat) SetHeight(value *float64) {
    obj.Height = value
}

func (obj *TableRowFormat) GetHeightRule() *string {
    return obj.HeightRule
}

func (obj *TableRowFormat) SetHeightRule(value *string) {
    obj.HeightRule = value
}

func (obj *TableRowFormat) GetAllowBreakAcrossPages() *bool {
    return obj.AllowBreakAcrossPages
}

func (obj *TableRowFormat) SetAllowBreakAcrossPages(value *bool) {
    obj.AllowBreakAcrossPages = value
}

func (obj *TableRowFormat) GetHeadingFormat() *bool {
    return obj.HeadingFormat
}

func (obj *TableRowFormat) SetHeadingFormat(value *bool) {
    obj.HeadingFormat = value
}

