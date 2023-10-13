/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="table_properties.go">
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

// DTO container with table properties.

type ITableProperties interface {
    IsTableProperties() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetAlignment() *string
    SetAlignment(value *string)
    GetAllowAutoFit() *bool
    SetAllowAutoFit(value *bool)
    GetBidi() *bool
    SetBidi(value *bool)
    GetBottomPadding() *float64
    SetBottomPadding(value *float64)
    GetCellSpacing() *float64
    SetCellSpacing(value *float64)
    GetLeftIndent() *float64
    SetLeftIndent(value *float64)
    GetLeftPadding() *float64
    SetLeftPadding(value *float64)
    GetPreferredWidth() IPreferredWidth
    SetPreferredWidth(value IPreferredWidth)
    GetRightPadding() *float64
    SetRightPadding(value *float64)
    GetStyleIdentifier() *string
    SetStyleIdentifier(value *string)
    GetStyleName() *string
    SetStyleName(value *string)
    GetStyleOptions() *string
    SetStyleOptions(value *string)
    GetTextWrapping() *string
    SetTextWrapping(value *string)
    GetTopPadding() *float64
    SetTopPadding(value *float64)
}

type TableProperties struct {
    // DTO container with table properties.
    Link IWordsApiLink `json:"Link,omitempty"`

    // DTO container with table properties.
    Alignment *string `json:"Alignment,omitempty"`

    // DTO container with table properties.
    AllowAutoFit *bool `json:"AllowAutoFit,omitempty"`

    // DTO container with table properties.
    Bidi *bool `json:"Bidi,omitempty"`

    // DTO container with table properties.
    BottomPadding *float64 `json:"BottomPadding,omitempty"`

    // DTO container with table properties.
    CellSpacing *float64 `json:"CellSpacing,omitempty"`

    // DTO container with table properties.
    LeftIndent *float64 `json:"LeftIndent,omitempty"`

    // DTO container with table properties.
    LeftPadding *float64 `json:"LeftPadding,omitempty"`

    // DTO container with table properties.
    PreferredWidth IPreferredWidth `json:"PreferredWidth,omitempty"`

    // DTO container with table properties.
    RightPadding *float64 `json:"RightPadding,omitempty"`

    // DTO container with table properties.
    StyleIdentifier *string `json:"StyleIdentifier,omitempty"`

    // DTO container with table properties.
    StyleName *string `json:"StyleName,omitempty"`

    // DTO container with table properties.
    StyleOptions *string `json:"StyleOptions,omitempty"`

    // DTO container with table properties.
    TextWrapping *string `json:"TextWrapping,omitempty"`

    // DTO container with table properties.
    TopPadding *float64 `json:"TopPadding,omitempty"`
}

func (TableProperties) IsTableProperties() bool {
    return true
}

func (TableProperties) IsLinkElement() bool {
    return true
}

func (obj *TableProperties) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.PreferredWidth != nil) {
        obj.PreferredWidth.Initialize()
    }


}

