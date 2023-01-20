/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="table_row_format.go">
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

// DTO container with formatting for a table row.
type TableRowFormatResult struct {
    // DTO container with formatting for a table row.
    Link WordsApiLinkResult `json:"Link,omitempty"`

    // DTO container with formatting for a table row.
    AllowBreakAcrossPages bool `json:"AllowBreakAcrossPages,omitempty"`

    // DTO container with formatting for a table row.
    HeadingFormat bool `json:"HeadingFormat,omitempty"`

    // DTO container with formatting for a table row.
    Height float64 `json:"Height,omitempty"`

    // DTO container with formatting for a table row.
    HeightRule string `json:"HeightRule,omitempty"`
}

type TableRowFormat struct {
    // DTO container with formatting for a table row.
    Link IWordsApiLink `json:"Link,omitempty"`

    // DTO container with formatting for a table row.
    AllowBreakAcrossPages *bool `json:"AllowBreakAcrossPages,omitempty"`

    // DTO container with formatting for a table row.
    HeadingFormat *bool `json:"HeadingFormat,omitempty"`

    // DTO container with formatting for a table row.
    Height *float64 `json:"Height,omitempty"`

    // DTO container with formatting for a table row.
    HeightRule *string `json:"HeightRule,omitempty"`
}

type ITableRowFormat interface {
    IsTableRowFormat() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
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

func (obj *TableRowFormat) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}


