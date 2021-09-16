/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="table_cell_format.go">
 *   Copyright (c) 2021 Aspose.Words for Cloud
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

// DTO container with all formatting for a table row.
type TableCellFormatResult struct {
    // DTO container with all formatting for a table row.
    Link WordsApiLinkResult `json:"Link,omitempty"`

    // DTO container with all formatting for a table row.
    BottomPadding float64 `json:"BottomPadding,omitempty"`

    // DTO container with all formatting for a table row.
    FitText bool `json:"FitText,omitempty"`

    // DTO container with all formatting for a table row.
    HorizontalMerge string `json:"HorizontalMerge,omitempty"`

    // DTO container with all formatting for a table row.
    LeftPadding float64 `json:"LeftPadding,omitempty"`

    // DTO container with all formatting for a table row.
    Orientation string `json:"Orientation,omitempty"`

    // DTO container with all formatting for a table row.
    PreferredWidth PreferredWidthResult `json:"PreferredWidth,omitempty"`

    // DTO container with all formatting for a table row.
    RightPadding float64 `json:"RightPadding,omitempty"`

    // DTO container with all formatting for a table row.
    TopPadding float64 `json:"TopPadding,omitempty"`

    // DTO container with all formatting for a table row.
    VerticalAlignment string `json:"VerticalAlignment,omitempty"`

    // DTO container with all formatting for a table row.
    VerticalMerge string `json:"VerticalMerge,omitempty"`

    // DTO container with all formatting for a table row.
    Width float64 `json:"Width,omitempty"`

    // DTO container with all formatting for a table row.
    WrapText bool `json:"WrapText,omitempty"`
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

type ITableCellFormat interface {
    IsTableCellFormat() bool
}
func (TableCellFormat) IsTableCellFormat() bool {
    return true
}

func (TableCellFormat) IsLinkElement() bool {
    return true
}


