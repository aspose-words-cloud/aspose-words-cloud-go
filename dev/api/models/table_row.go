/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="table_row.go">
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

// DTO container with a table row element.
type TableRowResult struct {
    // DTO container with a table row element.
    Link WordsApiLinkResult `json:"Link,omitempty"`

    // DTO container with a table row element.
    NodeId string `json:"NodeId,omitempty"`

    // DTO container with a table row element.
    RowFormat TableRowFormatResult `json:"RowFormat,omitempty"`

    // DTO container with a table row element.
    TableCellList []TableCellResult `json:"TableCellList,omitempty"`
}

type TableRow struct {
    // DTO container with a table row element.
    Link IWordsApiLink `json:"Link,omitempty"`

    // DTO container with a table row element.
    NodeId *string `json:"NodeId,omitempty"`

    // DTO container with a table row element.
    RowFormat ITableRowFormat `json:"RowFormat,omitempty"`

    // DTO container with a table row element.
    TableCellList []TableCell `json:"TableCellList,omitempty"`
}

type ITableRow interface {
    IsTableRow() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileContent) []FileContent
}

func (TableRow) IsTableRow() bool {
    return true
}

func (TableRow) IsNodeLink() bool {
    return true
}

func (TableRow) IsLinkElement() bool {
    return true
}

func (obj *TableRow) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.RowFormat != nil) {
        obj.RowFormat.Initialize()
    }

    if (obj.TableCellList != nil) {
        for _, element := range obj.TableCellList {
            element.Initialize()
        }
    }


}

func (obj *TableRow) CollectFilesContent(resultFilesContent []FileContent) []FileContent {
    if (obj.Link != nil) {
        resultFilesContent = obj.Link.CollectFilesContent(resultFilesContent)
    }


    if (obj.RowFormat != nil) {
        resultFilesContent = obj.RowFormat.CollectFilesContent(resultFilesContent)
    }

    if (obj.TableCellList != nil) {
        for _, element := range obj.TableCellList {
            resultFilesContent = element.CollectFilesContent(resultFilesContent)
        }
    }

    return resultFilesContent
}