func (obj *TableProperties) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["Alignment"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Alignment = &parsedValue
        }

    } else if jsonValue, exists := json["alignment"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Alignment = &parsedValue
        }

    }

    if jsonValue, exists := json["AllowAutoFit"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.AllowAutoFit = &parsedValue
        }

    } else if jsonValue, exists := json["allowAutoFit"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.AllowAutoFit = &parsedValue
        }

    }

    if jsonValue, exists := json["Bidi"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.Bidi = &parsedValue
        }

    } else if jsonValue, exists := json["bidi"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.Bidi = &parsedValue
        }

    }

    if jsonValue, exists := json["BottomPadding"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.BottomPadding = &parsedValue
        }

    } else if jsonValue, exists := json["bottomPadding"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.BottomPadding = &parsedValue
        }

    }

    if jsonValue, exists := json["CellSpacing"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.CellSpacing = &parsedValue
        }

    } else if jsonValue, exists := json["cellSpacing"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.CellSpacing = &parsedValue
        }

    }

    if jsonValue, exists := json["LeftIndent"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.LeftIndent = &parsedValue
        }

    } else if jsonValue, exists := json["leftIndent"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.LeftIndent = &parsedValue
        }

    }

    if jsonValue, exists := json["LeftPadding"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.LeftPadding = &parsedValue
        }

    } else if jsonValue, exists := json["leftPadding"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.LeftPadding = &parsedValue
        }

    }

    if jsonValue, exists := json["PreferredWidth"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IPreferredWidth = new(PreferredWidth)
            modelInstance.Deserialize(parsedValue)
            obj.PreferredWidth = modelInstance
        }

    } else if jsonValue, exists := json["preferredWidth"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IPreferredWidth = new(PreferredWidth)
            modelInstance.Deserialize(parsedValue)
            obj.PreferredWidth = modelInstance
        }

    }

    if jsonValue, exists := json["RightPadding"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.RightPadding = &parsedValue
        }

    } else if jsonValue, exists := json["rightPadding"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.RightPadding = &parsedValue
        }

    }

    if jsonValue, exists := json["StyleIdentifier"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.StyleIdentifier = &parsedValue
        }

    } else if jsonValue, exists := json["styleIdentifier"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.StyleIdentifier = &parsedValue
        }

    }

    if jsonValue, exists := json["StyleName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.StyleName = &parsedValue
        }

    } else if jsonValue, exists := json["styleName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.StyleName = &parsedValue
        }

    }

    if jsonValue, exists := json["StyleOptions"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.StyleOptions = &parsedValue
        }

    } else if jsonValue, exists := json["styleOptions"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.StyleOptions = &parsedValue
        }

    }

    if jsonValue, exists := json["TextWrapping"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.TextWrapping = &parsedValue
        }

    } else if jsonValue, exists := json["textWrapping"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.TextWrapping = &parsedValue
        }

    }

    if jsonValue, exists := json["TopPadding"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.TopPadding = &parsedValue
        }

    } else if jsonValue, exists := json["topPadding"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.TopPadding = &parsedValue
        }

    }
}

func (obj *TableProperties) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *TableProperties) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    return nil;
}

func (obj *TableProperties) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *TableProperties) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *TableProperties) GetAlignment() *string {
    return obj.Alignment
}

func (obj *TableProperties) SetAlignment(value *string) {
    obj.Alignment = value
}

func (obj *TableProperties) GetAllowAutoFit() *bool {
    return obj.AllowAutoFit
}

func (obj *TableProperties) SetAllowAutoFit(value *bool) {
    obj.AllowAutoFit = value
}

func (obj *TableProperties) GetBidi() *bool {
    return obj.Bidi
}

func (obj *TableProperties) SetBidi(value *bool) {
    obj.Bidi = value
}

func (obj *TableProperties) GetBottomPadding() *float64 {
    return obj.BottomPadding
}

func (obj *TableProperties) SetBottomPadding(value *float64) {
    obj.BottomPadding = value
}

func (obj *TableProperties) GetCellSpacing() *float64 {
    return obj.CellSpacing
}

func (obj *TableProperties) SetCellSpacing(value *float64) {
    obj.CellSpacing = value
}

func (obj *TableProperties) GetLeftIndent() *float64 {
    return obj.LeftIndent
}

func (obj *TableProperties) SetLeftIndent(value *float64) {
    obj.LeftIndent = value
}

func (obj *TableProperties) GetLeftPadding() *float64 {
    return obj.LeftPadding
}

func (obj *TableProperties) SetLeftPadding(value *float64) {
    obj.LeftPadding = value
}

func (obj *TableProperties) GetPreferredWidth() IPreferredWidth {
    return obj.PreferredWidth
}

func (obj *TableProperties) SetPreferredWidth(value IPreferredWidth) {
    obj.PreferredWidth = value
}

func (obj *TableProperties) GetRightPadding() *float64 {
    return obj.RightPadding
}

func (obj *TableProperties) SetRightPadding(value *float64) {
    obj.RightPadding = value
}

func (obj *TableProperties) GetStyleIdentifier() *string {
    return obj.StyleIdentifier
}

func (obj *TableProperties) SetStyleIdentifier(value *string) {
    obj.StyleIdentifier = value
}

func (obj *TableProperties) GetStyleName() *string {
    return obj.StyleName
}

func (obj *TableProperties) SetStyleName(value *string) {
    obj.StyleName = value
}

func (obj *TableProperties) GetStyleOptions() *string {
    return obj.StyleOptions
}

func (obj *TableProperties) SetStyleOptions(value *string) {
    obj.StyleOptions = value
}

func (obj *TableProperties) GetTextWrapping() *string {
    return obj.TextWrapping
}

func (obj *TableProperties) SetTextWrapping(value *string) {
    obj.TextWrapping = value
}

func (obj *TableProperties) GetTopPadding() *float64 {
    return obj.TopPadding
}

func (obj *TableProperties) SetTopPadding(value *float64) {
    obj.TopPadding = value
}

