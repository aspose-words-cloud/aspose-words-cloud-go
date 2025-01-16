/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="table_cell_format.go">
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

// DTO container with all formatting for a table row.

type ITableCellFormat interface {
    IsTableCellFormat() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetBottomPadding() *float64
    SetBottomPadding(value *float64)
    GetFitText() *bool
    SetFitText(value *bool)
    GetHorizontalMerge() *string
    SetHorizontalMerge(value *string)
    GetLeftPadding() *float64
    SetLeftPadding(value *float64)
    GetOrientation() *string
    SetOrientation(value *string)
    GetPreferredWidth() IPreferredWidth
    SetPreferredWidth(value IPreferredWidth)
    GetRightPadding() *float64
    SetRightPadding(value *float64)
    GetTopPadding() *float64
    SetTopPadding(value *float64)
    GetVerticalAlignment() *string
    SetVerticalAlignment(value *string)
    GetVerticalMerge() *string
    SetVerticalMerge(value *string)
    GetWidth() *float64
    SetWidth(value *float64)
    GetWrapText() *bool
    SetWrapText(value *bool)
}

type TableCellFormat struct {
    // DTO container with all formatting for a table row.
    Link IWordsApiLink `json:"Link,omitempty"`

    // DTO container with all formatting for a table row.
    BottomPadding *float64 `json:"BottomPadding,omitempty"`

    // DTO container with all formatting for a table row.
    FitText *bool `json:"FitText,omitempty"`

    // DTO container with all formatting for a table row.
    HorizontalMerge *string `json:"HorizontalMerge,omitempty"`

    // DTO container with all formatting for a table row.
    LeftPadding *float64 `json:"LeftPadding,omitempty"`

    // DTO container with all formatting for a table row.
    Orientation *string `json:"Orientation,omitempty"`

    // DTO container with all formatting for a table row.
    PreferredWidth IPreferredWidth `json:"PreferredWidth,omitempty"`

    // DTO container with all formatting for a table row.
    RightPadding *float64 `json:"RightPadding,omitempty"`

    // DTO container with all formatting for a table row.
    TopPadding *float64 `json:"TopPadding,omitempty"`

    // DTO container with all formatting for a table row.
    VerticalAlignment *string `json:"VerticalAlignment,omitempty"`

    // DTO container with all formatting for a table row.
    VerticalMerge *string `json:"VerticalMerge,omitempty"`

    // DTO container with all formatting for a table row.
    Width *float64 `json:"Width,omitempty"`

    // DTO container with all formatting for a table row.
    WrapText *bool `json:"WrapText,omitempty"`
}

func (TableCellFormat) IsTableCellFormat() bool {
    return true
}

func (TableCellFormat) IsLinkElement() bool {
    return true
}

func (obj *TableCellFormat) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.PreferredWidth != nil) {
        obj.PreferredWidth.Initialize()
    }


}

func (obj *TableCellFormat) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["BottomPadding"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.BottomPadding = &parsedValue
        }

    } else if jsonValue, exists := json["bottomPadding"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.BottomPadding = &parsedValue
        }

    }

    if jsonValue, exists := json["FitText"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.FitText = &parsedValue
        }

    } else if jsonValue, exists := json["fitText"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.FitText = &parsedValue
        }

    }

    if jsonValue, exists := json["HorizontalMerge"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.HorizontalMerge = &parsedValue
        }

    } else if jsonValue, exists := json["horizontalMerge"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.HorizontalMerge = &parsedValue
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

    if jsonValue, exists := json["Orientation"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Orientation = &parsedValue
        }

    } else if jsonValue, exists := json["orientation"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Orientation = &parsedValue
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

    if jsonValue, exists := json["TopPadding"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.TopPadding = &parsedValue
        }

    } else if jsonValue, exists := json["topPadding"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.TopPadding = &parsedValue
        }

    }

    if jsonValue, exists := json["VerticalAlignment"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.VerticalAlignment = &parsedValue
        }

    } else if jsonValue, exists := json["verticalAlignment"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.VerticalAlignment = &parsedValue
        }

    }

    if jsonValue, exists := json["VerticalMerge"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.VerticalMerge = &parsedValue
        }

    } else if jsonValue, exists := json["verticalMerge"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.VerticalMerge = &parsedValue
        }

    }

    if jsonValue, exists := json["Width"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Width = &parsedValue
        }

    } else if jsonValue, exists := json["width"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Width = &parsedValue
        }

    }

    if jsonValue, exists := json["WrapText"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.WrapText = &parsedValue
        }

    } else if jsonValue, exists := json["wrapText"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.WrapText = &parsedValue
        }

    }
}

func (obj *TableCellFormat) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *TableCellFormat) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Link != nil {
        if err := obj.Link.Validate(); err != nil {
            return err
        }
    }
    if obj.PreferredWidth != nil {
        if err := obj.PreferredWidth.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *TableCellFormat) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *TableCellFormat) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *TableCellFormat) GetBottomPadding() *float64 {
    return obj.BottomPadding
}

func (obj *TableCellFormat) SetBottomPadding(value *float64) {
    obj.BottomPadding = value
}

func (obj *TableCellFormat) GetFitText() *bool {
    return obj.FitText
}

func (obj *TableCellFormat) SetFitText(value *bool) {
    obj.FitText = value
}

func (obj *TableCellFormat) GetHorizontalMerge() *string {
    return obj.HorizontalMerge
}

func (obj *TableCellFormat) SetHorizontalMerge(value *string) {
    obj.HorizontalMerge = value
}

func (obj *TableCellFormat) GetLeftPadding() *float64 {
    return obj.LeftPadding
}

func (obj *TableCellFormat) SetLeftPadding(value *float64) {
    obj.LeftPadding = value
}

func (obj *TableCellFormat) GetOrientation() *string {
    return obj.Orientation
}

func (obj *TableCellFormat) SetOrientation(value *string) {
    obj.Orientation = value
}

func (obj *TableCellFormat) GetPreferredWidth() IPreferredWidth {
    return obj.PreferredWidth
}

func (obj *TableCellFormat) SetPreferredWidth(value IPreferredWidth) {
    obj.PreferredWidth = value
}

func (obj *TableCellFormat) GetRightPadding() *float64 {
    return obj.RightPadding
}

func (obj *TableCellFormat) SetRightPadding(value *float64) {
    obj.RightPadding = value
}

func (obj *TableCellFormat) GetTopPadding() *float64 {
    return obj.TopPadding
}

func (obj *TableCellFormat) SetTopPadding(value *float64) {
    obj.TopPadding = value
}

func (obj *TableCellFormat) GetVerticalAlignment() *string {
    return obj.VerticalAlignment
}

func (obj *TableCellFormat) SetVerticalAlignment(value *string) {
    obj.VerticalAlignment = value
}

func (obj *TableCellFormat) GetVerticalMerge() *string {
    return obj.VerticalMerge
}

func (obj *TableCellFormat) SetVerticalMerge(value *string) {
    obj.VerticalMerge = value
}

func (obj *TableCellFormat) GetWidth() *float64 {
    return obj.Width
}

func (obj *TableCellFormat) SetWidth(value *float64) {
    obj.Width = value
}

func (obj *TableCellFormat) GetWrapText() *bool {
    return obj.WrapText
}

func (obj *TableCellFormat) SetWrapText(value *bool) {
    obj.WrapText = value
}

