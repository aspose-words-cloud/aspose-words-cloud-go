/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="table_properties.go">
 *   Copyright (c) 2022 Aspose.Words for Cloud
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

// DTO container with table properties.
type TablePropertiesResult struct {
    // DTO container with table properties.
    Link WordsApiLinkResult `json:"Link,omitempty"`

    // DTO container with table properties.
    Alignment string `json:"Alignment,omitempty"`

    // DTO container with table properties.
    AllowAutoFit bool `json:"AllowAutoFit,omitempty"`

    // DTO container with table properties.
    Bidi bool `json:"Bidi,omitempty"`

    // DTO container with table properties.
    BottomPadding float64 `json:"BottomPadding,omitempty"`

    // DTO container with table properties.
    CellSpacing float64 `json:"CellSpacing,omitempty"`

    // DTO container with table properties.
    LeftIndent float64 `json:"LeftIndent,omitempty"`

    // DTO container with table properties.
    LeftPadding float64 `json:"LeftPadding,omitempty"`

    // DTO container with table properties.
    PreferredWidth PreferredWidthResult `json:"PreferredWidth,omitempty"`

    // DTO container with table properties.
    RightPadding float64 `json:"RightPadding,omitempty"`

    // DTO container with table properties.
    StyleIdentifier string `json:"StyleIdentifier,omitempty"`

    // DTO container with table properties.
    StyleName string `json:"StyleName,omitempty"`

    // DTO container with table properties.
    StyleOptions string `json:"StyleOptions,omitempty"`

    // DTO container with table properties.
    TextWrapping string `json:"TextWrapping,omitempty"`

    // DTO container with table properties.
    TopPadding float64 `json:"TopPadding,omitempty"`
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

type ITableProperties interface {
    IsTableProperties() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileContent) []FileContent
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

func (obj *TableProperties) CollectFilesContent(resultFilesContent []FileContent) []FileContent {
    if (obj.Link != nil) {
        resultFilesContent = obj.Link.CollectFilesContent(resultFilesContent)
    }








    if (obj.PreferredWidth != nil) {
        resultFilesContent = obj.PreferredWidth.CollectFilesContent(resultFilesContent)
    }







    return resultFilesContent
}


