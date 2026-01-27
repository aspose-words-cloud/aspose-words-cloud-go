/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="page_number.go">
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

// Class is used for insert page number request building.

type IPageNumber interface {
    IsPageNumber() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetAlignment() *string
    SetAlignment(value *string)
    GetFormat() *string
    SetFormat(value *string)
    GetIsTop() *bool
    SetIsTop(value *bool)
    GetPageStartingNumber() *int32
    SetPageStartingNumber(value *int32)
    GetSetPageNumberOnFirstPage() *bool
    SetSetPageNumberOnFirstPage(value *bool)
}

type PageNumber struct {
    // Class is used for insert page number request building.
    Alignment *string `json:"Alignment,omitempty"`

    // Class is used for insert page number request building.
    Format *string `json:"Format,omitempty"`

    // Class is used for insert page number request building.
    IsTop *bool `json:"IsTop,omitempty"`

    // Class is used for insert page number request building.
    PageStartingNumber *int32 `json:"PageStartingNumber,omitempty"`

    // Class is used for insert page number request building.
    SetPageNumberOnFirstPage *bool `json:"SetPageNumberOnFirstPage,omitempty"`
}

func (PageNumber) IsPageNumber() bool {
    return true
}


func (obj *PageNumber) Initialize() {
}

func (obj *PageNumber) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["Alignment"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Alignment = &parsedValue
        }

    } else if jsonValue, exists := json["alignment"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Alignment = &parsedValue
        }

    }

    if jsonValue, exists := json["Format"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Format = &parsedValue
        }

    } else if jsonValue, exists := json["format"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Format = &parsedValue
        }

    }

    if jsonValue, exists := json["IsTop"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsTop = &parsedValue
        }

    } else if jsonValue, exists := json["isTop"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsTop = &parsedValue
        }

    }

    if jsonValue, exists := json["PageStartingNumber"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.PageStartingNumber = new(int32)
            *obj.PageStartingNumber = int32(parsedValue)
        }

    } else if jsonValue, exists := json["pageStartingNumber"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.PageStartingNumber = new(int32)
            *obj.PageStartingNumber = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["SetPageNumberOnFirstPage"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.SetPageNumberOnFirstPage = &parsedValue
        }

    } else if jsonValue, exists := json["setPageNumberOnFirstPage"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.SetPageNumberOnFirstPage = &parsedValue
        }

    }
}

func (obj *PageNumber) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *PageNumber) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.IsTop == nil {
        return errors.New("Property IsTop in PageNumber is required.")
    }
    if obj.SetPageNumberOnFirstPage == nil {
        return errors.New("Property SetPageNumberOnFirstPage in PageNumber is required.")
    }
    return nil;
}

func (obj *PageNumber) GetAlignment() *string {
    return obj.Alignment
}

func (obj *PageNumber) SetAlignment(value *string) {
    obj.Alignment = value
}

func (obj *PageNumber) GetFormat() *string {
    return obj.Format
}

func (obj *PageNumber) SetFormat(value *string) {
    obj.Format = value
}

func (obj *PageNumber) GetIsTop() *bool {
    return obj.IsTop
}

func (obj *PageNumber) SetIsTop(value *bool) {
    obj.IsTop = value
}

func (obj *PageNumber) GetPageStartingNumber() *int32 {
    return obj.PageStartingNumber
}

func (obj *PageNumber) SetPageStartingNumber(value *int32) {
    obj.PageStartingNumber = value
}

func (obj *PageNumber) GetSetPageNumberOnFirstPage() *bool {
    return obj.SetPageNumberOnFirstPage
}

func (obj *PageNumber) SetSetPageNumberOnFirstPage(value *bool) {
    obj.SetPageNumberOnFirstPage = value
}

