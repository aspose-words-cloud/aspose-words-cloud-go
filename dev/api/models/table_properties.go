/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="table_properties.go">
 *   Copyright (c) 2020 Aspose.Words for Cloud
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

// Represents the table properties.
type TableProperties struct {
    // Represents the table properties.
    Link *WordsApiLink `json:"Link,omitempty"`

    // Represents the table properties.
    Alignment string `json:"Alignment,omitempty"`

    // Represents the table properties.
    AllowAutoFit *bool `json:"AllowAutoFit,omitempty"`

    // Represents the table properties.
    Bidi *bool `json:"Bidi,omitempty"`

    // Represents the table properties.
    BottomPadding *float64 `json:"BottomPadding,omitempty"`

    // Represents the table properties.
    CellSpacing *float64 `json:"CellSpacing,omitempty"`

    // Represents the table properties.
    LeftIndent *float64 `json:"LeftIndent,omitempty"`

    // Represents the table properties.
    LeftPadding *float64 `json:"LeftPadding,omitempty"`

    // Represents the table properties.
    PreferredWidth *PreferredWidth `json:"PreferredWidth,omitempty"`

    // Represents the table properties.
    RightPadding *float64 `json:"RightPadding,omitempty"`

    // Represents the table properties.
    StyleIdentifier string `json:"StyleIdentifier,omitempty"`

    // Represents the table properties.
    StyleName *string `json:"StyleName,omitempty"`

    // Represents the table properties.
    StyleOptions string `json:"StyleOptions,omitempty"`

    // Represents the table properties.
    TextWrapping string `json:"TextWrapping,omitempty"`

    // Represents the table properties.
    TopPadding *float64 `json:"TopPadding,omitempty"`
}

type ITableProperties interface {
    IsTableProperties() bool
}
func (TableProperties) IsTableProperties() bool {
    return true
}

func (TableProperties) IsLinkElement() bool {
    return true
}
