/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="table_cell_response.go">
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

// The REST response with a table cell.

type ITableCellResponse interface {
    IsTableCellResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetRequestId() *string
    SetRequestId(value *string)
    GetCell() ITableCell
    SetCell(value ITableCell)
}

type TableCellResponse struct {
    // The REST response with a table cell.
    RequestId *string

    // The REST response with a table cell.
    Cell ITableCell
}

func (TableCellResponse) IsTableCellResponse() bool {
    return true
}

func (TableCellResponse) IsWordsResponse() bool {
    return true
}

func (obj *TableCellResponse) Initialize() {
    if (obj.Cell != nil) {
        obj.Cell.Initialize()
    }


}

func (obj *TableCellResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["Cell"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ITableCell = new(TableCell)
            modelInstance.Deserialize(parsedValue)
            obj.Cell = modelInstance
        }

    } else if jsonValue, exists := json["cell"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ITableCell = new(TableCell)
            modelInstance.Deserialize(parsedValue)
            obj.Cell = modelInstance
        }

    }
}

func (obj *TableCellResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *TableCellResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *TableCellResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *TableCellResponse) GetCell() ITableCell {
    return obj.Cell
}

func (obj *TableCellResponse) SetCell(value ITableCell) {
    obj.Cell = value
}

